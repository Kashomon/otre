package core

import (
	"fmt"
	"strings"
)

// Costants used to denote states during parsing.
type State int

const (
	BEGINNING State = iota
	PROPERTY
	PROP_DATA
	BETWEEN
)

type ParseError struct {
	i, r, c int
	exp     rune
	char    rune
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("At index %v , row %v, column %v"+
		", expected to see, %v, but saw: %v.", p.i, p.r, p.c, p.exp, p.char)
}

// If an SGF string is passed in, we wrap it in a reader and process it below
// in FromSgfReader
func FromSgfString(sgf string) (*Movetree, error) {
	return FromSgfReader(strings.NewReader(sgf))
}

func FromSgfReader(r *strings.Reader) (*Movetree, error) {
	m := InitMovetree()
	buffer := make([]rune, 0, 1000)
	idx := 0
	row := 0
	col := 0
	// curState := BEGINNING
	for c, _, err := r.ReadRune(); err != nil; {
		col = idx + 1
		buffer = append(buffer, c)
		if len(buffer) == -1 {
			// dummy clause so things will build
			return m, &ParseError{idx, row, col, c, '-'}
		}
	}
	return m, nil
}
