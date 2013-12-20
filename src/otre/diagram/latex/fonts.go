package latex

import (
	"fmt"
	"strings"
)

type GoChar int

const (
	// Lines
	TL_CORNER GoChar = iota
	TR_CORNER
	BL_CORNER
	BR_CORNER
	TOP_EDGE
	BOT_EDGE
	LEFT_EDGE
	RIGHT_EDGE
	CENTER

	// Starpoint
	STARPOINT

	// Stones
	BSTONE
	WSTONE

	// Marks and StoneMarks
	BSTONE_TRIANGLE
	WSTONE_TRIANGLE
	TRIANGLE
	BSTONE_SQUARE
	WSTONE_SQUARE
	SQUARE
	BSTONE_CIRCLE
	WSTONE_CIRCLE
	CIRCLE
	BSTONE_XMARK
	WSTONE_XMARK
	XMARK

	// BigBoard Numbering
	// Example
	BSTONE_LABEL_BIG
	WSTONE_LABEL_BIG
	EMPTY_LABEL_BIG

	// SmallBoard Numbering
	BSTONE_LABEL
	WSTONE_LABEL
	EMPTY_LABEL

	// Formatting
	BSTONE_INLINE
	WSTONE_INLINE
	MISC_STONE_INLINE
)

// Font families
const (
	ComputerModernSansSerif = "cmss"
	ComputerModern = "cmr"
	ComputerModernFixedwidth = "cmtt"
)

type GoFont struct {
	// Font definition.  See GooemacsMap
	FontMap  map[GoChar]string
	LabelFont string
	Packages string
	Defs     string
	BoardDef string
}

func GooeFont() *GoFont {
	lf := ComputerModernSansSerif
	fontDef := fmt.Sprintf(gooeFontDefsBase, lf, lf, lf, lf)
	defs := strings.Join([]string{
		fontDef,
		gooeSizeDefs,
		gooeBigBoardDefs,
		gooeNormalBoardDefs,
		gooeInlineDefs}, "\n")
	return &GoFont{
		gooemacsMap,
		lf,
		gooePackages,
		defs,
		gooeBoardDef,
	}
}
