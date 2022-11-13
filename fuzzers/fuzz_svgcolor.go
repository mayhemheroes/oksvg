package myfuzz

import "github.com/srwiley/oksvg"

func Fuzz(data []byte) int {
	oksvg.ParseSVGColor(string(data))
	return 0
}
