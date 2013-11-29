package base

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func FromSgfCoord(coord string) (*Point, error) {
	if len(coord) != 2 {
		return nil, fmt.Errorf("Expected size 2, Found: %v", len(coord))
	}
	x := int(coord[0]) - int('a')
	y := int(coord[1]) - int('a')
	if x < 0 || y < 0 || x > 36 || y > 36  {
		return nil, fmt.Errorf("Unexpected values for pt coords. x:%v, y:%v", x, y)
	}
	return &Point{x, y}, nil
}
