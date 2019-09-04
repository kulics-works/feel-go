package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kulics/lite-go/parser/generate"
)

func (me *LiteVisitor) VisitImplementStatement(ctx *parser.ImplementStatementContext) any {
	obj := ""
	// if ctx.AnnotationSupport() >< () {
	// 	obj += Visit(context.annotationSupport())
	// }
	// 异步
	// if ctx.GetT().GetTokenType() == parser.XsLexerRight_Flow {
	// pout := Visit(ctx.ParameterClauseOut()).(string)
	// obj += ""id.permission" async static "pout" "id.text""
	// } else {
	// 	obj += Func + id.Text  + me.Visit(ctx.ParameterClauseOut()).(string)
	// }

	// 泛型
	// templateContract := ""
	// if context.templateDefine() >< () {
	// 	template := Visit(context.templateDefine()):TemplateItem
	// 	obj += template.Template
	// 	templateContract = template.Contract
	// }
	me.self = me.Visit(ctx.ParameterClauseSelf()).(Parameter)
	for _, item := range ctx.AllImplementSupportStatement() {
		if v, ok := me.Visit(item).(string); ok {
			obj += v
		}
	}
	me.self = Parameter{}
	return obj
}

func (me *LiteVisitor) VisitImplementSupportStatement(ctx *parser.ImplementSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *LiteVisitor) VisitImplementFunctionStatement(ctx *parser.ImplementFunctionStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	obj := ""
	me.add_current_set()
	obj += Func + "(" + me.self.Id + " " + me.self.Type + ")" +
		id.Text + me.Visit(ctx.ParameterClauseIn()).(string) +
		me.Visit(ctx.ParameterClauseOut()).(string) + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	me.delete_current_set()
	return obj
}
