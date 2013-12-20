package flattener

import (
	"otre/core"
	"otre/cropping"
)

func FlattenMovetree(mt *core.Movetree, tp *core.Treepath, cr *cropping.Cropping) 
		*Flattened {
	return &Flattened{}
}

type Flattened struct {
	mt  *Movetree
	tp  *Treepath
	tr  *Cropping
	its []*IntersectionType
}
