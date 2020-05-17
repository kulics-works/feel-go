// Code generated from FeelParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FeelParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFeelParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFeelParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitNamespaceSupportStatement(ctx *NamespaceSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitImportSubStatement(ctx *ImportSubStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeAliasStatement(ctx *TypeAliasStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeRedefineStatement(ctx *TypeRedefineStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeTagStatement(ctx *TypeTagStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitEnumStatement(ctx *EnumStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitEnumSupportStatement(ctx *EnumSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitNamespaceVariableStatement(ctx *NamespaceVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitNamespaceConstantStatement(ctx *NamespaceConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitNamespaceFunctionStatement(ctx *NamespaceFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageStatement(ctx *PackageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageStaticStatement(ctx *PackageStaticStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageStaticSupportStatement(ctx *PackageStaticSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageStaticVariableStatement(ctx *PackageStaticVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageStaticConstantStatement(ctx *PackageStaticConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageStaticFunctionStatement(ctx *PackageStaticFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageFieldStatement(ctx *PackageFieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageSupportStatement(ctx *PackageSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitIncludeStatement(ctx *IncludeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageNewStatement(ctx *PackageNewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageVariableStatement(ctx *PackageVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageConstantStatement(ctx *PackageConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageFunctionStatement(ctx *PackageFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageControlSubStatement(ctx *PackageControlSubStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPackageEventStatement(ctx *PackageEventStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitImplementStatement(ctx *ImplementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitOverrideVariableStatement(ctx *OverrideVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitOverrideConstantStatement(ctx *OverrideConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitOverrideFunctionStatement(ctx *OverrideFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitProtocolStatement(ctx *ProtocolStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitProtocolSubStatement(ctx *ProtocolSubStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitProtocolSupportStatement(ctx *ProtocolSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitProtocolVariableStatement(ctx *ProtocolVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitProtocolFunctionStatement(ctx *ProtocolFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitFunctionStatement(ctx *FunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitReturnAsyncStatement(ctx *ReturnAsyncStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitYieldReturnStatement(ctx *YieldReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitYieldBreakStatement(ctx *YieldBreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitParameterClauseIn(ctx *ParameterClauseInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitParameterClauseOut(ctx *ParameterClauseOutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitFunctionSupportStatement(ctx *FunctionSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeCaseStatement(ctx *JudgeCaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCaseExprStatement(ctx *CaseExprStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeStatement(ctx *JudgeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeElseStatement(ctx *JudgeElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeIfStatement(ctx *JudgeIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeElseIfStatement(ctx *JudgeElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopEachStatement(ctx *LoopEachStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopCaseStatement(ctx *LoopCaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopElseStatement(ctx *LoopElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopJumpStatement(ctx *LoopJumpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopContinueStatement(ctx *LoopContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCheckStatement(ctx *CheckStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitUsingStatement(ctx *UsingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCheckErrorStatement(ctx *CheckErrorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCheckFinallyStatment(ctx *CheckFinallyStatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCheckReportStatement(ctx *CheckReportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitIteratorStatement(ctx *IteratorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitVariableDeclaredStatement(ctx *VariableDeclaredStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitConstantDeclaredStatement(ctx *ConstantDeclaredStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitVarStatement(ctx *VarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBindStatement(ctx *BindStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAnnotationStatement(ctx *AnnotationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitVarId(ctx *VarIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitConstId(ctx *ConstIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAnnotationSupport(ctx *AnnotationSupportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAnnotationList(ctx *AnnotationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAnnotationItem(ctx *AnnotationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAnnotationString(ctx *AnnotationStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCallFunc(ctx *CallFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCallAsync(ctx *CallAsyncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCallAwait(ctx *CallAwaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCallChannel(ctx *CallChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTransfer(ctx *TransferContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCallElement(ctx *CallElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCallPkg(ctx *CallPkgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCallNew(ctx *CallNewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitOrElse(ctx *OrElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeConversion(ctx *TypeConversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeCheck(ctx *TypeCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPkgAssign(ctx *PkgAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPkgAssignElement(ctx *PkgAssignElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitListAssign(ctx *ListAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitDictionaryAssign(ctx *DictionaryAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitDictionary(ctx *DictionaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitDictionaryElement(ctx *DictionaryElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitSlice(ctx *SliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitSliceFull(ctx *SliceFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitSliceStart(ctx *SliceStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitSliceEnd(ctx *SliceEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitNameSpaceItem(ctx *NameSpaceItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTemplateDefine(ctx *TemplateDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTemplateDefineItem(ctx *TemplateDefineItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTemplateCall(ctx *TemplateCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLambdaIn(ctx *LambdaInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPkgAnonymous(ctx *PkgAnonymousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPkgAnonymousAssign(ctx *PkgAnonymousAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPkgAnonymousAssignElement(ctx *PkgAnonymousAssignElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPlusMinus(ctx *PlusMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitNegate(ctx *NegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBitwiseNotExpression(ctx *BitwiseNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLinq(ctx *LinqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLinqHeadItem(ctx *LinqHeadItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLinqItem(ctx *LinqItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeExpression(ctx *JudgeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeExpressionElseStatement(ctx *JudgeExpressionElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeExpressionIfStatement(ctx *JudgeExpressionIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeExpressionElseIfStatement(ctx *JudgeExpressionElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitJudgeCaseExpression(ctx *JudgeCaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCaseExpressionStatement(ctx *CaseExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopExpression(ctx *LoopExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopEachExpression(ctx *LoopEachExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLoopElseExpression(ctx *LoopElseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCheckExpression(ctx *CheckExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCheckErrorExpression(ctx *CheckErrorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitDataStatement(ctx *DataStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitStringContent(ctx *StringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitStringTemplate(ctx *StringTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitRawStringExpr(ctx *RawStringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitRawStringContent(ctx *RawStringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitRawStringTemplate(ctx *RawStringTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitIntegerExpr(ctx *IntegerExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeNotNull(ctx *TypeNotNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeType(ctx *TypeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeNullable(ctx *TypeNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeArray(ctx *TypeArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeSet(ctx *TypeSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeDictionary(ctx *TypeDictionaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeStack(ctx *TypeStackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeQueue(ctx *TypeQueueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeChannel(ctx *TypeChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypePackage(ctx *TypePackageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeFunction(ctx *TypeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeAny(ctx *TypeAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeFunctionParameterClause(ctx *TypeFunctionParameterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitTypeBasic(ctx *TypeBasicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBitwise(ctx *BitwiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBitwiseAnd(ctx *BitwiseAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBitwiseOr(ctx *BitwiseOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBitwiseNot(ctx *BitwiseNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBitwiseXor(ctx *BitwiseXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBitwiseLeftShift(ctx *BitwiseLeftShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitBitwiseRightShift(ctx *BitwiseRightShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCompareCombine(ctx *CompareCombineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCompare(ctx *CompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLogic(ctx *LogicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitAdd(ctx *AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitMul(ctx *MulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitPow(ctx *PowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitWave(ctx *WaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitIdItem(ctx *IdItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitEnd(ctx *EndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitMore(ctx *MoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLeft_brace(ctx *Left_braceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitRight_brace(ctx *Right_braceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLeft_paren(ctx *Left_parenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitRight_paren(ctx *Right_parenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitLeft_brack(ctx *Left_brackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeelParserVisitor) VisitRight_brack(ctx *Right_brackContext) interface{} {
	return v.VisitChildren(ctx)
}
