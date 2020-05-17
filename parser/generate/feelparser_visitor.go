// Code generated from FeelParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FeelParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FeelParser.
type FeelParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FeelParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by FeelParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by FeelParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#namespaceSupportStatement.
	VisitNamespaceSupportStatement(ctx *NamespaceSupportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#importSubStatement.
	VisitImportSubStatement(ctx *ImportSubStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#typeAliasStatement.
	VisitTypeAliasStatement(ctx *TypeAliasStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#typeRedefineStatement.
	VisitTypeRedefineStatement(ctx *TypeRedefineStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#typeTagStatement.
	VisitTypeTagStatement(ctx *TypeTagStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#enumStatement.
	VisitEnumStatement(ctx *EnumStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#enumSupportStatement.
	VisitEnumSupportStatement(ctx *EnumSupportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#namespaceVariableStatement.
	VisitNamespaceVariableStatement(ctx *NamespaceVariableStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#namespaceConstantStatement.
	VisitNamespaceConstantStatement(ctx *NamespaceConstantStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#namespaceFunctionStatement.
	VisitNamespaceFunctionStatement(ctx *NamespaceFunctionStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageStatement.
	VisitPackageStatement(ctx *PackageStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageStaticStatement.
	VisitPackageStaticStatement(ctx *PackageStaticStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageStaticSupportStatement.
	VisitPackageStaticSupportStatement(ctx *PackageStaticSupportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageStaticVariableStatement.
	VisitPackageStaticVariableStatement(ctx *PackageStaticVariableStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageStaticConstantStatement.
	VisitPackageStaticConstantStatement(ctx *PackageStaticConstantStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageStaticFunctionStatement.
	VisitPackageStaticFunctionStatement(ctx *PackageStaticFunctionStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageFieldStatement.
	VisitPackageFieldStatement(ctx *PackageFieldStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageSupportStatement.
	VisitPackageSupportStatement(ctx *PackageSupportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#includeStatement.
	VisitIncludeStatement(ctx *IncludeStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageNewStatement.
	VisitPackageNewStatement(ctx *PackageNewStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageVariableStatement.
	VisitPackageVariableStatement(ctx *PackageVariableStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageConstantStatement.
	VisitPackageConstantStatement(ctx *PackageConstantStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageFunctionStatement.
	VisitPackageFunctionStatement(ctx *PackageFunctionStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageControlSubStatement.
	VisitPackageControlSubStatement(ctx *PackageControlSubStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#packageEventStatement.
	VisitPackageEventStatement(ctx *PackageEventStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#implementStatement.
	VisitImplementStatement(ctx *ImplementStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#overrideVariableStatement.
	VisitOverrideVariableStatement(ctx *OverrideVariableStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#overrideConstantStatement.
	VisitOverrideConstantStatement(ctx *OverrideConstantStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#overrideFunctionStatement.
	VisitOverrideFunctionStatement(ctx *OverrideFunctionStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#protocolStatement.
	VisitProtocolStatement(ctx *ProtocolStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#protocolSubStatement.
	VisitProtocolSubStatement(ctx *ProtocolSubStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#protocolSupportStatement.
	VisitProtocolSupportStatement(ctx *ProtocolSupportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#protocolVariableStatement.
	VisitProtocolVariableStatement(ctx *ProtocolVariableStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#protocolFunctionStatement.
	VisitProtocolFunctionStatement(ctx *ProtocolFunctionStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#functionStatement.
	VisitFunctionStatement(ctx *FunctionStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#returnAsyncStatement.
	VisitReturnAsyncStatement(ctx *ReturnAsyncStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#yieldReturnStatement.
	VisitYieldReturnStatement(ctx *YieldReturnStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#yieldBreakStatement.
	VisitYieldBreakStatement(ctx *YieldBreakStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#parameterClauseIn.
	VisitParameterClauseIn(ctx *ParameterClauseInContext) interface{}

	// Visit a parse tree produced by FeelParser#parameterClauseOut.
	VisitParameterClauseOut(ctx *ParameterClauseOutContext) interface{}

	// Visit a parse tree produced by FeelParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by FeelParser#functionSupportStatement.
	VisitFunctionSupportStatement(ctx *FunctionSupportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeCaseStatement.
	VisitJudgeCaseStatement(ctx *JudgeCaseStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#caseExprStatement.
	VisitCaseExprStatement(ctx *CaseExprStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeStatement.
	VisitJudgeStatement(ctx *JudgeStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeElseStatement.
	VisitJudgeElseStatement(ctx *JudgeElseStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeIfStatement.
	VisitJudgeIfStatement(ctx *JudgeIfStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeElseIfStatement.
	VisitJudgeElseIfStatement(ctx *JudgeElseIfStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#loopEachStatement.
	VisitLoopEachStatement(ctx *LoopEachStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#loopCaseStatement.
	VisitLoopCaseStatement(ctx *LoopCaseStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#loopElseStatement.
	VisitLoopElseStatement(ctx *LoopElseStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#loopJumpStatement.
	VisitLoopJumpStatement(ctx *LoopJumpStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#loopContinueStatement.
	VisitLoopContinueStatement(ctx *LoopContinueStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#checkStatement.
	VisitCheckStatement(ctx *CheckStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#usingStatement.
	VisitUsingStatement(ctx *UsingStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#checkErrorStatement.
	VisitCheckErrorStatement(ctx *CheckErrorStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#checkFinallyStatment.
	VisitCheckFinallyStatment(ctx *CheckFinallyStatmentContext) interface{}

	// Visit a parse tree produced by FeelParser#checkReportStatement.
	VisitCheckReportStatement(ctx *CheckReportStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#iteratorStatement.
	VisitIteratorStatement(ctx *IteratorStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#variableDeclaredStatement.
	VisitVariableDeclaredStatement(ctx *VariableDeclaredStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#constantDeclaredStatement.
	VisitConstantDeclaredStatement(ctx *ConstantDeclaredStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#varStatement.
	VisitVarStatement(ctx *VarStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#bindStatement.
	VisitBindStatement(ctx *BindStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#annotationStatement.
	VisitAnnotationStatement(ctx *AnnotationStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#varId.
	VisitVarId(ctx *VarIdContext) interface{}

	// Visit a parse tree produced by FeelParser#constId.
	VisitConstId(ctx *ConstIdContext) interface{}

	// Visit a parse tree produced by FeelParser#tupleExpression.
	VisitTupleExpression(ctx *TupleExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#callExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by FeelParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by FeelParser#annotationSupport.
	VisitAnnotationSupport(ctx *AnnotationSupportContext) interface{}

	// Visit a parse tree produced by FeelParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by FeelParser#annotationList.
	VisitAnnotationList(ctx *AnnotationListContext) interface{}

	// Visit a parse tree produced by FeelParser#annotationItem.
	VisitAnnotationItem(ctx *AnnotationItemContext) interface{}

	// Visit a parse tree produced by FeelParser#annotationString.
	VisitAnnotationString(ctx *AnnotationStringContext) interface{}

	// Visit a parse tree produced by FeelParser#callFunc.
	VisitCallFunc(ctx *CallFuncContext) interface{}

	// Visit a parse tree produced by FeelParser#callAsync.
	VisitCallAsync(ctx *CallAsyncContext) interface{}

	// Visit a parse tree produced by FeelParser#callAwait.
	VisitCallAwait(ctx *CallAwaitContext) interface{}

	// Visit a parse tree produced by FeelParser#callChannel.
	VisitCallChannel(ctx *CallChannelContext) interface{}

	// Visit a parse tree produced by FeelParser#transfer.
	VisitTransfer(ctx *TransferContext) interface{}

	// Visit a parse tree produced by FeelParser#callElement.
	VisitCallElement(ctx *CallElementContext) interface{}

	// Visit a parse tree produced by FeelParser#callPkg.
	VisitCallPkg(ctx *CallPkgContext) interface{}

	// Visit a parse tree produced by FeelParser#callNew.
	VisitCallNew(ctx *CallNewContext) interface{}

	// Visit a parse tree produced by FeelParser#orElse.
	VisitOrElse(ctx *OrElseContext) interface{}

	// Visit a parse tree produced by FeelParser#typeConversion.
	VisitTypeConversion(ctx *TypeConversionContext) interface{}

	// Visit a parse tree produced by FeelParser#typeCheck.
	VisitTypeCheck(ctx *TypeCheckContext) interface{}

	// Visit a parse tree produced by FeelParser#pkgAssign.
	VisitPkgAssign(ctx *PkgAssignContext) interface{}

	// Visit a parse tree produced by FeelParser#pkgAssignElement.
	VisitPkgAssignElement(ctx *PkgAssignElementContext) interface{}

	// Visit a parse tree produced by FeelParser#listAssign.
	VisitListAssign(ctx *ListAssignContext) interface{}

	// Visit a parse tree produced by FeelParser#dictionaryAssign.
	VisitDictionaryAssign(ctx *DictionaryAssignContext) interface{}

	// Visit a parse tree produced by FeelParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by FeelParser#dictionary.
	VisitDictionary(ctx *DictionaryContext) interface{}

	// Visit a parse tree produced by FeelParser#dictionaryElement.
	VisitDictionaryElement(ctx *DictionaryElementContext) interface{}

	// Visit a parse tree produced by FeelParser#slice.
	VisitSlice(ctx *SliceContext) interface{}

	// Visit a parse tree produced by FeelParser#sliceFull.
	VisitSliceFull(ctx *SliceFullContext) interface{}

	// Visit a parse tree produced by FeelParser#sliceStart.
	VisitSliceStart(ctx *SliceStartContext) interface{}

	// Visit a parse tree produced by FeelParser#sliceEnd.
	VisitSliceEnd(ctx *SliceEndContext) interface{}

	// Visit a parse tree produced by FeelParser#nameSpaceItem.
	VisitNameSpaceItem(ctx *NameSpaceItemContext) interface{}

	// Visit a parse tree produced by FeelParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by FeelParser#templateDefine.
	VisitTemplateDefine(ctx *TemplateDefineContext) interface{}

	// Visit a parse tree produced by FeelParser#templateDefineItem.
	VisitTemplateDefineItem(ctx *TemplateDefineItemContext) interface{}

	// Visit a parse tree produced by FeelParser#templateCall.
	VisitTemplateCall(ctx *TemplateCallContext) interface{}

	// Visit a parse tree produced by FeelParser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by FeelParser#lambdaIn.
	VisitLambdaIn(ctx *LambdaInContext) interface{}

	// Visit a parse tree produced by FeelParser#pkgAnonymous.
	VisitPkgAnonymous(ctx *PkgAnonymousContext) interface{}

	// Visit a parse tree produced by FeelParser#pkgAnonymousAssign.
	VisitPkgAnonymousAssign(ctx *PkgAnonymousAssignContext) interface{}

	// Visit a parse tree produced by FeelParser#pkgAnonymousAssignElement.
	VisitPkgAnonymousAssignElement(ctx *PkgAnonymousAssignElementContext) interface{}

	// Visit a parse tree produced by FeelParser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#plusMinus.
	VisitPlusMinus(ctx *PlusMinusContext) interface{}

	// Visit a parse tree produced by FeelParser#negate.
	VisitNegate(ctx *NegateContext) interface{}

	// Visit a parse tree produced by FeelParser#bitwiseNotExpression.
	VisitBitwiseNotExpression(ctx *BitwiseNotExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#linq.
	VisitLinq(ctx *LinqContext) interface{}

	// Visit a parse tree produced by FeelParser#linqHeadItem.
	VisitLinqHeadItem(ctx *LinqHeadItemContext) interface{}

	// Visit a parse tree produced by FeelParser#linqItem.
	VisitLinqItem(ctx *LinqItemContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeExpression.
	VisitJudgeExpression(ctx *JudgeExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeExpressionElseStatement.
	VisitJudgeExpressionElseStatement(ctx *JudgeExpressionElseStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeExpressionIfStatement.
	VisitJudgeExpressionIfStatement(ctx *JudgeExpressionIfStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeExpressionElseIfStatement.
	VisitJudgeExpressionElseIfStatement(ctx *JudgeExpressionElseIfStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#judgeCaseExpression.
	VisitJudgeCaseExpression(ctx *JudgeCaseExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#caseExpressionStatement.
	VisitCaseExpressionStatement(ctx *CaseExpressionStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#loopExpression.
	VisitLoopExpression(ctx *LoopExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#loopEachExpression.
	VisitLoopEachExpression(ctx *LoopEachExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#loopElseExpression.
	VisitLoopElseExpression(ctx *LoopElseExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#checkExpression.
	VisitCheckExpression(ctx *CheckExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#checkErrorExpression.
	VisitCheckErrorExpression(ctx *CheckErrorExpressionContext) interface{}

	// Visit a parse tree produced by FeelParser#dataStatement.
	VisitDataStatement(ctx *DataStatementContext) interface{}

	// Visit a parse tree produced by FeelParser#stringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by FeelParser#stringContent.
	VisitStringContent(ctx *StringContentContext) interface{}

	// Visit a parse tree produced by FeelParser#stringTemplate.
	VisitStringTemplate(ctx *StringTemplateContext) interface{}

	// Visit a parse tree produced by FeelParser#rawStringExpr.
	VisitRawStringExpr(ctx *RawStringExprContext) interface{}

	// Visit a parse tree produced by FeelParser#rawStringContent.
	VisitRawStringContent(ctx *RawStringContentContext) interface{}

	// Visit a parse tree produced by FeelParser#rawStringTemplate.
	VisitRawStringTemplate(ctx *RawStringTemplateContext) interface{}

	// Visit a parse tree produced by FeelParser#floatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by FeelParser#integerExpr.
	VisitIntegerExpr(ctx *IntegerExprContext) interface{}

	// Visit a parse tree produced by FeelParser#typeNotNull.
	VisitTypeNotNull(ctx *TypeNotNullContext) interface{}

	// Visit a parse tree produced by FeelParser#typeType.
	VisitTypeType(ctx *TypeTypeContext) interface{}

	// Visit a parse tree produced by FeelParser#typeNullable.
	VisitTypeNullable(ctx *TypeNullableContext) interface{}

	// Visit a parse tree produced by FeelParser#typeArray.
	VisitTypeArray(ctx *TypeArrayContext) interface{}

	// Visit a parse tree produced by FeelParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by FeelParser#typeSet.
	VisitTypeSet(ctx *TypeSetContext) interface{}

	// Visit a parse tree produced by FeelParser#typeDictionary.
	VisitTypeDictionary(ctx *TypeDictionaryContext) interface{}

	// Visit a parse tree produced by FeelParser#typeStack.
	VisitTypeStack(ctx *TypeStackContext) interface{}

	// Visit a parse tree produced by FeelParser#typeQueue.
	VisitTypeQueue(ctx *TypeQueueContext) interface{}

	// Visit a parse tree produced by FeelParser#typeChannel.
	VisitTypeChannel(ctx *TypeChannelContext) interface{}

	// Visit a parse tree produced by FeelParser#typePackage.
	VisitTypePackage(ctx *TypePackageContext) interface{}

	// Visit a parse tree produced by FeelParser#typeFunction.
	VisitTypeFunction(ctx *TypeFunctionContext) interface{}

	// Visit a parse tree produced by FeelParser#typeAny.
	VisitTypeAny(ctx *TypeAnyContext) interface{}

	// Visit a parse tree produced by FeelParser#typeFunctionParameterClause.
	VisitTypeFunctionParameterClause(ctx *TypeFunctionParameterClauseContext) interface{}

	// Visit a parse tree produced by FeelParser#typeBasic.
	VisitTypeBasic(ctx *TypeBasicContext) interface{}

	// Visit a parse tree produced by FeelParser#nilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by FeelParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by FeelParser#bitwise.
	VisitBitwise(ctx *BitwiseContext) interface{}

	// Visit a parse tree produced by FeelParser#bitwiseAnd.
	VisitBitwiseAnd(ctx *BitwiseAndContext) interface{}

	// Visit a parse tree produced by FeelParser#bitwiseOr.
	VisitBitwiseOr(ctx *BitwiseOrContext) interface{}

	// Visit a parse tree produced by FeelParser#bitwiseNot.
	VisitBitwiseNot(ctx *BitwiseNotContext) interface{}

	// Visit a parse tree produced by FeelParser#bitwiseXor.
	VisitBitwiseXor(ctx *BitwiseXorContext) interface{}

	// Visit a parse tree produced by FeelParser#bitwiseLeftShift.
	VisitBitwiseLeftShift(ctx *BitwiseLeftShiftContext) interface{}

	// Visit a parse tree produced by FeelParser#bitwiseRightShift.
	VisitBitwiseRightShift(ctx *BitwiseRightShiftContext) interface{}

	// Visit a parse tree produced by FeelParser#compareCombine.
	VisitCompareCombine(ctx *CompareCombineContext) interface{}

	// Visit a parse tree produced by FeelParser#compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by FeelParser#logic.
	VisitLogic(ctx *LogicContext) interface{}

	// Visit a parse tree produced by FeelParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by FeelParser#add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by FeelParser#mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by FeelParser#pow.
	VisitPow(ctx *PowContext) interface{}

	// Visit a parse tree produced by FeelParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by FeelParser#wave.
	VisitWave(ctx *WaveContext) interface{}

	// Visit a parse tree produced by FeelParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by FeelParser#idItem.
	VisitIdItem(ctx *IdItemContext) interface{}

	// Visit a parse tree produced by FeelParser#end.
	VisitEnd(ctx *EndContext) interface{}

	// Visit a parse tree produced by FeelParser#more.
	VisitMore(ctx *MoreContext) interface{}

	// Visit a parse tree produced by FeelParser#left_brace.
	VisitLeft_brace(ctx *Left_braceContext) interface{}

	// Visit a parse tree produced by FeelParser#right_brace.
	VisitRight_brace(ctx *Right_braceContext) interface{}

	// Visit a parse tree produced by FeelParser#left_paren.
	VisitLeft_paren(ctx *Left_parenContext) interface{}

	// Visit a parse tree produced by FeelParser#right_paren.
	VisitRight_paren(ctx *Right_parenContext) interface{}

	// Visit a parse tree produced by FeelParser#left_brack.
	VisitLeft_brack(ctx *Left_brackContext) interface{}

	// Visit a parse tree produced by FeelParser#right_brack.
	VisitRight_brack(ctx *Right_brackContext) interface{}
}
