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
	t := transpiler.NewTranspiler()
	flag.Parse()
	for _, fn := range flag.Args() {
		glog.Infof("fn=%v", fn)
		r, err := os.Open(fn)
		if err != nil {
			glog.Errorf("failed to open %v %v", fn, err)
			continue
		}
		processFile(fn, r, t)
	}
	t.Output(os.Stdout)
}
