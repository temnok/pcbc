// Copyright © 2025 Alex Temnok. All rights reserved.

package lbrn

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

type XY = path.Point

type LightBurnProject struct {
	XMLName       xml.Name      `xml:"LightBurnProject"`
	AppVersion    string        `xml:"AppVersion,attr,omitempty"`
	FormatVersion string        `xml:"FormatVersion,attr,omitempty"`
	CutSetting    []*CutSetting `xml:"CutSetting"`
	CutSettingImg []*CutSetting `xml:"CutSetting_Img"`
	Shape         []*Shape      `xml:"Shape"`
}

type CutSetting struct {
	Type     string `xml:"type,attr"`
	Index    Param  `xml:"index"`
	Name     Param  `xml:"name"`
	Priority Param  `xml:"priority"`

	Speed        Param `xml:"speed"`
	NumPasses    Param `xml:"numPasses"`
	GlobalRepeat Param `xml:"globalRepeat"`
	MaxPower     Param `xml:"maxPower"`
	QPulseWidth  Param `xml:"QPulseWidth"`
	Frequency    Param `xml:"frequency"`

	Interval         Param `xml:"interval"`
	DPI              Param `xml:"dpi"`
	DitherMode       Param `xml:"ditherMode"`
	UseDotCorrection Param `xml:"useDotCorrection"`
	DotWidth         Param `xml:"dotWidth"`

	TabsEnabled Param `xml:"tabsEnabled"`
	TabSize     Param `xml:"tabSize"`

	CrossHatch Param `xml:"crossHatch"`
	Angle      Param `xml:"angle"`

	Negative Param `xml:"negative"`
}

type Param struct {
	Value string `xml:"Value,attr"`
}

type V struct {
	Vx  string `xml:"vx,attr"`
	Vy  string `xml:"vy,attr"`
	C0x string `xml:"c0x,attr,omitempty"`
	C0y string `xml:"c0y,attr,omitempty"`
	C1x string `xml:"c1x,attr,omitempty"`
	C1y string `xml:"c1y,attr,omitempty"`
}

type P struct {
	T  string `xml:"T,attr"`
	P0 string `xml:"p0,attr"`
	P1 string `xml:"p1,attr"`
}

// having this method is important because default Sprint(f)
// will result in different outputs on AMD and ARM platforms
func f2s(val float64) string {
	res := strconv.FormatFloat(val, 'f', 9, 64)
	for res[len(res)-1] == '0' {
		res = res[:len(res)-1]
	}
	if res[len(res)-1] == '.' {
		res = res[:len(res)-1]
	}
	return res
}

func xform(t transform.T) string {
	return fmt.Sprintf("%v %v %v %v %v %v", f2s(t.Ix), f2s(t.Iy), f2s(t.Jx), f2s(t.Jy), f2s(t.Kx), f2s(t.Ky))
}

func (p *LightBurnProject) SaveToFile(filename string) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0770); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	enc := xml.NewEncoder(file)
	enc.Indent("", "\t")
	return enc.Encode(p)
}
