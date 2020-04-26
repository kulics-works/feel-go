package visitor

import (
	parser "github.com/kulics-works/k-go/parser/generate"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *KVisitor) VisitCallExpression(ctx *parser.CallExpressionContext) any {
	r := me.Visit(ctx.Id()).(Result)
	r.Text = "." + r.Text
	if ctx.CallFunc() != nil {
		e2 := me.Visit(ctx.CallFunc()).(Result)
		r.Text = r.Text + e2.Text
	} else if ctx.CallElement() != nil {
		e2 := me.Visit(ctx.CallElement()).(Result)
		r.Text = r.Text + e2.Text
	}
	return r
}

func (me *KVisitor) VisitCallAsync(ctx *parser.CallAsyncContext) any {
	r := Result{Data: "var"}
	if ctx.Tuple() != nil {
		r.Text += "(" + me.Visit(ctx.Tuple()).(Result).Text + ")"
	} else {
		r.Text += "(" + me.Visit(ctx.Lambda()).(Result).Text + ")"
	}
	return r
}

func (me *KVisitor) VisitCallChannel(ctx *parser.CallChannelContext) any {
	r := Result{}
	r.Text = "<-" + me.Visit(ctx.Expression()).(Result).Text
	return r
}

func (me *KVisitor) VisitCallElement(ctx *parser.CallElementContext) any {
	if ctx.Expression() == nil {
		return Result{Text: me.Visit(ctx.Slice()).(string)}
	}
	r := me.Visit(ctx.Expression()).(Result)
	r.Text = "[" + r.Text + "]"
	return r
}

func (me *KVisitor) VisitSlice(ctx *parser.SliceContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
}

func (me *KVisitor) VisitSliceFull(ctx *parser.SliceFullContext) any {
	// order := true
	// if ctx.Tilde() == nil {
	// 	order = false
	// }
	expr1 := me.Visit(ctx.Expression(0)).(Result)
	expr2 := me.Visit(ctx.Expression(1)).(Result)
	return "[" + expr1.Text + ":" + expr2.Text + "]"
}

func (me *KVisitor) VisitSliceStart(ctx *parser.SliceStartContext) any {
	attach := ""
	expr := me.Visit(ctx.Expression()).(Result)
	return "[" + expr.Text + ":" + attach + "]"
}

func (me *KVisitor) VisitSliceEnd(ctx *parser.SliceEndContext) any {
	// order := true
	// if ctx.Tilde() == nil {
	// 	order = false
	// }
	expr := me.Visit(ctx.Expression()).(Result)
	return "[:" + expr.Text + "]"
}

func (me *KVisitor) VisitCallFunc(ctx *parser.CallFuncContext) any {
	r := Result{Data: "var"}
	if ctx.Tuple() != nil {
		r.Text += "(" + me.Visit(ctx.Tuple()).(Result).Text + ")"
	} else {
		r.Text += "(" + me.Visit(ctx.Lambda()).(Result).Text + ")"
	}
	return r
}

func (me *KVisitor) VisitCallNew(ctx *parser.CallNewContext) any {
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

func (me *KVisitor) VisitCallPkg(ctx *parser.CallPkgContext) any {
	r := Result{Data: me.Visit(ctx.TypeType())}
	r.Text = me.Visit(ctx.TypeType()).(string)
	if ctx.PkgAssign() != nil {
		r.Text += me.Visit(ctx.PkgAssign()).(string)
	} else if ctx.ListAssign() != nil {
		r.Text += me.Visit(ctx.ListAssign()).(string)
	} else if ctx.DictionaryAssign() != nil {
		r.Text += me.Visit(ctx.DictionaryAssign()).(string)
	} else {
		r.Text += "{}"
	}
	return r
}

func (me *KVisitor) VisitPkgAssign(ctx *parser.PkgAssignContext) any {
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

func (me *KVisitor) VisitListAssign(ctx *parser.ListAssignContext) any {
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

func (me *KVisitor) VisitDictionaryAssign(ctx *parser.DictionaryAssignContext) any {
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

func (me *KVisitor) VisitPkgAssignElement(ctx *parser.PkgAssignElementContext) any {
	obj := ""
	obj += me.Visit(ctx.Name()).(string) + " = " + me.Visit(ctx.Expression()).(Result).Text
	return obj
}

func (me *KVisitor) VisitList(ctx *parser.ListContext) any {
	typeName := "object"
	result := Result{}
	for index := 0; index < len(ctx.AllExpression()); index++ {
		r := me.Visit(ctx.Expression(index)).(Result)
		if index == 0 {
			typeName = r.Data.(str)
			result.Text += r.Text
		} else {
			if typeName != r.Data.(str) {
				typeName = Any
			}
			result.Text += "," + r.Text
		}
	}
	typeName = "[]" + typeName
	result.Data = typeName
	result.Text = result.Data.(str) + "{ " + result.Text + " }"
	return result
}

func (me *KVisitor) VisitDictionary(ctx *parser.DictionaryContext) any {
	key := Any
	value := Any
	result := Result{}
	for index := 0; index < len(ctx.AllDictionaryElement()); index++ {
		r := me.Visit(ctx.DictionaryElement(index)).(DicEle)
		if index == 0 {
			key = r.Key
			value = r.Value
			result.Text += r.Text
		} else {
			if key != r.Key {
				key = Any
			}
			if value != r.Value {
				value = Any
			}
			result.Text += "," + r.Text
		}
	}
	typeName := "map[" + key + "]" + value
	result.Data = typeName
	result.Text = typeName + "{ " + result.Text + " }"
	return result
}

func (me *KVisitor) VisitDictionaryElement(ctx *parser.DictionaryElementContext) any {
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
