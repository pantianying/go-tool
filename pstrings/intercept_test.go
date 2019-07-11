package pstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntercept(t *testing.T) {
	pattern := "/user/pan1/sayHello/{uid}"
	url := "/user/pan1/sayHello/1"
	result,_ := Intercept(pattern, url, "{", "}")
	assert.Equal(t, result["uid"], "1")
}
