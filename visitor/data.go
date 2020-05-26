package visitor

import (
	parser "github.com/kulics-works/feel-go/parser/generate"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *KVisitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) any {
	r := me.Visit(ctx.Expression()).(Result)
	return r.Text + Wrap
}

func (me *KVisitor) VisitExpression(ctx *parser.ExpressionContext) any {
	count := ctx.GetChildCount()
	r := Result{}
	if count == 3 {
		e1 := me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(Result)
		e2 := me.Visit(ctx.GetChild(2).(antlr.ParseTree))
		op := me.Visit(ctx.GetChild(1).(antlr.ParseTree))

		switch ctx.GetChild(1).(type) {
		case parser.TransferContext:
			op = "<-"
			r.Data = e1.Data
		case parser.CompareContext:
			// todo 如果左右不是bool类型值，报错
			r.Data = Bool
		case parser.LogicContext:
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
		} else if _, ok := ctx.GetChild(1).(*parser.CallAsyncContext); ok {
			e2 := me.Visit(ctx.GetChild(1).(antlr.ParseTree)).(Result)
			r.Text = "go " + r.Text + e2.Text
		} else if ctx.GetOp() != nil {
			if ctx.GetOp().GetTokenType() == parser.FeelLexerBang {
				r.Text = "*" + r.Text
			} else if ctx.GetOp().GetTokenType() == parser.FeelLexerQuestion {
				r.Text = "&" + r.Text
			}
		}
	} else if count == 1 {
		r = me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(Result)
	}
	return r
}

func (me *KVisitor) VisitVarStatement(ctx *parser.VarStatementContext) any {
	var obj = ""
	for i, v := range ctx.AllVarId() {
		if i != 0 {
			obj += "," + me.Visit(v).(str)
		} else {
			obj += me.Visit(v).(str)
		}
	}
	var r2 = me.Visit(ctx.TupleExpression()).(Result)
	obj += " := " + r2.Text + Wrap
	return obj
}

func (me *KVisitor) VisitBindStatement(ctx *parser.BindStatementContext) any {
	var obj = ""
	for i, v := range ctx.AllConstId() {
		if i != 0 {
			obj += "," + me.Visit(v).(str)
		} else {
			obj += me.Visit(v).(str)
		}
	}
	var r2 = me.Visit(ctx.TupleExpression()).(Result)
	obj += " := " + r2.Text + Wrap
	return obj
}

func (me *KVisitor) VisitVariableDeclaredStatement(ctx *parser.VariableDeclaredStatementContext) any {
	obj := ""
	Type := me.Visit(ctx.TypeType()).(string)
	r := me.Visit(ctx.Id()).(Result)
	obj = Var + r.Text + " " + Type + Wrap
	return obj
}

func (me *KVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) any {
	r1 := me.Visit(ctx.TupleExpression(0)).(Result)
	r2 := me.Visit(ctx.TupleExpression(1)).(Result)
	obj := r1.Text + me.Visit(ctx.Assign()).(string) + r2.Text + Wrap
	return obj
}

func (me *KVisitor) VisitAssign(ctx *parser.AssignContext) any {
	return ctx.GetOp().GetText()
}

func (me *KVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) any {
	if ctx.GetChildCount() == 1 {
		c := ctx.GetChild(0)
		if _, ok := c.(*parser.DataStatementContext); ok {
			return me.Visit(ctx.DataStatement())
		} else if _, ok := c.(*parser.IdContext); ok {
			return me.Visit(ctx.Id())
		} else if ctx.GetT().GetTokenType() == parser.FeelLexerDiscard {
			return Result{Text: "_", Data: "var"}
		}
	} else if ctx.GetChildCount() == 2 {
		id := me.Visit(ctx.Id()).(Result)
		template := me.Visit(ctx.TemplateCall()).(string)
		return Result{Text: id.Text + template, Data: id.Text + template}
	} else if ctx.GetChildCount() == 4 {
		id := me.Visit(ctx.Id()).(Result)
		var tmpArr = me.Visit(ctx.TemplateCall()).([]str)
		var lastId = ""
		switch id.Text {
		case "list":
			lastId = "[]" + tmpArr[0]
		case "dict":
			lastId = "map[" + tmpArr[0] + "]" + tmpArr[1]
		case "chan":
			lastId = "chan " + tmpArr[0]
		}
		return Result{
			Text: lastId,
			Data: lastId,
		}
	}
	r := me.Visit(ctx.Expression()).(Result)
	return Result{Text: "(" + r.Text + ")", Data: r.Data}
}

func (me *KVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) any {
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
func (me *KVisitor) VisitStringExpr(ctx *parser.StringExprContext) any {
	var text = ""
	if len(ctx.AllStringTemplate()) == 0 {
		for _, v := range ctx.AllStringContent() {
			text += me.Visit(v).(str)
		}
		return "\"" + text + "\""
	} else {
		text = "bytes.Buffer{}.WriteString("
		// 去除前后一个元素
		for i := 1; i <= ctx.GetChildCount()-2; i++ {
			var v = ctx.GetChild(i)
			var r = me.Visit(ctx.GetChild(i).(antlr.ParseTree)).(str)
			switch v.(type) {
			case parser.StringContentContext:
				text += ".WriteString(fmt.Sprint(" + r + ")"
			default:
				text += r
			}
		}
		text += ").String()"
		return text
	}
}

func (me *KVisitor) VisitStringContent(ctx *parser.StringContentContext) any {
	if ctx.TextLiteral().GetText() == "\\$" {
		return "$"
	}
	return ctx.TextLiteral().GetText()
}

func (me *KVisitor) VisitStringTemplate(ctx *parser.StringTemplateContext) any {
	var text = ""
	for _, v := range ctx.AllExpression() {
		var r = me.Visit(v).(Result)
		text += ".WriteString(fmt.Sprint(" + r.Text + ")"
	}
	return text
}

func (me *KVisitor) VisitDataStatement(ctx *parser.DataStatementContext) any {
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
	} else if ctx.StringExpr() != nil {
		r.Data = Str
		r.Text = me.Visit(ctx.StringExpr()).(string)
	} else if ctx.GetT().GetTokenType() == parser.FeelLexerCharLiteral {
		r.Data = Chr
		r.Text = ctx.CharLiteral().GetText()
	} else if ctx.GetT().GetTokenType() == parser.FeelLexerTrueLiteral {
		r.Data = Bool
		r.Text = T
	} else if ctx.GetT().GetTokenType() == parser.FeelLexerFalseLiteral {
		r.Data = Bool
		r.Text = F
	}
	return r
}

func (me *KVisitor) VisitFloatExpr(ctx *parser.FloatExprContext) any {
	number := ctx.FloatLiteral().GetText()
	return number
}

func (me *KVisitor) VisitIntegerExpr(ctx *parser.IntegerExprContext) any {
	number := ctx.GetChild(0).(antlr.ParseTree).GetText()
	return number
}

func (me *KVisitor) VisitTupleExpression(ctx *parser.TupleExpressionContext) any {
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
