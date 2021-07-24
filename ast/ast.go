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
	Stmts []*Stmt
}

// Stmt -
type Stmt struct {
	Decl *StmtNode
	Expr *Expr
	Go   *GoStmt
}

// GoStmt -
type GoStmt struct {
	Decl *GoStmtNode
}

// Expr -
type Expr struct {
	LHS *Expr
	RHS *Expr
	Num *int
	Str *string
	Op  string
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

	return p, nil
}
