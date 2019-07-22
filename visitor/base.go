package visitor

import (
	"fmt"

	"github.com/kulics/lite-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	Wrap      = "\r\n"
	Terminate = ";"

	Byte = "byte"
	Any  = "interface{}"
	Int  = "int"
	Num  = "float64"
	I8   = "int8"
	I16  = "int16"
	I32  = "int32"
	I64  = "int64"

	U8  = "uint8"
	U16 = "uint16"
	U32 = "uint32"
	U64 = "uint64"

	F32 = "float32"
	F64 = "float64"

	Bool = "bool"
	T    = "true"
	F    = "false"

	Chr = "rune"
	Str = "string"
	Lst = "Lst"
	Set = "Set"
	Dic = "Dic"

	BlockLeft  = "{"
	BlockRight = "}"

	Func  = "func "
	Var   = "var "
	Const = "const "
	Type  = "type "
	Chan  = "chan "
)

type any = interface{}
type str = string

type errorListener struct {
	*antlr.DefaultErrorListener

	file string
	Err  error
}

func NewErrorListener(file string) *errorListener {
	return &errorListener{
		file: file,
	}
}

func (me *errorListener) Error() string {
	if me.Err == nil {
		return ""
	}
	return me.Err.Error()
}

func (me *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	me.Err = fmt.Errorf("[ERR %s:%d:%d] %s", me.file, line, column, msg)
}

type LiteVisitor struct {
	parser.BaseLiteParserVisitor

	AllIDSet     *list_str
	CurrentIDSet *stack_str
}

type Result struct {
	Data       any
	Text       string
	Permission string
	IsVirtual  bool
	IsDefine   bool
}
type list_str struct {
	list []str
}

func (me *list_str) contains(id str) bool {
	return true
}
func (me *list_str) add(id str) {
	me.list = append(me.list, id)
}
func (me *list_str) except_with(t *list_str) {

}

type stack_str struct {
	stack []*list_str
}

func (me *stack_str) peek() *list_str {
	return me.stack[len(me.stack)-1]
}
func (me *stack_str) push(new *list_str) {
	me.stack = append(me.stack, new)
}
func (me *stack_str) pop() {
	me.stack = me.stack[:len(me.stack)-2]
}

func (me *LiteVisitor) has_id(id str) bool {
	return me.AllIDSet.contains(id) || me.CurrentIDSet.peek().contains(id)
}
func (me *LiteVisitor) add_id(id str) {
	me.CurrentIDSet.peek().add(id)
}
func (me *LiteVisitor) add_current_set() {
	for _, item := range me.CurrentIDSet.peek().list {
		me.AllIDSet.add(item)
	}
	me.CurrentIDSet.push(&list_str{})
}
func (me *LiteVisitor) delete_current_set() {
	me.AllIDSet.except_with(me.CurrentIDSet.peek())
	me.CurrentIDSet.pop()
}

func (me *LiteVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(me)
}

func (me *LiteVisitor) VisitChildren(tree antlr.RuleNode) any {
	return tree.Accept(me)
}

// func (me *LiteVisitor) VisitTerminal(tree antlr.TerminalNode) any {
// 	return tree.Accept(me)
// }

// func (me *LiteVisitor) VisitErrorNode(tree antlr.ErrorNode) any {
// 	return tree.Accept(me)
// }

func (me *LiteVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	obj := ""
	for _, item := range ctx.AllStatement() {
		obj += me.Visit(item).(string)
	}
	return obj
}

func (me *LiteVisitor) VisitId(ctx *parser.IdContext) any {
	r := Result{Data: "var"}
	first := me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(Result)
	r.Permission = first.Permission
	r.Text = first.Text
	r.IsVirtual = first.IsVirtual
	if ctx.GetChildCount() >= 2 {
		for i := 1; i < ctx.GetChildCount(); i++ {
			other := me.Visit(ctx.GetChild(i).(antlr.ParseTree)).(Result)
			r.Text += "_" + other.Text
		}
	}
	// todo
	// if keywords.Exists({t -> t == r.Text}) {
	// 	r.Text = "@" + r.Text
	// }
	return r
}

func (me *LiteVisitor) VisitIdItem(ctx *parser.IdItemContext) any {
	r := Result{Data: "var"}
	if ctx.TypeBasic() != nil {
		r.Permission = "public"
		r.Text += ctx.TypeBasic().GetText()
	} else if ctx.TypeAny() != nil {
		r.Permission = "public"
		r.Text += ctx.TypeAny().GetText()
	} else if ctx.LinqKeyword() != nil {
		r.Permission = "public"
		r.Text += me.Visit(ctx.LinqKeyword()).(string)
	} else if ctx.GetOp().GetTokenType() == parser.LiteLexerIDPublic {
		r.Permission = "public"
		r.Text += ctx.GetOp().GetText()
		// r.IsVirtual = r.Text[0].is Upper()
	} else if ctx.GetOp().GetTokenType() == parser.LiteLexerIDPrivate {
		r.Permission = "protected"
		r.Text += ctx.GetOp().GetText()
		// r.IsVirtual = r.Text[r.Text.find first({it -> it >< '_'})].is Upper()
	}
	return r
}

func (me *LiteVisitor) VisitIdExpression(ctx *parser.IdExpressionContext) any {
	r := Result{Data: "var"}
	if len(ctx.AllIdExprItem()) > 1 {
		r.Text = "("
		for i, v := range ctx.AllIdExprItem() {
			subID := me.Visit(v).(Result).Text
			if i != 0 {
				r.Text += ", " + subID
			} else {
				r.Text += subID
			}
			if me.has_id(subID) {
				r.IsDefine = true
			} else {
				me.add_id(subID)
			}
		}
		r.Text += ")"
	} else {
		r = me.Visit(ctx.IdExprItem(0)).(Result)
		if me.has_id(r.Text) {
			r.IsDefine = true
		} else {
			me.add_id(r.Text)
		}
	}
	return r
}

func (me *LiteVisitor) VisitIdExprItem(ctx *parser.IdExprItemContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}
