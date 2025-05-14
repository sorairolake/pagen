// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	flag.Parse()

	if opt.version {
		fmt.Printf("pagen %v\n", version)
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	outputPath := filepath.Clean(flag.Arg(0))

	width, height, blockSize := int(opt.width), int(opt.height), int(opt.blockSize)
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	for y := 0; y < height; y += blockSize {
		for x := 0; x < width; x += blockSize {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				r := uint8(rand.N(math.MaxUint8))
				g := uint8(rand.N(math.MaxUint8))
				b := uint8(rand.N(math.MaxUint8))
				a := uint8(math.MaxUint8)

				for by := 0; (by < blockSize) && (y+by < height); by++ {
					for bx := 0; (bx < blockSize) && (x+bx < width); bx++ {
						img.Set(x+bx, y+by, color.NRGBA{r, g, b, a})
					}
				}
			}(x, y)
		}
	}

	wg.Wait()

	output, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := output.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := png.Encode(output, img); err != nil {
		log.Print(err)
	}
}
