package transpiler

import "caster/ast"

// Module -
type Module struct {
	am     *ast.Module
	ports  []*Port
	stages []*Stage
}

// Port -
type Port struct {
	name string
	// "input", "output" and so on.
	kind  string
	width int
}

// Stage -
type Stage struct {
	as *ast.Stage
}

func (st *Stage) stateVariableName() string {
	return "_s_" + st.as.GetName()
}

func (st *Stage) isInitial() bool {
	return st.as.IsInitial()
}

func newModule(am *ast.Module) *Module {
	m := &Module{}
	m.am = am
	m.ports = make([]*Port, 0)
	m.preparePorts()
	m.stages = make([]*Stage, 0)
	m.prepareStages()
	return m
}

func (m *Module) preparePorts() {
	for _, p := range m.am.Ports {
		m.addPort(p)
	}
}

func (m *Module) addPort(ap *ast.Port) {
	p := &Port{}
	p.name = ap.Name
	p.kind = ap.Kind
	p.width = ap.Width
	m.ports = append(m.ports, p)
}

func (m *Module) getName() string {
	return m.am.Decl.Name
}

func (m *Module) prepareStages() {
	for _, s := range m.am.Stages {
		m.addStage(s)
	}
}

func (m *Module) addStage(as *ast.Stage) {
	s := new(Stage)
	s.as = as
	m.stages = append(m.stages, s)
}
