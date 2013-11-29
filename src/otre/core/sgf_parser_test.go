package core

import "testing"

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

func TestComplexProblemParsing(t *testing.T) {
	p := FromSgfString(complexproblem)
	mt, err := p.Parse()
	if err != nil {
		t.Errorf("Parsing failed, %v", err)
	}
	if mt == nil {
		t.Errorf("Movetree Nil, %v", p.ParseError("Movetree Nill"))
	}
	if p.curchar == '〜' {
		t.Errorf("Unexpected %v, Parser %v", string([]rune{p.curchar}),
			p.ParseError("Still at start state"))
	}
	if mt.Node() != mt.Root() {
		t.Errorf("Current node is not root node")
	}
	props, ok := mt.Node().GetProps(AP)
	if !ok {
		t.Errorf("No property: %v", string(AP))
	}
	if len(props) != 1 {
		t.Errorf("Found length %v for %v", len(props), AP)
	}
	if props[0] != "Glift" {
		t.Errorf("Expected value: %v", props[0])
	}
	if l := len(mt.Node().Children()); l != 4 {
		t.Errorf("Expected: 4 children. Saw, %v. %v", l, p.ParseError("--"))
	}
	mt.MoveDown(1)
	if props, ok := mt.Node().GetProps("B"); !ok || props[0] != "ma" {
		t.Errorf("Expected 'ma' -- instead, saw: %v", props[0])
	}
}
