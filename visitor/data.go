package visitor

import (
	"github.com/kulics/lite-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *LiteVisitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) any {
	r := me.Visit(ctx.Expression()).(Result)
	return r.Text + Wrap
}

func (me *LiteVisitor) VisitExpression(ctx *parser.ExpressionContext) any {
	count := ctx.GetChildCount()
	r := Result{}
	if count == 3 {
		e1 := me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(Result)
		e2 := me.Visit(ctx.GetChild(2).(antlr.ParseTree))
		op := me.Visit(ctx.GetChild(1).(antlr.ParseTree))

		switch ctx.GetChild(1).(type) {
		case parser.JudgeTypeContext:
			// r.data = Bool
			// e3 := me.Visit(ctx.GetChild(2)).(string)
			// if op == "==" {
			// 	r.text = "(" + e1.text + " is " + e3+")"
			// } else if op == "><" {
			// 	r.text = "!(" + e1.text +" is "+e3+")"
			// }
			return r
		case parser.JudgeContext:
			// todo 如果左右不是bool类型值，报错
			r.Data = Bool
		case parser.AddContext:
			// todo 如果左右不是number或text类型值，报错
			if e1.Data.(string) == Str || e2.(Result).Data.(string) == Str {
				r.Data = Str
			} else if e1.Data.(string) == I32 && e2.(Result).Data.(string) == I32 {
				r.Data = I32
			} else {
				r.Data = F64
			}
		case parser.MulContext:
			// todo 如果左右不是number类型值，报错
			if e1.Data.(string) == I32 && e2.(Result).Data.(string) == I32 {
				r.Data = I32
			} else {
				r.Data = F64
			}
		case parser.PowContext:
			// todo 如果左右部署number类型，报错
			r.Data = F64
			switch op {
			case "**":
				op = "Pow"
			case "//":
				op = "Root"
			case "%%":
				op = "Log"
			}
			r.Text = op.(string) + "(" + e1.Text + ", " + e2.(Result).Text + ")"
			return r
		}
		r.Text = e1.Text + op.(string) + e2.(Result).Text
	} else if count == 2 {
		r = me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(Result)
		if _, ok := ctx.GetChild(1).(*parser.TypeConversionContext); ok {
			e2 := me.Visit(ctx.GetChild(1).(antlr.ParseTree)).(string)
			r.Data = e2
			r.Text = e2 + "(" + r.Text + ")"
		} else if _, ok := ctx.GetChild(1).(*parser.CallExpressionContext); ok {
			e2 := me.Visit(ctx.GetChild(1).(antlr.ParseTree)).(Result)
			r.Text = r.Text + e2.Text
		} else if _, ok := ctx.GetChild(1).(*parser.CallFuncContext); ok {
			e2 := me.Visit(ctx.GetChild(1).(antlr.ParseTree)).(Result)
			r.Text = r.Text + e2.Text
		} else if _, ok := ctx.GetChild(1).(*parser.CallElementContext); ok {
			e2 := me.Visit(ctx.GetChild(1).(antlr.ParseTree)).(Result)
			r.Text = r.Text + e2.Text
		} else if _, ok := ctx.GetChild(1).(*parser.CallChannelContext); ok {
			e2 := me.Visit(ctx.GetChild(1).(antlr.ParseTree)).(Result)
			r.Text = e2.Text + r.Text
		} else if ctx.GetOp() != nil {
			if ctx.GetOp().GetTokenType() == parser.LiteLexerBang {
				r.Text = "*" + r.Text
			} else if ctx.GetOp().GetTokenType() == parser.LiteLexerQuestion {
				r.Text = "&" + r.Text
			} else if ctx.GetOp().GetTokenType() == parser.LiteLexerLeft_Flow {
				r.Text = "go " + r.Text
			}
		}
	} else if count == 1 {
		r = me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(Result)
	}
	return r
}

func (me *LiteVisitor) VisitVariableStatement(ctx *parser.VariableStatementContext) any {
	obj := ""
	r1 := me.Visit(ctx.IdExpression()).(Result)
	r2 := me.Visit(ctx.Expression()).(Result)
	Type := ""
	if ctx.TypeType() != nil {
		Type = me.Visit(ctx.TypeType()).(string)
		obj = Var + r1.Text + " " + Type + " = " + r2.Text + Wrap
	} else {
		if r1.IsDefine || r1.Text == self.Id {
			obj = r1.Text + " = " + r2.Text + Wrap
		} else {
			obj = r1.Text + " := " + r2.Text + Wrap
		}
	}
	return obj
}

func (me *LiteVisitor) VisitVariableDeclaredStatement(ctx *parser.VariableDeclaredStatementContext) any {
	obj := ""
	Type := me.Visit(ctx.TypeType()).(string)
	r := me.Visit(ctx.IdExpression()).(Result)
	obj = Var + r.Text + " " + Type + Wrap
	return obj
}

func (me *LiteVisitor) VisitChannelAssignStatement(ctx *parser.ChannelAssignStatementContext) any {
	r1 := me.Visit(ctx.Expression(0)).(Result)
	r2 := me.Visit(ctx.Expression(1)).(Result)
	obj := r1.Text + "<-" + r2.Text + Wrap
	return obj
}

func (me *LiteVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) any {
	r1 := me.Visit(ctx.TupleExpression(0)).(Result)
	r2 := me.Visit(ctx.TupleExpression(1)).(Result)
	obj := r1.Text + me.Visit(ctx.Assign()).(string) + r2.Text + Wrap
	return obj
}

func (me *LiteVisitor) VisitAssign(ctx *parser.AssignContext) any {
	return ctx.GetOp().GetText()
}

func (me *LiteVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) any {
	if ctx.GetChildCount() == 1 {
		c := ctx.GetChild(0)
		if _, ok := c.(*parser.DataStatementContext); ok {
			return me.Visit(ctx.DataStatement())
		} else if _, ok := c.(*parser.IdContext); ok {
			return me.Visit(ctx.Id())
		} else if ctx.GetT().GetTokenType() == parser.LiteLexerDot_Dot {
			return Result{Text: "me", Data: "var"}
		} else if ctx.GetT().GetTokenType() == parser.LiteLexerDiscard {
			return Result{Text: "_", Data: "var"}
		}
	} else if ctx.GetChildCount() == 2 {
		id := me.Visit(ctx.Id()).(Result)
		template := me.Visit(ctx.TemplateCall()).(string)
		return Result{Text: id.Text + template, Data: id.Text + template}
	}
	r := me.Visit(ctx.Expression()).(Result)
	return Result{Text: "(" + r.Text + ")", Data: r.Data}
}

func (me *LiteVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) any {
	r := Result{}
	obj := ""
	for i := 0; i < len(ctx.AllExpression()); i++ {
		temp := me.Visit(ctx.Expression(i)).(Result)
		if i == 0 {
			obj += temp.Text
		} else {
			obj += ", " + temp.Text
		}
	}
	r.Text = obj
	r.Data = "var"
	return r
}

func (me *LiteVisitor) VisitStringExpression(ctx *parser.StringExpressionContext) any {
	text := "bytes.Buffer{}.WriteString(" + ctx.TextLiteral().GetText() + ")"
	for _, item := range ctx.AllStringExpressionElement() {
		text += me.Visit(item).(string)
	}
	text += ".String()"
	return Result{
		Data: Str,
		Text: text,
	}
}

func (me *LiteVisitor) VisitStringExpressionElement(ctx *parser.StringExpressionElementContext) any {
	r := me.Visit(ctx.Expression()).(Result)
	text := ctx.TextLiteral().GetText()
	return ".WriteString(fmt.Sprint(" + r.Text + ").WriteString(" + text + ")"
}

func (me *LiteVisitor) VisitDataStatement(ctx *parser.DataStatementContext) any {
	r := Result{}
	if ctx.NilExpr() != nil {
		r.Data = Any
		r.Text = "nil"
	} else if ctx.FloatExpr() != nil {
		r.Data = Num
		r.Text = me.Visit(ctx.FloatExpr()).(string)
	} else if ctx.IntegerExpr() != nil {
		r.Data = Int
		r.Text = me.Visit(ctx.IntegerExpr()).(string)
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTextLiteral {
		r.Data = Str
		r.Text = ctx.TextLiteral().GetText()
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerCharLiteral {
		r.Data = Chr
		r.Text = ctx.CharLiteral().GetText()
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTrueLiteral {
		r.Data = Bool
		r.Text = T
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerFalseLiteral {
		r.Data = Bool
		r.Text = F
	}
	return r
}

func (me *LiteVisitor) VisitFloatExpr(ctx *parser.FloatExprContext) any {
	number := ""
	number += me.Visit(ctx.IntegerExpr(0)).(string) + "." + me.Visit(ctx.IntegerExpr(1)).(string)
	return number
}

func (me *LiteVisitor) VisitIntegerExpr(ctx *parser.IntegerExprContext) any {
	number := ""
	number += ctx.NumberLiteral().GetText()
	return number
}

func (me *LiteVisitor) VisitTupleExpression(ctx *parser.TupleExpressionContext) any {
	obj := ""
	for i := 0; i < len(ctx.AllExpression()); i++ {
		r := me.Visit(ctx.Expression(i)).(Result)
		if i == 0 {
			obj += r.Text
		} else {
			obj += ", " + r.Text
		}
	}
	obj += ""
	result := Result{Data: "var", Text: obj}
	return result
}
