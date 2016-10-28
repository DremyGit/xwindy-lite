package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func CopyFromMap(target interface{}, src map[string]interface{}) error {
	for k, v := range src {
		err := setField(target, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func CopyFromMapWithFilter(target interface{}, src map[string]interface{}, filter []string) error {

	srcCopy := FilterMap(src, filter)
	return CopyFromMap(target, srcCopy)
}

func FilterMap(src map[string]interface{}, filter []string) map[string]interface{} {
	filterSet := make(map[string]bool)
	filted := make(map[string]interface{})

	for _, field := range filter {
		filterSet[field] = true
	}
	for k := range src {
		if filterSet[k] {
			filted[k] = src[k]
		}
	}
	return filted
}

func setField(obj interface{}, name string, value interface{}) error {

	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {

		elem := reflect.TypeOf(obj).Elem()
		isFinded := false
		for i := 0; i < structValue.NumField(); i++ {
			jsonTag := elem.Field(i).Tag.Get("json")
			structJSONName := strings.Split(jsonTag, ",")[0]
			if structJSONName == name {
				name = elem.Field(i).Name
				structFieldValue = structValue.FieldByName(name)
				isFinded = true
				break
			}
		}
		if !isFinded {
			return nil
		}
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
