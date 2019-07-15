package transtype

import (
	"encoding/json"
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
