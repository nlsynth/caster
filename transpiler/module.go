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
		m.addPort(p.Decl)
	}
	if len(m.ports) == 0 {
		return
	}
	prev := m.ports[len(m.ports)-1]
	for i := len(m.ports) - 1; i >= 0; i-- {
		p := m.ports[i]
		if p.kind == "" {
			p.kind = prev.kind
			p.width = prev.width
		}
		if p.width < 0 {
			p.width = prev.width
		}
		prev = p
	}
}

func (m *Module) addPort(decl *ast.PortDecl) {
	if decl == nil {
		return
	}
	p := &Port{}
	p.name = decl.Name
	if decl.Kind != nil {
		p.kind = *decl.Kind
	}
	if decl.Width == nil {
		p.width = -1
	} else {
		p.width = decl.Width.Width
	}
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
