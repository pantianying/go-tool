package transtype

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Map2xx_yy(t *testing.T) {
	var testData struct {
		AaAa string
		BaBa string
		CaCa struct {
			AaAa string
			BaBa string
			XxYy struct {
				XxXx string
				Xx   string
			}
		}
	}
	testData.AaAa = "1"
	testData.BaBa = "1"
	testData.CaCa.BaBa = "2"
	testData.CaCa.AaAa = "2"
	testData.CaCa.XxYy.XxXx = "3"
	testData.CaCa.XxYy.Xx = "3"

	m := Struct2Map(testData)
	Map2x_y(m)
	s, e := json.Marshal(m)
	assert.Equal(t, e, nil)
	assert.Equal(t, string(s), `{"aa_aa":"1","ba_ba":"1","ca_ca":{"aa_aa":"2","ba_ba":"2","xx_yy":{"xx":"3","xx_xx":"3"}}}`)
}
func Test_struct2MapAll(t *testing.T) {
	var testData struct {
		AaAa string `m:"aaAa"`
		BaBa string
		CaCa struct {
			AaAa string
			BaBa string `m:"baBa"`
			XxYy struct {
				xxXx string `m:"xxXx"`
				Xx   string `m:"xx"`
			} `m:"xxYy"`
		} `m:"caCa"`
	}
	testData.AaAa = "1"
	testData.BaBa = "1"
	testData.CaCa.BaBa = "2"
	testData.CaCa.AaAa = "2"
	testData.CaCa.XxYy.xxXx = "3"
	testData.CaCa.XxYy.Xx = "3"
	m := Struct2MapAll(testData).(map[string]interface{})
	assert.Equal(t, "1", m["aaAa"].(string))
	assert.Equal(t, "1", m["baBa"].(string))
	assert.Equal(t, "2", m["caCa"].(map[string]interface{})["aaAa"].(string))
	assert.Equal(t, "3", m["caCa"].(map[string]interface{})["xxYy"].(map[string]interface{})["xx"].(string))

	assert.Equal(t, reflect.Map, reflect.TypeOf(m["caCa"]).Kind())
	assert.Equal(t, reflect.Map, reflect.TypeOf(m["caCa"].(map[string]interface{})["xxYy"]).Kind())
}

type testStruct struct {
	AaAa string
	BaBa string `m:"baBa"`
	XxYy struct {
		xxXx string `m:"xxXx"`
		Xx   string `m:"xx"`
	} `m:"xxYy"`
}

func Test_struct2MapAll_Slice(t *testing.T) {
	var testData struct {
		AaAa string `m:"aaAa"`
		BaBa string
		CaCa []testStruct `m:"caCa"`
	}
	testData.AaAa = "1"
	testData.BaBa = "1"
	var tmp testStruct
	tmp.BaBa = "2"
	tmp.AaAa = "2"
	tmp.XxYy.xxXx = "3"
	tmp.XxYy.Xx = "3"
	testData.CaCa = append(testData.CaCa, tmp)
	m := Struct2MapAll(testData).(map[string]interface{})

	assert.Equal(t, "1", m["aaAa"].(string))
	assert.Equal(t, "1", m["baBa"].(string))
	assert.Equal(t, "2", m["caCa"].([]interface{})[0].(map[string]interface{})["aaAa"].(string))
	assert.Equal(t, "3", m["caCa"].([]interface{})[0].(map[string]interface{})["xxYy"].(map[string]interface{})["xx"].(string))

	assert.Equal(t, reflect.Slice, reflect.TypeOf(m["caCa"]).Kind())
	assert.Equal(t, reflect.Map, reflect.TypeOf(m["caCa"].([]interface{})[0].(map[string]interface{})["xxYy"]).Kind())
}
