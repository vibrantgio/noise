package noise

import "math"

type Perlin2D struct {
	Permutations
}

func NewPerlin2D(seed float64) *Perlin2D {
	s := &Perlin2D{}
	s.Seedf(seed)
	return s
}

func (n *Perlin2D) Noise(x, y float64) float64 {
	// Find unit grid cell containing point
	fx := math.Floor(x)
	fy := math.Floor(y)

	// Get relative xy coordinates of point within that cell
	x -= fx
	y -= fy

	// Wrap the integer cells at 255 (smaller integer period can be introduced here)
	x0 := uint8(fx)
	y0 := uint8(fy)

	// Calculate noise contributions from each of the four corners
	n00 := n.gradP[x0+n.perm[y0]].dot2(x, y)
	n01 := n.gradP[x0+n.perm[y0+1]].dot2(x, y-1)
	n10 := n.gradP[x0+1+n.perm[y0]].dot2(x-1, y)
	n11 := n.gradP[x0+1+n.perm[y0+1]].dot2(x-1, y-1)

	// Compute the fade curve value for x
	u := fade(x)

	// Interpolate the four results
	return lerp(
		lerp(n00, n10, u),
		lerp(n01, n11, u), fade(y))
}
