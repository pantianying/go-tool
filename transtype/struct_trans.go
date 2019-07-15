package transtype

import (
	"net/http"
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

func HttpHeadtoMap(header http.Header) (h map[string]string) {
	h = make(map[string]string)
	for k, v := range header {
		if len(v) != 1 {
			continue
		}
		h[k] = v[0]
	}
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
			if reflect.TypeOf(v1).Kind() == reflect.Struct {
				m[x] = Struct2Map(v1)
				Map2x_y(m[x].(map[string]interface{}))
			} else {
				m[x] = v1
			}

		}

	}
}
