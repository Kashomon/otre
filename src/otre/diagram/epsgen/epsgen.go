package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"otre/core"
	"otre/diagram/eps"
	"strings"
)

// /ct {
// /Helvetica findfont 12 scalefont setfont setgray moveto 
// dup stringwidth pop -2 div 0 rmoveto
// show } def
// (A) 240 281.33 0 ct

// --Flags---
var startpath = flag.String("startpath", "", "Starting position in the SGF."+
	"This is a Treepath, so see treepath.go for more details")
var in = flag.String("in", "", "Path to input file (must be an SGF)")
var out = flag.String("out", "", "Path to output file (default: in.eps)")
var numbering = flag.String("numbering", "b+10", "How to number variations. "+
	"Examples:\n"+
	"    v+10: start of variation +10 (at most) moves.\n"+
	"    b+10: beginning +10 (at most) moves.\n"+
	"    v+n: start of variation +n (at most) moves.\n"+
	"    b+n: beginning +n (at most) moves.")

func main() {
	flag.Parse()
	if *in == "" || !strings.Contains(*in, ".sgf") {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *out == "" {
		*out = strings.Replace(*in, ".sgf", ".eps", 1)
	}
	inf, inerr := os.Open(*in)
	if inerr != nil {
		panic(inerr)
	}
	defer inf.Close()
	bts, err := ioutil.ReadAll(inf)
	if err != nil {
		panic(err)
	}
	mt, perr := core.FromSgfString(string(bts)).Parse()
	if perr != nil && perr != io.EOF {
		panic(perr)
	}

	fo, ferr := os.Create(*out)
	if ferr != nil {
		panic(ferr)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	fo.Write(eps.Generate(mt, 300))
}
