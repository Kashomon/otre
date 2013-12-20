package eps

import (
	"otre/point"
)

// Coordinates for stones.
type Coords struct {
	// Map from string form of the intersection to the coordinate
	CoordMap map[string]*point.CoordPt
	Radius   float64
}

// Calculate the coords for each point on the Go board.
func ConstructCoords(totalInts, side int) *Coords {
	spacing := float64(side) / float64(totalInts+1)
	radius := spacing / 2
	startX := spacing
	startY := spacing
	cmap := make(map[string]*point.CoordPt)
	for row := 0; row < totalInts; row++ {
		for col := 0; col < totalInts; col++ {
			intPt := &point.IntPt{col, row}
			coordPt := &point.CoordPt{startX + spacing*float64(col),
				startY + spacing*float64(row)}
			coordPt = coordPt.ReflectY(float64(side) / 2)
			strPt := intPt.String()
			cmap[strPt] = coordPt
		}
	}
	return &Coords{cmap, radius}
}
