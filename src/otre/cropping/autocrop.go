package core

import (
	"math"
	"otre/core"
)

func FromMovetree(mt *core.Movetree) *Cropping {
	ints := mt.getIntersections()
	middle := ints / 2
	return FromPreset(ALL)
}
