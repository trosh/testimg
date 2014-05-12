package main

import (
	"os"
	"log"
	"image"
	_ "image/png"
	"github.com/trosh/term"
)

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
	for y := bounds.Min.Y;
	    y < bounds.Max.Y;
	    y+=10 {
		for x := bounds.Min.X;
		    x < bounds.Max.X;
		    x+=10 {
			var r, g, b uint32
			for ix := x;
			    ix < x+10;
			    ix++ {
				for iy := y;
				    iy < y+10; iy++ {
					cr, cg, cb, _ := m.At(ix,
					                      iy).RGBA()
					r += cr
					g += cg
					b += cb
				}
			}
			size := uint32(81)
			r /= size
			g /= size
			b /= size
			scr.Plot(term.Pxl{x/10,
			                  y/10},
			         int((36*r+6*g+b)/10922+16))
		}
	}
}

