package visitor

import "github.com/kulics/lite-go/parser"

type Iterator struct {
	Begin  Result
	End    Result
	Step   Result
	Order  bool
	Attach bool
}

func (me *LiteVisitor) VisitIteratorStatement(ctx *parser.IteratorStatementContext) any {
	it := Iterator{Order: true, Attach: false}
	if ctx.GetOp().GetText() == ">=" || ctx.GetOp().GetText() == "<=" {
		it.Attach = true
	}
	if ctx.GetOp().GetText() == ">" || ctx.GetOp().GetText() == ">=" {
		it.Order = false
	}
	if len(ctx.AllExpression()) == 2 {
		it.Begin = me.Visit(ctx.Expression(0)).(Result)
		it.End = me.Visit(ctx.Expression(1)).(Result)
		it.Step = Result{Data: I32, Text: "1"}
	} else {
		it.Begin = me.Visit(ctx.Expression(0)).(Result)
		it.End = me.Visit(ctx.Expression(1)).(Result)
		it.Step = me.Visit(ctx.Expression(2)).(Result)
	}
	return it
}

func (me *LiteVisitor) VisitLoopStatement(ctx *parser.LoopStatementContext) any {
	obj := ""
	id := "ea"
	if ctx.Id() != nil {
		id = me.Visit(ctx.Id()).(Result).Text
	}
	it := me.Visit(ctx.IteratorStatement()).(Iterator)
	order := ""
	step := ""
	if it.Order {
		step = "+="
		if it.Attach {
			order = "<="
		} else {
			order = "<"
		}
	} else {
		step = "-="
		if it.Attach {
			order = ">="
		} else {
			order = ">"
		}
	}
	obj += "for " + id + " := " + it.Begin.Text + ";" + id + order + it.End.Text + ";" + id + step + it.Step.Text
	me.add_current_set()
	obj += BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	me.delete_current_set()
	return obj
}

func (me *LiteVisitor) VisitLoopInfiniteStatement(ctx *parser.LoopInfiniteStatementContext) any {
	me.add_current_set()
	obj := "for " + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	me.delete_current_set()
	return obj
}

func (me *LiteVisitor) VisitLoopEachStatement(ctx *parser.LoopEachStatementContext) any {
	obj := ""
	arr := me.Visit(ctx.Expression()).(Result)
	target := arr.Text
	id := "ea"
	if len(ctx.AllId()) == 2 {
		id = me.Visit(ctx.Id(0)).(Result).Text + "," + me.Visit(ctx.Id(1)).(Result).Text
	} else if len(ctx.AllId()) == 1 {
		id = "_," + me.Visit(ctx.Id(0)).(Result).Text
	}
	me.add_current_set()
	obj += "for " + id + " := range " + target
	obj += BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	me.delete_current_set()
	return obj
}

func (me *LiteVisitor) VisitLoopCaseStatement(ctx *parser.LoopCaseStatementContext) any {
	obj := ""
	expr := me.Visit(ctx.Expression()).(Result)
	me.add_current_set()
	obj += "for " + expr.Text
	obj += BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	me.delete_current_set()
	return obj
}

func (me *LiteVisitor) VisitLoopJumpStatement(ctx *parser.LoopJumpStatementContext) any {
	return "break" + Wrap
}

func (me *LiteVisitor) VisitLoopContinueStatement(ctx *parser.LoopContinueStatementContext) any {
	return "continue" + Wrap
}
