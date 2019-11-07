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
	pattern := "/user/pan1/sayHello/{uid}"
	url := "/user/pan1/sayHello/1"
	result, _ := Intercept(pattern, url, "{", "}")
	assert.Equal(t, result["uid"], "1")
}
func TestInterceptV2_1(t *testing.T) {
	pattern := "GET:/user/pan1/sayHello/{uid}"
	url := "/user/pan1/sayHello/1"
	result, err := InterceptV2(pattern, url, "{", "}")
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", result["uid"])
}
func TestInterceptV2_2(t *testing.T) {
	pattern := "GET:/{user}/pan1/sayHello/{uid}/"
	url := "xxx/pan/pan1/sayHello/1/"
	result, err := InterceptV2(pattern, url, "{", "}")
	assert.Equal(t, nil, err)
	assert.Equal(t, "1", result["uid"])
	assert.Equal(t, "pan", result["user"])
}
