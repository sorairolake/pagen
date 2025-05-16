// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	if opt.version {
		fmt.Printf("pagen %v\n", version)
		os.Exit(0)
	}

	n := flag.NArg()

	switch {
	case n == 0:
		fmt.Fprintln(os.Stderr, "pagen: a required argument was not provided")
		flag.Usage()
		os.Exit(1)
	case n > 1:
		fmt.Fprintf(os.Stderr, "pagen: unexpected argument '%v' found\n", flag.Arg(1))
		flag.Usage()
		os.Exit(1)
	}

	outputPath := flag.Arg(0)
	width, height, blockSize := int(opt.width), int(opt.height), int(opt.blockSize)

	if opt.svg {
		if err := generateSVG(outputPath, width, height, blockSize); err != nil {
			fmt.Fprintf(os.Stderr, "pagen: %v\n", err)
			os.Exit(1)
		}
	} else {
		if err := generatePNG(outputPath, width, height, blockSize); err != nil {
			fmt.Fprintf(os.Stderr, "pagen: %v\n", err)
			os.Exit(1)
		}
	}
}
