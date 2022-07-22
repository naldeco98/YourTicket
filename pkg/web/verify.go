package web

import (
	"fmt"
	"reflect"
)

func VerifyEmptyFields(obj interface{}) error {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" || v.Field(i).Interface() == 0 {
			return fmt.Errorf("field " + v.Type().Field(i).Name + " is empty")
		}
	}
	return nil
}
