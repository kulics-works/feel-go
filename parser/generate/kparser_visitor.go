// Code generated from KParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // KParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by KParser.
type KParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by KParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by KParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by KParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by KParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by KParser#namespaceSupportStatement.
	VisitNamespaceSupportStatement(ctx *NamespaceSupportStatementContext) interface{}

	// Visit a parse tree produced by KParser#typeAliasStatement.
	VisitTypeAliasStatement(ctx *TypeAliasStatementContext) interface{}

	// Visit a parse tree produced by KParser#typeRedefineStatement.
	VisitTypeRedefineStatement(ctx *TypeRedefineStatementContext) interface{}

	// Visit a parse tree produced by KParser#typeTagStatement.
	VisitTypeTagStatement(ctx *TypeTagStatementContext) interface{}

	// Visit a parse tree produced by KParser#enumStatement.
	VisitEnumStatement(ctx *EnumStatementContext) interface{}

	// Visit a parse tree produced by KParser#enumSupportStatement.
	VisitEnumSupportStatement(ctx *EnumSupportStatementContext) interface{}

	// Visit a parse tree produced by KParser#namespaceVariableStatement.
	VisitNamespaceVariableStatement(ctx *NamespaceVariableStatementContext) interface{}

	// Visit a parse tree produced by KParser#namespaceConstantStatement.
	VisitNamespaceConstantStatement(ctx *NamespaceConstantStatementContext) interface{}

	// Visit a parse tree produced by KParser#namespaceFunctionStatement.
	VisitNamespaceFunctionStatement(ctx *NamespaceFunctionStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageStatement.
	VisitPackageStatement(ctx *PackageStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageStaticStatement.
	VisitPackageStaticStatement(ctx *PackageStaticStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageStaticSupportStatement.
	VisitPackageStaticSupportStatement(ctx *PackageStaticSupportStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageStaticVariableStatement.
	VisitPackageStaticVariableStatement(ctx *PackageStaticVariableStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageStaticConstantStatement.
	VisitPackageStaticConstantStatement(ctx *PackageStaticConstantStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageStaticFunctionStatement.
	VisitPackageStaticFunctionStatement(ctx *PackageStaticFunctionStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageFieldStatement.
	VisitPackageFieldStatement(ctx *PackageFieldStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageSupportStatement.
	VisitPackageSupportStatement(ctx *PackageSupportStatementContext) interface{}

	// Visit a parse tree produced by KParser#includeStatement.
	VisitIncludeStatement(ctx *IncludeStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageNewStatement.
	VisitPackageNewStatement(ctx *PackageNewStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageVariableStatement.
	VisitPackageVariableStatement(ctx *PackageVariableStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageConstantStatement.
	VisitPackageConstantStatement(ctx *PackageConstantStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageFunctionStatement.
	VisitPackageFunctionStatement(ctx *PackageFunctionStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageControlSubStatement.
	VisitPackageControlSubStatement(ctx *PackageControlSubStatementContext) interface{}

	// Visit a parse tree produced by KParser#packageEventStatement.
	VisitPackageEventStatement(ctx *PackageEventStatementContext) interface{}

	// Visit a parse tree produced by KParser#implementStatement.
	VisitImplementStatement(ctx *ImplementStatementContext) interface{}

	// Visit a parse tree produced by KParser#overrideVariableStatement.
	VisitOverrideVariableStatement(ctx *OverrideVariableStatementContext) interface{}

	// Visit a parse tree produced by KParser#overrideConstantStatement.
	VisitOverrideConstantStatement(ctx *OverrideConstantStatementContext) interface{}

	// Visit a parse tree produced by KParser#overrideFunctionStatement.
	VisitOverrideFunctionStatement(ctx *OverrideFunctionStatementContext) interface{}

	// Visit a parse tree produced by KParser#protocolStatement.
	VisitProtocolStatement(ctx *ProtocolStatementContext) interface{}

	// Visit a parse tree produced by KParser#protocolSubStatement.
	VisitProtocolSubStatement(ctx *ProtocolSubStatementContext) interface{}

	// Visit a parse tree produced by KParser#protocolSupportStatement.
	VisitProtocolSupportStatement(ctx *ProtocolSupportStatementContext) interface{}

	// Visit a parse tree produced by KParser#protocolVariableStatement.
	VisitProtocolVariableStatement(ctx *ProtocolVariableStatementContext) interface{}

	// Visit a parse tree produced by KParser#protocolFunctionStatement.
	VisitProtocolFunctionStatement(ctx *ProtocolFunctionStatementContext) interface{}

	// Visit a parse tree produced by KParser#functionStatement.
	VisitFunctionStatement(ctx *FunctionStatementContext) interface{}

	// Visit a parse tree produced by KParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by KParser#yieldReturnStatement.
	VisitYieldReturnStatement(ctx *YieldReturnStatementContext) interface{}

	// Visit a parse tree produced by KParser#yieldBreakStatement.
	VisitYieldBreakStatement(ctx *YieldBreakStatementContext) interface{}

	// Visit a parse tree produced by KParser#parameterClauseIn.
	VisitParameterClauseIn(ctx *ParameterClauseInContext) interface{}

	// Visit a parse tree produced by KParser#parameterClauseOut.
	VisitParameterClauseOut(ctx *ParameterClauseOutContext) interface{}

	// Visit a parse tree produced by KParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by KParser#functionSupportStatement.
	VisitFunctionSupportStatement(ctx *FunctionSupportStatementContext) interface{}

	// Visit a parse tree produced by KParser#judgeCaseStatement.
	VisitJudgeCaseStatement(ctx *JudgeCaseStatementContext) interface{}

	// Visit a parse tree produced by KParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by KParser#caseExprStatement.
	VisitCaseExprStatement(ctx *CaseExprStatementContext) interface{}

	// Visit a parse tree produced by KParser#judgeStatement.
	VisitJudgeStatement(ctx *JudgeStatementContext) interface{}

	// Visit a parse tree produced by KParser#judgeElseStatement.
	VisitJudgeElseStatement(ctx *JudgeElseStatementContext) interface{}

	// Visit a parse tree produced by KParser#judgeIfStatement.
	VisitJudgeIfStatement(ctx *JudgeIfStatementContext) interface{}

	// Visit a parse tree produced by KParser#judgeElseIfStatement.
	VisitJudgeElseIfStatement(ctx *JudgeElseIfStatementContext) interface{}

	// Visit a parse tree produced by KParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by KParser#loopEachStatement.
	VisitLoopEachStatement(ctx *LoopEachStatementContext) interface{}

	// Visit a parse tree produced by KParser#loopCaseStatement.
	VisitLoopCaseStatement(ctx *LoopCaseStatementContext) interface{}

	// Visit a parse tree produced by KParser#loopElseStatement.
	VisitLoopElseStatement(ctx *LoopElseStatementContext) interface{}

	// Visit a parse tree produced by KParser#loopJumpStatement.
	VisitLoopJumpStatement(ctx *LoopJumpStatementContext) interface{}

	// Visit a parse tree produced by KParser#loopContinueStatement.
	VisitLoopContinueStatement(ctx *LoopContinueStatementContext) interface{}

	// Visit a parse tree produced by KParser#checkStatement.
	VisitCheckStatement(ctx *CheckStatementContext) interface{}

	// Visit a parse tree produced by KParser#usingStatement.
	VisitUsingStatement(ctx *UsingStatementContext) interface{}

	// Visit a parse tree produced by KParser#checkErrorStatement.
	VisitCheckErrorStatement(ctx *CheckErrorStatementContext) interface{}

	// Visit a parse tree produced by KParser#checkFinallyStatment.
	VisitCheckFinallyStatment(ctx *CheckFinallyStatmentContext) interface{}

	// Visit a parse tree produced by KParser#checkReportStatement.
	VisitCheckReportStatement(ctx *CheckReportStatementContext) interface{}

	// Visit a parse tree produced by KParser#iteratorStatement.
	VisitIteratorStatement(ctx *IteratorStatementContext) interface{}

	// Visit a parse tree produced by KParser#variableDeclaredStatement.
	VisitVariableDeclaredStatement(ctx *VariableDeclaredStatementContext) interface{}

	// Visit a parse tree produced by KParser#constantDeclaredStatement.
	VisitConstantDeclaredStatement(ctx *ConstantDeclaredStatementContext) interface{}

	// Visit a parse tree produced by KParser#varStatement.
	VisitVarStatement(ctx *VarStatementContext) interface{}

	// Visit a parse tree produced by KParser#bindStatement.
	VisitBindStatement(ctx *BindStatementContext) interface{}

	// Visit a parse tree produced by KParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by KParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by KParser#varId.
	VisitVarId(ctx *VarIdContext) interface{}

	// Visit a parse tree produced by KParser#constId.
	VisitConstId(ctx *ConstIdContext) interface{}

	// Visit a parse tree produced by KParser#tupleExpression.
	VisitTupleExpression(ctx *TupleExpressionContext) interface{}

	// Visit a parse tree produced by KParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by KParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by KParser#callExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by KParser#tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by KParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by KParser#annotationSupport.
	VisitAnnotationSupport(ctx *AnnotationSupportContext) interface{}

	// Visit a parse tree produced by KParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by KParser#annotationList.
	VisitAnnotationList(ctx *AnnotationListContext) interface{}

	// Visit a parse tree produced by KParser#annotationItem.
	VisitAnnotationItem(ctx *AnnotationItemContext) interface{}

	// Visit a parse tree produced by KParser#callFunc.
	VisitCallFunc(ctx *CallFuncContext) interface{}

	// Visit a parse tree produced by KParser#callAwaitFunc.
	VisitCallAwaitFunc(ctx *CallAwaitFuncContext) interface{}

	// Visit a parse tree produced by KParser#callAwait.
	VisitCallAwait(ctx *CallAwaitContext) interface{}

	// Visit a parse tree produced by KParser#callAsync.
	VisitCallAsync(ctx *CallAsyncContext) interface{}

	// Visit a parse tree produced by KParser#callChannel.
	VisitCallChannel(ctx *CallChannelContext) interface{}

	// Visit a parse tree produced by KParser#transfer.
	VisitTransfer(ctx *TransferContext) interface{}

	// Visit a parse tree produced by KParser#callElement.
	VisitCallElement(ctx *CallElementContext) interface{}

	// Visit a parse tree produced by KParser#callPkg.
	VisitCallPkg(ctx *CallPkgContext) interface{}

	// Visit a parse tree produced by KParser#callNew.
	VisitCallNew(ctx *CallNewContext) interface{}

	// Visit a parse tree produced by KParser#orElse.
	VisitOrElse(ctx *OrElseContext) interface{}

	// Visit a parse tree produced by KParser#typeConversion.
	VisitTypeConversion(ctx *TypeConversionContext) interface{}

	// Visit a parse tree produced by KParser#typeCheck.
	VisitTypeCheck(ctx *TypeCheckContext) interface{}

	// Visit a parse tree produced by KParser#pkgAssign.
	VisitPkgAssign(ctx *PkgAssignContext) interface{}

	// Visit a parse tree produced by KParser#pkgAssignElement.
	VisitPkgAssignElement(ctx *PkgAssignElementContext) interface{}

	// Visit a parse tree produced by KParser#listAssign.
	VisitListAssign(ctx *ListAssignContext) interface{}

	// Visit a parse tree produced by KParser#dictionaryAssign.
	VisitDictionaryAssign(ctx *DictionaryAssignContext) interface{}

	// Visit a parse tree produced by KParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by KParser#dictionary.
	VisitDictionary(ctx *DictionaryContext) interface{}

	// Visit a parse tree produced by KParser#dictionaryElement.
	VisitDictionaryElement(ctx *DictionaryElementContext) interface{}

	// Visit a parse tree produced by KParser#slice.
	VisitSlice(ctx *SliceContext) interface{}

	// Visit a parse tree produced by KParser#sliceFull.
	VisitSliceFull(ctx *SliceFullContext) interface{}

	// Visit a parse tree produced by KParser#sliceStart.
	VisitSliceStart(ctx *SliceStartContext) interface{}

	// Visit a parse tree produced by KParser#sliceEnd.
	VisitSliceEnd(ctx *SliceEndContext) interface{}

	// Visit a parse tree produced by KParser#nameSpaceItem.
	VisitNameSpaceItem(ctx *NameSpaceItemContext) interface{}

	// Visit a parse tree produced by KParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by KParser#templateDefine.
	VisitTemplateDefine(ctx *TemplateDefineContext) interface{}

	// Visit a parse tree produced by KParser#templateDefineItem.
	VisitTemplateDefineItem(ctx *TemplateDefineItemContext) interface{}

	// Visit a parse tree produced by KParser#templateCall.
	VisitTemplateCall(ctx *TemplateCallContext) interface{}

	// Visit a parse tree produced by KParser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by KParser#lambdaIn.
	VisitLambdaIn(ctx *LambdaInContext) interface{}

	// Visit a parse tree produced by KParser#pkgAnonymous.
	VisitPkgAnonymous(ctx *PkgAnonymousContext) interface{}

	// Visit a parse tree produced by KParser#pkgAnonymousAssign.
	VisitPkgAnonymousAssign(ctx *PkgAnonymousAssignContext) interface{}

	// Visit a parse tree produced by KParser#pkgAnonymousAssignElement.
	VisitPkgAnonymousAssignElement(ctx *PkgAnonymousAssignElementContext) interface{}

	// Visit a parse tree produced by KParser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by KParser#plusMinus.
	VisitPlusMinus(ctx *PlusMinusContext) interface{}

	// Visit a parse tree produced by KParser#negate.
	VisitNegate(ctx *NegateContext) interface{}

	// Visit a parse tree produced by KParser#bitwiseNotExpression.
	VisitBitwiseNotExpression(ctx *BitwiseNotExpressionContext) interface{}

	// Visit a parse tree produced by KParser#linq.
	VisitLinq(ctx *LinqContext) interface{}

	// Visit a parse tree produced by KParser#linqHeadItem.
	VisitLinqHeadItem(ctx *LinqHeadItemContext) interface{}

	// Visit a parse tree produced by KParser#linqItem.
	VisitLinqItem(ctx *LinqItemContext) interface{}

	// Visit a parse tree produced by KParser#judgeExpression.
	VisitJudgeExpression(ctx *JudgeExpressionContext) interface{}

	// Visit a parse tree produced by KParser#judgeExpressionElseStatement.
	VisitJudgeExpressionElseStatement(ctx *JudgeExpressionElseStatementContext) interface{}

	// Visit a parse tree produced by KParser#judgeExpressionIfStatement.
	VisitJudgeExpressionIfStatement(ctx *JudgeExpressionIfStatementContext) interface{}

	// Visit a parse tree produced by KParser#judgeExpressionElseIfStatement.
	VisitJudgeExpressionElseIfStatement(ctx *JudgeExpressionElseIfStatementContext) interface{}

	// Visit a parse tree produced by KParser#judgeCaseExpression.
	VisitJudgeCaseExpression(ctx *JudgeCaseExpressionContext) interface{}

	// Visit a parse tree produced by KParser#caseExpressionStatement.
	VisitCaseExpressionStatement(ctx *CaseExpressionStatementContext) interface{}

	// Visit a parse tree produced by KParser#loopExpression.
	VisitLoopExpression(ctx *LoopExpressionContext) interface{}

	// Visit a parse tree produced by KParser#loopEachExpression.
	VisitLoopEachExpression(ctx *LoopEachExpressionContext) interface{}

	// Visit a parse tree produced by KParser#loopElseExpression.
	VisitLoopElseExpression(ctx *LoopElseExpressionContext) interface{}

	// Visit a parse tree produced by KParser#checkExpression.
	VisitCheckExpression(ctx *CheckExpressionContext) interface{}

	// Visit a parse tree produced by KParser#checkErrorExpression.
	VisitCheckErrorExpression(ctx *CheckErrorExpressionContext) interface{}

	// Visit a parse tree produced by KParser#dataStatement.
	VisitDataStatement(ctx *DataStatementContext) interface{}

	// Visit a parse tree produced by KParser#stringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by KParser#stringContent.
	VisitStringContent(ctx *StringContentContext) interface{}

	// Visit a parse tree produced by KParser#stringTemplate.
	VisitStringTemplate(ctx *StringTemplateContext) interface{}

	// Visit a parse tree produced by KParser#rawStringExpr.
	VisitRawStringExpr(ctx *RawStringExprContext) interface{}

	// Visit a parse tree produced by KParser#rawStringContent.
	VisitRawStringContent(ctx *RawStringContentContext) interface{}

	// Visit a parse tree produced by KParser#rawStringTemplate.
	VisitRawStringTemplate(ctx *RawStringTemplateContext) interface{}

	// Visit a parse tree produced by KParser#floatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by KParser#integerExpr.
	VisitIntegerExpr(ctx *IntegerExprContext) interface{}

	// Visit a parse tree produced by KParser#typeNotNull.
	VisitTypeNotNull(ctx *TypeNotNullContext) interface{}

	// Visit a parse tree produced by KParser#typeType.
	VisitTypeType(ctx *TypeTypeContext) interface{}

	// Visit a parse tree produced by KParser#typeNullable.
	VisitTypeNullable(ctx *TypeNullableContext) interface{}

	// Visit a parse tree produced by KParser#typeArray.
	VisitTypeArray(ctx *TypeArrayContext) interface{}

	// Visit a parse tree produced by KParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by KParser#typeSet.
	VisitTypeSet(ctx *TypeSetContext) interface{}

	// Visit a parse tree produced by KParser#typeDictionary.
	VisitTypeDictionary(ctx *TypeDictionaryContext) interface{}

	// Visit a parse tree produced by KParser#typeStack.
	VisitTypeStack(ctx *TypeStackContext) interface{}

	// Visit a parse tree produced by KParser#typeQueue.
	VisitTypeQueue(ctx *TypeQueueContext) interface{}

	// Visit a parse tree produced by KParser#typeChannel.
	VisitTypeChannel(ctx *TypeChannelContext) interface{}

	// Visit a parse tree produced by KParser#typePackage.
	VisitTypePackage(ctx *TypePackageContext) interface{}

	// Visit a parse tree produced by KParser#typeFunction.
	VisitTypeFunction(ctx *TypeFunctionContext) interface{}

	// Visit a parse tree produced by KParser#typeAny.
	VisitTypeAny(ctx *TypeAnyContext) interface{}

	// Visit a parse tree produced by KParser#typeFunctionParameterClause.
	VisitTypeFunctionParameterClause(ctx *TypeFunctionParameterClauseContext) interface{}

	// Visit a parse tree produced by KParser#typeBasic.
	VisitTypeBasic(ctx *TypeBasicContext) interface{}

	// Visit a parse tree produced by KParser#nilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by KParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by KParser#bitwise.
	VisitBitwise(ctx *BitwiseContext) interface{}

	// Visit a parse tree produced by KParser#bitwiseAnd.
	VisitBitwiseAnd(ctx *BitwiseAndContext) interface{}

	// Visit a parse tree produced by KParser#bitwiseOr.
	VisitBitwiseOr(ctx *BitwiseOrContext) interface{}

	// Visit a parse tree produced by KParser#bitwiseNot.
	VisitBitwiseNot(ctx *BitwiseNotContext) interface{}

	// Visit a parse tree produced by KParser#bitwiseXor.
	VisitBitwiseXor(ctx *BitwiseXorContext) interface{}

	// Visit a parse tree produced by KParser#bitwiseLeftShift.
	VisitBitwiseLeftShift(ctx *BitwiseLeftShiftContext) interface{}

	// Visit a parse tree produced by KParser#bitwiseRightShift.
	VisitBitwiseRightShift(ctx *BitwiseRightShiftContext) interface{}

	// Visit a parse tree produced by KParser#compareCombine.
	VisitCompareCombine(ctx *CompareCombineContext) interface{}

	// Visit a parse tree produced by KParser#compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by KParser#logic.
	VisitLogic(ctx *LogicContext) interface{}

	// Visit a parse tree produced by KParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by KParser#add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by KParser#mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by KParser#pow.
	VisitPow(ctx *PowContext) interface{}

	// Visit a parse tree produced by KParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by KParser#wave.
	VisitWave(ctx *WaveContext) interface{}

	// Visit a parse tree produced by KParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by KParser#idItem.
	VisitIdItem(ctx *IdItemContext) interface{}

	// Visit a parse tree produced by KParser#end.
	VisitEnd(ctx *EndContext) interface{}

	// Visit a parse tree produced by KParser#more.
	VisitMore(ctx *MoreContext) interface{}

	// Visit a parse tree produced by KParser#left_brace.
	VisitLeft_brace(ctx *Left_braceContext) interface{}

	// Visit a parse tree produced by KParser#right_brace.
	VisitRight_brace(ctx *Right_braceContext) interface{}

	// Visit a parse tree produced by KParser#left_paren.
	VisitLeft_paren(ctx *Left_parenContext) interface{}

	// Visit a parse tree produced by KParser#right_paren.
	VisitRight_paren(ctx *Right_parenContext) interface{}

	// Visit a parse tree produced by KParser#left_brack.
	VisitLeft_brack(ctx *Left_brackContext) interface{}

	// Visit a parse tree produced by KParser#right_brack.
	VisitRight_brack(ctx *Right_brackContext) interface{}
}
