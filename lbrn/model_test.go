package lbrn

import (
	"encoding/base64"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"os"
	"temnok/lab/contour"
	"temnok/lab/geom"
	"temnok/lab/util"
	"testing"
)

func TestPCBEtch(t *testing.T) {
	imageData, err := os.ReadFile("../pcbc/tmp/cu.png")
	assert.NoError(t, err)
	imageBase64 := base64.StdEncoding.EncodeToString(imageData)

	p := LightBurnProject{
		CutSettingImg: []CutSetting{
			{
				Type:     "Image",
				Index:    Param{"0"},
				Name:     Param{"C00"},
				Priority: Param{"0"},

				MaxPower:    Param{"20"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				Speed:            Param{"600"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},
			},
			{
				Type:     "Image",
				Index:    Param{"1"},
				Name:     Param{"C01"},
				Priority: Param{"3"},

				MaxPower:    Param{"50"},
				QPulseWidth: Param{"2"},
				Frequency:   Param{"280000"},

				Speed:            Param{"2000"},
				Interval:         Param{"0.01"},
				DPI:              Param{"2540"},
				UseDotCorrection: Param{"1"},
				DotWidth:         Param{"0.05"},
			},
		},
		CutSetting: []CutSetting{
			{
				Type: "Cut",

				Index:    Param{"2"},
				Name:     Param{"C02"},
				Priority: Param{"2"},

				Speed:        Param{"100"},
				GlobalRepeat: Param{"100"},

				MaxPower:    Param{"90"},
				QPulseWidth: Param{"200"},
				Frequency:   Param{"20000"},

				TabsEnabled: Param{"1"},
				TabSize:     Param{"0.2"},
			},
		},
		Shape: []Shape{
			{
				Type:     "Bitmap",
				CutIndex: "0",
				Data:     imageBase64,
				W:        "38",
				H:        "48",
				XForm:    "1 0 0 -1 55 55",
			},
			{
				Type:     "Bitmap",
				CutIndex: "1",
				Data:     imageBase64,
				W:        "38",
				H:        "48",
				XForm:    "1 0 0 -1 55 55",
			},
			{
				Type:     "Ellipse",
				CutIndex: "2",
				Rx:       "1",
				Ry:       "1",
				XForm:    "1 0 0 1 41 36",
			},
			{
				Type:     "Ellipse",
				CutIndex: "2",
				Rx:       "1",
				Ry:       "1",
				XForm:    "1 0 0 1 69 36",
			},
			{
				Type:     "Ellipse",
				CutIndex: "2",
				Rx:       "1",
				Ry:       "1",
				XForm:    "1 0 0 1 69 74",
			},
			{
				Type:     "Ellipse",
				CutIndex: "2",
				Rx:       "1",
				Ry:       "1",
				XForm:    "1 0 0 1 41 73",
			},
			{
				Type:     "Path",
				CutIndex: "2",
				XForm:    "1 0 0 1 55 55",
			},
			{
				Type:     "Path",
				CutIndex: "2",
				XForm:    "1 0 0 1 55 46",
			},
			{
				Type:     "Path",
				CutIndex: "2",
				XForm:    "1 0 0 1 55 64",
			},
		},
	}
	p.Shape[len(p.Shape)-3].SetPath(contour.RoundRect(36, 46, 4))
	p.Shape[len(p.Shape)-3].SetTabs([]geom.XY{
		{-18, 0}, {18, 0}, {0, -23}, {0, 23},
	})

	for i := 0; i < 2; i++ {
		p.Shape[len(p.Shape)-2+i].SetPath(contour.RoundRect(24, 16, 2))
		p.Shape[len(p.Shape)-2+i].SetTabs([]geom.XY{
			{-12, 0}, {12, 0}, {0, -8}, {0, 8},
		})
	}

	bytes, err := xml.MarshalIndent(&p, "", "\t")
	assert.NoError(t, err)

	util.SaveTmpFile("image.lbrn", bytes)
}
