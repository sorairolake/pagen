// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

// Pagen generates pixel art. The value of each pixel is determined from random
// numbers.
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
//		Set the width of pixel art.
//	-height <SIZE>
//		Set the height of pixel art.
//	-block-size <SIZE>
//		Set the block size of pixel art.
//	-svg
//		Generate SVG file instead of PNG file.
//	-version
//		Print version number.
package main
