package main

import (
	"os"
	"log"
	"image"
	_ "image/png"
	"github.com/trosh/term"
)

type img struct {
	m image.Image
}

func (i img) meangray(p image.Point, pr image.Point) (gray int) {
	var r, g, b uint32
	for iy := p.Y; iy < p.Y+pr.Y; iy += 1 {
		for ix := p.X; ix < p.X+pr.X; ix += 1 {
			cr, cg, cb, _ := i.m.At(ix, iy).RGBA()
			r += cr
			g += cg
			b += cb
		}
	}
	size := uint32(pr.X*pr.Y)
	r /= size
	g /= size
	b /= size
	return int((r+g+b)/8548)+232
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
	//scr.Flush()
	pr := image.Point{1, 2}
	if len(os.Args) == 3 {
		for i := 0; i < len(os.Args[2]); i++ {
			pr.X*=2
			pr.Y*=2
		}
	}
	for y := bounds.Min.Y; y < bounds.Min.Y+30*pr.Y; y += pr.Y {
		for x := bounds.Min.X; x < bounds.Min.X+110*pr.X; x += pr.X {
			gray := img{m}.meangray(image.Point{x, y}, pr)
			scr.Plot(term.Pxl{x/pr.X, y/pr.Y}, gray)
		}
	}
	/*
	for y := 0; y < 30; y += 1 {
		for x := 0; x < 110; x += 1 {
			r, g, b, _ := m.At(x, y).RGBA()
			scr.Plot(term.Pxl{x+1, y+1}, int((r+g+b)/8548)+232)
		}
	}
	*/
}
