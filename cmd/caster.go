package main

import (
	"flag"
	"os"

	"caster/ast"
	"caster/transpiler"

	"github.com/golang/glog"
)

func processFile(fn string, file *os.File, t *transpiler.Transpiler) {
	p, err := ast.GetProgram(fn, file)
	if err != nil {
		glog.Errorf("failed to parse %v %v", fn, err)
		return
	}
	t.TranspileProgram(p)
}

func main() {
	ofn := flag.String("o", "", "output file name")
	flag.Parse()
	t := transpiler.NewTranspiler()
	for _, fn := range flag.Args() {
		glog.Infof("fn=%v", fn)
		r, err := os.Open(fn)
		if err != nil {
			glog.Errorf("failed to open %v %v", fn, err)
			continue
		}
		processFile(fn, r, t)
	}
	ow := os.Stdout
	if *ofn != "" {
		var err error
		ow, err = os.Create(*ofn)
		if err != nil {
			glog.Errorf("Failed to open [%v]", *ofn)
			return
		}
		defer ow.Close()
	}
	t.Output(ow)
}
