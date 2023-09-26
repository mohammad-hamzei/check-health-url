package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	errs "git.tashilcar.com/core/tashilcar/core/yerrors"
)

var InputConverter InputConverterInterface = &inputConverter{}

type InputConverterInterface interface {
	StringToUnsignedInt(i string) (int, error)
}

type inputConverter struct{}

func (*inputConverter) StringToUnsignedInt(i string) (int, error) {
	i = strings.Split(strings.TrimSpace(i), ".")[0]
	if i == "" {
		return 0, nil
	}
	const op errs.Op = "inputConverter.stringToUnsignedInt"
	number, err := strconv.Atoi(i)
	if errs.IsNotNil(err) {
		return 0, errs.E(op, errs.KindUnexpected, err, errs.LevelInfo)
	}
	if number < 0 {
		return 0, errs.E(
			op, errs.KindUnexpected, errs.LevelInfo,
			errors.New(fmt.Sprintf("negative number: %d", number)),
		)
	}
	return number, nil
}
