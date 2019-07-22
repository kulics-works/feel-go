package visitor

import (
	"github.com/kulics/lite-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *LiteVisitor) VisitCallExpression(ctx *parser.CallExpressionContext) any {
	r := me.Visit(ctx.Id()).(Result)
	r.Text = "." + r.Text
	if ctx.CallFunc() != nil {
		e2 := me.Visit(ctx.CallFunc()).(Result)
		r.Text = r.Text + e2.Text
	} else if ctx.CallElement() != nil {
		e2 := me.Visit(ctx.CallElement()).(Result)
		r.Text = r.Text + e2.Text
	} else if ctx.CallChannel() != nil {
		e2 := me.Visit(ctx.CallChannel()).(Result)
		r.Text = e2.Text + r.Text
	}
	return r
}

func (me *LiteVisitor) VisitCallChannel(ctx *parser.CallChannelContext) any {
	r := Result{}
	r.Text = "<-"
	return r
}

func (me *LiteVisitor) VisitCallElement(ctx *parser.CallElementContext) any {
	if ctx.Expression() == nil {
		return Result{Text: me.Visit(ctx.Slice()).(string)}
	}
	r := me.Visit(ctx.Expression()).(Result)
	r.Text = "[" + r.Text + "]"
	return r
}

func (me *LiteVisitor) VisitSlice(ctx *parser.SliceContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
}

func (me *LiteVisitor) VisitSliceFull(ctx *parser.SliceFullContext) any {
	attach := ""
	switch ctx.GetOp().GetText() {
	case "<=":
		attach = "+1"
	}
	expr1 := me.Visit(ctx.Expression(0)).(Result)
	expr2 := me.Visit(ctx.Expression(1)).(Result)
	return "[" + expr1.Text + ":" + expr2.Text + attach + "]"
}

func (me *LiteVisitor) VisitSliceStart(ctx *parser.SliceStartContext) any {
	attach := ""
	expr := me.Visit(ctx.Expression()).(Result)
	return "[" + expr.Text + ":" + attach + "]"
}

func (me *LiteVisitor) VisitSliceEnd(ctx *parser.SliceEndContext) any {
	attach := ""
	switch ctx.GetOp().GetText() {
	case "<=":
		attach = "+1"
	}
	expr := me.Visit(ctx.Expression()).(Result)
	return "[:" + expr.Text + attach + "]"
}

func (me *LiteVisitor) VisitCallFunc(ctx *parser.CallFuncContext) any {
	r := Result{Data: "var"}
	if ctx.Tuple() != nil {
		r.Text += "(" + me.Visit(ctx.Tuple()).(Result).Text + ")"
	} else {
		r.Text += "(" + me.Visit(ctx.Lambda()).(Result).Text + ")"
	}
	return r
}

func (me *LiteVisitor) VisitCallNew(ctx *parser.CallNewContext) any {
	r := Result{Data: me.Visit(ctx.TypeType())}
	param := ""
	if ctx.ExpressionList() != nil {
		param = me.Visit(ctx.ExpressionList()).(Result).Text
	}
	r.Text = "make(" + me.Visit(ctx.TypeType()).(string)
	if param != "" {
		r.Text += "," + param
	}
	r.Text += ")"
	return r
}

func (me *LiteVisitor) VisitCallPkg(ctx *parser.CallPkgContext) any {
	r := Result{Data: me.Visit(ctx.TypeType())}
	r.Text = me.Visit(ctx.TypeType()).(string)
	if ctx.PkgAssign() != nil {
		r.Text += me.Visit(ctx.PkgAssign()).(string)
	} else if ctx.ListAssign() != nil {
		r.Text += me.Visit(ctx.ListAssign()).(string)
	} else if ctx.SetAssign() != nil {
		r.Text += me.Visit(ctx.SetAssign()).(string)
	} else if ctx.DictionaryAssign() != nil {
		r.Text += me.Visit(ctx.DictionaryAssign()).(string)
	} else {
		r.Text += "{}"
	}
	return r
}

func (me *LiteVisitor) VisitPkgAssign(ctx *parser.PkgAssignContext) any {
	obj := ""
	obj += "{"
	for i := 0; i < len(ctx.AllPkgAssignElement()); i++ {
		if i == 0 {
			obj += me.Visit(ctx.PkgAssignElement(i)).(string)
		} else {
			obj += "," + me.Visit(ctx.PkgAssignElement(i)).(string)
		}
	}
	obj += "}"
	return obj
}

func (me *LiteVisitor) VisitListAssign(ctx *parser.ListAssignContext) any {
	obj := ""
	obj += "{"
	for i := 0; i < len(ctx.AllExpression()); i++ {
		r := me.Visit(ctx.Expression(i)).(Result)
		if i == 0 {
			obj += r.Text
		} else {
			obj += "," + r.Text
		}
	}
	obj += "}"
	return obj
}

func (me *LiteVisitor) VisitSetAssign(ctx *parser.SetAssignContext) any {
	obj := ""
	obj += "{"
	for i := 0; i < len(ctx.AllExpression()); i++ {
		r := me.Visit(ctx.Expression(i)).(Result)
		if i == 0 {
			obj += r.Text
		} else {
			obj += "," + r.Text
		}
	}
	obj += "}"
	return obj
}

func (me *LiteVisitor) VisitDictionaryAssign(ctx *parser.DictionaryAssignContext) any {
	obj := ""
	obj += "{"
	for i := 0; i < len(ctx.AllDictionaryElement()); i++ {
		r := me.Visit(ctx.DictionaryElement(i)).(DicEle)
		if i == 0 {
			obj += r.Text
		} else {
			obj += "," + r.Text
		}
	}
	obj += "}"
	return obj
}

func (me *LiteVisitor) VisitPkgAssignElement(ctx *parser.PkgAssignElementContext) any {
	obj := ""
	obj += me.Visit(ctx.Name()).(string) + " = " + me.Visit(ctx.Expression()).(Result).Text
	return obj
}

func (me *LiteVisitor) VisitDictionaryElement(ctx *parser.DictionaryElementContext) any {
	r1 := me.Visit(ctx.Expression(0)).(Result)
	r2 := me.Visit(ctx.Expression(1)).(Result)
	r := DicEle{
		Key:   r1.Data.(string),
		Value: r2.Data.(string),
		Text:  r1.Text + ":" + r2.Text,
	}
	return r
}

type DicEle struct {
	Key   string
	Value string
	Text  string
}
