package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/kulics-works/k-go/parser/generate"
)

func (me *KVisitor) VisitImplementStatement(ctx *parser.ImplementStatementContext) any {
	obj := ""
	var methed = ""
	id := me.Visit(ctx.Id()).(Result)
	me.self.Type = id.Text

	for _, item := range ctx.AllPackageFieldStatement() {
		var r = me.Visit(item).(Result)
		obj += r.Text
		methed += r.Data.(string)
	}
	for _, item := range ctx.AllPackageImplementStatement() {
		methed += me.Visit(item).(Result).Text
	}
	obj += methed
	me.self = Parameter{}
	return obj
}

func (me *KVisitor) VisitImplementSupportStatement(ctx *parser.ImplementSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *KVisitor) VisitImplementFunctionStatement(ctx *parser.ImplementFunctionStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	obj := ""
	me.add_current_set()
	obj += Func + "(" + me.self.Id + " " + me.self.Type + ")" +
		id.Text + me.Visit(ctx.ParameterClauseIn()).(string) +
		me.Visit(ctx.ParameterClauseOut()).(string) + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	me.delete_current_set()
	return obj
}
