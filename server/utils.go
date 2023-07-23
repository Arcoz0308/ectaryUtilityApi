package server

func validSize(s int) bool {
	if s == 0 || (s >= 8 && s <= 2048 && s%8 == 0) {
		return true
	}
	return false
}
