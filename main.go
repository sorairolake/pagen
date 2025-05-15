// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	outputPath := flag.Arg(0)
	width, height, blockSize := int(opt.width), int(opt.height), int(opt.blockSize)

	if opt.svg {
		if err := generateSVG(outputPath, width, height, blockSize); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := generatePNG(outputPath, width, height, blockSize); err != nil {
			log.Fatal(err)
		}
	}
}
