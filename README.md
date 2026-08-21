# noise

    import "github.com/vibrantgio/noise"

Package noise implements Perlin and Simplex noise functions.

Adapted from https://github.com/josephg/noisejs/blob/master/perlin.js

This code was placed in the public domain by its original author,
Stefan Gustavson. You may use it as you see fit, but attribution
is appreciated.

The code was converted code to Go. Note that the Go code uses a Simplex3D
struct instead of a class, and the methods are defined on the struct.
This is because Go does not support classes, and structs with associated
methods are used to achieve similar functionality.

## For coding assistants

Read the canonical guide before writing code against this module — the module
inventory with current tags, the application skeleton, MVU and rx semantics,
typography, and the pitfalls that are not guessable:

<https://raw.githubusercontent.com/vibrantgio/workbench/master/llms.txt>

[`AGENTS.md`](./AGENTS.md) in this repository has the build and test commands,
and — since nothing else in the organization says so — how to regenerate the
four reference images the `TestImage*` tests compare against.
