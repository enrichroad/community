package reflect

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func IsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

func Contains(search interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == search {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(search)).IsValid() {
			return true
		}
	}
	return false
}

func ContainsIgnoreCase(search string, target []string) bool {
	if len(search) == 0 {
		return false
	}
	if len(target) == 0 {
		return false
	}
	search = strings.ToLower(search)
	for i := 0; i < len(target); i++ {
		if strings.ToLower(target[i]) == search {
			return true
		}
	}
	return false
}

func StructToMap(obj interface{}, excludes ...string) map[string]interface{} {
	var data = make(map[string]interface{})
	keys := reflect.TypeOf(obj)
	values := reflect.ValueOf(obj)
	fillMap(data, keys, values, excludes...)
	return data
}

func fillMap(data map[string]interface{}, keys reflect.Type, values reflect.Value, excludes ...string) {
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	if keys.Kind() == reflect.Ptr {
		keys = keys.Elem()
	}

	for i := 0; i < keys.NumField(); i++ {
		keyField := keys.Field(i)
		valueField := values.Field(i)

		if keyField.Anonymous {
			fillMap(data, keyField.Type, valueField, excludes...)
		} else {
			if !ContainsIgnoreCase(keyField.Name, excludes) {
				jsonTag := keyField.Tag.Get("json")
				if len(jsonTag) > 0 {
					data[jsonTag] = valueField.Interface()
				} else {
					data[keyField.Name] = valueField.Interface()
				}
			}
		}
	}
}

func MapToStruct(obj interface{}, data map[string]interface{}) error {
	for k, v := range data {
		err := setField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj ", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value ", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type ")
	}
	structFieldValue.Set(val)
	return nil
}

func MD5(str string) string {
	return MD5Bytes([]byte(str))
}

func MD5Bytes(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// 获取struct字段
func StructFields(s interface{}) []reflect.StructField {
	t := StructTypeOf(s)
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}

	var results []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		results = append(results, f)
		// if f.Anonymous {
		// 	fields := StructFields(f.Type)
		// 	results = append(results, fields...)
		// }
	}
	return results
}

// 获取struct name
func StructName(s interface{}) string {
	t := StructTypeOf(s)
	return t.Name()
}

func StructTypeOf(s interface{}) reflect.Type {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// ToInterfaceArray used to convert a array to array of interface
func ToInterfaceArray(array interface{}) []interface{} {
	var result []interface{}
	arrType := reflect.TypeOf(array)
	if arrType.Kind() != reflect.Slice && arrType.Kind() != reflect.Array {
		result = append(result, array)
		return result
	}

	arrVal := reflect.ValueOf(array)
	for i := 0; i < arrVal.Len(); i++ {
		result = append(result, arrVal.Index(i).Interface())
	}
	return result
}

// IsHashAble check the object whether it can be used as the key of the map
func IsHashAble(obj interface{}) bool {
	if obj == nil {
		return false
	}
	return isHashAbleHelper(reflect.TypeOf(obj))
}

func isHashAbleHelper(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Slice, reflect.Func, reflect.Map:
		return false
	case reflect.Struct:
		for i := 0; i < typ.NumField(); i++ {
			if !isHashAbleHelper(typ.Field(i).Type) {
				return false
			}
		}
	}
	return true
}

// Append for interface
func Append(source []interface{}, desc interface{}) []interface{} {
	descValue := reflect.ValueOf(desc)
	descType := reflect.TypeOf(desc)
	switch descType.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < descValue.Len(); i++ {
			source = append(source, descValue.Index(i).Interface())
		}
	default:
		source = append(source, descValue.Interface())
	}

	return source
}
