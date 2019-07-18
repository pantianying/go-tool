package transtype

import "errors"

func MapIItoMapSI(in interface{}) (out interface{}, err error) {
	if _, ok := in.(map[interface{}]interface{}); !ok {
		out = in
		return
	}
	in_map := make(map[interface{}]interface{})
	out_map := make(map[string]interface{})
	in_map = in.(map[interface{}]interface{})
	out = make(map[string]interface{})
	for k, v := range in_map {
		if s, ok := k.(string); ok {
			if s == "class" {
				//class字段不返回
				continue
			}
			if _, ok := v.(map[interface{}]interface{}); ok {
				v, err = MapIItoMapSI(v)
				if err != nil {
					return
				}
			}
			if vv, ok := v.([]interface{}); ok {
				var os []interface{}
				for i := range vv {
					osv, e := MapIItoMapSI(vv[i])
					if e != nil {
						err = e
						return
					}
					os = append(os, osv)
				}
				v = os
			}
			out_map[s] = v
		} else {
			err = errors.New("key's type not string")
			return
		}
	}
	out = out_map
	return
}
