// Copyright © 2025 Alex Temnok. All rights reserved.

package lbrn

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
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

func (s *Shape) SetTabs(tabs []path.Point) {
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

func NewPath(i int, t transform.T, path path.Path) *Shape {
	s := &Shape{
		Type:     "Path",
		CutIndex: fmt.Sprint(i),
		XForm:    XForm(t),
	}

	s.SetPath(path)

	return s
}

func NewPathWithTabs(index int, t transform.T, p path.Path) *Shape {
	s := NewPath(index, t, p)

	var tabs path.Points
	for i := 0; i < len(p); i += 3 {
		if isLine := i > 0 && p[i-3] == p[i-2] && p[i-1] == p[i]; isLine {
			u, v := p[i-2], p[i-1]
			tabs = append(tabs, path.Point{X: (u.X + v.X) / 2, Y: (u.Y + v.Y) / 2})
		}
	}
	s.SetTabs(tabs)

	return s
}

func NewBitmap(i int, t transform.T, im image.Image) *Shape {
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
