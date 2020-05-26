package visitor

import (
	parser "github.com/kulics-works/feel-go/parser/generate"
)

func (me *KVisitor) VisitJudgeCaseStatement(ctx *parser.JudgeCaseStatementContext) any {
	obj := ""
	expr := me.Visit(ctx.Expression()).(Result)
	obj += "switch " + expr.Text + BlockLeft + Wrap
	for _, item := range ctx.AllCaseStatement() {
		r := me.Visit(item).(string)
		obj += r + Wrap
	}
	if item := ctx.CaseElseStatement(); item != nil {
		obj += me.Visit(item).(str)
	}
	obj += BlockRight + Wrap
	return obj
}

func (me *KVisitor) VisitJudgeCase(ctx *parser.JudgeCaseContext) any {
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
	}
	return obj
}

func (me *KVisitor) VisitCaseStatement(ctx *parser.CaseStatementContext) any {
	obj := ""
	for _, item := range ctx.AllJudgeCase() {
		r := me.Visit(item).(string)
		me.add_current_set()
		process := BlockLeft + me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement()) + BlockRight
		obj += r + process + Wrap
		me.delete_current_set()
	}
	return obj
}

func (me *KVisitor) VisitCaseElseStatement(ctx *parser.CaseElseStatementContext) any {
	obj := ""

	me.add_current_set()
	process := BlockLeft + me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement()) + BlockRight
	obj += "default:" + process + Wrap
	me.delete_current_set()

	return obj
}

func (me *KVisitor) VisitJudgeStatement(ctx *parser.JudgeStatementContext) any {
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

func (me *KVisitor) VisitJudgeIfStatement(ctx *parser.JudgeIfStatementContext) any {
	b := me.Visit(ctx.Expression()).(Result)
	me.add_current_set()
	obj := "if " + b.Text + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	me.delete_current_set()
	return obj
}

func (me *KVisitor) VisitJudgeElseIfStatement(ctx *parser.JudgeElseIfStatementContext) any {
	b := me.Visit(ctx.Expression()).(Result)
	obj := "else if " + b.Text + BlockLeft + Wrap
	me.add_current_set()
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	me.delete_current_set()
	return obj
}

func (me *KVisitor) VisitJudgeElseStatement(ctx *parser.JudgeElseStatementContext) any {
	me.add_current_set()
	obj := "else " + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	me.delete_current_set()
	return obj
}
