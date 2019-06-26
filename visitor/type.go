package visitor

import (
	"github.com/kulics/lite-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (sf *LiteVisitor) VisitTypeAliasStatement(ctx *parser.TypeAliasStatementContext) interface{} {
	obj := ""
	obj = Type + sf.Visit(ctx.Id()).(Result).Text + "=" + sf.Visit(ctx.TypeType()).(string) + Wrap
	return obj
}

func (sf *LiteVisitor) VisitTypeRedefineStatement(ctx *parser.TypeRedefineStatementContext) interface{} {
	obj := ""
	obj = Type + sf.Visit(ctx.Id()).(Result).Text + " " + sf.Visit(ctx.TypeType()).(string) + Wrap
	return obj
}

func (sf *LiteVisitor) VisitTypeType(ctx *parser.TypeTypeContext) interface{} {
	obj := ""
	obj = sf.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
	return obj
}

func (sf *LiteVisitor) VisitTypeReference(ctx *parser.TypeReferenceContext) interface{} {
	obj := "*"
	if ctx.TypeNullable() != nil {
		obj += sf.Visit(ctx.TypeNullable()).(string)
	} else if ctx.TypeNotNull() != nil {
		obj += sf.Visit(ctx.TypeNotNull()).(string)
	}
	return obj
}

func (sf *LiteVisitor) VisitTypeNullable(ctx *parser.TypeNullableContext) interface{} {
	obj := ""
	obj = "*" + sf.Visit(ctx.TypeNotNull()).(string)
	return obj
}

func (sf *LiteVisitor) VisitTypeNotNull(ctx *parser.TypeNotNullContext) interface{} {
	obj := ""
	obj = sf.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
	return obj
}

func (sf *LiteVisitor) VisitTypeTuple(ctx *parser.TypeTupleContext) interface{} {
	obj := ""
	obj += "("
	for i := 0; i < len(ctx.AllTypeType()); i++ {
		if i == 0 {
			obj += sf.Visit(ctx.TypeType(i)).(string)
		} else {
			obj += "," + sf.Visit(ctx.TypeType(i)).(string)
		}
	}
	obj += ")"
	return obj
}

func (sf *LiteVisitor) VisitGetType(ctx *parser.GetTypeContext) interface{} {
	r := Result{Data: "System.Type"}
	// if ctx.TypeType() == nil {
	// 	r.Text = sf.Visit(ctx.Expression()).(Result).Text+".GetType()"
	// } else {
	// 	r.Text = "typeof("+sf.Visit(ctx.typeType())+")"
	// }
	return r
}

func (sf *LiteVisitor) VisitTypeArray(ctx *parser.TypeArrayContext) interface{} {
	return "[]" + sf.Visit(ctx.TypeType()).(string)
}

func (sf *LiteVisitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	return "[]" + sf.Visit(ctx.TypeType()).(string)
}

func (sf *LiteVisitor) VisitTypeSet(ctx *parser.TypeSetContext) interface{} {
	return "Set[]" + sf.Visit(ctx.TypeType()).(string)
}

func (sf *LiteVisitor) VisitTypeDictionary(ctx *parser.TypeDictionaryContext) interface{} {
	return "map[" + sf.Visit(ctx.TypeType(0)).(string) + "]" + sf.Visit(ctx.TypeType(1)).(string)
}

func (sf *LiteVisitor) VisitTypeChannel(ctx *parser.TypeChannelContext) interface{} {
	return Chan + sf.Visit(ctx.TypeType()).(string)
}

func (sf *LiteVisitor) VisitTypePackage(ctx *parser.TypePackageContext) interface{} {
	obj := ""
	obj += sf.Visit(ctx.NameSpaceItem()).(string)
	if ctx.TemplateCall() != nil {
		obj += sf.Visit(ctx.TemplateCall()).(string)
	}
	return obj
}

func (sf *LiteVisitor) VisitTypeFunction(ctx *parser.TypeFunctionContext) interface{} {
	obj := ""
	in := sf.Visit(ctx.TypeFunctionParameterClause(0)).(string)
	out := sf.Visit(ctx.TypeFunctionParameterClause(1)).(string)
	if ctx.GetT().GetTokenType() == parser.LiteLexerRight_Arrow {
		obj = Func + "(" + in + ")" + "(" + out + ")"
	} else {
		obj = Func + "(" + in + ")" + "(" + out + ")"
	}
	return obj
}

func (sf *LiteVisitor) VisitTypeAny(ctx *parser.TypeAnyContext) interface{} {
	return Any
}

func (sf *LiteVisitor) VisitTypeFunctionParameterClause(ctx *parser.TypeFunctionParameterClauseContext) interface{} {
	obj := ""
	for i := 0; i <= len(ctx.AllTypeType())-1; i++ {
		p := sf.Visit(ctx.TypeType(i)).(string)
		if i == 0 {
			obj += p
		} else {
			obj += ", " + p
		}
	}
	return obj
}

func (sf *LiteVisitor) VisitTypeBasic(ctx *parser.TypeBasicContext) interface{} {
	obj := ""
	if ctx.GetT().GetTokenType() == parser.LiteLexerTypeI8 {
		obj = I8
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeU8 {
		obj = U8
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeI16 {
		obj = I16
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeU16 {
		obj = U16
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeI32 {
		obj = I32
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeU32 {
		obj = U32
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeI64 {
		obj = I64
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeU64 {
		obj = U64
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeF32 {
		obj = F32
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeF64 {
		obj = F64
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeChr {
		obj = Chr
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeStr {
		obj = Str
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeBool {
		obj = Bool
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeInt {
		obj = Int
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeNum {
		obj = Num
	} else if ctx.GetT().GetTokenType() == parser.LiteLexerTypeByte {
		obj = Byte
	} else {
		obj = Any
	}
	return obj
}
