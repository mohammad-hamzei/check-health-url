package utils

import (
	"errors"
	"reflect"
	"strings"

	errs "git.tashilcar.com/core/tashilcar/core/yerrors"
)

// Tags
// This method takes an interface and extracts the fields with 'tagName' tag.
// Collects the tags value and fields value in a map and returns that.
func Tags(request interface{}, tagName string) ([]string, error) {
	const op errs.Op = "utils.Tags"
	if tagName == "" {
		tagName = "json"
	}

	// Checking the type of input to ensure that it is a struct type.
	if reflect.TypeOf(request).Elem().Kind() != reflect.Struct {
		return nil, errs.E(op, errors.New("input should be a struct"))
	}

	// Getting Value of  the input
	s := reflect.ValueOf(request).Elem()

	tags := make([]string, 0)
	// Loop through the struct fields to process the tags.
	for i := 0; i < s.NumField(); i++ {
		field := reflect.TypeOf(request).Elem().Field(i)
		t := field.Tag.Get(tagName)
		// Value of each field, that is ignored here.
		//v := cast.ToString(f.Interface())
		if t != "" {
			// If the tag contain any other value plus name, remove them.
			if strings.Contains(t, ",") {
				_tags := strings.Split(t, ",")
				tags = append(tags, _tags[0])
			} else {
				tags = append(tags, t)
			}
		}
	}

	return tags, nil
}
