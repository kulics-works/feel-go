package visitor

import (
	parser "github.com/kulics-works/feel-go/parser/generate"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *KVisitor) VisitTypeAliasStatement(ctx *parser.TypeAliasStatementContext) any {
	obj := ""
	obj = Type + me.Visit(ctx.Id()).(Result).Text + "=" + me.Visit(ctx.TypeType()).(string) + Wrap
	return obj
}

func (me *KVisitor) VisitTypeRedefineStatement(ctx *parser.TypeRedefineStatementContext) any {
	obj := ""
	obj = Type + me.Visit(ctx.Id()).(Result).Text + " " + me.Visit(ctx.TypeType()).(string) + Wrap
	return obj
}

func (me *KVisitor) VisitTypeType(ctx *parser.TypeTypeContext) any {
	obj := ""
	obj = me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
	return obj
}

func (me *KVisitor) VisitTypeNullable(ctx *parser.TypeNullableContext) any {
	obj := ""
	obj = "*" + me.Visit(ctx.TypeNotNull()).(string)
	return obj
}

func (me *KVisitor) VisitTypeNotNull(ctx *parser.TypeNotNullContext) any {
	obj := ""
	obj = me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
	return obj
}

func (me *KVisitor) VisitTypeArray(ctx *parser.TypeArrayContext) any {
	return "[]" + me.Visit(ctx.TypeType()).(string)
}

func (me *KVisitor) VisitTypeList(ctx *parser.TypeListContext) any {
	return "[]" + me.Visit(ctx.TypeType()).(string)
}

func (me *KVisitor) VisitTypeSet(ctx *parser.TypeSetContext) any {
	return "Set[]" + me.Visit(ctx.TypeType()).(string)
}

func (me *KVisitor) VisitTypeDictionary(ctx *parser.TypeDictionaryContext) any {
	return "map[" + me.Visit(ctx.TypeType(0)).(string) + "]" + me.Visit(ctx.TypeType(1)).(string)
}

func (me *KVisitor) VisitTypeChannel(ctx *parser.TypeChannelContext) any {
	return Chan + me.Visit(ctx.TypeType()).(string)
}

func (me *KVisitor) VisitTypePackage(ctx *parser.TypePackageContext) any {
	obj := ""
	obj += me.Visit(ctx.NameSpaceItem()).(string)
	if ctx.TemplateCall() != nil {
		obj += me.Visit(ctx.TemplateCall()).(string)
	}
	return obj
}

func (me *KVisitor) VisitTypeFunction(ctx *parser.TypeFunctionContext) any {
	obj := ""
	in := me.Visit(ctx.TypeFunctionParameterClause(0)).(string)
	out := me.Visit(ctx.TypeFunctionParameterClause(1)).(string)
	if ctx.GetT().GetTokenType() == parser.FeelLexerRight_Arrow {
		obj = Func + "(" + in + ")" + "(" + out + ")"
	} else {
		obj = Func + "(" + in + ")" + "(" + out + ")"
	}
	return obj
}

func (me *KVisitor) VisitTypeAny(ctx *parser.TypeAnyContext) any {
	return Any
}

func (me *KVisitor) VisitTypeFunctionParameterClause(ctx *parser.TypeFunctionParameterClauseContext) any {
	obj := ""
	for i := 0; i <= len(ctx.AllTypeType())-1; i++ {
		p := me.Visit(ctx.TypeType(i)).(string)
		if i == 0 {
			obj += p
		} else {
			obj += ", " + p
		}
	}
	return obj
}

func (me *KVisitor) VisitTypeBasic(ctx *parser.TypeBasicContext) any {
	obj := ""
	switch ctx.GetT().GetTokenType() {
	case parser.FeelLexerTypeI8:
		obj = I8
	case parser.FeelLexerTypeU8:
		obj = U8
	case parser.FeelLexerTypeI16:
		obj = I16
	case parser.FeelLexerTypeU16:
		obj = U16
	case parser.FeelLexerTypeI32:
		obj = I32
	case parser.FeelLexerTypeU32:
		obj = U32
	case parser.FeelLexerTypeI64:
		obj = I64
	case parser.FeelLexerTypeU64:
		obj = U64
	case parser.FeelLexerTypeF32:
		obj = F32
	case parser.FeelLexerTypeF64:
		obj = F64
	case parser.FeelLexerTypeChr:
		obj = Chr
	case parser.FeelLexerTypeStr:
		obj = Str
	case parser.FeelLexerTypeBool:
		obj = Bool
	case parser.FeelLexerTypeInt:
		obj = Int
	case parser.FeelLexerTypeNum:
		obj = Num
	case parser.FeelLexerTypeByte:
		obj = Byte
	default:
		obj = Any
	}
	return obj
}
