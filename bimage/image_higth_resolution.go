package bimage

import (
	"image"
	"image/color"
)

type ImageWithBetterResolution struct {
	M, RW, RH, FW, FH, B int
	Pix                  []uint8
}

func NewImageWithBetterResolution(m, w, h, b, fw, fh int) *ImageWithBetterResolution {
	return &ImageWithBetterResolution{
		M:   m,
		RW:  w,
		RH:  h,
		B:   b,
		FW:  fw,
		FH:  fh,
		Pix: make([]uint8, w*h*4),
	}
}
func (img *ImageWithBetterResolution) ColorModel() color.Model { return color.RGBAModel }

func (img *ImageWithBetterResolution) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.FW+(img.B*2), img.FH+(img.B*2))
}

func (img *ImageWithBetterResolution) At(x, y int) color.Color {
	if x < img.B || x > img.FW+img.B-1 || y < img.B || y > img.FH+img.B-1 {
		return color.RGBA{}
	}
	i := img.PixOffset(x, y)
	c := img.Pix[i : i+4 : i+4]
	return color.RGBA{
		R: c[0],
		G: c[1],
		B: c[2],
		A: c[3],
	}
}
func (img *ImageWithBetterResolution) PixOffset(x, y int) int {
	return ((((y - img.B) / img.M) * img.RW) + ((x - img.B) / img.M)) * 4
}

func (img *ImageWithBetterResolution) Set(x, y int, r, g, b, a uint8) {
	i := ((img.RW * y) + x) * 4
	img.Pix[i] = r
	img.Pix[i+1] = g
	img.Pix[i+2] = b
	img.Pix[i+3] = a
}
