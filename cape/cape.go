package cape

import (
	"image"
	"log"
	"math"
)

type Cape struct {
	Pix  []uint8
	W, H int
}

func FromData(data []uint8) (*Cape, error) {
	w, h := wAndHBySize(len(data) / 4)
	log.Println(len(data), w, h)
	return &Cape{
		Pix: data,
		W:   w,
		H:   h,
	}, nil
}

func (c *Cape) Image() image.Image {
	return &image.RGBA{
		Pix:    c.Pix,
		Stride: c.W * 4,
		Rect:   image.Rect(0, 0, c.W, c.H),
	}
}

func wAndHBySize(size int) (int, int) {
	w := int(math.Sqrt(float64(size / 2)))
	return w * 2, w
}
