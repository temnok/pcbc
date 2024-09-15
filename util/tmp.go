package util

import (
	"image"
	"image/png"
	"os"
)

func SaveTmpPng(name string, im image.Image) error {
	_ = os.Mkdir("tmp", 0770)

	f, err := os.Create("tmp/" + name)
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

func SaveTmpFile(name string, data []byte) error {
	_ = os.Mkdir("tmp", 0770)

	return os.WriteFile("tmp/"+name, data, 0770)
}
