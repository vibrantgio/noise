# noise

    import "github.com/vibrantgio/noise"

Package noise implements Perlin and Simplex noise functions.

Adapted from https://github.com/josephg/noisejs/blob/master/perlin.js

This code was placed in the public domain by its original developer,
Stefan Gustavson. You may use it as you see fit, but attribution
is appreciated.

The code was converted code to Go. Note that the Go code uses a Simplex3D
struct instead of a class, and the methods are defined on the struct.
This is because Go does not support classes, and structs with associated
methods are used to achieve similar functionality.

## For coding assistants

Read the org guide before you write code against this module: the plan
root's [`AGENTS.md`](https://github.com/vibrantgio/.github/blob/master/AGENTS.md).

[`AGENTS.md`](./AGENTS.md) in this repository has the build and test commands,
and — since nothing else in the organization says so — how to regenerate the
four reference images the `TestImage*` tests compare against.
