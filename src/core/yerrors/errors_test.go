package yerrors_test

import (
	"errors"
	"testing"

	err "git.tashilcar.com/core/tashilcar/core/yerrors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestE(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *err.Error
	}{
		{
			"Simple",
			args{args: []interface{}{err.Op("GL.findByID"), err.KindNotFound, errors.New("not found error")}},
			&err.Error{
				Op:   "GL.findByID",
				Kind: err.KindNotFound,
				Err:  errors.New("not found error"),
			},
		},
		{
			"Nested",
			args{args: []interface{}{
				err.Op("GL.found"),
				err.KindUnauthorized,
				&err.Error{
					Op:   "GL.found",
					Kind: err.KindNotFound,
					Err:  errors.New("GL not found error"),
				},
			}},
			&err.Error{
				Op:   "GL.found",
				Kind: err.KindUnauthorized,
				Err: &err.Error{
					Op:   "GL.found",
					Kind: err.KindNotFound,
					Err:  errors.New("GL not found error"),
				},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, err.E(tt.args.args...), tt.name)
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		Op   err.Op
		Kind codes.Code
		Err  error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Error with nested error",
			fields{
				Op:   "GL.findByID",
				Kind: err.KindNotFound,
				Err:  errors.New("GL not found"),
			},
			"GL not found",
		},
		{
			"Error with nested Error",
			fields{
				Op:   "GL.findByID",
				Kind: err.KindNotFound,
				Err:  err.E(err.Op("GL.findByID"), errors.New("unexpected error")),
			},
			"unexpected error",
		},
	}
	for _, tt := range tests {
		e := &err.Error{
			Op:   tt.fields.Op,
			Kind: tt.fields.Kind,
			Err:  tt.fields.Err,
		}
		assert.Equal(t, tt.want, e.Error(), tt.name)
	}
}

func TestOps(t *testing.T) {
	type args struct {
		e *err.Error
	}
	tests := []struct {
		name string
		args args
		want []err.Op
	}{
		{
			"Nested Errors",
			args{e: err.E(err.Op("GL.findByID"), err.E(err.Op("GL.getAll")))},
			[]err.Op{"GL.findByID", "GL.getAll"},
		},
		{
			"Error with nested error",
			args{e: err.E(err.Op("GL.findByID"), errors.New("unexpected error"))},
			[]err.Op{"GL.findByID"},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, err.Ops(tt.args.e), tt.name)
	}
}
