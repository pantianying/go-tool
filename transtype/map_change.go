package transtype

func MapIItoMapSI(in interface{}) interface{} {
	in_map := make(map[interface{}]interface{})
	if v, ok := in.(map[interface{}]interface{}); !ok {
		return in
	} else {
		in_map = v
	}

	out_map := make(map[string]interface{}, len(in_map))
	for k, v := range in_map {
		if s, ok := k.(string); ok {
			if s == "class" || v == nil {
				continue
			}
			if _, ok := v.(map[interface{}]interface{}); ok {
				v = MapIItoMapSI(v)
			}
			if vv, ok := v.([]interface{}); ok {
				var os = make([]interface{}, 0, len(vv))
				for i := range vv {
					osv := MapIItoMapSI(vv[i])
					os = append(os, osv)
				}
				v = os
			}
			out_map[s] = v
		}
	}
	return out_map
}
