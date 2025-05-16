// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

// Pagen generates pixel art. The value of each pixel is determined from random
// numbers.
//
// By default, this will generate a PNG image with the size of 256×256 pixels
// and the block size of 1×1 pixel.
//
// Usage:
//
//	pagen [OPTIONS] <FILE>
//
// Arguments:
//
//	<FILE>
//		The path of the generated image file.
//
// Options:
//
//	-width <SIZE>
//		Set the width of pixel art. Default is 256.
//	-height <SIZE>
//		Set the height of pixel art. Default is 256.
//	-block-size <SIZE>
//		Set the block size of pixel art. Default is 1.
//	-svg
//		Generate SVG file instead of PNG file.
//	-version
//		Print version number.
package main
