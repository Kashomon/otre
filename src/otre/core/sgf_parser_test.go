package core

import "io"
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
	if err != nil && err != io.EOF {
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
	props, ok := mt.Props().Get(AP)
	if !ok {
		t.Errorf("No property: %v", string(AP))
	}
	if len(props) != 1 {
		t.Errorf("Found length %v for %v", len(props), AP)
	}
	if props[0] != "Glift" {
		t.Errorf("Expected value: %v", props[0])
	}
	if ints, _ := mt.Props().GetFirst(SZ); ints != "19" {
		t.Errorf("Expected number of intersections to be 19")
	}
	if l := len(mt.Node().Children()); l != 4 {
		t.Errorf("Expected: 4 children. Saw, %v. %v", l, p.ParseError("--"))
	}
	mt.MoveDown(1)
	if props, ok := mt.Props().Get("B"); !ok || props[0] != "ma" {
		t.Errorf("Expected 'ma' -- instead, saw: %v", props[0])
	}
}

var marktest string = `
(;GM[1]FF[4]CA[UTF-8]AP[CGoban:3]ST[2]
RU[Japanese]SZ[19]KM[0.00]
PW[White]PB[Black]
AW[na][oa][pa][qa][ra][sa][ka][la][ma][ja]
AB[nb][ob][pb][qb][rb][sb][kb][lb][mb][jb]
LB[pa:A][ob:2][pb:B][pc:C][pd:D]
[oa:1][oc:3][ne:9][oe:8][pe:7][qe:6][re:5][se:4]
[nf:15][of:14][pf:13][qf:11][rf:12][sf:10]
[ng:22][og:44][pg:100]
[ka:a][kb:b][kc:c][kd:d][ke:e][kf:f][kg:g]
[ma:\u4e00][mb:\u4e8c][mc:\u4e09][md:\u56db][me:\u4e94]
[la:\u516d][lb:\u4e03][lc:\u516b][ld:\u4e5d][le:\u5341]
MA[na][nb][nc]
CR[qa][qb][qc]
TR[sa][sb][sc]
SQ[ra][rb][rc]
)`

func TestMarksParsing(t *testing.T) {
	p := FromSgfString(marktest)
	mt, err := p.Parse()
	if err != nil && err != io.EOF {
		t.Errorf("Parsing failed, %v", err)
	}
	if p.curstate != BETWEEN {
		t.Errorf("Expected state BETWEEN, found %v", p.curstate)
	}
	if _, ok := mt.Props().Get("SQ"); !ok {
		t.Errorf("Couldn't find SQ")
	}
}
