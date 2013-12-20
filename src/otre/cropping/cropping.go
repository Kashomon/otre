package core

import (
	"otre/point"
)

// Convenience enum for specifying a cropping direction.
type CroppingPreset int

const (
	// X -
	// - -
	TOP_LEFT CroppingPreset = iota

	// - X
	// - -
	TOP_RIGHT

	// - -
	// - X
	BOTTOM_RIGHT
	
	// - -
	// X -
	BOTTOM_LEFT

	// X X
	// - -
	TOP

	// X -
	// X -
	LEFT

	// - X
	// - X
	RIGHT

	// - -
	// X X
	BOTTOM

	// X X
	// X X
	ALL
)

// The cropping is a bounding box, specified by intersection points.
type Cropping struct {
	TopLeft  *point.IntPt
	BotRight *point.IntPt
}

// Create a cropping box from the maxInts. Note that the integer points in the
// crop box are 0 indexed, but maxInts is 1-indexed.  In other words, we would
// typically expect the max ints to range from 9 to 19.
// 
// Following the SGF covention, we consider the topleft to be 0,0
func FromPreset(p CroppingPreset, maxInts int) *Cropping {
	halfInts := maxInts / 2
	minInts := 0
	top := minInts
	left := minInts
	bot := maxInts
	right := maxInts
	switch(p) {
	case ALL: // nothing to change 
	case LEFT: 
		right = halfInts + 1
	case RIGHT:
		left = halfInts - 1
	case TOP:
		bot = halfInts + 1
	case BOTTOM:
		top = halfInts - 1 
	case TOP_LEFT:
		bot = halfInts + 1
		right = halfInts + 2
	case TOP_RIGHT:
		bot = halfInts + 1
		left = halfInts - 2
	case BOTTOM_LEFT:
		top = halfInts - 1
		right = halfInts + 2
	case BOTTOM_RIGHT:
		top = halfInts - 1
		left = halfInts - 2
	}
	return &Cropping{
		&point.IntPt{top, left},
		&point.IntPt{bot, right},
	}
}
