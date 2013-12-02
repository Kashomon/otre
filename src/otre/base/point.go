package base

import (
	"fmt"
)

type Coord interface {
	Xf() float64
	Yf() float64
}

const strtemplate = "%v,%v"

func FromSgfCoord(coord string) (*IntPt, error) {
	if len(coord) != 2 {
		return nil, fmt.Errorf("Expected size 2, Found: %v", len(coord))
	}
	x := int(coord[0]) - int('a')
	y := int(coord[1]) - int('a')
	if x < 0 || y < 0 || x > 36 || y > 36 {
		return nil, fmt.Errorf("Unexpected values for pt coords. x:%v, y:%v", x, y)
	}
	return &IntPt{x, y}, nil
}

type IntPt struct {
	X, Y int
}

func (p *IntPt) Xf() float64    { return float64(p.X) }
func (p *IntPt) Yf() float64    { return float64(p.Y) }
func (p *IntPt) String() string { return fmt.Sprintf(strtemplate, p.X, p.Y) }

type CoordPt struct {
	X, Y float64
}

func (p *CoordPt) Xf() float64    { return p.X }
func (p *CoordPt) Yf() float64    { return p.Y }
func (p *CoordPt) String() string { return fmt.Sprintf(strtemplate, p.X, p.Y) }

// Reflect along an X axis, which is, perhaps confusingly specified as a y
// value. This is an immutable operation on a point.
func (p *CoordPt) ReflectY(xaxis float64) *CoordPt {
	return &CoordPt{p.X, -p.Y + 2 * xaxis}
}

// Reflect along an Y axis, which is, perhaps confusingly specified as an x
// value. this is an immutable operation on a point.
func (p *CoordPt) ReflectX(yaxis float64) *CoordPt {
	return &CoordPt{-p.X + 2 * yaxis, p.Y}
}

// Translate a point by some x / y value.  This is an immutable operation,
// returning a new Coordinate point.
func (p *CoordPt) Translate(x, y float64) *CoordPt {
	return &CoordPt{p.X + x, p.Y + y}
}
