package core

import (
	"fmt"
	"strings"
)

// Special chars , used to delimite sections of the SGF.
const (
	LPAREN rune = '('
	RPAREN rune = ')'
	LBRACE rune = '['
	RBRACE rune = ']'
	SCOLON rune = ';'

	// WhiteSpace
	TAB     rune = '\t'
	SPACE   rune = ' '
	NEWLINE rune = '\n'
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
	s       State
	i, r, c int
	char    rune
	msg     string
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("Error during state %v, at index %v , row %v, column %v"+
		", curchar %v. %v", p.s, p.i, p.r, p.c, p.char, p.msg)
}

type Parser struct {
	r             *strings.Reader
	idx, row, col int
	curchar       rune
	curstate      State
	mt            *Movetree
	err           error
}

// If an SGF string is passed in, we wrap it in a reader and process it below
// in FromSgfReader
func FromSgfString(sgf string) *Parser {
	return FromSgfReader(strings.NewReader(sgf))
}

func FromSgfReader(r *strings.Reader) *Parser {
	return &Parser{r, 0, 0, 0, '0', BEGINNING, nil, nil}
}

// Helper function for parsing.
func isWhitespace(c rune) bool {
	return c == NEWLINE || c == TAB || c == SPACE
}

func (p *Parser) ParseError(msg string) *ParseError {
	return &ParseError{p.curstate, p.idx, p.row, p.col, p.curchar, msg}
}

func (p *Parser) Parse() (*Movetree, error) {
	if p.mt != nil {
		// We've already performed the parsing.
		return p.mt, p.err
	}
	p.mt = InitMovetree()
	branches := make([]*Node, 0, 1000)

	// buffer := make([]rune, 0, 1000)
	for c, _, err := p.r.ReadRune(); err != nil; {
		p.curchar = c
		p.idx++
		p.col++
		if p.curchar == NEWLINE {
			p.row++
			p.col = 0
			if p.curstate != PROP_DATA {
				continue // White space only matters in property data
			}
		}
		switch p.curstate {
		case BEGINNING:
			if p.curchar == LPAREN {
				branches = append(branches, p.mt.CurrentNode) // Append the root node.
			} else if p.curchar == SCOLON {
				p.curstate = BETWEEN // The SGF Begins!
			} else if isWhitespace(p.curchar) {
				// We can safely ignore whitespace here.
			} else {
				p.ParseError(fmt.Sprintf("Unexpected char"))
			}
			break
		case PROPERTY: // TODO(kashomon): fill out
		case PROP_DATA: // TODO(kashomon): fill out
		case BETWEEN: // TODO(kashomon): fill out
		default: // We shouldn't ever get here
			p.ParseError("Unexpected state -- couldn't match state")
		}
	}
	return p.mt, nil
}
