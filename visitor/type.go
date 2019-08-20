package visitor

import (
	"github.com/kulics/lite-go/parser/generate"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (me *LiteVisitor) VisitTypeAliasStatement(ctx *parser.TypeAliasStatementContext) any {
	obj := ""
	obj = Type + me.Visit(ctx.Id()).(Result).Text + "=" + me.Visit(ctx.TypeType()).(string) + Wrap
	return obj
}

func (me *LiteVisitor) VisitTypeRedefineStatement(ctx *parser.TypeRedefineStatementContext) any {
	obj := ""
	obj = Type + me.Visit(ctx.Id()).(Result).Text + " " + me.Visit(ctx.TypeType()).(string) + Wrap
	return obj
}

func (me *LiteVisitor) VisitTypeType(ctx *parser.TypeTypeContext) any {
	obj := ""
	obj = me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
	return obj
}

func (me *LiteVisitor) VisitTypeReference(ctx *parser.TypeReferenceContext) any {
	obj := "*"
	if ctx.TypeNullable() != nil {
		obj += me.Visit(ctx.TypeNullable()).(string)
	} else if ctx.TypeNotNull() != nil {
		obj += me.Visit(ctx.TypeNotNull()).(string)
	}
	return obj
}

func (me *LiteVisitor) VisitTypeNullable(ctx *parser.TypeNullableContext) any {
	obj := ""
	obj = "*" + me.Visit(ctx.TypeNotNull()).(string)
	return obj
}

func (me *LiteVisitor) VisitTypeNotNull(ctx *parser.TypeNotNullContext) any {
	obj := ""
	obj = me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
	return obj
}

func (me *LiteVisitor) VisitTypeTuple(ctx *parser.TypeTupleContext) any {
	obj := ""
	obj += "("
	for i := 0; i < len(ctx.AllTypeType()); i++ {
		if i == 0 {
			obj += me.Visit(ctx.TypeType(i)).(string)
		} else {
			obj += "," + me.Visit(ctx.TypeType(i)).(string)
		}
	}
	obj += ")"
	return obj
}

func (me *LiteVisitor) VisitTypeArray(ctx *parser.TypeArrayContext) any {
	return "[]" + me.Visit(ctx.TypeType()).(string)
}

func (me *LiteVisitor) VisitTypeList(ctx *parser.TypeListContext) any {
	return "[]" + me.Visit(ctx.TypeType()).(string)
}

func (me *LiteVisitor) VisitTypeSet(ctx *parser.TypeSetContext) any {
	return "Set[]" + me.Visit(ctx.TypeType()).(string)
}

func (me *LiteVisitor) VisitTypeDictionary(ctx *parser.TypeDictionaryContext) any {
	return "map[" + me.Visit(ctx.TypeType(0)).(string) + "]" + me.Visit(ctx.TypeType(1)).(string)
}

func (me *LiteVisitor) VisitTypeChannel(ctx *parser.TypeChannelContext) any {
	return Chan + me.Visit(ctx.TypeType()).(string)
}

func (me *LiteVisitor) VisitTypePackage(ctx *parser.TypePackageContext) any {
	obj := ""
	obj += me.Visit(ctx.NameSpaceItem()).(string)
	if ctx.TemplateCall() != nil {
		obj += me.Visit(ctx.TemplateCall()).(string)
	}
	return obj
}

func (me *LiteVisitor) VisitTypeFunction(ctx *parser.TypeFunctionContext) any {
	obj := ""
	in := me.Visit(ctx.TypeFunctionParameterClause(0)).(string)
	out := me.Visit(ctx.TypeFunctionParameterClause(1)).(string)
	if ctx.GetT().GetTokenType() == parser.LiteLexerRight_Arrow {
		obj = Func + "(" + in + ")" + "(" + out + ")"
	} else {
		obj = Func + "(" + in + ")" + "(" + out + ")"
	}
	return obj
}

func (me *LiteVisitor) VisitTypeAny(ctx *parser.TypeAnyContext) any {
	return Any
}

func (me *LiteVisitor) VisitTypeFunctionParameterClause(ctx *parser.TypeFunctionParameterClauseContext) any {
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

func (me *LiteVisitor) VisitTypeBasic(ctx *parser.TypeBasicContext) any {
	obj := ""
	switch ctx.GetT().GetTokenType() {
	case parser.LiteLexerTypeI8:
		obj = I8
	case parser.LiteLexerTypeU8:
		obj = U8
	case parser.LiteLexerTypeI16:
		obj = I16
	case parser.LiteLexerTypeU16:
		obj = U16
	case parser.LiteLexerTypeI32:
		obj = I32
	case parser.LiteLexerTypeU32:
		obj = U32
	case parser.LiteLexerTypeI64:
		obj = I64
	case parser.LiteLexerTypeU64:
		obj = U64
	case parser.LiteLexerTypeF32:
		obj = F32
	case parser.LiteLexerTypeF64:
		obj = F64
	case parser.LiteLexerTypeChr:
		obj = Chr
	case parser.LiteLexerTypeStr:
		obj = Str
	case parser.LiteLexerTypeBool:
		obj = Bool
	case parser.LiteLexerTypeInt:
		obj = Int
	case parser.LiteLexerTypeNum:
		obj = Num
	case parser.LiteLexerTypeByte:
		obj = Byte
	default:
		obj = Any
	}
	return obj
}
