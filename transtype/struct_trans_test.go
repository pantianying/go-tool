package transtype

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
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
func Test_Struct2Map(t *testing.T) {
	var testData struct {
		AaAa string `m:"AaAa"`
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
	m := Struct2MapAll(testData)
	fmt.Printf("%v", m)
	expect := `map[AaAa:1 BaBa:1 CaCa:map[AaAa:2 BaBa:2 XxYy:map[Xx:3 XxXx:3]]]`
	get := fmt.Sprintf("%v", m)
	assert.Equal(t, expect, get)
}
func Test_Struct2MapAll(t *testing.T) {
	var testData struct {
		AaAa string `m:"aaAa"`
		BaBa string `m:"baBa"`
		CaCa struct {
			AaAa string `m:"aaAa"`
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
	m := Struct2MapAll(testData)
	fmt.Printf("%v", m)
	expect := `map[aaAa:1 baBa:1 caCa:map[aaAa:2 baBa:2 xxYy:map[xx:3]]]`
	get := fmt.Sprintf("%v", m)
	assert.Equal(t, expect, get)
}
