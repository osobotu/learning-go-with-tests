package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.RGBA{0xAA, 0xAA, 0x03, 0xff}, color.Black, color.RGBA{0xbb, 0xbb, 0xbb, 0xff}}

const (
	whiteIndex   = 0 // first color in palette
	coloredIndex = 1 // next color in palette
	blackIndex   = 2 // black color in palette
	lastIndex    = 3 // last color in palette
)

func Lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), coloredIndex)
			img.SetColorIndex(size+int(x*size+0.25), size+int(y*size+0.25), blackIndex)
			img.SetColorIndex(size+int(x*size+1), size+int(y*size+1), lastIndex)
		}
		phase += 0.2
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
