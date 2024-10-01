package lbrn

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"temnok/lab/geom"
	"temnok/lab/path"
)

type Shape struct {
	Type     string `xml:"Type,attr"`
	CutIndex string `xml:"CutIndex,attr"`
	CutOrder string `xml:"CutOrder,attr,omitempty"`
	Cr       string `xml:"Cr,attr,omitempty"`
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

func (s *Shape) SetTabs(tabs []XY) {
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

func (s *Shape) SetPath(path path.Path) {
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

func (s *Shape) SetCutOrder(order int) *Shape {
	s.CutOrder = fmt.Sprint(order)

	return s
}

func NewPath(i int, t geom.Transform, path path.Path) *Shape {
	s := &Shape{
		Type:     "Path",
		CutIndex: fmt.Sprint(i),
		XForm:    XForm(t),
	}

	s.SetPath(path)

	return s
}

func NewPathWithTabs(index int, t geom.Transform, path path.Path) *Shape {
	s := NewPath(index, t, path)

	var tabs []XY
	for i := 0; i < len(path); i += 3 {
		if isLine := i > 0 && path[i-3] == path[i-2] && path[i-1] == path[i]; isLine {
			u, v := path[i-2], path[i-1]
			tabs = append(tabs, XY{(u.X + v.X) / 2, (u.Y + v.Y) / 2})
		}
	}
	s.SetTabs(tabs)

	return s
}

func NewCircle(i int, t geom.Transform, r float64) *Shape {
	return &Shape{
		Type:     "Ellipse",
		CutIndex: fmt.Sprint(i),
		XForm:    XForm(t),
		Rx:       fmt.Sprint(r),
		Ry:       fmt.Sprint(r),
	}
}

func NewRect(i int, t geom.Transform, w, h, r float64) *Shape {
	return &Shape{
		Type:     "Rect",
		CutIndex: fmt.Sprint(i),
		XForm:    XForm(t),
		W:        fmt.Sprint(w),
		H:        fmt.Sprint(h),
		Cr:       fmt.Sprint(r),
	}
}

func NewBitmap(i int, t geom.Transform, im image.Image) *Shape {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, im); err != nil {
		panic(err)
	}

	return &Shape{
		Type:     "Bitmap",
		CutIndex: fmt.Sprint(i),
		XForm:    XForm(t),
		W:        fmt.Sprint(im.Bounds().Dx()),
		H:        fmt.Sprint(im.Bounds().Dy()),
		Data:     base64.StdEncoding.EncodeToString(buf.Bytes()),
	}
}
