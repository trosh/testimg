package main

import (
	"os"
	"log"
	"image"
	_ "image/png"
	"github.com/trosh/term"
)

func mean100to666(m image.Image, p image.Point) (or, og, ob int) {
	var r, g, b uint32
	for iy := p.Y; iy < p.Y+10; iy += 1 {
		for ix := p.X; ix < p.X+10; ix += 1 {
			cr, cg, cb, _ := m.At(ix, iy).RGBA()
			r += cr
			g += cg
			b += cb
		}
	}
	size := uint32(100)
	r /= size
	g /= size
	b /= size
	or, og, ob = int(r/10922), int(g/10922), int(b/10922)
	return or, og, ob
}

func main() {
	reader, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	scr := term.Scr{term.Pxl{1, 1},
	                term.Pxl{bounds.Dx()/10,
	                         bounds.Dy()/10},
	                0}
	scr.Flush()
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 10 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 10 {
			r, g, b := mean100to666(m, image.Point{x, y})
			scr.Plot(term.Pxl{x/10, y/10}, 36*r + 6*g + b + 16)
		}
	}
}

