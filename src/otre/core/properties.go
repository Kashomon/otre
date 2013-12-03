package core

import (
	"fmt"
	"otre/point"
)

// The SGF Properties.
type Props struct {
	properties map[SgfProperty][]string
}

func NewProps() *Props {
	return &Props{make(map[SgfProperty][]string)}
}

func (p *Props) Set(prop SgfProperty, data ...string) *Props {
	p.properties[prop] = data
	return p
}

func (p *Props) AddTo(property SgfProperty, data ...string) {
	props, ok := p.properties[property]
	if !ok {
		props = make([]string, 0, 10)
	}
	for _, d := range data {
		props = append(props, d)
	}
	p.properties[property] = props
}

func (p *Props) Get(prop SgfProperty) ([]string, bool) {
	props, ok := p.properties[prop]
	return props, ok
}

func (p *Props) GetFirst(prop SgfProperty) (string, bool) {
	props, ok := p.properties[prop]
	if !ok {
		return "", ok
	}
	return props[0], ok
}

func (p *Props) GetAsPoints(prop SgfProperty) ([]*point.IntPt, error) {
	pts := make([]*point.IntPt, 0, 10)
	data, ok := p.Get(prop)
	if !ok {
		return pts, fmt.Errorf("No Data for prop %v", prop)
	}
	for _, v := range data {
		p, err := point.FromSgfCoord(v)
		if err != nil {
			return pts, err
		}
		pts = append(pts, p)
	}
	return pts, nil
}
