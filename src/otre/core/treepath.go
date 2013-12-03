package core

import (
	"strconv"
	"strings"
)

// A Treepath dictates how to traverse a movetree to get to a position.  In
// other words, it is a list of variation numbers.
//
// Examples:
// 	[] // The 0th move (The root node)
// 	[0,0,0] // The 3rd move (4th node)
type Treepath struct {
	Vars []int
}

// Because specifying all the list of variations is tedious, users can supply a
// string shorthand to indicate the treepath. The structure looks like:
// 	MoveNumber.VariationNumber.VariationNumber-MoveNumber
//
// Examples:
// 	Str   -> Treepath
// 	0     -> []
// 	1     -> [0]
// 	0.1   -> [1]
// 	53    -> [0,0,0...(53 times]
// 	3.1   -> [0,0,0,1]
// 	2.0.1 -> [0.0,0,1]
func ParseTreepath(tp string) (*Treepath, error) {
	out := make([]int, 0, 10)
	if len(tp) == 0 {
		return &Treepath{out}, nil
	}
	s := strings.Split(tp, "-")
	idx := 0
	for i := 0; i < len(s); i++ {
		subs := strings.Split(s[i], ".")
		for j := 0; j < len(subs); j++ {
			num, err := strconv.Atoi(subs[j])
			if err != nil {
				return &Treepath{make([]int, 0)}, err
			}
			if j == 0 {
				for idx < num {
					out = append(out, 0)
					idx++
				}
			} else {
				out = append(out, num)
				idx++
			}
		}
	}
	return &Treepath{out}, nil
}
