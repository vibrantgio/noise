package noise

import (
	"bytes"
	_ "embed"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"
)

//go:embed ref_p3d.png
var ref_p3d []byte

func TestImagePerlin3D(t *testing.T) {
	const GRID_SIZE = 4.0
	const RESOLUTION = 128.0
	const COLOR_SCALE = 250.0

	num_pixels := GRID_SIZE / RESOLUTION // 4/128 = 1/32 = 0.03125
	pixel_size := int(1024 / RESOLUTION) // 8

	rgba := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	s := NewPerlin3D(0.0)
	for y := 0.0; y < GRID_SIZE; y += num_pixels / GRID_SIZE {
		for x := 0.0; x < GRID_SIZE; x += num_pixels / GRID_SIZE {
			src := image.NewUniform(HSLToRGB(s.Noise(x, y, 0)*COLOR_SCALE, 50, 50))
			left := int((x / GRID_SIZE) * 1024)
			top := int((y / GRID_SIZE) * 1024)
			dst := image.Rect(left, top, left+pixel_size, top+pixel_size)
			draw.Draw(rgba, dst, src, image.ZP, draw.Src)
		}
	}

	// Write reference image
	if write_reference_image {
		f, err := os.Create(SourceFile("ref_p3d.png"))
		if err != nil {
			t.Errorf("failed to create reference image: %v", err)
		}
		defer f.Close()
		png.Encode(f, rgba)
	}

	// Compare with reference image
	f := bytes.NewBuffer(nil)
	png.Encode(f, rgba)
	if !bytes.Equal(f.Bytes(), ref_p3d) {
		t.Errorf("generated noise image does not match stored image")
	}
}

func FuzzPerlin3D(f *testing.F) {
	testcases := [][]float64{
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 1.0},
		{0.0, 1.0, 0.0},
		{0.0, 1.0, 1.0},
		{1.0, 0.0, 0.0},
		{1.0, 0.0, 1.0},
		{1.0, 1.0, 0.0},
		{1.0, 1.0, 1.0},
		// {1.0, 170.0, 0.0},
	}
	for _, tc := range testcases {
		f.Add(tc[0], tc[1], tc[2])
	}
	s := NewPerlin3D(0.0)
	count := 0
	f.Fuzz(func(t *testing.T, x, y, z float64) {
		count++
		generate := func(x, y, z float64) float64 {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Panic on {%f,%f,%f} : %v : %d", x, y, z, r, count)
				}
			}()
			return s.Noise(x, y, z)
		}
		n := generate(x, y, z)
		if n < -1 || n > 1 {
			t.Errorf("Return value out of range %f", n)
		}
	})
}
