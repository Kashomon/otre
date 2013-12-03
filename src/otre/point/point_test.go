package point

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

func TestReflect(t *testing.T) {
	pt := &CoordPt{50, 50}
	n1 := pt.ReflectY(150)
	if pt.X != n1.X {
		t.Errorf("Expected x vals to be equal")
	}
	if n1.Y != 250 {
		t.Errorf("Expected Y reflection around 150")
	}
	n2 := pt.ReflectX(150)
	if pt.Y != n2.Y {
		t.Errorf("Expected no change in Y vals")
	}
	if n2.X != 250 {
		t.Errorf("Expected X reflection around 150")
	}
}
