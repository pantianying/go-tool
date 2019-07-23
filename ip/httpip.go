package ip

import (
	"net/http"
	"strconv"
	"strings"
)

const IPaBegin = 10 * 256 * 256 * 256
const IPaEnd = 10*256*256*256 + 255*256*256 + 255*256 + 255
const IPbBegin = 172*256*256*256 + 16*256*256
const IPbEnd = 172*256*256*256 + 31*256*256 + 255*256 + 255
const IPcBegin = 192*256*256*256 + 168*256*256
const IPcEnd = 192*256*256*256 + 168*256*256 + 255*256 + 255

func GetRemoteIp(req *http.Request) string {
	fromNginxIP := req.Header.Get("X-Real-IP")
	if len(fromNginxIP) != 0 {
		if isInnerIP(fromNginxIP) {
			if len(req.Header.Get("RealIP")) != 0 {
				return (addr(req.Header.Get("RealIP"))).RemoteIP()
			}
		}
		return addr(fromNginxIP).RemoteIP()
	} else {
		directIP := (addr(req.RemoteAddr)).RemoteIP()
		if isInnerIP(directIP) {
			if len(req.Header.Get("RealIP")) != 0 {
				return (addr(req.Header.Get("RealIP"))).RemoteIP()
			}
		}
		return addr(directIP).RemoteIP()
	}
}
func isInnerIP(IP string) bool {
	IPNum := IPToInt(IP)

	return (IPNum >= IPaBegin && IPNum <= IPaEnd) ||
		(IPNum >= IPbBegin && IPNum <= IPbEnd) ||
		(IPNum >= IPcBegin && IPNum <= IPcEnd) ||
		IPNum == 127*256*256*256+1
}

type addr string

func (ad addr) RemoteIP() string {
	s := strings.Split(string(ad), ":")
	//Debug( s[0])
	if len(s) != 0 {
		return s[0]
	}
	return ""
}

func IPToInt(ip string) int64 {
	tmpArray := strings.Split(ip, ".")
	if len(tmpArray) != 4 {
		return -1
	}

	IPa, _ := strconv.ParseInt(tmpArray[0], 10, 64)
	IPb, _ := strconv.ParseInt(tmpArray[1], 10, 64)
	IPc, _ := strconv.ParseInt(tmpArray[2], 10, 64)
	IPd, _ := strconv.ParseInt(tmpArray[3], 10, 64)
	return IPa*256*256*256 + IPb*256*256 + IPc*256 + IPd
}
