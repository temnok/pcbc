package lbrn

import (
	"encoding/xml"
	"fmt"
	"temnok/lab/geom"
)

type LightBurnProject struct {
	XMLName       xml.Name     `xml:"LightBurnProject"`
	AppVersion    string       `xml:"AppVersion,attr,omitempty"`
	FormatVersion string       `xml:"FormatVersion,attr,omitempty"`
	CutSetting    []CutSetting `xml:"CutSetting"`
	CutSettingImg []CutSetting `xml:"CutSetting_Img"`
	Shape         []Shape      `xml:"Shape"`
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
}

type Shape struct {
	Type     string `xml:"Type,attr"`
	CutIndex string `xml:"CutIndex,attr"`
	W        string `xml:"W,attr,omitempty"`
	H        string `xml:"H,attr,omitempty"`
	Rx       string `xml:"Rx,attr,omitempty"`
	Ry       string `xml:"Ry,attr,omitempty"`
	XForm    string `xml:"XForm"`
	Tabs     string `xml:"Tabs"`

	V []V `xml:"V"`
	P []P `xml:"P"`

	File string `xml:"File,attr,omitempty"`
	Data string `xml:"Data,attr,omitempty"`
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

func (s *Shape) SetTabs(tabs []geom.XY) {
	var buf []byte

	for _, xy := range tabs {
		if len(buf) > 0 {
			buf = append(buf, ' ')
		}
		buf = append(buf, fmt.Sprint(xy.X)...)
		buf = append(buf, ',')
		buf = append(buf, fmt.Sprint(xy.Y)...)
	}

	s.Tabs = string(buf)
}

func (s *Shape) SetPath(path []geom.XY) {
	s.Type = "Path"
	s.V = nil
	s.P = nil

	for i := 0; i < len(path); i += 3 {
		xy := path[i]

		v := V{
			Vx: fmt.Sprint(xy.X),
			Vy: fmt.Sprint(xy.Y),
		}
		if i > 0 && path[i-1] != xy {
			v.C1x = fmt.Sprint(path[i-1].X)
			v.C1y = fmt.Sprint(path[i-1].Y)
		}
		if i+1 < len(path) && path[i+1] != xy {
			v.C0x = fmt.Sprint(path[i+1].X)
			v.C0y = fmt.Sprint(path[i+1].Y)
		}

		s.V = append(s.V, v)

		if n := len(s.V); n > 1 {
			u := s.V[n-2]
			p := P{
				P0: fmt.Sprint(n - 2),
				P1: fmt.Sprint(n - 1),
			}
			if isLine := u.C0x == "" && v.C1x == ""; isLine {
				p.T = "L"
			} else {
				p.T = "B"
			}
			s.P = append(s.P, p)
		}
	}
}
