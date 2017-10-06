package form

import (
	"fmt"
	"errors"
	"reflect"
	"strings"
)


func Validate(form interface{}) (bool, error) {
	tagName := "binding"
	t := reflect.ValueOf(form)
	for i := 0; i < t.NumField(); i++ {
		field := t.Type().Field(i)
		tag := field.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}
		args := strings.Split(tag, ",")
		if args[0] == "required" {
			val := t.Field(i).Interface()
			if reflect.DeepEqual(val, reflect.Zero(field.Type).Interface()) {
				return false, errors.New(fmt.Sprintf("%s is required", field.Name))
			}
		}
	}
	return true, nil
}
