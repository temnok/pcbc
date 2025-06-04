// Copyright © 2025 Alex Temnok. All rights reserved.

package lbrn

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

type LightBurnProject struct {
	XMLName       xml.Name `xml:"LightBurnProject"`
	AppVersion    string   `xml:"AppVersion,attr,omitempty"`
	FormatVersion string   `xml:"FormatVersion,attr,omitempty"`

	UIPrefs       *UIPrefs
	CutSetting    []*CutSetting `xml:"CutSetting"`
	CutSettingImg []*CutSetting `xml:"CutSetting_Img"`
	Shape         []*Shape      `xml:"Shape"`
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
