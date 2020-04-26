package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/kulics-works/k-go/parser/generate"
)

func (me *KVisitor) VisitTypeConversion(ctx *parser.TypeConversionContext) any {
	return me.Visit(ctx.TypeType()).(string)
}

func (me *KVisitor) VisitCall(ctx *parser.CallContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitWave(ctx *parser.WaveContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitAdd(ctx *parser.AddContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitMul(ctx *parser.MulContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitPow(ctx *parser.PowContext) any {
	return "^"
}

func (me *KVisitor) VisitTransfer(ctx *parser.TransferContext) any {
	return "<-"
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

func (me *KVisitor) VisitBitwise(ctx *parser.BitwiseContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(str)
}

func (me *KVisitor) VisitBitwiseAnd(ctx *parser.BitwiseAndContext) any {
	return "&"
}

func (me *KVisitor) VisitBitwiseOr(ctx *parser.BitwiseOrContext) any {
	return "|"
}

func (me *KVisitor) VisitBitwiseXor(ctx *parser.BitwiseXorContext) any {
	return "^"
}

func (me *KVisitor) VisitBitwiseLeftShift(ctx *parser.BitwiseLeftShiftContext) any {
	return "<<"
}

func (me *KVisitor) VisitBitwiseRightShift(ctx *parser.BitwiseRightShiftContext) any {
	return ">>"
}

func (me *KVisitor) VisitCompare(ctx *parser.CompareContext) any {
	if ctx.GetOp().GetTokenType() == parser.KParserNot_Equal {
		return "!="
	}
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitLogic(ctx *parser.LogicContext) any {
	if ctx.GetOp().GetTokenType() == parser.KParserAnd {
		return "&&"
	} else if ctx.GetOp().GetTokenType() == parser.KParserOr {
		return "||"
	}
	return ctx.GetOp().GetText()
}
