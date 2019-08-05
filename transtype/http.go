package transtype

import "net/http"

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
