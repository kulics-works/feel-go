// Code generated from LiteParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // LiteParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LiteParser.
type LiteParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LiteParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by LiteParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by LiteParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#namespaceSupportStatement.
	VisitNamespaceSupportStatement(ctx *NamespaceSupportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#typeAliasStatement.
	VisitTypeAliasStatement(ctx *TypeAliasStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#typeRedefineStatement.
	VisitTypeRedefineStatement(ctx *TypeRedefineStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#enumStatement.
	VisitEnumStatement(ctx *EnumStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#enumSupportStatement.
	VisitEnumSupportStatement(ctx *EnumSupportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#namespaceVariableStatement.
	VisitNamespaceVariableStatement(ctx *NamespaceVariableStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#namespaceControlStatement.
	VisitNamespaceControlStatement(ctx *NamespaceControlStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#namespaceConstantStatement.
	VisitNamespaceConstantStatement(ctx *NamespaceConstantStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#namespaceFunctionStatement.
	VisitNamespaceFunctionStatement(ctx *NamespaceFunctionStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#packageStatement.
	VisitPackageStatement(ctx *PackageStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#packageSupportStatement.
	VisitPackageSupportStatement(ctx *PackageSupportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#includeStatement.
	VisitIncludeStatement(ctx *IncludeStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#packageNewStatement.
	VisitPackageNewStatement(ctx *PackageNewStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#packageVariableStatement.
	VisitPackageVariableStatement(ctx *PackageVariableStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#packageControlSubStatement.
	VisitPackageControlSubStatement(ctx *PackageControlSubStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#packageEventStatement.
	VisitPackageEventStatement(ctx *PackageEventStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#implementStatement.
	VisitImplementStatement(ctx *ImplementStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#implementSupportStatement.
	VisitImplementSupportStatement(ctx *ImplementSupportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#implementFunctionStatement.
	VisitImplementFunctionStatement(ctx *ImplementFunctionStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#implementControlStatement.
	VisitImplementControlStatement(ctx *ImplementControlStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#overrideStatement.
	VisitOverrideStatement(ctx *OverrideStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#overrideSupportStatement.
	VisitOverrideSupportStatement(ctx *OverrideSupportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#overrideFunctionStatement.
	VisitOverrideFunctionStatement(ctx *OverrideFunctionStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#overrideControlStatement.
	VisitOverrideControlStatement(ctx *OverrideControlStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#protocolStatement.
	VisitProtocolStatement(ctx *ProtocolStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#protocolSupportStatement.
	VisitProtocolSupportStatement(ctx *ProtocolSupportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#protocolControlStatement.
	VisitProtocolControlStatement(ctx *ProtocolControlStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#protocolControlSubStatement.
	VisitProtocolControlSubStatement(ctx *ProtocolControlSubStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#protocolFunctionStatement.
	VisitProtocolFunctionStatement(ctx *ProtocolFunctionStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#functionStatement.
	VisitFunctionStatement(ctx *FunctionStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#parameterClauseIn.
	VisitParameterClauseIn(ctx *ParameterClauseInContext) interface{}

	// Visit a parse tree produced by LiteParser#parameterClauseOut.
	VisitParameterClauseOut(ctx *ParameterClauseOutContext) interface{}

	// Visit a parse tree produced by LiteParser#parameterClauseSelf.
	VisitParameterClauseSelf(ctx *ParameterClauseSelfContext) interface{}

	// Visit a parse tree produced by LiteParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by LiteParser#functionSupportStatement.
	VisitFunctionSupportStatement(ctx *FunctionSupportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#judgeCaseStatement.
	VisitJudgeCaseStatement(ctx *JudgeCaseStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#caseExprStatement.
	VisitCaseExprStatement(ctx *CaseExprStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#judgeStatement.
	VisitJudgeStatement(ctx *JudgeStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#judgeElseStatement.
	VisitJudgeElseStatement(ctx *JudgeElseStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#judgeIfStatement.
	VisitJudgeIfStatement(ctx *JudgeIfStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#judgeElseIfStatement.
	VisitJudgeElseIfStatement(ctx *JudgeElseIfStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#loopEachStatement.
	VisitLoopEachStatement(ctx *LoopEachStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#loopCaseStatement.
	VisitLoopCaseStatement(ctx *LoopCaseStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#loopInfiniteStatement.
	VisitLoopInfiniteStatement(ctx *LoopInfiniteStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#loopJumpStatement.
	VisitLoopJumpStatement(ctx *LoopJumpStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#loopContinueStatement.
	VisitLoopContinueStatement(ctx *LoopContinueStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#checkStatement.
	VisitCheckStatement(ctx *CheckStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#usingStatement.
	VisitUsingStatement(ctx *UsingStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#checkErrorStatement.
	VisitCheckErrorStatement(ctx *CheckErrorStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#checkFinallyStatment.
	VisitCheckFinallyStatment(ctx *CheckFinallyStatmentContext) interface{}

	// Visit a parse tree produced by LiteParser#reportStatement.
	VisitReportStatement(ctx *ReportStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#iteratorStatement.
	VisitIteratorStatement(ctx *IteratorStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#variableDeclaredStatement.
	VisitVariableDeclaredStatement(ctx *VariableDeclaredStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#channelAssignStatement.
	VisitChannelAssignStatement(ctx *ChannelAssignStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#idExpression.
	VisitIdExpression(ctx *IdExpressionContext) interface{}

	// Visit a parse tree produced by LiteParser#idExprItem.
	VisitIdExprItem(ctx *IdExprItemContext) interface{}

	// Visit a parse tree produced by LiteParser#tupleExpression.
	VisitTupleExpression(ctx *TupleExpressionContext) interface{}

	// Visit a parse tree produced by LiteParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by LiteParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by LiteParser#callExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by LiteParser#tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by LiteParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by LiteParser#annotationSupport.
	VisitAnnotationSupport(ctx *AnnotationSupportContext) interface{}

	// Visit a parse tree produced by LiteParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by LiteParser#annotationList.
	VisitAnnotationList(ctx *AnnotationListContext) interface{}

	// Visit a parse tree produced by LiteParser#annotationItem.
	VisitAnnotationItem(ctx *AnnotationItemContext) interface{}

	// Visit a parse tree produced by LiteParser#annotationAssign.
	VisitAnnotationAssign(ctx *AnnotationAssignContext) interface{}

	// Visit a parse tree produced by LiteParser#callFunc.
	VisitCallFunc(ctx *CallFuncContext) interface{}

	// Visit a parse tree produced by LiteParser#callChannel.
	VisitCallChannel(ctx *CallChannelContext) interface{}

	// Visit a parse tree produced by LiteParser#callElement.
	VisitCallElement(ctx *CallElementContext) interface{}

	// Visit a parse tree produced by LiteParser#callPkg.
	VisitCallPkg(ctx *CallPkgContext) interface{}

	// Visit a parse tree produced by LiteParser#callNew.
	VisitCallNew(ctx *CallNewContext) interface{}

	// Visit a parse tree produced by LiteParser#getType.
	VisitGetType(ctx *GetTypeContext) interface{}

	// Visit a parse tree produced by LiteParser#typeConversion.
	VisitTypeConversion(ctx *TypeConversionContext) interface{}

	// Visit a parse tree produced by LiteParser#pkgAssign.
	VisitPkgAssign(ctx *PkgAssignContext) interface{}

	// Visit a parse tree produced by LiteParser#pkgAssignElement.
	VisitPkgAssignElement(ctx *PkgAssignElementContext) interface{}

	// Visit a parse tree produced by LiteParser#listAssign.
	VisitListAssign(ctx *ListAssignContext) interface{}

	// Visit a parse tree produced by LiteParser#setAssign.
	VisitSetAssign(ctx *SetAssignContext) interface{}

	// Visit a parse tree produced by LiteParser#dictionaryAssign.
	VisitDictionaryAssign(ctx *DictionaryAssignContext) interface{}

	// Visit a parse tree produced by LiteParser#callAwait.
	VisitCallAwait(ctx *CallAwaitContext) interface{}

	// Visit a parse tree produced by LiteParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by LiteParser#set.
	VisitSet(ctx *SetContext) interface{}

	// Visit a parse tree produced by LiteParser#dictionary.
	VisitDictionary(ctx *DictionaryContext) interface{}

	// Visit a parse tree produced by LiteParser#dictionaryElement.
	VisitDictionaryElement(ctx *DictionaryElementContext) interface{}

	// Visit a parse tree produced by LiteParser#slice.
	VisitSlice(ctx *SliceContext) interface{}

	// Visit a parse tree produced by LiteParser#sliceFull.
	VisitSliceFull(ctx *SliceFullContext) interface{}

	// Visit a parse tree produced by LiteParser#sliceStart.
	VisitSliceStart(ctx *SliceStartContext) interface{}

	// Visit a parse tree produced by LiteParser#sliceEnd.
	VisitSliceEnd(ctx *SliceEndContext) interface{}

	// Visit a parse tree produced by LiteParser#nameSpaceItem.
	VisitNameSpaceItem(ctx *NameSpaceItemContext) interface{}

	// Visit a parse tree produced by LiteParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by LiteParser#templateDefine.
	VisitTemplateDefine(ctx *TemplateDefineContext) interface{}

	// Visit a parse tree produced by LiteParser#templateDefineItem.
	VisitTemplateDefineItem(ctx *TemplateDefineItemContext) interface{}

	// Visit a parse tree produced by LiteParser#templateCall.
	VisitTemplateCall(ctx *TemplateCallContext) interface{}

	// Visit a parse tree produced by LiteParser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by LiteParser#lambdaIn.
	VisitLambdaIn(ctx *LambdaInContext) interface{}

	// Visit a parse tree produced by LiteParser#pkgAnonymous.
	VisitPkgAnonymous(ctx *PkgAnonymousContext) interface{}

	// Visit a parse tree produced by LiteParser#pkgAnonymousAssign.
	VisitPkgAnonymousAssign(ctx *PkgAnonymousAssignContext) interface{}

	// Visit a parse tree produced by LiteParser#pkgAnonymousAssignElement.
	VisitPkgAnonymousAssignElement(ctx *PkgAnonymousAssignElementContext) interface{}

	// Visit a parse tree produced by LiteParser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by LiteParser#plusMinus.
	VisitPlusMinus(ctx *PlusMinusContext) interface{}

	// Visit a parse tree produced by LiteParser#negate.
	VisitNegate(ctx *NegateContext) interface{}

	// Visit a parse tree produced by LiteParser#bitwiseNotExpression.
	VisitBitwiseNotExpression(ctx *BitwiseNotExpressionContext) interface{}

	// Visit a parse tree produced by LiteParser#linq.
	VisitLinq(ctx *LinqContext) interface{}

	// Visit a parse tree produced by LiteParser#linqItem.
	VisitLinqItem(ctx *LinqItemContext) interface{}

	// Visit a parse tree produced by LiteParser#linqKeyword.
	VisitLinqKeyword(ctx *LinqKeywordContext) interface{}

	// Visit a parse tree produced by LiteParser#linqHeadKeyword.
	VisitLinqHeadKeyword(ctx *LinqHeadKeywordContext) interface{}

	// Visit a parse tree produced by LiteParser#linqBodyKeyword.
	VisitLinqBodyKeyword(ctx *LinqBodyKeywordContext) interface{}

	// Visit a parse tree produced by LiteParser#stringExpression.
	VisitStringExpression(ctx *StringExpressionContext) interface{}

	// Visit a parse tree produced by LiteParser#stringExpressionElement.
	VisitStringExpressionElement(ctx *StringExpressionElementContext) interface{}

	// Visit a parse tree produced by LiteParser#dataStatement.
	VisitDataStatement(ctx *DataStatementContext) interface{}

	// Visit a parse tree produced by LiteParser#floatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by LiteParser#integerExpr.
	VisitIntegerExpr(ctx *IntegerExprContext) interface{}

	// Visit a parse tree produced by LiteParser#typeNotNull.
	VisitTypeNotNull(ctx *TypeNotNullContext) interface{}

	// Visit a parse tree produced by LiteParser#typeType.
	VisitTypeType(ctx *TypeTypeContext) interface{}

	// Visit a parse tree produced by LiteParser#typeReference.
	VisitTypeReference(ctx *TypeReferenceContext) interface{}

	// Visit a parse tree produced by LiteParser#typeNullable.
	VisitTypeNullable(ctx *TypeNullableContext) interface{}

	// Visit a parse tree produced by LiteParser#typeTuple.
	VisitTypeTuple(ctx *TypeTupleContext) interface{}

	// Visit a parse tree produced by LiteParser#typeArray.
	VisitTypeArray(ctx *TypeArrayContext) interface{}

	// Visit a parse tree produced by LiteParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by LiteParser#typeSet.
	VisitTypeSet(ctx *TypeSetContext) interface{}

	// Visit a parse tree produced by LiteParser#typeDictionary.
	VisitTypeDictionary(ctx *TypeDictionaryContext) interface{}

	// Visit a parse tree produced by LiteParser#typeChannel.
	VisitTypeChannel(ctx *TypeChannelContext) interface{}

	// Visit a parse tree produced by LiteParser#typePackage.
	VisitTypePackage(ctx *TypePackageContext) interface{}

	// Visit a parse tree produced by LiteParser#typeFunction.
	VisitTypeFunction(ctx *TypeFunctionContext) interface{}

	// Visit a parse tree produced by LiteParser#typeAny.
	VisitTypeAny(ctx *TypeAnyContext) interface{}

	// Visit a parse tree produced by LiteParser#typeFunctionParameterClause.
	VisitTypeFunctionParameterClause(ctx *TypeFunctionParameterClauseContext) interface{}

	// Visit a parse tree produced by LiteParser#typeBasic.
	VisitTypeBasic(ctx *TypeBasicContext) interface{}

	// Visit a parse tree produced by LiteParser#nilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by LiteParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by LiteParser#judgeType.
	VisitJudgeType(ctx *JudgeTypeContext) interface{}

	// Visit a parse tree produced by LiteParser#bitwise.
	VisitBitwise(ctx *BitwiseContext) interface{}

	// Visit a parse tree produced by LiteParser#bitwiseAnd.
	VisitBitwiseAnd(ctx *BitwiseAndContext) interface{}

	// Visit a parse tree produced by LiteParser#bitwiseOr.
	VisitBitwiseOr(ctx *BitwiseOrContext) interface{}

	// Visit a parse tree produced by LiteParser#bitwiseNot.
	VisitBitwiseNot(ctx *BitwiseNotContext) interface{}

	// Visit a parse tree produced by LiteParser#bitwiseXor.
	VisitBitwiseXor(ctx *BitwiseXorContext) interface{}

	// Visit a parse tree produced by LiteParser#bitwiseLeftShift.
	VisitBitwiseLeftShift(ctx *BitwiseLeftShiftContext) interface{}

	// Visit a parse tree produced by LiteParser#bitwiseRightShift.
	VisitBitwiseRightShift(ctx *BitwiseRightShiftContext) interface{}

	// Visit a parse tree produced by LiteParser#judge.
	VisitJudge(ctx *JudgeContext) interface{}

	// Visit a parse tree produced by LiteParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by LiteParser#add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by LiteParser#mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by LiteParser#pow.
	VisitPow(ctx *PowContext) interface{}

	// Visit a parse tree produced by LiteParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by LiteParser#wave.
	VisitWave(ctx *WaveContext) interface{}

	// Visit a parse tree produced by LiteParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by LiteParser#idItem.
	VisitIdItem(ctx *IdItemContext) interface{}

	// Visit a parse tree produced by LiteParser#end.
	VisitEnd(ctx *EndContext) interface{}

	// Visit a parse tree produced by LiteParser#more.
	VisitMore(ctx *MoreContext) interface{}

	// Visit a parse tree produced by LiteParser#left_brace.
	VisitLeft_brace(ctx *Left_braceContext) interface{}

	// Visit a parse tree produced by LiteParser#right_brace.
	VisitRight_brace(ctx *Right_braceContext) interface{}

	// Visit a parse tree produced by LiteParser#left_paren.
	VisitLeft_paren(ctx *Left_parenContext) interface{}

	// Visit a parse tree produced by LiteParser#right_paren.
	VisitRight_paren(ctx *Right_parenContext) interface{}

	// Visit a parse tree produced by LiteParser#left_brack.
	VisitLeft_brack(ctx *Left_brackContext) interface{}

	// Visit a parse tree produced by LiteParser#right_brack.
	VisitRight_brack(ctx *Right_brackContext) interface{}
}
