package main

import (
	"os"
	"otre/core"
	"otre/eps"
)

var complexproblem string = `
(;GM[1]FF[4]CA[UTF-8]AP[Glift]ST[2]
RU[Japanese]SZ[19]KM[0.00]
C[Black to play. There aren't many options 
to choose from, but you might be surprised at the answer!]
PW[White]PB[Black]AW[pa][qa][nb][ob][qb][oc][pc][md][pd][ne][oe]
AB[na][ra][mb][rb][lc][qc][ld][od][qd][le][pe][qe][mf][nf][of][pg]
(;B[mc]
	;W[nc]C[White lives.])
(;B[ma]
	(;W[oa]
		;B[nc]
		;W[nd]
		;B[mc]C[White dies.]GB[1])
	(;W[mc]
		(;B[oa]
		;W[nd]
		;B[pb]C[White lives])
		(;B[nd]
			;W[nc]
			;B[oa]C[White dies.]GB[1]))
	(;W[nd]
		;B[mc]
		;W[oa]
		;B[nc]C[White dies.]GB[1]))
(;B[nc]
	;W[mc]C[White lives])
(;B[]C[A default consideration]
	;W[mc]C[White lives easily]))`

func main() {
	mt, perr := core.FromSgfString(complexproblem).Parse()
	if perr != nil {
		panic(perr)
	}

	fo, ferr := os.Create("output.eps")
	if ferr != nil {
		panic(ferr)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	fo.Write(eps.Generate(mt))
}
