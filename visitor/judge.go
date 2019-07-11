package visitor

import (
	"github.com/kulics/lite-go/parser"
)

func (sf *LiteVisitor) VisitJudgeCaseStatement(ctx *parser.JudgeCaseStatementContext) interface{} {
	obj := ""
	expr := sf.Visit(ctx.Expression()).(Result)
	obj += "switch " + expr.Text + BlockLeft + Wrap
	for _, item := range ctx.AllCaseStatement() {
		r := sf.Visit(item).(string)
		obj += r + Wrap
	}
	obj += BlockRight + Wrap
	return obj
}

func (sf *LiteVisitor) VisitCaseExprStatement(ctx *parser.CaseExprStatementContext) interface{} {
	obj := ""
	if ctx.Expression() != nil {
		expr := sf.Visit(ctx.Expression()).(Result)
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

func (sf *LiteVisitor) VisitCaseStatement(ctx *parser.CaseStatementContext) interface{} {
	obj := ""
	process := BlockLeft + sf.ProcessFunctionSupport(ctx.AllFunctionSupportStatement()) + BlockRight
	for _, item := range ctx.AllCaseExprStatement() {
		r := sf.Visit(item).(string)
		obj += r + process + Wrap
	}
	return obj
}

func (sf *LiteVisitor) VisitJudgeStatement(ctx *parser.JudgeStatementContext) interface{} {
	obj := ""
	obj += sf.Visit(ctx.JudgeIfStatement()).(string)
	for _, it := range ctx.AllJudgeElseIfStatement() {
		obj += sf.Visit(it).(string)
	}
	if ctx.JudgeElseStatement() != nil {
		obj += sf.Visit(ctx.JudgeElseStatement()).(string)
	}
	obj += Wrap
	return obj
}

func (sf *LiteVisitor) VisitJudgeIfStatement(ctx *parser.JudgeIfStatementContext) interface{} {
	b := sf.Visit(ctx.Expression()).(Result)
	obj := "if " + b.Text + BlockLeft + Wrap
	obj += sf.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	return obj
}

func (sf *LiteVisitor) VisitJudgeElseIfStatement(ctx *parser.JudgeElseIfStatementContext) interface{} {
	b := sf.Visit(ctx.Expression()).(Result)
	obj := "else if " + b.Text + BlockLeft + Wrap
	obj += sf.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	return obj
}

func (sf *LiteVisitor) VisitJudgeElseStatement(ctx *parser.JudgeElseStatementContext) interface{} {
	obj := "else " + BlockLeft + Wrap
	obj += sf.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	return obj
}
