package util

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func SaveTmpPng(name string, im image.Image) error {
	return SavePng("tmp/"+name, im)
}

func SavePng(name string, im image.Image) error {
	if err := os.MkdirAll(filepath.Dir(name), 0400); err != nil {
		return err
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	if err = png.Encode(f, im); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}
