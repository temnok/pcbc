package util

import (
	"image"
	"image/png"
	"os"
)

func SaveTmpPng(name string, im image.Image) error {
	if err := os.MkdirAll("tmp", 0770); err != nil {
		return err
	}

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
