package eps


import (
	"bytes"
	"fmt"
	"otre/point"
	"otre/core"
)

// A generator is a struct that contains information about diagram gevneration
type Generator struct {
	// The output buffer.
	b    bytes.Buffer
	// The length of a side (only applies to some diagram types)
	side int
	// Coordinates tracker. Doesn't apply to font-based diagrams.
	c    *Coords
	// Set to record which intersections have black stones.  In these cases, we
	// need to flip the color of certain marks and labels.
	bstones map[string]bool
	wstones map[string]bool
}

// Generate an EPS image for a baduk Movetree.
func Generate(mt *core.Movetree, side int) []byte {
	g := &Generator{
		side:    side,
		c:       ConstructCoords(mt.Intersections(), side),
		bstones: make(map[string]bool),
		wstones: make(map[string]bool),
	}
	g.header().content(mt).footer()
	return g.b.Bytes()
}

func (g *Generator) content(mt *core.Movetree) *Generator {
	return g.genDefs().genLines(mt).genStarPoints(mt).genStones(mt).genMarks(mt.Props())
}

func (g *Generator) genDefs() *Generator {
	g.b.WriteString(PsDefs())
	return g
}

func (g *Generator) genLines(mt *core.Movetree) *Generator {
	ints := mt.Intersections()
	for i := 0; i < ints; i++ {
		rowStart, _ := g.c.CoordMap[(&point.IntPt{i, 0}).String()]
		rowEnd, _ := g.c.CoordMap[(&point.IntPt{i, ints - 1}).String()]
		g.b.WriteString(Line(rowStart, rowEnd))
		colStart := g.c.CoordMap[(&point.IntPt{0, i}).String()]
		colEnd := g.c.CoordMap[(&point.IntPt{ints - 1, i}).String()]
		g.b.WriteString(Line(colStart, colEnd))
	}
	g.b.WriteString(Stroke(0.5))
	return g
}

func (g *Generator) genStarPoints(mt *core.Movetree) *Generator {
	sp := getStarPoints(mt.Intersections())
	for _, pt := range sp {
		c := g.c.CoordMap[pt.String()]
		g.b.WriteString(StarPoint(c, g.c.Radius*.15))
	}
	return g
}

func getStarPoints(ints int) []*point.IntPt {
	var pts [][]int
	midway := ints / 2
	if ints == 19 {
		pts = [][]int{[]int{3, 9, 15}}
	} else if ints == 13 {
		pts = [][]int{[]int{3, 9}, []int{6}}
	} else if ints == 9 {
		pts = [][]int{[]int{2, 6}, []int{4}}
	} else {
		pts = [][]int{[]int{midway}}
	}
	var outpts []*point.IntPt
	for i := 0; i < len(pts); i++ {
		subslice := pts[i]
		for j := 0; j < len(subslice); j++ {
			for k := 0; k < len(subslice); k++ {
				outpts = append(outpts, &point.IntPt{subslice[j], subslice[k]})
			}
		}
	}
	return outpts
}

func (g *Generator) genStones(mt *core.Movetree) *Generator {
	r := g.c.Radius - 0.2
	bpts, _ := mt.Props().GetAsPoints(core.AB)
	wpts, _ := mt.Props().GetAsPoints(core.AW)
	for i, _ := range bpts {
		strForm := bpts[i].String()
		coord, ok := g.c.CoordMap[strForm]
		if ok {
			g.b.WriteString(Stone(coord, r, BLACK))
			g.bstones[strForm] = true
		}
	}
	for i, _ := range wpts {
		strForm := wpts[i].String()
		coord, ok := g.c.CoordMap[wpts[i].String()]
		if ok {
			g.b.WriteString(Stone(coord, r, WHITE))
			g.wstones[strForm] = true
		}
	}
	return g
}

var headerStr = `%%!PS-Adobe-3.0 EPSF-3.0
%%%%Creator: Otre 0.1
%%%%Pages: 1
%%%%Orientation: Portrait
%%%%BoundingBox: 0 0 %v %v
%%%%HiResBoundingBox: 0 0 %v %v
%%%%EndComments
%%%%Page: 1 1
`

func (g *Generator) header() *Generator {
	g.b.Write([]byte(
		fmt.Sprintf(headerStr, g.side, g.side, g.side, g.side)))
	return g
}

func (g *Generator) footer() *Generator {
	g.b.Write([]byte("%%EOF"))
	return g
}
