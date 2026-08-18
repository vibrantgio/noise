# AGENTS.md — noise

Perlin and simplex noise: `NewPerlin2D`, `NewPerlin3D`, `NewSimplex2D` and
`NewSimplex3D`, each seeded and each answering `Noise(x, y)` or `Noise(x,
y, z)` over a shared `Permutations` table. Converted from Stefan
Gustavson's public-domain implementation by way of josephg/noisejs.

**Layer.** Outside ADR-001's tier table: a support library, which the rule
binds in one direction only — every tier may import it, and it may import
nothing in the table itself. One module, one package, standard library
only. Its root module imports nothing else in the organization. That
direction is measured rather than typed — `scripts/check-layers.sh --edges`
reports the graph and `scripts/sync-agents.sh` renders these sentences from
it — so correcting them here changes nothing. The other direction is
measured too and deliberately not written down: the gate checks the graph
both ways, but a public API's consumers are unknowable, so this file says
what its module needs and never who needs it.

**Read the canonical guide before you write code against this module.** It is
the organization's one agent guide — the module inventory with current tags,
the application skeleton, the MVU loop and rx semantics, typography, and the
pitfalls that are not guessable. It lives exactly once, in `vibrantgio/.github`,
and this file links it rather than copying it:

    https://raw.githubusercontent.com/vibrantgio/.github/master/llms.txt

**Module.** `github.com/vibrantgio/noise`, one module at the repository
root.

**Build and test.** From the repository root:

    go build ./... && go test ./...

**The reference images are golden images without a flag.** The four
`TestImage*` tests render a 1024×1024 PNG and compare its bytes against
`ref_p2d.png`, `ref_p3d.png`, `ref_s2d.png` and `ref_s3d.png`, embedded from
the repository root. No command line regenerates them; `noise_test.go`
declares

    const write_reference_image = false

and each test writes its PNG only while that constant is `true`.

Regenerating therefore takes two runs, because the bytes a test compares
against were embedded when the binary was built — before that same run
overwrote the file. Flip the constant to `true`: `go test ./...` rewrites the
four PNGs and still fails; run it again and it passes against what it just
wrote. Look at the images, flip the constant back to `false`, confirm the
tests are still green, and say in the commit that you moved them. Leaving it
`true` makes the tests self-approving.
