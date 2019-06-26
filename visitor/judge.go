package visitor

import (
	"github.com/kulics/lite-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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

func (sf *LiteVisitor) VisitCaseDefaultStatement(ctx *parser.CaseDefaultStatementContext) interface{} {
	obj := ""
	obj += "default:" + BlockLeft + Wrap
	obj += sf.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight
	return obj
}

func (sf *LiteVisitor) VisitCaseExprStatement(ctx *parser.CaseExprStatementContext) interface{} {
	obj := ""
	if ctx.TypeType() == nil {
		expr := sf.Visit(ctx.Expression()).(Result)
		obj += "case " + expr.Text + " :" + Wrap
	} else {
		// id := "it"
		// ? ctx.id() >< () {
		// 	id = Visit(ctx.id()).(Result).text
		// }
		// type := Visit(ctx.typeType()):Str
		// obj += "case "type" "id" :"Wrap""
	}

	obj += BlockLeft + sf.ProcessFunctionSupport(ctx.AllFunctionSupportStatement()) + BlockRight
	return obj
}

func (sf *LiteVisitor) VisitCaseStatement(ctx *parser.CaseStatementContext) interface{} {
	obj := sf.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
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
