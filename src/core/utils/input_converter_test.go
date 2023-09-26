package utils_test

import (
	"testing"

	"git.tashilcar.com/core/tashilcar/core/utils"
	"github.com/stretchr/testify/assert"
)

func TestInputConverter(t *testing.T) {
	assert.NotNil(t, utils.InputConverter, "InputConverter should not return nil")
}

func TestStringToUnsignedInt(t *testing.T) {
	tests := []struct {
		msg  string
		in   string
		want int
		pass bool
	}{
		{"should return normal int", "5432", 5432, true},
		{"should return 0 when receive empty string", "", 0, true},
		{"should return decimals when receive float", "12.63", 12, true},
		{"should return error when receive negative number", "-425", 0, false},
		{"should return error when receive malformed number", "6d8s5", 0, false},
	}

	for _, tt := range tests {
		out, err := utils.InputConverter.StringToUnsignedInt(tt.in)
		if tt.pass {
			assert.Equal(t, tt.want, out, tt.msg)
		} else {
			assert.EqualValues(t, 0, tt.want, "when failed should return 0")
			assert.NotNil(t, err, "when failed should return error")
		}
	}
}
