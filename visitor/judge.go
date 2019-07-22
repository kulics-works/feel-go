package visitor

import (
	"github.com/kulics/lite-go/parser"
)

func (me *LiteVisitor) VisitJudgeCaseStatement(ctx *parser.JudgeCaseStatementContext) any {
	obj := ""
	expr := me.Visit(ctx.Expression()).(Result)
	obj += "switch " + expr.Text + BlockLeft + Wrap
	for _, item := range ctx.AllCaseStatement() {
		r := me.Visit(item).(string)
		obj += r + Wrap
	}
	obj += BlockRight + Wrap
	return obj
}

func (me *LiteVisitor) VisitCaseExprStatement(ctx *parser.CaseExprStatementContext) any {
	obj := ""
	if ctx.Expression() != nil {
		expr := me.Visit(ctx.Expression()).(Result)
		obj += "case " + expr.Text + " :" + Wrap
	} else if ctx.TypeType() != nil {
		// id := "it"
		// ? ctx.id() >< () {
		// 	id = Visit(ctx.id()).(Result).text
		// }
		// type := Visit(ctx.typeType()):Str
		// obj += "case "type" "id" :"Wrap""
	} else {
		obj += "default:" + Wrap
	}
	return obj
}

func (me *LiteVisitor) VisitCaseStatement(ctx *parser.CaseStatementContext) any {
	obj := ""
	process := BlockLeft + me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement()) + BlockRight
	for _, item := range ctx.AllCaseExprStatement() {
		r := me.Visit(item).(string)
		obj += r + process + Wrap
	}
	return obj
}

func (me *LiteVisitor) VisitJudgeStatement(ctx *parser.JudgeStatementContext) any {
	obj := ""
	obj += me.Visit(ctx.JudgeIfStatement()).(string)
	for _, it := range ctx.AllJudgeElseIfStatement() {
		obj += me.Visit(it).(string)
	}
	if ctx.JudgeElseStatement() != nil {
		obj += me.Visit(ctx.JudgeElseStatement()).(string)
	}
	obj += Wrap
	return obj
}

func (me *LiteVisitor) VisitJudgeIfStatement(ctx *parser.JudgeIfStatementContext) any {
	b := me.Visit(ctx.Expression()).(Result)
	obj := "if " + b.Text + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	return obj
}

func (me *LiteVisitor) VisitJudgeElseIfStatement(ctx *parser.JudgeElseIfStatementContext) any {
	b := me.Visit(ctx.Expression()).(Result)
	obj := "else if " + b.Text + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	return obj
}

func (me *LiteVisitor) VisitJudgeElseStatement(ctx *parser.JudgeElseStatementContext) any {
	obj := "else " + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	return obj
}
