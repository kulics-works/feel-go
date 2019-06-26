package visitor

import "github.com/kulics/lite-go/parser"

func (sf *LiteVisitor) VisitTypeConversion(ctx *parser.TypeConversionContext) interface{} {
	return sf.Visit(ctx.TypeType()).(string)
}

func (sf *LiteVisitor) VisitCall(ctx *parser.CallContext) interface{} {
	return ctx.GetOp().GetText()
}

func (sf *LiteVisitor) VisitWave(ctx *parser.WaveContext) interface{} {
	return ctx.GetOp().GetText()
}

func (sf *LiteVisitor) VisitJudgeType(ctx *parser.JudgeTypeContext) interface{} {
	return ctx.GetOp().GetText()
}

func (sf *LiteVisitor) VisitJudge(ctx *parser.JudgeContext) interface{} {
	if ctx.GetOp().GetText() == "><" {
		return "!="
	} else if ctx.GetOp().GetText() == "&" {
		return "&&"
	} else if ctx.GetOp().GetText() == "|" {
		return "||"
	}
	return ctx.GetOp().GetText()
}

func (sf *LiteVisitor) VisitAdd(ctx *parser.AddContext) interface{} {
	return ctx.GetOp().GetText()
}

func (sf *LiteVisitor) VisitMul(ctx *parser.MulContext) interface{} {
	return ctx.GetOp().GetText()
}

func (sf *LiteVisitor) VisitPow(ctx *parser.PowContext) interface{} {
	return ctx.GetOp().GetText()
}

func (sf *LiteVisitor) VisitPlusMinus(ctx *parser.PlusMinusContext) interface{} {
	r := Result{}
	expr := sf.Visit(ctx.Expression()).(Result)
	op := sf.Visit(ctx.Add()).(string)
	r.Data = expr.Data
	r.Text = op + expr.Text
	return r
}

func (sf *LiteVisitor) VisitNegate(ctx *parser.NegateContext) interface{} {
	r := Result{}
	expr := sf.Visit(ctx.Expression()).(Result)
	r.Data = expr.Data
	r.Text = "!" + expr.Text
	return r
}
