package core

import (
	"fmt"
	"strings"
	"unicode"
)

// Special chars , used to delimite sections of the SGF.
const (
	LPAREN  rune = '('
	RPAREN  rune = ')'
	LBRACE  rune = '['
	RBRACE  rune = ']'
	SCOLON  rune = ';'
	NEWLINE rune = '\n'
)

// Used to denote states during parsing.
type State int

const (
	// In BEGINNING, we have not yet begun parsing the SGF.
	BEGINNING State = iota

	// In PROPERTY, we are looking for data to complete a property token, such as
	// AW or B.  A property is considered complete when we see a left brace '['.
	PROPERTY

	// In PROP_DATA, we are looking for all data associated with a property
	// token. We are finished with the Property data when we see a right brace ']'
	PROP_DATA

	// In BETWEEN, we are not accumulating data, we are just trying to figure out
	// where to go next.  Thus, we could find a new property, we could find more
	// property data, or we colud find a new variation.
	BETWEEN
)

type ParseError struct {
	s       State
	i, r, c int
	char    rune
	msg     string
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("Error during state %v, at index %v, line %v, column %v"+
		", curchar %v. %v", p.s, p.i, p.r, p.c, string(p.char), p.msg)
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
	return &Parser{r, 0, 0, 0, '〜', BEGINNING, nil, nil}
}

func (p *Parser) ParseError(msg string) *ParseError {
	return &ParseError{p.curstate, p.idx, p.row, p.col, p.curchar, msg}
}

type TrackingData struct {
	branches []*Node
	charbuf  []rune
	propdata []string
	curprop  SgfProperty
}

func NewTrackingData() *TrackingData {
	return &TrackingData{
		make([]*Node, 0, 100),
		make([]rune, 0, 100),
		make([]string, 0, 100),
		SgfProperty(""),
	}
}

func (t *TrackingData) FlushBuffer() string {
	out := string(t.charbuf)
	t.charbuf = make([]rune, 0, 100)
	return out
}

func (t *TrackingData) FlushPropData(mt *Movetree) {
	if len(t.curprop) > 0 {
		mt.Node().Props().AddTo(t.curprop, t.propdata...)
		t.propdata = make([]string, 0, 100)
		t.curprop = SgfProperty("")
	}
}

func (p *Parser) Parse() (*Movetree, error) {
	if p.mt != nil {
		// We've already performed the parsing.
		return p.mt, p.err
	}
	p.mt = EmptyMovetree()
	t := NewTrackingData()

	// buffer := make([]rune, 0, 1000)
	c, _, err := p.r.ReadRune()
	p.curchar = c
	for err == nil && p.r.Len() >= 0 {
		p.idx++
		p.col++
		if p.curchar == NEWLINE {
			p.row++
			p.col = 0
			if p.curstate != PROP_DATA {
				p.curchar, _, err = p.r.ReadRune()
				continue // White space only matters in property data
			} else {
				// We will apply this character to the propdata
			}
		}
		switch p.curstate {
		case BEGINNING:
			if p.curchar == LPAREN {
				t.branches = append(t.branches, p.mt.Root())
			} else if p.curchar == SCOLON {
				p.curstate = BETWEEN // The SGF Begins!
			} else if unicode.IsSpace(p.curchar) {
				// We can safely ignore whitespace here.
			} else {
				return nil, p.ParseError(fmt.Sprintf("Unexpected char"))
			}
			break
		case PROPERTY:
			if unicode.IsUpper(p.curchar) {
				t.charbuf = append(t.charbuf, p.curchar)
			} else if p.curchar == LBRACE {
				testprop := t.FlushBuffer()
				if !IsValidProperty(testprop) {
					return nil, p.ParseError("Unknown property: " + string(t.curprop))
				}
				t.curprop = SgfProperty(testprop)
				p.curstate = PROP_DATA
			} else if unicode.IsSpace(p.curchar) {
				// Should space be allowed?
				return nil, p.ParseError("Unexpected whitespace in Property")
			} else {
				return nil, p.ParseError("Unexpected character")
			}
			break
		case PROP_DATA:
			if p.curchar == RBRACE &&
				len(t.charbuf) > 0 &&
				t.charbuf[len(t.charbuf)-1] == '\\' {
				t.charbuf = append(t.charbuf, p.curchar)
			} else if p.curchar == RBRACE {
				t.propdata = append(t.propdata, t.FlushBuffer())
				p.curstate = BETWEEN
			} else {
				t.charbuf = append(t.charbuf, p.curchar)
			}
			break
		case BETWEEN:
			if unicode.IsUpper(p.curchar) {
				// A new property
				t.FlushPropData(p.mt)
				t.charbuf = append(t.charbuf, p.curchar)
				p.curstate = PROPERTY
			} else if p.curchar == LBRACE {
				// More property data
				p.curstate = PROP_DATA // more data to process
			} else if p.curchar == LPAREN {
				// A New Variation
				t.FlushPropData(p.mt)
				t.branches = append(t.branches, p.mt.Node())
			} else if p.curchar == RPAREN {
				// Finish a variation
				t.FlushPropData(p.mt)
				parent := t.branches[len(t.branches)-1]
				t.branches = t.branches[:len(t.branches)-1]
				for n := p.mt.Node(); n != parent; {
					n = p.mt.MoveUp().Node()
				}
			} else if p.curchar == SCOLON {
				// Finish a node, if necessary
				t.FlushPropData(p.mt)
				p.mt.NewNode()
			} else if unicode.IsSpace(p.curchar) {
				// Do nothing.  Whitespace can be ignored here.
			} else {
				return nil, p.ParseError("Unknown token")
			}
			break
		default: // We shouldn't ever get here
			return nil, p.ParseError("Unexpected state -- couldn't match state")
		}
		p.curchar, _, err = p.r.ReadRune()
	}
	return setMovetreeDefaults(p.mt.FromRoot()), err
}
