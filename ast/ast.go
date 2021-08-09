package ast

import "os"

// Program -
type Program struct {
	Modules []*Module
}

// Module -
type Module struct {
	Decl   *ModuleDecl
	Stages []*Stage
}

// Stage -
type Stage struct {
	Decl  *StageDecl
	Stmts StmtList
}

// StmtList -
type StmtList struct {
	Stmts []*Stmt
}

// Stmt -
type Stmt struct {
	Decl *StmtNode
	Expr *Expr
	Go   *GoStmt
	If   *IfStmt
}

// GoStmt -
type GoStmt struct {
	Stage *Stage
	Decl  *GoStmtNode
}

// Expr -
type Expr struct {
	LHS *Expr
	RHS *Expr
	Num *int
	Str *string
	Op  string
}

// IfStmt -
type IfStmt struct {
	Decl *IfStmtNode
}

// GetName -
func (st *Stage) GetName() string {
	return st.Decl.Name
}

// IsInitial -
func (st *Stage) IsInitial() bool {
	return st.Decl.Initial != nil
}

// GetProgram -
func GetProgram(fn string, file *os.File) (*Program, error) {
	pf := ProgramFile{}
	err := getParser().Parse(fn, file, &pf)
	if err != nil {
		return nil, err
	}

	cv := newConverter(&pf)
	p := cv.convert()
	rs := newResolver(p)
	rs.resolve()

	return p, nil
}
