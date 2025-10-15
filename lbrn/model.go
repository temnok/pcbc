// Copyright © 2025 Alex Temnok. All rights reserved.

package lbrn

import (
	"fmt"
	"github.com/temnok/pcbc/transform"
)

type CutSetting struct {
	Type     string `xml:"type,attr"`
	Index    *Param `xml:"index"`
	Name     *Param `xml:"name"`
	Priority *Param `xml:"priority"`
	DoOutput *Param `xml:"doOutput"`

	Speed        *Param `xml:"speed"`
	NumPasses    *Param `xml:"numPasses"`
	GlobalRepeat *Param `xml:"globalRepeat"`
	MaxPower     *Param `xml:"maxPower"`
	QPulseWidth  *Param `xml:"QPulseWidth"`
	Frequency    *Param `xml:"frequency"`

	Interval         *Param `xml:"interval"`
	DPI              *Param `xml:"dpi"`
	DitherMode       *Param `xml:"ditherMode"`
	UseDotCorrection *Param `xml:"useDotCorrection"`
	DotWidth         *Param `xml:"dotWidth"`

	TabsEnabled *Param `xml:"tabsEnabled"`
	TabSize     *Param `xml:"tabSize"`

	CrossHatch *Param `xml:"crossHatch"`
	Angle      *Param `xml:"angle"`

	Negative *Param `xml:"negative"`

	CleanupPass **Param   `xml:"cleanupPass"`
	SubLayer    *SubLayer `xml:"SubLayer"`
}

type SubLayer struct {
	Type      string `xml:"type,attr"`
	Index     string `xml:"index,attr"`
	Subname   *Param `xml:"subname"`
	IsCleanup *Param `xml:"isCleanup"`

	Speed        *Param `xml:"speed"`
	Angle        *Param `xml:"angle"`
	AnglePerPass *Param `xml:"anglePerPass"`
	FloodFill    *Param `xml:"floodFill"`

	NumPasses   *Param `xml:"numPasses"`
	MaxPower    *Param `xml:"maxPower"`
	QPulseWidth *Param `xml:"QPulseWidth"`
	Frequency   *Param `xml:"frequency"`

	CrossHatch *Param `xml:"crossHatch"`

	Interval *Param `xml:"interval"`
	DPI      *Param `xml:"dpi"`
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

func xform(t transform.T) string {
	return fmt.Sprintf("%v %v %v %v %v %v",
		f2s(t.Ix), f2s(t.Iy),
		f2s(t.Jx), f2s(t.Jy),
		f2s(t.Kx), f2s(t.Ky),
	)
}
