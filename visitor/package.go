package visitor

import (
	"github.com/kulics/lite-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *LiteVisitor) VisitIncludeStatement(ctx *parser.IncludeStatementContext) any {
	return me.Visit(ctx.TypeType()).(string) + Wrap
}

func (me *LiteVisitor) VisitPackageStatement(ctx *parser.PackageStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	obj := ""
	Init := ""

	// # 处理构造函数 #
	// ctx.packageNewStatement() @ item {
	// 	Init += "public " id.text " " Visit(item):Str ""
	// }
	for _, item := range ctx.AllPackageSupportStatement() {
		obj += me.Visit(item).(string)
	}
	obj = Init + obj
	obj += BlockRight + Wrap
	header := ""
	if ctx.AnnotationSupport() != nil {
		header += me.Visit(ctx.AnnotationSupport()).(string)
	}
	header += "type " + id.Text + " struct"
	// # 泛型 #
	// template := ""
	templateContract := ""
	// ? ctx.templateDefine() >< () {
	// 	item := Visit(ctx.templateDefine()):TemplateItem
	// 	template += item.Template
	// 	templateContract = item.Contract
	// 	header += template;
	// }

	header += templateContract + BlockLeft + Wrap
	obj = header + obj

	return obj
}

func (me *LiteVisitor) VisitPackageSupportStatement(ctx *parser.PackageSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *LiteVisitor) VisitPackageVariableStatement(ctx *parser.PackageVariableStatementContext) any {
	r1 := me.Visit(ctx.Id()).(Result)
	typ := ""
	// r2:= Result{}
	// if ctx.Expression() != nil {
	// 	r2 = Visit(ctx.expression()).(Result)
	// 	typ = r2.data:Str
	// }
	if ctx.TypeType() != nil {
		typ = me.Visit(ctx.TypeType()).(string)
	}
	obj := ""
	if ctx.AnnotationSupport() != nil {
		obj += me.Visit(ctx.AnnotationSupport()).(string)
	}

	obj += r1.Text + " " + typ
	// if r2 != nil {
	// 	obj += " = " + r2.text
	// }
	obj += Wrap
	return obj
}
