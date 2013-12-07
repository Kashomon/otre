package eps

import (
	"otre/point"
	"otre/core"
	"regexp"
	"strings"
)

const (
	ROOT_TWO   float64 = 1.41421
	ROOT_THREE float64 = 1.73205
)

func (g *Generator) genMarks(p *core.Props) *Generator {
	if pts, err := p.GetAsPoints(core.MA); err == nil {
		g.xMarks(pts)
	}
	if pts, err := p.GetAsPoints(core.TR); err == nil {
		g.triangleMarks(pts)
	}
	if pts, err := p.GetAsPoints(core.CR); err == nil {
		g.circleMarks(pts)
	}
	if pts, err := p.GetAsPoints(core.SQ); err == nil {
		g.squareMarks(pts)
	}
	if data, ok := p.Get(core.LB); ok {
		g.labels(data)
	}
	return g
}

// Draw X Marks
func (g *Generator) xMarks(pts []*point.IntPt) *Generator {
	r := g.c.Radius
	fudge := float64(r / 5) // so we stay within the circle
	a := ROOT_TWO/2*r - fudge
	g.b.WriteString(bCrossDef + wCrossDef)
	g.b.WriteString(Stroke(1))
	for i := range pts {
		s := pts[i].String()
		c := g.c.CoordMap[s]
		tl := c.Translate(-a, a)
		bl := c.Translate(-a, -a)
		br := c.Translate(a, -a)
		tr := c.Translate(a, a)
		if _, ok := g.bstones[s]; ok {
			g.b.WriteString(Cross(tl, bl, br, tr, WHITE))
		} else if _, ok := g.wstones[s]; ok {
			g.b.WriteString(Cross(tl, bl, br, tr, BLACK))
		} else {
			g.b.WriteString(WhiteDisc(c, a))
			g.b.WriteString(Cross(tl, bl, br, tr, BLACK))
		}
	}
	return g
}

func (g *Generator) triangleMarks(pts []*point.IntPt) *Generator {
	fudge := g.c.Radius / 4. // so we stay within the circle
	r := g.c.Radius - fudge
	g.b.WriteString(bTriDef + wTriDef)
	g.b.WriteString(Stroke(1))
	for i := range pts {
		s := pts[i].String()
		c := g.c.CoordMap[s]
		right := c.Translate(r*ROOT_THREE/2., -r*0.5)
		left := c.Translate(r*-ROOT_THREE/2., -r*0.5)
		top := c.Translate(0, r)
		if _, ok := g.bstones[s]; ok {
			g.b.WriteString(Triangle(right, left, top, WHITE))
		} else if _, ok := g.wstones[s]; ok {
			g.b.WriteString(Triangle(right, left, top, BLACK))
		} else {
			g.b.WriteString(WhiteDisc(c, g.c.Radius * .95))
			g.b.WriteString(Triangle(right, left, top, BLACK))
		}
	}
	return g
}

func (g *Generator) circleMarks(pts []*point.IntPt) *Generator {
	r := g.c.Radius / 2
	g.b.WriteString(circleDef)
	g.b.WriteString(Stroke(1)) // Why is this here?
	for i := range pts {
		s := pts[i].String()
		c := g.c.CoordMap[s]
		if _, ok := g.bstones[s]; ok {
			g.b.WriteString(Circle(c, r, WHITE))
		} else if _, ok := g.wstones[s]; ok {
			g.b.WriteString(Circle(c, r, BLACK))
		} else {
			g.b.WriteString(WhiteDisc(c, r))
			g.b.WriteString(Circle(c, r, BLACK))
		}
	}
	return g
}

func (g *Generator) squareMarks(pts []*point.IntPt) *Generator {
	r := g.c.Radius
	fudge := float64(r / 6) // so we stay within the circle
	a := ROOT_TWO/2*r - fudge
	g.b.WriteString(bSquareDef + wSquareDef)
	g.b.WriteString(Stroke(1))
	for i := range pts {
		s := pts[i].String()
		c := g.c.CoordMap[s]
		tl := c.Translate(-a, a)
		bl := c.Translate(-a, -a)
		br := c.Translate(a, -a)
		tr := c.Translate(a, a)
		if _, ok := g.bstones[s]; ok {
			g.b.WriteString(Square(tl, bl, br, tr, WHITE))
		} else if _, ok := g.wstones[s]; ok {
			g.b.WriteString(Square(tl, bl, br, tr, BLACK))
		} else {
			g.b.WriteString(WhiteDisc(c, a))
			g.b.WriteString(Square(tl, bl, br, tr, BLACK))
		}
	}
	return g
}

func (g *Generator) labels(data []string) *Generator {
	unicodeRegex, err := regexp.Compile("\\\\u")
	g.b.WriteString(textCenterDef)
	if err != nil {
		panic(err)
	}
	for i := range data {
		splat := strings.Split(data[i], ":")
		if len(splat) != 2 {
			continue
		}
		lh, txt := splat[0], splat[1]
		pt, err := point.FromSgfCoord(lh)
		if err != nil {
			continue
		}
		if unicodeRegex.MatchString(txt) {
			// General unicode support is difficult =(
			continue
		}
		s := pt.String()
		c := g.c.CoordMap[s]
		if _, ok := g.bstones[s]; ok {
			g.b.WriteString(Text(c, g.c.Radius*2-2, txt, WHITE))
		} else if _, ok := g.wstones[s]; ok {
			g.b.WriteString(Text(c, g.c.Radius*2-2, txt, BLACK))
		} else {
			g.b.WriteString(WhiteDisc(c, g.c.Radius * .95))
			g.b.WriteString(Text(c, g.c.Radius*2-2, txt, BLACK))
		}
	}
	return g
}
