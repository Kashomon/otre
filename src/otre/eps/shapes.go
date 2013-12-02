package eps

import (
	"fmt"
	"math"
	"otre/base"
)

// circleTemplate = "%v %v %v 0 360 arc closepath gsave %v fill" +
// "\n grestore stroke\n"
const (
	// PostScript Definitions. These must come before template use
	wStoneDef = "/ws {\n0 360 arc closepath gsave\n" +
		"1 setgray fill grestore 0 setgray stroke\n} def\n"
	bStoneDef = "/bs {\n0 360 arc closepath gsave\n" +
		"0 setgray fill grestore 0 setgray stroke\n} def\n"
	wDiscDef = "/dd {\n0 360 arc closepath gsave\n" +
		"1 setgray fill grestore 1 setgray stroke\n} def\n"
	rayDef     = "/rd { moveto lineto 1 setlinecap } def\n"
	wTriDef    = "/tw { moveto lineto lineto closepath 1 setlinecap 1 setgray stroke} def\n"
	bTriDef    = "/tb { moveto lineto lineto closepath 1 setlinecap 0 setgray stroke} def\n"
	bCircleDef = "/cb { 0 360 arc closepath 0 setgray stroke } def\n"
	wCircleDef = "/cw { 0 360 arc closepath 1 setgray stroke } def\n"
	bSquareDef = "/qb { moveto lineto lineto lineto closepath 0 setgray stroke } def\n"
	wSquareDef = "/qw { moveto lineto lineto lineto closepath 1 setgray stroke } def\n"
	bCrossDef  = "/xb { moveto lineto moveto lineto 0 setgray stroke } def\n"
	wCrossDef  = "/xw { moveto lineto moveto lineto 1 setgray stroke } def\n"

	// Templates
	linejoin       = "1 linejoin\n"
	stoneTemplate  = "%v %v %v %v\n" // last entry must be either ws or bs
	lineTemplate   = "%v %v %v %v rd\n"
	triTemplate    = "%v %v %v %v %v %v %v\n"
	squareTemplate = "%v %v %v %v %v %v %v %v %v\n"
	crossTemplate  = "%v %v %v %v %v %v %v %v %v\n"
	circleTemplate = "%v %v %v %v \n"
	stroke         = "%v setlinewidth stroke\n"
	textTemplate2  = "gsave /Helvetica findfont %v scalefont setfont\n" +
		"%v %v moveto (%v) dup stringwidth pop 2 div neg 0 " +
		"rmoveto %v setgray show grestore\n"
)

func PsDefs() string {
	return wStoneDef + bStoneDef + rayDef + wDiscDef
}

func Stone(pt base.Coord, r float64, c ColorDef) string {
	sType := "ws"
	if c == BLACK {
		sType = "bs"
	}
	return fmt.Sprintf(stoneTemplate, twop(pt.Xf()), twop(pt.Yf()), twop(r), sType)
}

func StarPoint(pt base.Coord, r float64) string {
	return fmt.Sprintf(stoneTemplate, twop(pt.Xf()), twop(pt.Yf()), twop(r), "bs")
}

func Line(s, e base.Coord) string {
	return fmt.Sprintf(lineTemplate,
		twop(s.Xf()), twop(s.Yf()), twop(e.Xf()), twop(e.Yf()))
}

func Cross(tl, bl, br, tr base.Coord, c ColorDef) string {
	cType := "xw"
	if c == BLACK {
		cType = "xb"
	}
	return fmt.Sprintf(crossTemplate,
		twop(tl.Xf()), twop(tl.Yf()),
		twop(br.Xf()), twop(br.Yf()),
		twop(bl.Xf()), twop(bl.Yf()),
		twop(tr.Xf()), twop(tr.Yf()),
		cType)
}

func Triangle(right, left, top base.Coord, c ColorDef) string {
	cType := "tw"
	if c == BLACK {
		cType = "tb"
	}
	return fmt.Sprintf(triTemplate,
		twop(right.Xf()), twop(right.Yf()),
		twop(left.Xf()), twop(left.Yf()),
		twop(top.Xf()), twop(top.Yf()),
		cType)
}

func Circle(pt base.Coord, r float64, c ColorDef) string {
	cType := "cw"
	if c == BLACK {
		cType = "cb"
	}
	return fmt.Sprintf(circleTemplate, twop(pt.Xf()), twop(pt.Yf()), twop(r), cType)
}

func Square(tl, bl, br, tr base.Coord, c ColorDef) string {
	cType := "qw"
	if c == BLACK {
		cType = "qb"
	}
	return fmt.Sprintf(squareTemplate,
		twop(tl.Xf()), twop(tl.Yf()),
		twop(bl.Xf()), twop(bl.Yf()),
		twop(br.Xf()), twop(br.Yf()),
		twop(tr.Xf()), twop(tr.Yf()),
		cType)
}

func Text(pt base.Coord, size float64, chars string, c ColorDef) string {
	fsize := twop(size) - 2
	return fmt.Sprintf(textTemplate2,
		fsize,
		twop(pt.Xf()), twop(pt.Yf()-fsize/3),
		chars,
		c)
}

func WhiteDisc(pt base.Coord, r float64) string {
	return fmt.Sprintf(stoneTemplate, twop(pt.Xf()), twop(pt.Yf()), twop(r), "dd")
}

func Stroke(linewidth float64) string {
	return fmt.Sprintf(stroke, linewidth)
}

// Convert a float to two decimal places
func twop(f float64) float64 {
	return math.Floor(f*100.0) / 100.0
}
