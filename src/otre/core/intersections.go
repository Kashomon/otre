package core

type Intersection struct {
	Type  IntersectionType
	Label string
}

type IntersectionType int

const (
	TL_CORNER IntersectionType = iota
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

	// Labeling
	BSTONE_LABEL
	WSTONE_LABEL
	EMPTY_LABEL
)
