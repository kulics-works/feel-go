package visitor

import parser "github.com/kulics-works/k-go/parser/generate"

func (me *KVisitor) VisitTypeConversion(ctx *parser.TypeConversionContext) any {
	return me.Visit(ctx.TypeType()).(string)
}

func (me *KVisitor) VisitCall(ctx *parser.CallContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitWave(ctx *parser.WaveContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitJudge(ctx *parser.JudgeContext) any {
	if ctx.GetOp().GetText() == "><" {
		return "!="
	} else if ctx.GetOp().GetText() == "&" {
		return "&&"
	} else if ctx.GetOp().GetText() == "|" {
		return "||"
	}
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitAdd(ctx *parser.AddContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitMul(ctx *parser.MulContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitPow(ctx *parser.PowContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitPlusMinus(ctx *parser.PlusMinusContext) any {
	r := Result{}
	expr := me.Visit(ctx.Expression()).(Result)
	op := me.Visit(ctx.Add()).(string)
	r.Data = expr.Data
	r.Text = op + expr.Text
	return r
}

func (me *KVisitor) VisitNegate(ctx *parser.NegateContext) any {
	r := Result{}
	expr := me.Visit(ctx.Expression()).(Result)
	r.Data = expr.Data
	r.Text = "!" + expr.Text
	return r
}
