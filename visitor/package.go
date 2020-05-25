package visitor

import (
	parser "github.com/kulics-works/feel-go/parser/generate"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *KVisitor) VisitIncludeStatement(ctx *parser.IncludeStatementContext) any {
	return me.Visit(ctx.TypeType()).(string) + Wrap
}

func (me *KVisitor) VisitPackageStatement(ctx *parser.PackageStatementContext) any {
	id := me.Visit(ctx.Id()).(Result)
	me.self.Type = id.Text

	var obj = ""
	var Init = ""
	var methed = ""

	// # 处理构造函数 #
	// ctx.packageNewStatement() @ item {
	// 	Init += "public " id.text " " Visit(item):Str ""
	// }
	if item := ctx.PackageFieldStatement(); item != nil {
		var r = me.Visit(item).(Result)
		obj += r.Text
		methed += r.Data.(string)
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
	obj += methed
	return obj
}

func (me *KVisitor) VisitPackageSupportStatement(ctx *parser.PackageSupportStatementContext) any {
	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (me *KVisitor) VisitPackageFieldStatement(ctx *parser.PackageFieldStatementContext) any {
	var obj = ""
	var method = ""
	if ctx.Id(0) != nil {
		me.self.Id = me.Visit(ctx.Id(0)).(Result).Text
	}
	if ctx.GetP() != nil {
		me.self.Type = "*" + me.self.Type
	}
	for _, item := range ctx.AllPackageSupportStatement() {
		switch item.GetChild(0).(type) {
		case *parser.PackageFunctionStatementContext:
			method += me.Visit(item).(string)
		case *parser.PackageVariableStatementContext,
			*parser.IncludeStatementContext:
			obj += me.Visit(item).(string)
		}
	}
	me.self.Id = ""
	me.self.Type = me.self.Type[1:]
	return Result{Text: obj, Data: method}
}

func (me *KVisitor) VisitPackageVariableStatement(ctx *parser.PackageVariableStatementContext) any {
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

func (me *KVisitor) VisitPackageFunctionStatement(ctx *parser.PackageFunctionStatementContext) any {
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
