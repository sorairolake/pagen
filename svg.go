// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

package main

import (
	"os"
	"path/filepath"

	svg "github.com/ajstarks/svgo"
)

func generateSVG(outputPath string, width, height, blockSize int) (err error) {
	output, err := os.Create(filepath.Clean(outputPath))
	if err != nil {
		return err
	}

	defer func() {
		if e := output.Close(); e != nil {
			err = e
		}
	}()

	canvas := svg.New(output)
	canvas.Start(width, height)
	defer canvas.End()

	for y := 0; y < height; y += blockSize {
		for x := 0; x < width; x += blockSize {
			c := generateNRGBA()
			fill := canvas.RGB(int(c.R), int(c.G), int(c.B))

			canvas.Rect(x, y, blockSize, blockSize, fill)
		}
	}

	return nil
}
