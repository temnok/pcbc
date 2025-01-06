// Copyright © 2025 Alex Temnok. All rights reserved.

package util

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
)

// SavePNG creates a PNG file out of provided image,
// making all file parent directories if necessary.
func SavePNG(name string, im image.Image) error {
	if err := os.MkdirAll(filepath.Dir(name), 0700); err != nil {
		return err
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	if err = png.Encode(f, im); err != nil {
		return err
	}

	return f.Close()
}
