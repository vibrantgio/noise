package noise

import (
	"image/color"
	"math"
	"path"
	"runtime"
)

func SourceFile(filename string) string {
	_, sourcepath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(sourcepath), filename)
}

// HSLToRGB converts the color from HSL to RGB with a lossy algorithm.
//
// reference: https://www.ginifab.com.tw/tools/colors/js/colorconverter.js
// h 0-360
// s 0-100
// l 0-100
func HSLToRGB(h float64, s float64, l float64) color.RGBA {
	h = h / 360
	s = s / 100
	l = l / 100

	var r, g, b float64
	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = HueToRGB(p, q, h+float64(float64(1)/float64(3)))
		g = HueToRGB(p, q, h)
		b = HueToRGB(p, q, h-float64(float64(1)/float64(3)))
	}

	return color.RGBA{
		R: uint8((math.Round(r * 255))),
		G: uint8((math.Round(g * 255))),
		B: uint8((math.Round(b * 255))),
		A: 255,
	}
}

// HueToRGB converts the color from Hue to RGB.
//
// reference: https://www.ginifab.com.tw/tools/colors/js/colorconverter.js
func HueToRGB(p float64, q float64, t float64) float64 {

	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < float64(1)/float64(6) {
		return p + (q-p)*6*t
	}
	if t < float64(1)/float64(2) {
		return q
	}
	if t < float64(2)/float64(3) {
		return p + (q-p)*(float64(2)/float64(3)-t)*6
	}
	return p
}
