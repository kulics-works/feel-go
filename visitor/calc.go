package visitor

import "github.com/kulics/lite-go/parser"

func (me *LiteVisitor) VisitTypeConversion(ctx *parser.TypeConversionContext) any {
	return me.Visit(ctx.TypeType()).(string)
}

func (me *LiteVisitor) VisitCall(ctx *parser.CallContext) any {
	return ctx.GetOp().GetText()
}

func (me *LiteVisitor) VisitWave(ctx *parser.WaveContext) any {
	return ctx.GetOp().GetText()
}

func (me *LiteVisitor) VisitJudgeType(ctx *parser.JudgeTypeContext) any {
	return ctx.GetOp().GetText()
}

func (me *LiteVisitor) VisitJudge(ctx *parser.JudgeContext) any {
	if ctx.GetOp().GetText() == "><" {
		return "!="
	} else if ctx.GetOp().GetText() == "&" {
		return "&&"
	} else if ctx.GetOp().GetText() == "|" {
		return "||"
	}
	return ctx.GetOp().GetText()
}

func (me *LiteVisitor) VisitAdd(ctx *parser.AddContext) any {
	return ctx.GetOp().GetText()
}

func (me *LiteVisitor) VisitMul(ctx *parser.MulContext) any {
	return ctx.GetOp().GetText()
}

func (me *LiteVisitor) VisitPow(ctx *parser.PowContext) any {
	return ctx.GetOp().GetText()
}

func (me *LiteVisitor) VisitPlusMinus(ctx *parser.PlusMinusContext) any {
	r := Result{}
	expr := me.Visit(ctx.Expression()).(Result)
	op := me.Visit(ctx.Add()).(string)
	r.Data = expr.Data
	r.Text = op + expr.Text
	return r
}

func (me *LiteVisitor) VisitNegate(ctx *parser.NegateContext) any {
	r := Result{}
	expr := me.Visit(ctx.Expression()).(Result)
	r.Data = expr.Data
	r.Text = "!" + expr.Text
	return r
}
