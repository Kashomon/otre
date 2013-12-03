package core

import (
	"testing"
)

func tpSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBadTpParse(t *testing.T) {
	tp, err := ParseTreepath("aoeu")
	if err == nil {
		t.Errorf("Expected parsing error")
	}
	if !tpSlicesEqual(tp.Vars, make([]int, 0)) {
		t.Errorf("Expected no paths")
	}
}

func TestSimpleTpParse(t *testing.T) {
	tp, err := ParseTreepath("1")
	if err != nil {
		t.Errorf("Expected no parsing error")
	}
	if !tpSlicesEqual(tp.Vars, []int{0}) {
		t.Errorf("Expected 1 path")
	}
}

func TestTpParseVars(t *testing.T) {
	tp, err := ParseTreepath("3.1")
	if err != nil {
		t.Errorf("Expected no parsing error")
	}
	if !tpSlicesEqual(tp.Vars, []int{0, 0, 0, 1}) {
		t.Errorf("Unexpected path %v", tp.Vars)
	}
}

func TestTpParseVarsV2(t *testing.T) {
	tp, err := ParseTreepath("2.2.1")
	if err != nil {
		t.Errorf("Expected no parsing error")
	}
	if !tpSlicesEqual(tp.Vars, []int{0, 0, 2, 1}) {
		t.Errorf("Unexpected path %v", tp.Vars)
	}
}

func TestTpFromBeginning(t *testing.T) {
	tp, err := ParseTreepath("0.2.1")
	if err != nil {
		t.Errorf("Expected no parsing error")
	}
	if !tpSlicesEqual(tp.Vars, []int{2, 1}) {
		t.Errorf("Unexpected path %v", tp.Vars)
	}
}

func TestTpFromBeginningV2(t *testing.T) {
	tp, err := ParseTreepath("0.0.0.0")
	if err != nil {
		t.Errorf("Expected no parsing error")
	}
	if !tpSlicesEqual(tp.Vars, []int{0, 0, 0}) {
		t.Errorf("Unexpected path %v", tp.Vars)
	}
}

func TestTpComplexParse(t *testing.T) {
	tp, err := ParseTreepath("0.2.1-5.3")
	if err != nil {
		t.Errorf("Expected no parsing error")
	}
	if l := len(tp.Vars); l != 6 {
		t.Errorf("Expected to be on the 6th move, but on %v", l)
	}
	if !tpSlicesEqual(tp.Vars, []int{2, 1, 0, 0, 0, 3}) {
		t.Errorf("Unexpected path %v", tp.Vars)
	}
}

func TestTpLongParse(t *testing.T) {
	tp, err := ParseTreepath("53")
	if err != nil {
		t.Errorf("Expected no parsing error")
	}
	if l := len(tp.Vars); l != 53 {
		t.Errorf("Expected to be on the 53rd move, but on %v", l)
	}
}
