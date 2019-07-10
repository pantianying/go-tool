package pstrings

import (
	"fmt"
	"strings"
)

func Intercept(pattern, in, leftSep, rightSep string) (result map[string]string) {
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
				panic(" len(cs) != 2")
			}
		} else {
			if k != "" {
				sss = append(sss, k)
			}
		}
	}
	for i, _ := range kkk {
		if i == len(sss)-1 {
			indexStart := strings.Index(in, sss[i]) + len(sss[i])
			result[kkk[i]] = in[indexStart:]
		} else {
			indexStart := strings.Index(in, sss[i]) + len(sss[i])
			indexStop := strings.Index(in, sss[i+1]) + 1
			fmt.Println(indexStart, indexStop, sss[i], sss[i+1])
			result[kkk[i]] = in[indexStart : indexStop-1]
		}
	}
	return
}
