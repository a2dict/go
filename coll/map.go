package coll

import (
	"fmt"
	"reflect"
	"strings"
)

// MustToMapWithTag ...
func MustToMapWithTag(in interface{}, tag string) map[string]interface{} {
	m, err := ToMapWithTag(in, tag)
	if err != nil {
		panic(err)
	}
	return m
}

// ToMapWithTag converts a struct to a map using the struct's tags.
//
// ToMapWithTag uses tags on struct fields to decide which fields to add to the
// returned map.
func ToMapWithTag(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMapWithTag only accepts structs; got %T", v)
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		k := fi.Name

		if tag != "" {
			tagv := fi.Tag.Get(tag)
			if tagv == "-" {
				continue
			}
			if tagv != "" {
				k = strings.Split(tagv, ",")[0]
			}
		}
		out[k] = v.Field(i).Interface()
	}
	return out, nil
}
