package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand/v2"
	"os"
)

// Exercise 1.5 & 1.6
// 1.5: change the colors to green on black
// 1.6: modify to produce images in multiple colors
// NOTE: focusing on 1.6 as it extends 1.5 anyway

func RandomColor() color.Color {
	r := uint8(rand.IntN(256))
	g := uint8(rand.IntN(256))
	b := uint8(rand.IntN(256))
	return color.RGBA{r, g, b, 255}
}

func RandomPalette(numColors int) []color.Color {
	colors := []color.Color{color.Black}
	for i := 0; i < numColors; i++ {
		colors = append(colors, RandomColor())
	}
	return colors
}

func LissajousV2(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 512
		delay   = 1
	)

	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.2
	palette := RandomPalette(4)
	var randomColor uint8 = 0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		if i%4 == 0 {
			// Random Palette always has black as the first index
			// so don't use index 0 (black) as a foreground color
			randomColor = uint8(rand.IntN(len(palette)-1) + 1)
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), randomColor)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LJ: %v\n", err)
	}
}
