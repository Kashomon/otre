package eps

import (
	"bytes"
	"otre/base"
	"otre/core"
)

type EpsWriter struct {
	b bytes.Buffer
}

func Generate(mt *core.Movetree) []byte {
	e := &EpsWriter{}
	e.header()
	e.content(mt)
	e.footer()
	return e.b.Bytes()
}

func (e *EpsWriter) content(mt *core.Movetree) {
	pt := &base.Point{50, 50}
	// writer := e.b.(io.Writer)
	Circle(&e.b, pt, 10, BLACK)
}

var headerStr = `%!PS-Adobe-3.0 EPSF-3.0
%%Creator: Otre 0.1
%%Pages: 1
%%Orientation: Portrait
%%BoundingBox: 0 0 300 300
%%HiResBoundingBox: 0 0 300 300
%%EndComments
%%Page: 1 1
`

func (e *EpsWriter) header() *EpsWriter {
	e.b.Write([]byte(headerStr))
	return e
}

func (e *EpsWriter) footer() *EpsWriter {
	e.b.Write([]byte("\n%%EOF"))
	return e
}
