package skin

import (
	"math"
)

func sizeToWAndH(size int) (int, int) {
	if size < 64*32 {
		return 0, 0
	}
	if size/64 == 32 {
		return 64, 32
	}
	sq := math.Sqrt(float64(size))
	if sq == float64(int(sq)) {
		return int(sq), int(sq)
	}
	sq = math.Sqrt(float64(size / 2))
	if sq == float64(int(sq)) {
		return int(sq), int(sq) / 2
	}
	return 0, 0
}
