package ast

// Resolver -
type Resolver struct {
	p *Program
}

type moduleResolver struct {
	m           *Module
	nameToStage map[string]*Stage
}

func newResolver(p *Program) *Resolver {
	rs := new(Resolver)
	rs.p = p
	return rs
}

func (rs *Resolver) resolve() {
	for _, m := range rs.p.Modules {
		mr := newModuleResolver(m)
		mr.resolve()
	}
}

func newModuleResolver(m *Module) *moduleResolver {
	mr := new(moduleResolver)
	mr.m = m
	mr.nameToStage = make(map[string]*Stage)
	return mr
}

func (mr *moduleResolver) resolve() {
	for _, s := range mr.m.Stages {
		mr.nameToStage[s.GetName()] = s
	}
	for _, s := range mr.m.Stages {
		mr.resolveStage(s)
	}
}

func (mr *moduleResolver) resolveStage(stg *Stage) {
	for _, s := range stg.Stmts {
		mr.resolveStmt(s)
	}
}

func (mr *moduleResolver) resolveStmt(stmt *Stmt) {
	if stmt.Go != nil {
		stmt.Go.Stage = mr.nameToStage[*stmt.Go.Decl.Go]
	}
}
