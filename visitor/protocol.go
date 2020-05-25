package visitor

import (
	parser "github.com/kulics-works/feel-go/parser/generate"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *KVisitor) VisitProtocolStatement(ctx *parser.ProtocolStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	obj := ""
	interfaceProtocol := ""
	ptclName := id.Text
	if ctx.AnnotationSupport() != nil {
		obj += me.Visit(ctx.AnnotationSupport()).(string)
	}
	if item := ctx.ProtocolSubStatement(); item != nil {
		interfaceProtocol += me.Visit(item).(string)
	}
	obj += "type " + ptclName + " interface"
	// 泛型
	templateContract := ""
	// ? ctx.templateDefine() >< () {
	// 	template := me.Visit(ctx.templateDefine()):TemplateItem
	// 	obj += template.Template
	// 	templateContract = template.Contract
	// }
	obj += templateContract + BlockLeft + Wrap
	obj += interfaceProtocol
	obj += BlockRight + Wrap
	return obj
}

func (me *KVisitor) VisitProtocolSubStatement(ctx *parser.ProtocolSubStatementContext) any {
	var obj = ""
	for _, item := range ctx.AllProtocolSupportStatement() {
		obj += me.Visit(item).(str)
	}
	return obj
}

func (me *KVisitor) VisitProtocolSupportStatement(ctx *parser.ProtocolSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *KVisitor) VisitProtocolFunctionStatement(ctx *parser.ProtocolFunctionStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	obj := ""
	if ctx.AnnotationSupport() != nil {
		obj += me.Visit(ctx.AnnotationSupport()).(string)
	}
	// # 异步 #
	// ? ctx.t.Type == Right Flow {
	// 	pout := me.Visit(ctx.parameterClauseOut()):Str
	// 	? pout >< "void" {
	// 		pout = ""Task"<"pout">"
	// 	} _ {
	// 		pout = Task
	// 	}
	// 	r.text += pout + " " + id.text
	// } _ {
	// 	r.text += me.Visit(ctx.parameterClauseOut()) + " " + id.text
	// }
	// 泛型
	templateContract := ""
	// ? ctx.templateDefine() >< () {
	// 	template := me.Visit(ctx.templateDefine()):TemplateItem
	// 	r.text += template.Template
	// 	templateContract = template.Contract
	// }
	obj += id.Text + me.Visit(ctx.ParameterClauseIn()).(string) + templateContract + me.Visit(ctx.ParameterClauseOut()).(string) + Wrap
	return obj
}
