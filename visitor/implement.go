package visitor

import (
	parser "github.com/kulics-works/feel-go/parser/generate"
)

func (me *KVisitor) VisitImplementStatement(ctx *parser.ImplementStatementContext) any {
	obj := ""
	var methed = ""
	id := me.Visit(ctx.Id()).(Result)
	me.self.Type = id.Text

	if item := ctx.PackageFieldStatement(); item != nil {
		var r = me.Visit(item).(Result)
		obj += r.Text
		methed += r.Data.(string)
	}
	// for _, item := range ctx.AllPackageImplementStatement() {
	// 	methed += me.Visit(item).(Result).Text
	// }
	obj += methed
	me.self = Parameter{}
	return obj
}
