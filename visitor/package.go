package visitor

import (
	"github.com/kulics/lite-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (sf *LiteVisitor) VisitIncludeStatement(ctx *parser.IncludeStatementContext) interface{} {
	return sf.Visit(ctx.TypeType()).(string) + Wrap
}

func (sf *LiteVisitor) VisitPackageStatement(ctx *parser.PackageStatementContext) interface{} {
	id := sf.Visit(ctx.Id()).(Result)
	obj := ""
	Init := ""

	// # 处理构造函数 #
	// ctx.packageNewStatement() @ item {
	// 	Init += "public " id.text " " Visit(item):Str ""
	// }
	for _, item := range ctx.AllPackageSupportStatement() {
		obj += sf.Visit(item).(string)
	}
	obj = Init + obj
	obj += BlockRight + Wrap
	header := ""
	if ctx.AnnotationSupport() != nil {
		header += sf.Visit(ctx.AnnotationSupport()).(string)
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

func (sf *LiteVisitor) VisitPackageSupportStatement(ctx *parser.PackageSupportStatementContext) interface{} {
	return sf.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (sf *LiteVisitor) VisitPackageVariableStatement(ctx *parser.PackageVariableStatementContext) interface{} {
	r1 := sf.Visit(ctx.Id()).(Result)
	typ := ""
	// r2:= Result{}
	// if ctx.Expression() != nil {
	// 	r2 = Visit(ctx.expression()).(Result)
	// 	typ = r2.data:Str
	// }
	if ctx.TypeType() != nil {
		typ = sf.Visit(ctx.TypeType()).(string)
	}
	obj := ""
	if ctx.AnnotationSupport() != nil {
		obj += sf.Visit(ctx.AnnotationSupport()).(string)
	}

	obj += r1.Text + " " + typ
	// if r2 != nil {
	// 	obj += " = " + r2.text
	// }
	obj += Wrap
	return obj
}
