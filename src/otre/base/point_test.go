package base

import (
	"testing"
)

func TestSgfCoordHappyPath(t *testing.T) {
	c, err := FromSgfCoord("ca")
	if err != nil {
		t.Errorf("%v", err)
	}
	if c.X != 2 {
		t.Errorf("Expected 2, saw %v", c.X)
	}
	if c.Y != 0 {
		t.Errorf("Expected 0, saw %v", c.Y)
	}
}
