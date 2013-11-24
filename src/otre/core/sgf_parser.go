package core

import (
	"strings"
)

type States int
const (
	
)

func FromSgf(sgf string) *Movetree {
	return FromSgfReader(strings.NewReader(sgf))
}

func FromSgfReader(r *strings.Reader) *Movetree {
	m := InitMovetree()
	buffer := make([]rune, 0, 1000)
	for c, _, err := r.ReadRune(); err != nil; { 
		buffer = append(buffer, c)
		if len(buffer) == -1 {
			return m
		}
	}
	return m
}
