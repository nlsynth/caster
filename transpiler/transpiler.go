package transpiler

import (
	"caster/ast"
	"fmt"
	"io"
)

// ModuleTranspiler -
type ModuleTranspiler struct {
	m *Module
}

type Transpiler struct {
	mt []*ModuleTranspiler
}

// NewTranspiler creates a Transpiler.
func NewTranspiler() *Transpiler {
	t := new(Transpiler)
	t.mt = make([]*ModuleTranspiler, 0)
	return t
}

// TranspileProgram processes a file.
func (t *Transpiler) TranspileProgram(p *ast.Program) {
	for _, am := range p.Modules {
		m := newModule(am)
		mt := newModuleTranspiler(m)
		t.mt = append(t.mt, mt)
	}
}

// Output outputs transpiled code.
func (t *Transpiler) Output(w io.Writer) {
	i := newIndenter(w)
	for _, mt := range t.mt {
		mt.output(i)
	}
}

func newModuleTranspiler(m *Module) *ModuleTranspiler {
	mt := new(ModuleTranspiler)
	mt.m = m
	return mt
}

func (mt *ModuleTranspiler) output(w io.Writer) {
	fmt.Fprintf(w, "module %v(\n", mt.m.getName())
	mt.writePorts(w)
	fmt.Fprint(w, ");\n\n")
	mt.declareStageVariables(w)
	fmt.Fprint(w, " always @(posedge clk) begin\n")
	fmt.Fprint(w, " if (rst) begin\n")
	mt.resetStageVariables(w)
	fmt.Fprint(w, " end else begin\n")
	for _, s := range mt.m.stages {
		mt.writeStage(s, w)
	}
	fmt.Fprint(w, " end\n")
	fmt.Fprint(w, " end\n")
	fmt.Fprint(w, "endmodule\n")
}

func (mt *ModuleTranspiler) declareStageVariables(w io.Writer) {
	for _, s := range mt.m.stages {
		fmt.Fprintf(w, " reg %s;\n", s.stateVariableName())
	}
	fmt.Fprintf(w, "\n")
}

func (mt *ModuleTranspiler) resetStageVariables(w io.Writer) {
	for _, s := range mt.m.stages {
		fmt.Fprintf(w, "  %s <= ", s.stateVariableName())
		if s.isInitial() {
			fmt.Fprintf(w, "1")
		} else {
			fmt.Fprintf(w, "0")
		}
		fmt.Fprint(w, ";\n")
	}
}

func (mt *ModuleTranspiler) writePorts(w io.Writer) {
	mt.writeBasicPorts(w)
	isFirst := true
	for _, p := range mt.m.ports {
		if !isFirst {
			fmt.Fprintf(w, ",\n")
		}
		fmt.Fprintf(w, " ")
		mt.writePort(p, w)
		isFirst = false
	}
}

func (mt *ModuleTranspiler) writeBasicPorts(w io.Writer) {
	fmt.Fprintf(w, " input clk,\n")
	fmt.Fprintf(w, " input rst,\n")
}

func (mt *ModuleTranspiler) writePort(p *Port, w io.Writer) {
	fmt.Fprintf(w, "%v", p.kind)
	if p.width > 0 {
		fmt.Fprintf(w, " [%d:0]", p.width-1)
	}
	fmt.Fprintf(w, " %v", p.name)
}

func (mt *ModuleTranspiler) writeStage(s *Stage, w io.Writer) {
	fmt.Fprintf(w, "  // stage %v\n", s.as.GetName())
	fmt.Fprintf(w, "  if (%s) begin\n", s.stateVariableName())
	fmt.Fprintf(w, "   %s <= 0;\n", s.stateVariableName())
	for _, st := range s.as.Stmts.Stmts {
		stp := newStmtTranspiler(mt, st)
		stp.write(w)
	}
	fmt.Fprintf(w, "  end\n")
}

func (mt *ModuleTranspiler) getStageFromAst(as *ast.Stage) *Stage {
	// linear search for now.
	for _, s := range mt.m.stages {
		if s.as == as {
			return s
		}
	}
	return nil
}
