// Package interpreter implements execution of an AST.
package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/pseidemann/tik/ast"
)

const maxStackSize = 1000

// Interpreter can execute an AST.
type Interpreter struct {
	stdout io.Writer
	stack  contextStack
}

type context struct {
	vars  map[string]*variable
	funcs map[string]*ast.FuncDef
}

type varType int

const (
	varNumber varType = iota
	varString
)

type variable struct {
	varType varType
	intVal  int
	strVal  string
}

func newContext() *context {
	return &context{
		vars:  make(map[string]*variable),
		funcs: make(map[string]*ast.FuncDef),
	}
}

func copyContext(prev *context) *context {
	ctx := newContext()
	for k, v := range prev.vars {
		ctx.vars[k] = v
	}
	for k, v := range prev.funcs {
		ctx.funcs[k] = v
	}
	return ctx
}

// New creates an Interpreter.
func New(stdout io.Writer) *Interpreter {
	in := &Interpreter{
		stdout: stdout,
	}
	in.stack.push(newContext())
	return in
}

func (in *Interpreter) addContext() {
	in.stack.push(copyContext(in.context()))
	if in.stack.size() > maxStackSize {
		panic("max stack size exceeded")
	}
}

func (in *Interpreter) removeContext() {
	in.stack.pop()
}

func (in *Interpreter) context() *context {
	return in.stack.peek()
}

func (in *Interpreter) setVar(name string, variable *variable) {
	in.context().vars[name] = variable
}

func (in *Interpreter) getVar(name string) *variable {
	v, ok := in.context().vars[name]
	if !ok {
		panic(fmt.Sprintf("undefined variable %q", name))
	}
	return v
}

func (in *Interpreter) setFunc(f *ast.FuncDef) {
	in.context().funcs[f.Name] = f
}

func (in *Interpreter) getFunc(name string) *ast.FuncDef {
	f, ok := in.context().funcs[name]
	if !ok {
		panic(fmt.Sprintf("undefined function %q", name))
	}
	return f
}

// Execute interprets the given AST.
func (in *Interpreter) Execute(root ast.Node) {
	in.execAst(root)
}

func (in *Interpreter) execAst(n ast.Node) {
	switch v := n.(type) {
	case *ast.FuncDef:
		in.setFunc(v)
	case *ast.FuncCall:
		in.execFuncCall(v)
	case *ast.Assign:
		in.execAssign(v)
	default:
		for _, child := range n.Children() {
			in.execAst(child)
		}
	}
}

func (in *Interpreter) execFuncCall(funcCall *ast.FuncCall) {
	switch funcCall.Name {
	case "print":
		buf := bufio.NewWriter(in.stdout)
		lastIdx := len(funcCall.Args) - 1
		for i, child := range funcCall.Args {
			var str string
			switch v := child.(type) {
			case *ast.String:
				str = v.Str
			case *ast.Operation:
				num := in.evalExpr(v)
				str = strconv.Itoa(num.intVal)
			case *ast.Ident:
				vari := in.getVar(v.Name)
				switch vari.varType {
				case varNumber:
					str = strconv.Itoa(vari.intVal)
				case varString:
					str = vari.strVal
				default:
					panic("unknown variable type")
				}
			}
			buf.WriteString(str)
			if i < lastIdx {
				buf.WriteRune(' ')
			}
		}
		buf.WriteRune('\n')
		buf.Flush()
	default:
		f := in.getFunc(funcCall.Name)
		in.addContext()
		if len(funcCall.Args) != len(f.Params) {
			panic("number of defined args and passed args don't match")
		}
		for i, arg := range funcCall.Args {
			name := f.Params[i].Name
			in.setVar(name, in.evalExpr(arg))
		}
		in.execAst(f.Body)
		in.removeContext()
	}
}

func (in *Interpreter) evalExpr(n ast.Node) *variable {
	switch v := n.(type) {
	case *ast.Operation:
		return in.evalOp(v)
	case *ast.Number:
		n, err := strconv.Atoi(v.Num)
		if err != nil {
			panic(err)
		}
		return &variable{varType: varNumber, intVal: n}
	case *ast.Ident:
		return in.getVar(v.Name)
	case *ast.String:
		return &variable{varType: varString, strVal: v.Str}
	default:
		panic(fmt.Sprintf("unknown expression %v", n))
	}
}

func (in *Interpreter) evalOp(op *ast.Operation) *variable {
	switch op.OpType {
	case ast.OpAdd:
		v := in.evalExpr(op.Left).intVal + in.evalExpr(op.Right).intVal
		return &variable{varType: varNumber, intVal: v}
	case ast.OpSub:
		v := in.evalExpr(op.Left).intVal - in.evalExpr(op.Right).intVal
		return &variable{varType: varNumber, intVal: v}
	case ast.OpMul:
		v := in.evalExpr(op.Left).intVal * in.evalExpr(op.Right).intVal
		return &variable{varType: varNumber, intVal: v}
	case ast.OpDiv:
		v := in.evalExpr(op.Left).intVal / in.evalExpr(op.Right).intVal
		return &variable{varType: varNumber, intVal: v}
	default:
		panic(fmt.Sprintf("unknown operation %v", op))
	}
}

func (in *Interpreter) execAssign(n *ast.Assign) {
	ident, ok := n.Left.(*ast.Ident)
	if !ok {
		panic("expected identifier on left side of assignment")
	}
	in.setVar(ident.Name, in.evalExpr(n.Right))
}
