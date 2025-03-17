// Copyright © 2025 Alex Temnok. All rights reserved.

package lbrn

import (
	"fmt"
	"image"
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
		buf = append(buf, f2s(xy.X)...)
		buf = append(buf, ',')
		buf = append(buf, f2s(xy.Y)...)
	}

	s.Tabs = string(buf)
}

func (s *Shape) SetPath(path path.Path) {
	s.V = nil
	s.P = nil

	for i := 0; i < len(path); i += 3 {
		xy := path[i]

		v := V{
			Vx: f2s(xy.X),
			Vy: f2s(xy.Y),
		}
		if i > 0 && path[i-1] != xy {
			v.C1x = f2s(path[i-1].X)
			v.C1y = f2s(path[i-1].Y)
		}
		if i+1 < len(path) && path[i+1] != xy {
			v.C0x = f2s(path[i+1].X)
			v.C0y = f2s(path[i+1].Y)
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

func NewRect(i int, t transform.T, w, h float64) *Shape {
	return &Shape{
		Type:     "Rect",
		CutIndex: fmt.Sprint(i),
		W:        f2s(w),
		H:        f2s(h),
		XForm:    xform(t),
	}
}

func NewPath(i int, t transform.T, path path.Path) *Shape {
	s := &Shape{
		Type:     "Path",
		CutIndex: fmt.Sprint(i),
		XForm:    xform(t),
	}

	s.SetPath(path)

	return s
}

func NewPathWithTabs(index int, t transform.T, p path.Path) *Shape {
	s := NewPath(index, t, p)

	var tabs []path.Point
	for i := 0; i < len(p); i += 3 {
		if isLine := i > 0 && p[i-3] == p[i-2] && p[i-1] == p[i]; isLine {
			u, v := p[i-2], p[i-1]
			tabs = append(tabs, path.Point{X: (u.X + v.X) / 2, Y: (u.Y + v.Y) / 2})
		}
	}
	s.SetTabs(tabs)

	return s
}

func NewBitmapShapeFromImage(i int, t transform.T, im image.Image) *Shape {
	return NewBitmapShape(i, t, NewBase64Bitmap(im))
}

func NewBitmapShape(i int, t transform.T, bm *Base64Bitmap) *Shape {
	return &Shape{
		Type:     "Bitmap",
		CutIndex: fmt.Sprint(i),
		XForm:    xform(t),
		W:        fmt.Sprint(bm.W),
		H:        fmt.Sprint(bm.H),
		Data:     bm.Data,
	}
}
