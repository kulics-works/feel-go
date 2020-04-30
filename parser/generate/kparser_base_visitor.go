// Code generated from KParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // KParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseKParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseKParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitNamespaceSupportStatement(ctx *NamespaceSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitImportSubStatement(ctx *ImportSubStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeAliasStatement(ctx *TypeAliasStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeRedefineStatement(ctx *TypeRedefineStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeTagStatement(ctx *TypeTagStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitEnumStatement(ctx *EnumStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitEnumSupportStatement(ctx *EnumSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitNamespaceVariableStatement(ctx *NamespaceVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitNamespaceConstantStatement(ctx *NamespaceConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitNamespaceFunctionStatement(ctx *NamespaceFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageStatement(ctx *PackageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageStaticStatement(ctx *PackageStaticStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageStaticSupportStatement(ctx *PackageStaticSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageStaticVariableStatement(ctx *PackageStaticVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageStaticConstantStatement(ctx *PackageStaticConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageStaticFunctionStatement(ctx *PackageStaticFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageFieldStatement(ctx *PackageFieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageSupportStatement(ctx *PackageSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitIncludeStatement(ctx *IncludeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageNewStatement(ctx *PackageNewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageVariableStatement(ctx *PackageVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageConstantStatement(ctx *PackageConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageFunctionStatement(ctx *PackageFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageControlSubStatement(ctx *PackageControlSubStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPackageEventStatement(ctx *PackageEventStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitImplementStatement(ctx *ImplementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitOverrideVariableStatement(ctx *OverrideVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitOverrideConstantStatement(ctx *OverrideConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitOverrideFunctionStatement(ctx *OverrideFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitProtocolStatement(ctx *ProtocolStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitProtocolSubStatement(ctx *ProtocolSubStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitProtocolSupportStatement(ctx *ProtocolSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitProtocolVariableStatement(ctx *ProtocolVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitProtocolFunctionStatement(ctx *ProtocolFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitFunctionStatement(ctx *FunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitReturnAsyncStatement(ctx *ReturnAsyncStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitYieldReturnStatement(ctx *YieldReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitYieldBreakStatement(ctx *YieldBreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitParameterClauseIn(ctx *ParameterClauseInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitParameterClauseOut(ctx *ParameterClauseOutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitFunctionSupportStatement(ctx *FunctionSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeCaseStatement(ctx *JudgeCaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCaseExprStatement(ctx *CaseExprStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeStatement(ctx *JudgeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeElseStatement(ctx *JudgeElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeIfStatement(ctx *JudgeIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeElseIfStatement(ctx *JudgeElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopEachStatement(ctx *LoopEachStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopCaseStatement(ctx *LoopCaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopElseStatement(ctx *LoopElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopJumpStatement(ctx *LoopJumpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopContinueStatement(ctx *LoopContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCheckStatement(ctx *CheckStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitUsingStatement(ctx *UsingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCheckErrorStatement(ctx *CheckErrorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCheckFinallyStatment(ctx *CheckFinallyStatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCheckReportStatement(ctx *CheckReportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitIteratorStatement(ctx *IteratorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitVariableDeclaredStatement(ctx *VariableDeclaredStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitConstantDeclaredStatement(ctx *ConstantDeclaredStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitVarStatement(ctx *VarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBindStatement(ctx *BindStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAnnotationStatement(ctx *AnnotationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitVarId(ctx *VarIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitConstId(ctx *ConstIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAnnotationSupport(ctx *AnnotationSupportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAnnotationList(ctx *AnnotationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAnnotationItem(ctx *AnnotationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAnnotationString(ctx *AnnotationStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCallFunc(ctx *CallFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCallAsync(ctx *CallAsyncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCallAwait(ctx *CallAwaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCallChannel(ctx *CallChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTransfer(ctx *TransferContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCallElement(ctx *CallElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCallPkg(ctx *CallPkgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCallNew(ctx *CallNewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitOrElse(ctx *OrElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeConversion(ctx *TypeConversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeCheck(ctx *TypeCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPkgAssign(ctx *PkgAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPkgAssignElement(ctx *PkgAssignElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitListAssign(ctx *ListAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitDictionaryAssign(ctx *DictionaryAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitDictionary(ctx *DictionaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitDictionaryElement(ctx *DictionaryElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitSlice(ctx *SliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitSliceFull(ctx *SliceFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitSliceStart(ctx *SliceStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitSliceEnd(ctx *SliceEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitNameSpaceItem(ctx *NameSpaceItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTemplateDefine(ctx *TemplateDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTemplateDefineItem(ctx *TemplateDefineItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTemplateCall(ctx *TemplateCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLambdaIn(ctx *LambdaInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPkgAnonymous(ctx *PkgAnonymousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPkgAnonymousAssign(ctx *PkgAnonymousAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPkgAnonymousAssignElement(ctx *PkgAnonymousAssignElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPlusMinus(ctx *PlusMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitNegate(ctx *NegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBitwiseNotExpression(ctx *BitwiseNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLinq(ctx *LinqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLinqHeadItem(ctx *LinqHeadItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLinqItem(ctx *LinqItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeExpression(ctx *JudgeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeExpressionElseStatement(ctx *JudgeExpressionElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeExpressionIfStatement(ctx *JudgeExpressionIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeExpressionElseIfStatement(ctx *JudgeExpressionElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitJudgeCaseExpression(ctx *JudgeCaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCaseExpressionStatement(ctx *CaseExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopExpression(ctx *LoopExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopEachExpression(ctx *LoopEachExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLoopElseExpression(ctx *LoopElseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCheckExpression(ctx *CheckExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCheckErrorExpression(ctx *CheckErrorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitDataStatement(ctx *DataStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitStringContent(ctx *StringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitStringTemplate(ctx *StringTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitRawStringExpr(ctx *RawStringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitRawStringContent(ctx *RawStringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitRawStringTemplate(ctx *RawStringTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitIntegerExpr(ctx *IntegerExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeNotNull(ctx *TypeNotNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeType(ctx *TypeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeNullable(ctx *TypeNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeArray(ctx *TypeArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeSet(ctx *TypeSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeDictionary(ctx *TypeDictionaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeStack(ctx *TypeStackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeQueue(ctx *TypeQueueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeChannel(ctx *TypeChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypePackage(ctx *TypePackageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeFunction(ctx *TypeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeAny(ctx *TypeAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeFunctionParameterClause(ctx *TypeFunctionParameterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitTypeBasic(ctx *TypeBasicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBitwise(ctx *BitwiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBitwiseAnd(ctx *BitwiseAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBitwiseOr(ctx *BitwiseOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBitwiseNot(ctx *BitwiseNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBitwiseXor(ctx *BitwiseXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBitwiseLeftShift(ctx *BitwiseLeftShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitBitwiseRightShift(ctx *BitwiseRightShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCompareCombine(ctx *CompareCombineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCompare(ctx *CompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLogic(ctx *LogicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitAdd(ctx *AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitMul(ctx *MulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitPow(ctx *PowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitWave(ctx *WaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitIdItem(ctx *IdItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitEnd(ctx *EndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitMore(ctx *MoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLeft_brace(ctx *Left_braceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitRight_brace(ctx *Right_braceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLeft_paren(ctx *Left_parenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitRight_paren(ctx *Right_parenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitLeft_brack(ctx *Left_brackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKParserVisitor) VisitRight_brack(ctx *Right_brackContext) interface{} {
	return v.VisitChildren(ctx)
}
