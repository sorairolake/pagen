// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

package main

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"sync"
)

func generatePNG(outputPath string, width, height, blockSize int) (err error) {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	for y := 0; y < height; y += blockSize {
		for x := 0; x < width; x += blockSize {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()

				c := generateNRGBA()

				for by := 0; (by < blockSize) && (y+by < height); by++ {
					for bx := 0; (bx < blockSize) && (x+bx < width); bx++ {
						img.Set(x+bx, y+by, c)
					}
				}
			}(x, y)
		}
	}

	wg.Wait()

	output, err := os.Create(filepath.Clean(outputPath))
	if err != nil {
		return err
	}

	defer func() {
		if e := output.Close(); e != nil {
			err = e
		}
	}()

	return png.Encode(output, img)
}
