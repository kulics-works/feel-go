package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kulics/lite-go/parser"
)

var self Parameter

func (sf *LiteVisitor) VisitImplementStatement(ctx *parser.ImplementStatementContext) interface{} {
	obj := ""
	// if ctx.AnnotationSupport() >< () {
	// 	obj += Visit(context.annotationSupport())
	// }
	// 异步
	// if ctx.GetT().GetTokenType() == parser.XsLexerRight_Flow {
	// pout := Visit(ctx.ParameterClauseOut()).(string)
	// obj += ""id.permission" async static "pout" "id.text""
	// } else {
	// 	obj += Func + id.Text  + sf.Visit(ctx.ParameterClauseOut()).(string)
	// }

	// 泛型
	// templateContract := ""
	// if context.templateDefine() >< () {
	// 	template := Visit(context.templateDefine()):TemplateItem
	// 	obj += template.Template
	// 	templateContract = template.Contract
	// }
	self = sf.Visit(ctx.ParameterClauseSelf()).(Parameter)
	for _, item := range ctx.AllImplementSupportStatement() {
		obj += sf.Visit(item).(string)
	}
	self = Parameter{}
	return obj
}

func (sf *LiteVisitor) VisitImplementSupportStatement(ctx *parser.ImplementSupportStatementContext) interface{} {
	return sf.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (sf *LiteVisitor) VisitImplementFunctionStatement(ctx *parser.ImplementFunctionStatementContext) interface{} {
	id := sf.Visit(ctx.Id()).(Result)
	obj := ""
	obj += Func + "(" + self.Id + " " + self.Type + ")" +
		id.Text + sf.Visit(ctx.ParameterClauseIn()).(string) +
		sf.Visit(ctx.ParameterClauseOut()).(string) + BlockLeft + Wrap
	obj += sf.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	return obj
}
