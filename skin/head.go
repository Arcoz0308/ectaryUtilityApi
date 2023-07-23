package skin

import (
	"ectary/bimage"
	"image"
	"image/color"
)

func (s *Skin) Head(border int, resolution int) image.Image {
	// if there don't are a border and don't are resizing
	if border == 0 && (resolution == 0 || resolution == s.H/8) {
		return s.headByDefault()
		// if there are a border but not a resizing
	} else if resolution == 0 || resolution <= s.H/8 {
		return s.headWithBorder(border)
		// if there are resizing and optional border
	} else {
		return s.headWithResizing(border, resolution)
	}
}
func (s *Skin) headByDefault() image.Image {
	size := s.W / 8
	r := image.Rect(0, 0, size, size)
	img := image.NewRGBA(r)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			i1 := (((y + size) * s.W) + x + size) * 4
			i2 := (((y + size) * s.W) + (x + size + (s.W / 2))) * 4
			rgba := s.Pix[i1 : i1+4 : i1+4]
			rgba2 := s.Pix[i2 : i2+4 : i2+4]
			if rgba2[3] > 0 {
				rgba = rgba2
			}
			c := color.RGBA{
				R: rgba[0],
				G: rgba[1],
				B: rgba[2],
				A: rgba[3],
			}
			img.Set(x, y, c)
		}
	}
	return img
}
func (s *Skin) headWithResizing(border int, resolution int) image.Image {
	size := s.W / 8
	m := resolution / size
	img := bimage.NewImageWithBetterResolution(m, size, size, border, resolution, resolution)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			i1 := (((y + size) * s.W) + x + size) * 4
			i2 := (((y + size) * s.W) + (x + size + (s.W / 2))) * 4
			rgba := s.Pix[i1 : i1+4 : i1+4]
			rgba2 := s.Pix[i2 : i2+4 : i2+4]
			if rgba2[3] > 0 {
				rgba = rgba2
			}
			img.Set(x, y, rgba[0], rgba[1], rgba[2], rgba[3])
		}
	}
	return img
}
func (s *Skin) headWithBorder(border int) image.Image {
	size := s.W / 8
	r := image.Rect(0, 0, size+(border*2), size+(border*2))
	img := image.NewRGBA(r)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			i1 := (((y + size) * s.W) + x + size) * 4
			i2 := (((y + size) * s.W) + (x + size + (s.W / 2))) * 4
			rgba := s.Pix[i1 : i1+4 : i1+4]
			rgba2 := s.Pix[i2 : i2+4 : i2+4]
			if rgba2[3] > 0 {
				rgba = rgba2
			}
			c := color.RGBA{
				R: rgba[0],
				G: rgba[1],
				B: rgba[2],
				A: rgba[3],
			}
			img.Set(x+border, y+border, c)
		}
	}
	return img
}
