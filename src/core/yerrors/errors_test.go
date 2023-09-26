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
			args{args: []interface{}{err.Op("TC.findByID"), err.KindNotFound, errors.New("not found error")}},
			&err.Error{
				Op:   "TC.findByID",
				Kind: err.KindNotFound,
				Err:  errors.New("not found error"),
			},
		},
		{
			"Nested",
			args{args: []interface{}{
				err.Op("TC.found"),
				err.KindUnauthorized,
				&err.Error{
					Op:   "TC.found",
					Kind: err.KindNotFound,
					Err:  errors.New("TC not found error"),
				},
			}},
			&err.Error{
				Op:   "TC.found",
				Kind: err.KindUnauthorized,
				Err: &err.Error{
					Op:   "TC.found",
					Kind: err.KindNotFound,
					Err:  errors.New("TC not found error"),
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
				Op:   "TC.findByID",
				Kind: err.KindNotFound,
				Err:  errors.New("TC not found"),
			},
			"TC not found",
		},
		{
			"Error with nested Error",
			fields{
				Op:   "TC.findByID",
				Kind: err.KindNotFound,
				Err:  err.E(err.Op("TC.findByID"), errors.New("unexpected error")),
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
			args{e: err.E(err.Op("TC.findByID"), err.E(err.Op("TC.getAll")))},
			[]err.Op{"TC.findByID", "TC.getAll"},
		},
		{
			"Error with nested error",
			args{e: err.E(err.Op("TC.findByID"), errors.New("unexpected error"))},
			[]err.Op{"TC.findByID"},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, err.Ops(tt.args.e), tt.name)
	}
}
