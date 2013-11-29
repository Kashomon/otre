package eps

import (
	"fmt"
)

var BLACK = &Grayscale{0}
var WHITE = &Grayscale{1}

type Color interface {
	ColorString() string
}

// EPS Color struct.
type RGB struct {
	r, g, b float32
}

func (c *RGB) ColorString() string {
	return fmt.Sprintf("%v %v %v setrgbcolor", c.r, c.g, c.b)
}

type Grayscale struct {
	p float32
}

func (c *Grayscale) ColorString() string {
	return fmt.Sprintf("%v setgray", c.p)
}
