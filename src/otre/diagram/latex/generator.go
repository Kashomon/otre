// Code to generate Latex based diagrams.
package latex

import (
	"bytes"
	"otre/core"
)

type Generator struct {
	// The output buffer.
	b bytes.Buffer

	// Set to record which intersections have black stones.  In these cases, we
	// need to flip the color of certain marks and labels.
	bstones map[string]bool
	wstones map[string]bool
}

func Generate(mt *core.Movetree, tp core.Treepath) {
	
}
