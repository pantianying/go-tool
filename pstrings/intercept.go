package pstrings

import (
	"errors"
	"strings"
)

var conErr = errors.New("pattern is not as required")

func Intercept(pattern, in, leftSep, rightSep string) (result map[string]string, err error) {
	result = make(map[string]string)
	ss1 := strings.Split(pattern, leftSep)
	var sss []string //多余字符串数组
	var kkk []string //key数组
	for _, k := range ss1 {
		if strings.Contains(k, rightSep) {
			cs := strings.Split(k, rightSep)
			if len(cs) == 2 {
				kkk = append(kkk, cs[0])
				if cs[1] != "" {
					sss = append(sss, cs[1])
				}
			} else {
				err = conErr
				return
			}
		} else {
			if k != "" {
				sss = append(sss, k)
			}
		}
	}
	for i := range kkk {
		if i == len(sss)-1 {
			indexStart := strings.Index(in, sss[i]) + len(sss[i])
			if indexStart > len(in) {
				err = conErr
				return
			}
			result[kkk[i]] = in[indexStart:]
		} else {
			indexStart := strings.Index(in, sss[i]) + len(sss[i])
			if indexStart > len(in) {
				err = conErr
				return
			}
			indexStop := strings.Index(in, sss[i+1]) + 1
			if indexStop-1 > len(in) {
				err = conErr
				return
			}
			if indexStop <= indexStart {
				err = conErr
				return
			}
			result[kkk[i]] = in[indexStart : indexStop-1]
		}
	}
	return
}
func InterceptV2(pattern, in, leftSep, rightSep string) (result map[string]string, err error) {
	result = make(map[string]string, 1)
	ssp := strings.Split(pattern, "/") //按pattern切割
	ssi := strings.Split(in, "/")
	if len(ssp) != len(ssi) {
		err = conErr
		return
	}
	for i := range ssp {
		if len(ssp[i]) > 2 {
			if ssp[i][:1] == leftSep {
				//确认是{}
				result[ssp[i][1:len(ssp[i])-1]] = ssi[i]
			}
		}
	}
	return
}
