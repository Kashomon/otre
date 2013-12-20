package eps

import (
	"fmt"
	"math"
	"otre/point"
)

// circleTemplate = "%v %v %v 0 360 arc closepath gsave %v fill" +
// "\n grestore stroke\n"
const (
	// PostScript Definitions. These must come before template use

	// Format: grayval, x, y, r,
	stoneDef = "/sd {\n 0 360 arc closepath setgray gsave\n" +
		"fill grestore 0 setgray stroke\n} def\n"

	wDiscDef = "/dd {\n0 360 arc closepath gsave\n" +
		"1 setgray fill grestore 1 setgray stroke\n} def\n"
	rayDef     = "/rd { moveto lineto 1 setlinecap } def\n"
	wTriDef    = "/tw { moveto lineto lineto closepath 1 setlinecap 1 setgray stroke} def\n"
	bTriDef    = "/tb { moveto lineto lineto closepath 1 setlinecap 0 setgray stroke} def\n"

	// Format: grayval, x, y, r 
	circleDef = "/cd { 0 360 arc closepath setgray stroke } def\n"
	bSquareDef = "/qb { moveto lineto lineto lineto closepath 0 setgray stroke } def\n"
	wSquareDef = "/qw { moveto lineto lineto lineto closepath 1 setgray stroke } def\n"
	bCrossDef  = "/xb { moveto lineto moveto lineto 0 setgray stroke } def\n"
	wCrossDef  = "/xw { moveto lineto moveto lineto 1 setgray stroke } def\n"

	// Format: string, grayvalue, x, y, fontsize
	textCenterDef = "/ct {\n" +
		"gsave /Helvetica findfont 12 scalefont setfont setgray moveto \n" +
		"dup stringwidth pop -2 div 0 rmoveto\n show grestore} def\n"

	// Templates
	linejoin       = "1 linejoin\n"
	// Format: x, y, r, grayvalue
	stoneTemplate  = "%v %v %v %v sd\n"
	discTemplate = "%v %v %v dd\n"
	lineTemplate   = "%v %v %v %v rd\n"
	triTemplate    = "%v %v %v %v %v %v %v\n"
	squareTemplate = "%v %v %v %v %v %v %v %v %v\n"
	crossTemplate  = "%v %v %v %v %v %v %v %v %v\n"
	// Format: grayval, x, y, r
	circleTemplate = "%v %v %v %v cd\n"
	stroke         = "%v setlinewidth stroke\n"

	// Format: string, x, y, grayvalue
	textTemplate   = "(%v) %v %v %v ct\n"
	textTemplate2  = "gsave /Helvetica findfont %v scalefont setfont\n" +
		"%v %v moveto (%v) dup stringwidth pop 2 div neg 0 " +
		"rmoveto %v setgray show grestore\n"
)

func PsDefs() string {
	return stoneDef + rayDef + wDiscDef
}

func Stone(pt point.Coord, r float64, c ColorDef) string {
	return fmt.Sprintf(stoneTemplate, c, twop(pt.Xf()), twop(pt.Yf()), twop(r))
}

func StarPoint(pt point.Coord, r float64) string {
	return Stone(pt, r, BLACK)
}

func Line(s, e point.Coord) string {
	return fmt.Sprintf(lineTemplate,
		twop(s.Xf()), twop(s.Yf()), twop(e.Xf()), twop(e.Yf()))
}

func Cross(tl, bl, br, tr point.Coord, c ColorDef) string {
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

func Triangle(right, left, top point.Coord, c ColorDef) string {
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

func Circle(pt point.Coord, r float64, c ColorDef) string {
	return fmt.Sprintf(circleTemplate, c, twop(pt.Xf()), twop(pt.Yf()), twop(r))
}

func Square(tl, bl, br, tr point.Coord, c ColorDef) string {
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

func Text(pt point.Coord, size float64, chars string, c ColorDef) string {
	fsize := twop(size) - 2
	return fmt.Sprintf(textTemplate, chars,
		twop(pt.Xf()), twop(pt.Yf()-fsize/3), c)
}

func WhiteDisc(pt point.Coord, r float64) string {
	return fmt.Sprintf(discTemplate, twop(pt.Xf()), twop(pt.Yf()), twop(r))
}

func Stroke(linewidth float64) string {
	return fmt.Sprintf(stroke, linewidth)
}

// Convert a float to two decimal places
func twop(f float64) float64 {
	return math.Floor(f*100.0) / 100.0
}
