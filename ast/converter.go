package ast

import "fmt"

// Converter converts parse tree to AST.
type Converter struct {
	pf *ProgramFile
}

func newConverter(pf *ProgramFile) *Converter {
	cv := new(Converter)
	cv.pf = pf
	return cv
}

func (cv *Converter) convert() *Program {
	p := Program{}
	p.Modules = make([]*Module, len(cv.pf.Modules))
	for i, d := range cv.pf.Modules {
		m := cv.convertModule(d)
		p.Modules[i] = m
	}
	return &p
}

func (cv *Converter) convertModule(md *ModuleDecl) *Module {
	m := new(Module)
	m.Decl = md
	m.Stages = make([]*Stage, len(md.Stages))
	for i, d := range md.Stages {
		s := cv.convertStage(d)
		m.Stages[i] = s
	}
	return m
}

func (cv *Converter) convertStage(sd *StageDecl) *Stage {
	st := new(Stage)
	st.Decl = sd
	st.Stmts = *cv.convertStmtList(sd.Stmts)
	return st
}

func (cv *Converter) convertStmtList(sln *StmtListNode) *StmtList {
	sl := new(StmtList)
	sl.Stmts = make([]*Stmt, len(sln.Stmts))
	for i, n := range sln.Stmts {
		s := cv.convertStmt(n)
		sl.Stmts[i] = s
	}
	return sl
}

func (cv *Converter) convertStmt(sn *StmtNode) *Stmt {
	s := new(Stmt)
	s.Decl = sn
	if s.Decl.Expr != nil {
		s.Expr = cv.convertExpr0(s.Decl.Expr)
	}
	if s.Decl.Go != nil {
		s.Go = cv.convertGo(s.Decl.Go)
	}
	if s.Decl.If != nil {
		s.If = cv.convertIf(s.Decl.If)
	}
	return s
}

func (cv *Converter) convertGo(gn *GoStmtNode) *GoStmt {
	gs := new(GoStmt)
	gs.Decl = gn
	return gs
}

func (cv *Converter) convertIf(in *IfStmtNode) *IfStmt {
	is := new(IfStmt)
	is.Decl = in
	is.Cond = cv.convertExpr0(in.Cond)
	is.True = cv.convertStmtList(in.Stmts)
	is.False = nil
	return is
}

func (cv *Converter) convertExpr0(e0 *Expr0Node) *Expr {
	if e0.RHS == nil {
		return cv.convertExpr1(e0.LHS)
	}
	e := new(Expr)
	e.Op = "<="
	e.LHS = cv.convertExpr1(e0.LHS)
	e.RHS = cv.convertExpr1(e0.RHS.Node)
	return e
}

func (cv *Converter) convertExpr1(e1 *Expr1Node) *Expr {
	if e1.RHS == nil {
		return cv.convertExpr9(e1.LHS)
	}
	var e *Expr
	rhs := *e1.RHS
	for i := 0; i < len(rhs); i++ {
		if e == nil {
			e = new(Expr)
			e.LHS = cv.convertExpr9(e1.LHS)
		} else {
			t := new(Expr)
			t.LHS = e
			e = t
		}
		e.Op = fmt.Sprintf("%s", rhs[i].Op)
		e.RHS = cv.convertExpr9(rhs[i].Node)
	}
	return e
}

func (cv *Converter) convertExpr9(e9 *Expr9Node) *Expr {
	e := new(Expr)
	e.Num = e9.NumExpr
	e.Str = e9.StrExpr
	return e
}
