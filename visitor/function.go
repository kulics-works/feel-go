package visitor

import (
	"github.com/kulics/lite-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Parameter struct {
	Id         string
	Type       string
	Value      string
	Annotation string
	Permission string
}

func (me *LiteVisitor) ProcessFunctionSupport(items []parser.IFunctionSupportStatementContext) string {
	obj := ""
	content := ""
	// lazy := []string{}
	// for _, item := range items {
	// if item.GetChild(0) == :UsingStatementContext {
	// 	lazy.add("}")
	// 	content += "using (" + me.Visit(item).(string) + ") " + BlockLeft + Wrap
	// } else {
	// content += me.Visit(item).(string)
	// }
	// }
	// if lazy.Count > 0 {
	// 	for i := lazy.Count - 1; i >= 0; i-- {
	// 		content += BlockRight
	// 	}
	// }
	for _, item := range items {
		if v, ok := me.Visit(item).(string); ok {
			content += v
		}
	}
	obj += content
	return obj
}

func (me *LiteVisitor) VisitFunctionStatement(ctx *parser.FunctionStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	obj := ""
	// 异步
	// ? context.t.Type == Right Flow {
	// 	pout := Visit(context.parameterClauseOut()):Str
	// 	? pout >< "void" {
	// 		pout = ""Task"<"pout">"
	// 	} _ {
	// 		pout = Task
	// 	}
	// 	obj += " async "pout" "id.text""
	// } _ {
	// 	obj += ""Visit(context.parameterClauseOut())" "id.text""
	// }
	// # 泛型 #
	templateContract := ""
	// ? context.templateDefine() >< () {
	// 	template := Visit(context.templateDefine()):TemplateItem
	// 	obj += template.Template
	// 	templateContract = template.Contract
	// }
	me.add_current_set()
	obj += id.Text + ":=" + Func + me.Visit(ctx.ParameterClauseIn()).(string) + templateContract +
		me.Visit(ctx.ParameterClauseOut()).(string) + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	me.delete_current_set()
	return obj
}

func (me *LiteVisitor) VisitFunctionSupportStatement(ctx *parser.FunctionSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *LiteVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) any {
	if ctx.TupleExpression() != nil {
		r := me.Visit(ctx.TupleExpression()).(Result)
		return "return " + r.Text + Wrap
	}
	return "return " + Wrap
}

func (me *LiteVisitor) VisitParameterClauseIn(ctx *parser.ParameterClauseInContext) any {
	obj := "("
	temp := []string{}
	for i := len(ctx.AllParameter()) - 1; i >= 0; i-- {
		p := me.Visit(ctx.Parameter(i)).(Parameter)
		temp = append(temp, p.Annotation+" "+p.Id+" "+p.Type)
		me.add_id(p.Id)
	}
	for i := len(temp) - 1; i >= 0; i-- {
		if i == len(temp)-1 {
			obj += temp[i]
		} else {
			obj += ", " + temp[i]
		}
	}
	obj += ")"
	return obj
}

func (me *LiteVisitor) VisitParameterClauseOut(ctx *parser.ParameterClauseOutContext) any {
	obj := "("
	temp := []string{}
	for i := len(ctx.AllParameter()) - 1; i >= 0; i-- {
		p := me.Visit(ctx.Parameter(i)).(Parameter)
		temp = append(temp, p.Annotation+" "+p.Id+" "+p.Type)
	}
	for i := len(temp) - 1; i >= 0; i-- {
		if i == len(temp)-1 {
			obj += temp[i]
		} else {
			obj += ", " + temp[i]
		}
	}
	obj += ")"
	return obj
}

func (me *LiteVisitor) VisitParameterClauseSelf(ctx *parser.ParameterClauseSelfContext) any {
	p := Parameter{}
	id := me.Visit(ctx.Id()).(Result)
	p.Id = id.Text
	p.Permission = id.Permission
	p.Type = me.Visit(ctx.TypeType()).(string)
	return p
}

func (me *LiteVisitor) VisitParameter(ctx *parser.ParameterContext) any {
	p := Parameter{}
	id := me.Visit(ctx.Id()).(Result)
	p.Id = id.Text
	p.Permission = id.Permission
	if ctx.AnnotationSupport() != nil {
		p.Annotation = me.Visit(ctx.AnnotationSupport()).(string)
	}
	if ctx.Expression() != nil {
		p.Value = "=" + me.Visit(ctx.Expression()).(Result).Text
	}
	p.Type = me.Visit(ctx.TypeType()).(string)
	return p
}

func (me *LiteVisitor) VisitTuple(ctx *parser.TupleContext) any {
	obj := ""
	for i := 0; i < len(ctx.AllExpression()); i++ {
		r := me.Visit(ctx.Expression(i)).(Result)
		if i == 0 {
			obj += r.Text
		} else {
			obj += ", " + r.Text
		}
	}
	result := Result{Data: "var", Text: obj}
	return result
}

func (me *LiteVisitor) VisitFunctionExpression(ctx *parser.FunctionExpressionContext) any {
	r := Result{}
	r.Data = "var"
	// 异步
	// ? context.t.Type == Right Flow {
	// 	pout := Visit(context.parameterClauseOut()):Str
	// 	? pout >< "void" {
	// 		pout = ""Task"<"pout">"
	// 	} _ {
	// 		pout = Task
	// 	}
	// 	obj += " async "pout" "id.text""
	// } _ {
	// 	obj += ""Visit(context.parameterClauseOut())" "id.text""
	// }
	// # 泛型 #
	templateContract := ""
	// ? context.templateDefine() >< () {
	// 	template := Visit(context.templateDefine()):TemplateItem
	// 	obj += template.Template
	// 	templateContract = template.Contract
	// }
	me.add_current_set()
	r.Text += Func + me.Visit(ctx.ParameterClauseIn()).(string) + templateContract +
		me.Visit(ctx.ParameterClauseOut()).(string) + BlockLeft + Wrap
	r.Text += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	r.Text += BlockRight
	me.delete_current_set()
	return r
}
