package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kulics/lite-go/parser/generate"
)

func (me *LiteVisitor) VisitImplementStatement(ctx *parser.ImplementStatementContext) any {
	obj := ""
	var methed = ""
	id := me.Visit(ctx.Id(0)).(Result)
	me.self.Type = id.Text
	if ctx.Id(1) != nil {
		me.self.Id = me.Visit(ctx.Id(1)).(Result).Text
	}
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

func (me *LiteVisitor) VisitImplementSupportStatement(ctx *parser.ImplementSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *LiteVisitor) VisitImplementFunctionStatement(ctx *parser.ImplementFunctionStatementContext) any {
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
