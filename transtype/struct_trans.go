package transtype

import (
	"reflect"
	"strings"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
func Struct2MapAll(obj interface{}) interface{} {
	if obj == nil {
		return obj
	}
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Struct {
		result := make(map[string]interface{})
		for i := 0; i < t.NumField(); i++ {
			if v.Field(i).Kind() == reflect.Struct {
				if v.Field(i).CanInterface() {
					setInMap(result, t.Field(i), Struct2MapAll(v.Field(i).Interface()))
				}
			} else if v.Field(i).Kind() == reflect.Slice {
				if v.Field(i).CanInterface() {
					setInMap(result, t.Field(i), Struct2MapAll(v.Field(i).Interface()))
				}
			} else {
				if v.Field(i).CanInterface() {
					setInMap(result, t.Field(i), v.Field(i).Interface())
				}
			}
		}
		return result
	} else if t.Kind() == reflect.Slice {
		value := reflect.ValueOf(obj)
		var newTemps []interface{}
		for i := 0; i < value.Len(); i++ {
			newTemp := Struct2MapAll(value.Index(i).Interface())
			newTemps = append(newTemps, newTemp)
		}
		return newTemps
	} else {
		return obj
	}
}
func setInMap(m map[string]interface{}, structField reflect.StructField, value interface{}) (result map[string]interface{}) {
	result = m
	if tagName := structField.Tag.Get("m"); tagName == "" {
		result[headerAtoa(structField.Name)] = value
	} else {
		result[tagName] = value
	}
	return
}
func headerAtoa(a string) (b string) {
	b = strings.ToLower(a[:1]) + a[1:]
	return
}

/* 驼峰转下划线 */
func XY2_x_y(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

/* map的key驼峰转下划线，会遍历所有的key */
func Map2x_y(m map[string]interface{}) {
	for k1, v1 := range m {
		x := XY2_x_y(k1)
		if k1 != x {
			delete(m, k1)
			if v1 == nil {
				m[x] = v1
			} else if reflect.TypeOf(v1).Kind() == reflect.Struct {
				m[x] = Struct2Map(v1)
				Map2x_y(m[x].(map[string]interface{}))
			} else {
				m[x] = v1
			}
		}

	}
}
