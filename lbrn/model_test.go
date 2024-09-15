package lbrn

import (
	"encoding/base64"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"os"
	"temnok/lab/util"
	"testing"
)

func TestModel(t *testing.T) {
	imageData, err := os.ReadFile("../pcbc/tmp/cu.png")
	assert.NoError(t, err)
	imageBase64 := base64.StdEncoding.EncodeToString(imageData)

	p := (&LightBurnProject{
		CutSetting_Img: []*CutSetting{
			(&CutSetting{}).SetDefaults(0),
		},
		Shape: []*Shape{
			(&Shape{
				Data: imageBase64,
			}).SetSize(26, 18).SetDefaults(0),
		},
	}).SetDefaults()

	bytes, err := xml.MarshalIndent(p, "", "    ")
	assert.NoError(t, err)

	util.SaveTmpFile("model.lbrn2", bytes)
}
