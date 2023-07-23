package skin

import (
	error2 "ectary/utils/error"
	"image"
	"image/draw"
)

type Skin struct {
	W, H int
	Pix  []uint8
}

func FromData(data []uint8) (*Skin, *error2.ApiError) {
	w, h := sizeToWAndH(len(data) / 4)
	if w == 256 {
		return nil, nil
	}
	if w == 0 {
		return nil, &error2.ApiError{
			Code: error2.CodeErrInvalidSkinSize,
			AdditionalData: map[string]interface{}{
				"px": len(data) / 4,
			},
			Err: error2.ErrInvalidSkinSize,
		}
	}
	return &Skin{W: w, H: h, Pix: data}, nil
}

func FromImage(img image.Image) *Skin {
	if img, ok := img.(*image.RGBA); ok {
		return &Skin{Pix: img.Pix, W: img.Rect.Size().X, H: img.Rect.Size().Y}
	}
	b := img.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Size().X, b.Size().Y))
	draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)
	return &Skin{Pix: m.Pix, W: m.Rect.Dx(), H: m.Rect.Dy()}
}

func (s *Skin) FullSkin() image.Image {
	return &image.RGBA{
		Pix:    s.Pix,
		Stride: 4 * s.W,
		Rect:   image.Rect(0, 0, s.W, s.H),
	}
}
