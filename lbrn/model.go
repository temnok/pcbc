package lbrn

import (
	"encoding/xml"
	"fmt"
)

type LightBurnProject struct {
	XMLName        xml.Name      `xml:"LightBurnProject"`
	AppVersion     string        `xml:"AppVersion,attr"`
	FormatVersion  string        `xml:"FormatVersion,attr"`
	CutSetting_Img []*CutSetting `xml:"CutSetting_Img"`
	Shape          []*Shape      `xml:"Shape"`
}

type CutSetting struct {
	Type     string `xml:"type,attr"`
	Index    Param  `xml:"index"`
	Name     Param  `xml:"name"`
	Priority Param  `xml:"priority"`

	NumPasses   Param `xml:"numPasses"`
	MaxPower    Param `xml:"maxPower"`
	QPulseWidth Param `xml:"QPulseWidth"`
	Frequency   Param `xml:"frequency"`

	Speed      Param `xml:"speed"`
	Interval   Param `xml:"interval"`
	CrossHatch Param `xml:"crossHatch"`

	DPI              Param `xml:"dpi"`
	DitherMode       Param `xml:"ditherMode"`
	UseDotCorrection Param `xml:"useDotCorrection"`
	DotWidth         Param `xml:"dotWidth"`
}

type Shape struct {
	Type     string `xml:"Type,attr"`
	CutIndex string `xml:"CutIndex,attr"`
	W        string `xml:"W,attr"`
	H        string `xml:"H,attr"`
	File     string `xml:"File,attr"`
	Data     string `xml:"Data,attr"`
	XForm    string `xml:"XForm"`
}

type Param struct {
	Value string `xml:"Value,attr"`
}

func (p *LightBurnProject) SetDefaults() *LightBurnProject {
	p.AppVersion = "1.6.03"
	p.FormatVersion = "1"

	return p
}

func (c *CutSetting) SetDefaults(i int) *CutSetting {
	c.Type = "Image"
	c.Index.SetDefault(fmt.Sprint(i))
	c.Name.SetDefault(fmt.Sprintf("C%02d", i))
	c.Priority.SetDefault(c.Index.Value)

	c.NumPasses.SetDefault("1")
	c.MaxPower.SetDefault("20")
	c.QPulseWidth.SetDefault("200")
	c.Frequency.SetDefault("20000")

	c.Speed.SetDefault("600")
	c.Interval.SetDefault("0.01")
	c.CrossHatch.SetDefault("1")

	c.DPI.SetDefault("2540")
	c.DitherMode.SetDefault("threshold")
	c.UseDotCorrection.SetDefault("1")
	c.DotWidth.SetDefault("0.05")

	return c
}

func (s *Shape) SetDefaults(i int) *Shape {
	s.Type = "Bitmap"
	s.CutIndex = fmt.Sprint(i)
	if s.W == "" && s.H == "" {
		s.SetSize(10, 10)
	}
	if s.XForm == "" {
		s.SetPosition(55, 55)
	}

	return s
}

func (s *Shape) SetSize(w, h float64) *Shape {
	s.W = fmt.Sprint(w)
	s.H = fmt.Sprint(h)

	return s
}

func (s *Shape) SetPosition(x, y float64) *Shape {
	s.XForm = fmt.Sprintf("1 0 0 -1 %v %v", x, y)

	return s
}

func (p *Param) SetDefault(val string) {
	if p.Value == "" {
		p.Value = val
	}
}
