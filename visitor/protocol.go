package visitor

import (
	"github.com/kulics/lite-go/parser/generate"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *LiteVisitor) VisitProtocolStatement(ctx *parser.ProtocolStatementContext) any {
	id := me.Visit(ctx.Id(0)).(Result)
	obj := ""
	interfaceProtocol := ""
	ptclName := id.Text
	if ctx.AnnotationSupport() != nil {
		obj += me.Visit(ctx.AnnotationSupport()).(string)
	}
	for _, item := range ctx.AllProtocolSupportStatement() {
		if r, ok := me.Visit(item).(Result); ok {
			interfaceProtocol += r.Text
		} else {
			interfaceProtocol += me.Visit(item).(string)
		}
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

func (me *LiteVisitor) VisitProtocolSupportStatement(ctx *parser.ProtocolSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *LiteVisitor) VisitProtocolFunctionStatement(ctx *parser.ProtocolFunctionStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	r := Result{}
	if ctx.AnnotationSupport() != nil {
		r.Text += me.Visit(ctx.AnnotationSupport()).(string)
	}
	r.Permission = "public"
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
	r.Text += id.Text + me.Visit(ctx.ParameterClauseIn()).(string) + templateContract + me.Visit(ctx.ParameterClauseOut()).(string) + Wrap
	return r
}
