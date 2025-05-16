// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

package main

import (
	"image/color"
	"math"
	"math/rand/v2"
)

func generateNRGBA() color.NRGBA {
	r := uint8(rand.N(math.MaxUint8))
	g := uint8(rand.N(math.MaxUint8))
	b := uint8(rand.N(math.MaxUint8))
	a := uint8(math.MaxUint8)

	return color.NRGBA{r, g, b, a}
}
