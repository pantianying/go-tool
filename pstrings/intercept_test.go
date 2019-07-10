package pstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntercept(t *testing.T) {
	pattern := "/test/cccc/{id}/get/{name}"
	url := "/test/cccc/22222/get/pan"
	result,_ := Intercept(pattern, url, "{", "}")
	assert.Equal(t, result["name"], "pan")
}
