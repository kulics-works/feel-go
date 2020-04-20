package visitor

import (
	"fmt"

	parser "github.com/kulics-works/k-go/parser/generate"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Namespace struct {
	Name    string
	Imports string
	Alias   string
}

func (me *KVisitor) VisitStatement(ctx *parser.StatementContext) any {
	obj := ""
	ns, ok := me.Visit(ctx.ExportStatement()).(Namespace)
	if !ok {
		return ""
	}
	obj += fmt.Sprintf("package %s%s%s%s%s", ns.Name, Wrap, ns.Imports, Wrap, ns.Alias)
	for _, item := range ctx.AllNamespaceSupportStatement() {
		if v, ok := me.Visit(item).(string); ok {
			obj += v
		}
	}
	return obj
}

func (me *KVisitor) VisitExportStatement(ctx *parser.ExportStatementContext) any {
	name := ctx.TextLiteral().GetText()
	obj := Namespace{
		Name: name[1 : len(name)-1],
	}
	for _, item := range ctx.AllImportStatement() {
		obj.Imports += me.Visit(item).(string)
	}
	for _, item := range ctx.AllTypeAliasStatement() {
		obj.Alias += me.Visit(item).(string)
	}
	return obj
}

func (me *KVisitor) VisitImportStatement(ctx *parser.ImportStatementContext) any {
	obj := "import "
	if ctx.AnnotationSupport() != nil {
		obj += me.Visit(ctx.AnnotationSupport()).(string)
	}
	ns := ctx.TextLiteral().GetText()
	if ctx.Call() != nil {
		obj += ". " + ns
	} else if ctx.Id() != nil {
		obj += me.Visit(ctx.Id()).(Result).Text + " " + ns
	} else {
		obj += ns
	}
	obj += Wrap
	return obj
}

func (me *KVisitor) VisitNameSpaceItem(ctx *parser.NameSpaceItemContext) any {
	obj := ""
	for i := 0; i < len(ctx.AllId()); i++ {
		id := me.Visit(ctx.Id(i)).(Result)
		if i == 0 {
			obj += "" + id.Text
		} else {
			obj += "." + id.Text
		}
	}
	return obj
}

func (me *KVisitor) VisitNamespaceSupportStatement(ctx *parser.NamespaceSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *KVisitor) VisitNamespaceFunctionStatement(ctx *parser.NamespaceFunctionStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
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
	me.add_current_set()
	obj += Func + id.Text + me.Visit(ctx.ParameterClauseIn()).(string) + me.Visit(ctx.ParameterClauseOut()).(string) + BlockLeft + Wrap
	obj += me.ProcessFunctionSupport(ctx.AllFunctionSupportStatement())
	obj += BlockRight + Wrap
	me.delete_current_set()
	return obj
}

func (me *KVisitor) VisitNamespaceConstantStatement(ctx *parser.NamespaceConstantStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	expr := me.Visit(ctx.Expression()).(Result)
	typ := ""
	if ctx.TypeType() != nil {
		typ = me.Visit(ctx.TypeType()).(string)
	}

	obj := ""
	if ctx.AnnotationSupport() != nil {
		obj += me.Visit(ctx.AnnotationSupport()).(string)
	}

	obj += Const + id.Text + " " + typ + " = " + expr.Text + Wrap
	return obj
}

func (me *KVisitor) VisitNamespaceVariableStatement(ctx *parser.NamespaceVariableStatementContext) any {
	r1 := me.Visit(ctx.Id()).(Result)
	me.add_id(r1.Text)
	typ := ""

	if ctx.TypeType() != nil {
		typ = me.Visit(ctx.TypeType()).(string)
	}
	obj := ""
	if ctx.AnnotationSupport() != nil {
		obj += me.Visit(ctx.AnnotationSupport()).(string)
	}

	obj += Var + r1.Text + " " + typ
	if ctx.Expression() != nil {
		obj += " = " + me.Visit(ctx.Expression()).(Result).Text + Wrap
	}
	return obj
}
