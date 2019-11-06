package pstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntercept(t *testing.T) {
	pattern := "/user/pan1/sayHello/{uid}"
	url := "/user/pan1/sayHello/1"
	result, _ := Intercept(pattern, url, "{", "}")
	assert.Equal(t, result["uid"], "1")
}
func TestIntercept1(t *testing.T) {
	pattern := "/user/pan1/sayHello/{uid}/"
	url := "/user/pan1/sayHello/1/"
	result, _ := Intercept(pattern, url, "{", "}")
	assert.Equal(t, result["uid"], "1")
}
func TestIntercept2(t *testing.T) {
	pattern := "xxxx{uid}xxxx{aaa}aaa{bbb}ccc"
	url := "xxxxx23xxxxcccaaaBBBccc"
	result, _ := Intercept(pattern, url, "{", "}")
	assert.Equal(t, result["uid"], "x23")
}
func TestInterceptV2_1(t *testing.T) {
	pattern := "/user/pan1/sayHello/{uid}"
	url := "/user/pan1/sayHello/1"
	result, err := InterceptV2(pattern, url, "{", "}")
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", result["uid"])
}
func TestInterceptV2_2(t *testing.T) {
	pattern := "/user/pan1/sayHello/{uid}/"
	url := "/user/pan1/sayHello/1/xxx"
	result, _ := InterceptV2(pattern, url, "{", "}")
	assert.Equal(t, result["uid"], "1")
}
