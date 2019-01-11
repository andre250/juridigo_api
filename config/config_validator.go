package config

import (
	"reflect"
)

func configValidator() ([]string, bool) {
	config := &globaConfig
	var errList []string
	val := reflect.ValueOf(config).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Type.Kind() == reflect.Struct {
			subStructLength := val.Field(i).NumField()
			for k := 0; k < subStructLength; k++ {
				value := val.Field(i).Field(k)
				if value.String() == "" {
					errList = append(errList, val.Field(i).Type().Field(k).Name)
				}
			}
		} else {
			value := val.Field(i)
			if value.String() == "" {
				errList = append(errList, val.Type().Field(i).Name)
			}
		}
	}
	if len(errList) != 0 {
		return errList, false
	}
	return nil, true

}
