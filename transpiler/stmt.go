package transpiler

import (
	"caster/ast"
	"fmt"
	"io"
)

type stmtTranspiler struct {
	as *ast.Stmt
}

func newStmtTranspiler(as *ast.Stmt) *stmtTranspiler {
	stp := new(stmtTranspiler)
	stp.as = as
	return stp
}

func (stp *stmtTranspiler) write(w io.Writer) {
	fmt.Fprint(w, "   ")
	if stp.as.Expr != nil {
		stp.writeExpr(stp.as.Expr, w)
	}
	if stp.as.Go != nil {
		stp.writeGo(stp.as.Go, w)
	}
	fmt.Fprint(w, ";\n")
}

func (stp *stmtTranspiler) writeExpr(e *ast.Expr, w io.Writer) {
	if e.LHS != nil {
		stp.writeExpr(e.LHS, w)
		fmt.Fprintf(w, " %s ", e.Op)
		stp.writeExpr(e.RHS, w)
	}
	if e.Num != nil {
		fmt.Fprintf(w, "%d", *e.Num)
	}
	if e.Str != nil {
		fmt.Fprintf(w, "%s", *e.Str)
	}
}

func (stp *stmtTranspiler) writeGo(g *ast.GoStmt, w io.Writer) {
	fmt.Fprintf(w, "_s_%s <= 1", *g.Decl.Go)
}
