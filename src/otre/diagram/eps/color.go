package eps

import (
	"fmt"
)

// ColorDefault
type ColorDef int

const (
	BLACK ColorDef = ColorDef(iota)
	WHITE
)

// EPS Color interface
type Color interface {
	ColorString() string
}

// Postscript Color container.
type RGB struct {
	r, g, b float64
}

// Generate the PostScript color declaration.
func (c *RGB) ColorString() string {
	return fmt.Sprintf("%v %v %v setrgbcolor", c.r, c.g, c.b)
}

// Postscript color container. A simplification of RGB.
type Grayscale struct {
	p float64
}

// Generate the PostScript color declaration.
func (c *Grayscale) ColorString() string {
	return fmt.Sprintf("%v setgray", c.p)
}
