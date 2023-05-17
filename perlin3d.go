package noise

import (
	"fmt"
	"math"
)

type Perlin3D struct {
	Permutations
}

// Seed the noise functions. Only 65536 different seeds are supported.
// Use a float between 0 and 1 or an integer from 1 to 65536.
func NewPerlin3D(seed float64) *Perlin3D {
	s := &Perlin3D{}
	s.Seedf(seed)
	return s
}

func (n *Perlin3D) Noise(x, y, z float64) float64 {
	fx := math.Floor(x)
	fy := math.Floor(y)
	fz := math.Floor(z)

	x -= fx
	y -= fy
	z -= fz

	x0 := uint8(fx)
	y0 := uint8(fy)
	z0 := uint8(fz)

	n000 := n.gradP[x0+n.perm[y0+n.perm[z0]]].dot3(x, y, z)
	n001 := n.gradP[x0+n.perm[y0+n.perm[z0+1]]].dot3(x, y, z-1)
	n010 := n.gradP[x0+n.perm[y0+1+n.perm[z0]]].dot3(x, y-1, z)
	n011 := n.gradP[x0+n.perm[y0+1+n.perm[z0+1]]].dot3(x, y-1, z-1)
	n100 := n.gradP[x0+1+n.perm[y0+n.perm[z0]]].dot3(x-1, y, z)
	n101 := n.gradP[x0+1+n.perm[y0+n.perm[z0+1]]].dot3(x-1, y, z-1)
	n110 := n.gradP[x0+1+n.perm[y0+1+n.perm[z0]]].dot3(x-1, y-1, z)
	n111 := n.gradP[x0+1+n.perm[y0+1+n.perm[z0+1]]].dot3(x-1, y-1, z-1)

	u := fade(x)
	v := fade(y)
	w := fade(z)

	result := lerp(
		lerp(
			lerp(n000, n100, u),
			lerp(n001, n101, u), w),
		lerp(
			lerp(n010, n110, u),
			lerp(n011, n111, u), w),
		v)
	if result < -1.0 || result > 1.0 {
		panic(fmt.Sprintf("Perlin3D: invalid noise value %f", result))
	}
	return result
}
