// Code generated from TiqParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TiqParser
import "github.com/antlr4-go/antlr/v4"

// BaseTiqParserListener is a complete listener for a parse tree produced by TiqParser.
type BaseTiqParserListener struct{}

var _ TiqParserListener = &BaseTiqParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTiqParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTiqParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTiqParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTiqParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseTiqParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseTiqParserListener) ExitScript(ctx *ScriptContext) {}

// EnterPackageDef is called when production packageDef is entered.
func (s *BaseTiqParserListener) EnterPackageDef(ctx *PackageDefContext) {}

// ExitPackageDef is called when production packageDef is exited.
func (s *BaseTiqParserListener) ExitPackageDef(ctx *PackageDefContext) {}

// EnterConstantsDef is called when production constantsDef is entered.
func (s *BaseTiqParserListener) EnterConstantsDef(ctx *ConstantsDefContext) {}

// ExitConstantsDef is called when production constantsDef is exited.
func (s *BaseTiqParserListener) ExitConstantsDef(ctx *ConstantsDefContext) {}

// EnterConstantDef is called when production constantDef is entered.
func (s *BaseTiqParserListener) EnterConstantDef(ctx *ConstantDefContext) {}

// ExitConstantDef is called when production constantDef is exited.
func (s *BaseTiqParserListener) ExitConstantDef(ctx *ConstantDefContext) {}

// EnterVariablesDef is called when production variablesDef is entered.
func (s *BaseTiqParserListener) EnterVariablesDef(ctx *VariablesDefContext) {}

// ExitVariablesDef is called when production variablesDef is exited.
func (s *BaseTiqParserListener) ExitVariablesDef(ctx *VariablesDefContext) {}

// EnterVariableDef is called when production variableDef is entered.
func (s *BaseTiqParserListener) EnterVariableDef(ctx *VariableDefContext) {}

// ExitVariableDef is called when production variableDef is exited.
func (s *BaseTiqParserListener) ExitVariableDef(ctx *VariableDefContext) {}

// EnterFunctionDef is called when production functionDef is entered.
func (s *BaseTiqParserListener) EnterFunctionDef(ctx *FunctionDefContext) {}

// ExitFunctionDef is called when production functionDef is exited.
func (s *BaseTiqParserListener) ExitFunctionDef(ctx *FunctionDefContext) {}

// EnterClass is called when production class is entered.
func (s *BaseTiqParserListener) EnterClass(ctx *ClassContext) {}

// ExitClass is called when production class is exited.
func (s *BaseTiqParserListener) ExitClass(ctx *ClassContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseTiqParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseTiqParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *BaseTiqParserListener) EnterParameterDecl(ctx *ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *BaseTiqParserListener) ExitParameterDecl(ctx *ParameterDeclContext) {}

// EnterResults is called when production results is entered.
func (s *BaseTiqParserListener) EnterResults(ctx *ResultsContext) {}

// ExitResults is called when production results is exited.
func (s *BaseTiqParserListener) ExitResults(ctx *ResultsContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseTiqParserListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseTiqParserListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterEnumItem is called when production enumItem is entered.
func (s *BaseTiqParserListener) EnterEnumItem(ctx *EnumItemContext) {}

// ExitEnumItem is called when production enumItem is exited.
func (s *BaseTiqParserListener) ExitEnumItem(ctx *EnumItemContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseTiqParserListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseTiqParserListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseTiqParserListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseTiqParserListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseTiqParserListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseTiqParserListener) ExitProperty(ctx *PropertyContext) {}

// EnterType is called when production type is entered.
func (s *BaseTiqParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseTiqParserListener) ExitType(ctx *TypeContext) {}

// EnterTimeSeriesType is called when production timeSeriesType is entered.
func (s *BaseTiqParserListener) EnterTimeSeriesType(ctx *TimeSeriesTypeContext) {}

// ExitTimeSeriesType is called when production timeSeriesType is exited.
func (s *BaseTiqParserListener) ExitTimeSeriesType(ctx *TimeSeriesTypeContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseTiqParserListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseTiqParserListener) ExitListType(ctx *ListTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseTiqParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseTiqParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BaseTiqParserListener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BaseTiqParserListener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseTiqParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseTiqParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTiqParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTiqParserListener) ExitStatement(ctx *StatementContext) {}

// EnterVarsDeclaration is called when production varsDeclaration is entered.
func (s *BaseTiqParserListener) EnterVarsDeclaration(ctx *VarsDeclarationContext) {}

// ExitVarsDeclaration is called when production varsDeclaration is exited.
func (s *BaseTiqParserListener) ExitVarsDeclaration(ctx *VarsDeclarationContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseTiqParserListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseTiqParserListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterVarsAssignmentOrFunctionCall is called when production varsAssignmentOrFunctionCall is entered.
func (s *BaseTiqParserListener) EnterVarsAssignmentOrFunctionCall(ctx *VarsAssignmentOrFunctionCallContext) {
}

// ExitVarsAssignmentOrFunctionCall is called when production varsAssignmentOrFunctionCall is exited.
func (s *BaseTiqParserListener) ExitVarsAssignmentOrFunctionCall(ctx *VarsAssignmentOrFunctionCallContext) {
}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseTiqParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseTiqParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterElseIfBlock is called when production elseIfBlock is entered.
func (s *BaseTiqParserListener) EnterElseIfBlock(ctx *ElseIfBlockContext) {}

// ExitElseIfBlock is called when production elseIfBlock is exited.
func (s *BaseTiqParserListener) ExitElseIfBlock(ctx *ElseIfBlockContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseTiqParserListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseTiqParserListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseTiqParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseTiqParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTiqParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTiqParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseTiqParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseTiqParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterExpressionInParenthesis is called when production expressionInParenthesis is entered.
func (s *BaseTiqParserListener) EnterExpressionInParenthesis(ctx *ExpressionInParenthesisContext) {}

// ExitExpressionInParenthesis is called when production expressionInParenthesis is exited.
func (s *BaseTiqParserListener) ExitExpressionInParenthesis(ctx *ExpressionInParenthesisContext) {}

// EnterChainedExpression is called when production chainedExpression is entered.
func (s *BaseTiqParserListener) EnterChainedExpression(ctx *ChainedExpressionContext) {}

// ExitChainedExpression is called when production chainedExpression is exited.
func (s *BaseTiqParserListener) ExitChainedExpression(ctx *ChainedExpressionContext) {}

// EnterChainItem is called when production chainItem is entered.
func (s *BaseTiqParserListener) EnterChainItem(ctx *ChainItemContext) {}

// ExitChainItem is called when production chainItem is exited.
func (s *BaseTiqParserListener) ExitChainItem(ctx *ChainItemContext) {}

// EnterParamsExpression is called when production paramsExpression is entered.
func (s *BaseTiqParserListener) EnterParamsExpression(ctx *ParamsExpressionContext) {}

// ExitParamsExpression is called when production paramsExpression is exited.
func (s *BaseTiqParserListener) ExitParamsExpression(ctx *ParamsExpressionContext) {}

// EnterAccessorExpression is called when production accessorExpression is entered.
func (s *BaseTiqParserListener) EnterAccessorExpression(ctx *AccessorExpressionContext) {}

// ExitAccessorExpression is called when production accessorExpression is exited.
func (s *BaseTiqParserListener) ExitAccessorExpression(ctx *AccessorExpressionContext) {}

// EnterBarExpression is called when production barExpression is entered.
func (s *BaseTiqParserListener) EnterBarExpression(ctx *BarExpressionContext) {}

// ExitBarExpression is called when production barExpression is exited.
func (s *BaseTiqParserListener) ExitBarExpression(ctx *BarExpressionContext) {}

// EnterListExpression is called when production listExpression is entered.
func (s *BaseTiqParserListener) EnterListExpression(ctx *ListExpressionContext) {}

// ExitListExpression is called when production listExpression is exited.
func (s *BaseTiqParserListener) ExitListExpression(ctx *ListExpressionContext) {}

// EnterMapExpression is called when production mapExpression is entered.
func (s *BaseTiqParserListener) EnterMapExpression(ctx *MapExpressionContext) {}

// ExitMapExpression is called when production mapExpression is exited.
func (s *BaseTiqParserListener) ExitMapExpression(ctx *MapExpressionContext) {}

// EnterKeyValueCouple is called when production keyValueCouple is entered.
func (s *BaseTiqParserListener) EnterKeyValueCouple(ctx *KeyValueCoupleContext) {}

// ExitKeyValueCouple is called when production keyValueCouple is exited.
func (s *BaseTiqParserListener) ExitKeyValueCouple(ctx *KeyValueCoupleContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseTiqParserListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseTiqParserListener) ExitKeyValue(ctx *KeyValueContext) {}

// EnterConstantValueExpression is called when production constantValueExpression is entered.
func (s *BaseTiqParserListener) EnterConstantValueExpression(ctx *ConstantValueExpressionContext) {}

// ExitConstantValueExpression is called when production constantValueExpression is exited.
func (s *BaseTiqParserListener) ExitConstantValueExpression(ctx *ConstantValueExpressionContext) {}

// EnterTimeValue is called when production timeValue is entered.
func (s *BaseTiqParserListener) EnterTimeValue(ctx *TimeValueContext) {}

// ExitTimeValue is called when production timeValue is exited.
func (s *BaseTiqParserListener) ExitTimeValue(ctx *TimeValueContext) {}

// EnterDateValue is called when production dateValue is entered.
func (s *BaseTiqParserListener) EnterDateValue(ctx *DateValueContext) {}

// ExitDateValue is called when production dateValue is exited.
func (s *BaseTiqParserListener) ExitDateValue(ctx *DateValueContext) {}

// EnterBoolValue is called when production boolValue is entered.
func (s *BaseTiqParserListener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production boolValue is exited.
func (s *BaseTiqParserListener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterErrorValue is called when production errorValue is entered.
func (s *BaseTiqParserListener) EnterErrorValue(ctx *ErrorValueContext) {}

// ExitErrorValue is called when production errorValue is exited.
func (s *BaseTiqParserListener) ExitErrorValue(ctx *ErrorValueContext) {}

// EnterFqIdentifier is called when production fqIdentifier is entered.
func (s *BaseTiqParserListener) EnterFqIdentifier(ctx *FqIdentifierContext) {}

// ExitFqIdentifier is called when production fqIdentifier is exited.
func (s *BaseTiqParserListener) ExitFqIdentifier(ctx *FqIdentifierContext) {}
