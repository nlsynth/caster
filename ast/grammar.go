package ast

// ProgramFile -
type ProgramFile struct {
	Modules []*ModuleDecl `parser:"@@*"`
}

// ModuleDecl -
type ModuleDecl struct {
	Name      string         `parser:"'module' @Ident"`
	Interface *InterfaceDecl `parser:"@@"`
	Start     string         `parser:"'{'"`
	Regs      []*RegDecl     `parser:"@@*"`
	Stages    []*StageDecl   `parser:"@@*"`
	End       string         `parser:"'}'"`
}

// InterfaceDecl -
type InterfaceDecl struct {
	Empty0 string        `parser:"'('"`
	Ports  *PortDeclList `parser:"@@?"`
	Empty1 string        `parser:"')'"`
}

// PortDeclList -
type PortDeclList struct {
	Port0 *PortDecl       `parser:"@@"`
	Tail  []*PortDeclTail `parser:"@@*"`
}

// PortDeclTail -
type PortDeclTail struct {
	Empty string    `parser:"','"`
	Port  *PortDecl `parser:"@@"`
}

// PortDecl -
type PortDecl struct {
	Name  string     `parser:"@Ident"`
	Width *WidthDecl `parser:"@@?"`
	Kind  *string    `parser:"@Ident?"`
}

// WidthDecl -
type WidthDecl struct {
	Width int `parser:"'#' @Number"`
}

// RegDecl -
type RegDecl struct {
	Reg   string     `parser:"'reg' @Ident"`
	Tail  []*RegTail `parser:"@@*"`
	Width *WidthDecl `parser:"@@?"`
}

// RegTail -
type RegTail struct {
	Name string `parser:"',' @Ident"`
}

// StageDecl -
type StageDecl struct {
	Initial *string       `parser:"@Initial?"`
	Empty   string        `parser:"'stage'"`
	Name    string        `parser:"@Ident"`
	Stmts   *StmtListNode `parser:"@@"`
}

// StmtListNode -
type StmtListNode struct {
	Stmts []*StmtNode `parser:"'{' @@* '}'"`
}

// StmtNode -
type StmtNode struct {
	Go   *GoStmtNode `parser:"@@ |"`
	Expr *Expr0Node  `parser:"@@ |"`
	If   *IfStmtNode `parser:"@@"`
}

// GoStmtNode -
type GoStmtNode struct {
	Go *string `parser:"'go' @Ident"`
}

// IfStmtNode -
type IfStmtNode struct {
	Cond      *Expr0Node    `parser:"'if' @@"`
	ThenStmts *StmtListNode `parser:"@@"`
	ElseStmts *ElseNode     `parser:"@@?"`
}

// ElseNode -
type ElseNode struct {
	Stmts *StmtListNode `parser:"'else' @@"`
}

// Operator -
type Operator string

const (
	// OpAdd -
	OpAdd Operator = "+"
	// OpSub -
	OpSub Operator = "-"
)

var operatorMap = map[string]Operator{"+": OpAdd, "-": OpSub}

// Capture -
func (o *Operator) Capture(s []string) error {
	*o = operatorMap[s[0]]
	return nil
}

// ExprNode0 := ExprNode1 '<=' ExprNode1
// ExprNode1 := ExprNode9 '+'/'-' ExprNode9
//

// Expr0Node -
type Expr0Node struct {
	LHS *Expr1Node    `parser:"@@"`
	RHS *Expr0RHSNode `parser:"@@?"`
}

// Expr0RHSNode -
type Expr0RHSNode struct {
	Node *Expr1Node `parser:"'<=' @@"`
}

// Expr1Node -
type Expr1Node struct {
	LHS *Expr9Node      `parser:"@@"`
	RHS *[]Expr1RHSNode `parser:"@@*"`
}

// Expr1RHSNode -
type Expr1RHSNode struct {
	Op   Operator   `parser:"@('+' | '-')"`
	Node *Expr9Node `parser:"@@"`
}

// Expr9Node -
type Expr9Node struct {
	NumExpr *int    `parser:"@Number |"`
	StrExpr *string `parser:"@Ident"`
}
