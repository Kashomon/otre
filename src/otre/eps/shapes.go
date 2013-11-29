package eps

import (
	"fmt"
	"io"
	"otre/base"
)

func (e *EpsWriter) Board() {
	for i := 0; i < 19; i++ {
		fmt.Println("foo")
	}
}

var linejoin = "1 linejoin"

const (
	circleTempl = "%v %v %v 0 360 arc closepath stroke"
	lineTempl = "%v %v moveto %v %v lineto"
)


func Circle(w io.Writer, pt *base.Point, r int, color Color) {
	fmt.Fprintf(w, "%v %v %v 0 360 arc closepath stroke\n",
		pt.X, pt.Y, r)
}

func (e *EpsWriter) Line(start, finish *base.Point) {

}
