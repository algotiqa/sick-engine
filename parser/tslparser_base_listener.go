// Code generated from TslParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TslParser
import "github.com/antlr4-go/antlr/v4"

// BaseTslParserListener is a complete listener for a parse tree produced by TslParser.
type BaseTslParserListener struct{}

var _ TslParserListener = &BaseTslParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTslParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTslParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTslParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTslParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseTslParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseTslParserListener) ExitScript(ctx *ScriptContext) {}

// EnterPackageDef is called when production packageDef is entered.
func (s *BaseTslParserListener) EnterPackageDef(ctx *PackageDefContext) {}

// ExitPackageDef is called when production packageDef is exited.
func (s *BaseTslParserListener) ExitPackageDef(ctx *PackageDefContext) {}

// EnterConstantsDef is called when production constantsDef is entered.
func (s *BaseTslParserListener) EnterConstantsDef(ctx *ConstantsDefContext) {}

// ExitConstantsDef is called when production constantsDef is exited.
func (s *BaseTslParserListener) ExitConstantsDef(ctx *ConstantsDefContext) {}

// EnterConstantDef is called when production constantDef is entered.
func (s *BaseTslParserListener) EnterConstantDef(ctx *ConstantDefContext) {}

// ExitConstantDef is called when production constantDef is exited.
func (s *BaseTslParserListener) ExitConstantDef(ctx *ConstantDefContext) {}

// EnterVariablesDef is called when production variablesDef is entered.
func (s *BaseTslParserListener) EnterVariablesDef(ctx *VariablesDefContext) {}

// ExitVariablesDef is called when production variablesDef is exited.
func (s *BaseTslParserListener) ExitVariablesDef(ctx *VariablesDefContext) {}

// EnterVariableDef is called when production variableDef is entered.
func (s *BaseTslParserListener) EnterVariableDef(ctx *VariableDefContext) {}

// ExitVariableDef is called when production variableDef is exited.
func (s *BaseTslParserListener) ExitVariableDef(ctx *VariableDefContext) {}

// EnterFunctionDef is called when production functionDef is entered.
func (s *BaseTslParserListener) EnterFunctionDef(ctx *FunctionDefContext) {}

// ExitFunctionDef is called when production functionDef is exited.
func (s *BaseTslParserListener) ExitFunctionDef(ctx *FunctionDefContext) {}

// EnterClass is called when production class is entered.
func (s *BaseTslParserListener) EnterClass(ctx *ClassContext) {}

// ExitClass is called when production class is exited.
func (s *BaseTslParserListener) ExitClass(ctx *ClassContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseTslParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseTslParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *BaseTslParserListener) EnterParameterDecl(ctx *ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *BaseTslParserListener) ExitParameterDecl(ctx *ParameterDeclContext) {}

// EnterResults is called when production results is entered.
func (s *BaseTslParserListener) EnterResults(ctx *ResultsContext) {}

// ExitResults is called when production results is exited.
func (s *BaseTslParserListener) ExitResults(ctx *ResultsContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseTslParserListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseTslParserListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterEnumItem is called when production enumItem is entered.
func (s *BaseTslParserListener) EnterEnumItem(ctx *EnumItemContext) {}

// ExitEnumItem is called when production enumItem is exited.
func (s *BaseTslParserListener) ExitEnumItem(ctx *EnumItemContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseTslParserListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseTslParserListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseTslParserListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseTslParserListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseTslParserListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseTslParserListener) ExitProperty(ctx *PropertyContext) {}

// EnterType is called when production type is entered.
func (s *BaseTslParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseTslParserListener) ExitType(ctx *TypeContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseTslParserListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseTslParserListener) ExitListType(ctx *ListTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseTslParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseTslParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BaseTslParserListener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BaseTslParserListener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseTslParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseTslParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTslParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTslParserListener) ExitStatement(ctx *StatementContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseTslParserListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseTslParserListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseTslParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseTslParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterElseIfBlock is called when production elseIfBlock is entered.
func (s *BaseTslParserListener) EnterElseIfBlock(ctx *ElseIfBlockContext) {}

// ExitElseIfBlock is called when production elseIfBlock is exited.
func (s *BaseTslParserListener) ExitElseIfBlock(ctx *ElseIfBlockContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseTslParserListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseTslParserListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseTslParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseTslParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseTslParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseTslParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTslParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTslParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseTslParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseTslParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterIdentifierExpression is called when production identifierExpression is entered.
func (s *BaseTslParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production identifierExpression is exited.
func (s *BaseTslParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterFqIdentifier is called when production fqIdentifier is entered.
func (s *BaseTslParserListener) EnterFqIdentifier(ctx *FqIdentifierContext) {}

// ExitFqIdentifier is called when production fqIdentifier is exited.
func (s *BaseTslParserListener) ExitFqIdentifier(ctx *FqIdentifierContext) {}

// EnterAccessorExpression is called when production accessorExpression is entered.
func (s *BaseTslParserListener) EnterAccessorExpression(ctx *AccessorExpressionContext) {}

// ExitAccessorExpression is called when production accessorExpression is exited.
func (s *BaseTslParserListener) ExitAccessorExpression(ctx *AccessorExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseTslParserListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseTslParserListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterParamsExpression is called when production paramsExpression is entered.
func (s *BaseTslParserListener) EnterParamsExpression(ctx *ParamsExpressionContext) {}

// ExitParamsExpression is called when production paramsExpression is exited.
func (s *BaseTslParserListener) ExitParamsExpression(ctx *ParamsExpressionContext) {}

// EnterExpressionInParenthesis is called when production expressionInParenthesis is entered.
func (s *BaseTslParserListener) EnterExpressionInParenthesis(ctx *ExpressionInParenthesisContext) {}

// ExitExpressionInParenthesis is called when production expressionInParenthesis is exited.
func (s *BaseTslParserListener) ExitExpressionInParenthesis(ctx *ExpressionInParenthesisContext) {}

// EnterListExpression is called when production listExpression is entered.
func (s *BaseTslParserListener) EnterListExpression(ctx *ListExpressionContext) {}

// ExitListExpression is called when production listExpression is exited.
func (s *BaseTslParserListener) ExitListExpression(ctx *ListExpressionContext) {}

// EnterInitialListValues is called when production initialListValues is entered.
func (s *BaseTslParserListener) EnterInitialListValues(ctx *InitialListValuesContext) {}

// ExitInitialListValues is called when production initialListValues is exited.
func (s *BaseTslParserListener) ExitInitialListValues(ctx *InitialListValuesContext) {}

// EnterMapExpression is called when production mapExpression is entered.
func (s *BaseTslParserListener) EnterMapExpression(ctx *MapExpressionContext) {}

// ExitMapExpression is called when production mapExpression is exited.
func (s *BaseTslParserListener) ExitMapExpression(ctx *MapExpressionContext) {}

// EnterInitialMapValues is called when production initialMapValues is entered.
func (s *BaseTslParserListener) EnterInitialMapValues(ctx *InitialMapValuesContext) {}

// ExitInitialMapValues is called when production initialMapValues is exited.
func (s *BaseTslParserListener) ExitInitialMapValues(ctx *InitialMapValuesContext) {}

// EnterKeyValueCouple is called when production keyValueCouple is entered.
func (s *BaseTslParserListener) EnterKeyValueCouple(ctx *KeyValueCoupleContext) {}

// ExitKeyValueCouple is called when production keyValueCouple is exited.
func (s *BaseTslParserListener) ExitKeyValueCouple(ctx *KeyValueCoupleContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseTslParserListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseTslParserListener) ExitKeyValue(ctx *KeyValueContext) {}

// EnterSimplifiedExpression is called when production simplifiedExpression is entered.
func (s *BaseTslParserListener) EnterSimplifiedExpression(ctx *SimplifiedExpressionContext) {}

// ExitSimplifiedExpression is called when production simplifiedExpression is exited.
func (s *BaseTslParserListener) ExitSimplifiedExpression(ctx *SimplifiedExpressionContext) {}

// EnterConstantValueExpression is called when production constantValueExpression is entered.
func (s *BaseTslParserListener) EnterConstantValueExpression(ctx *ConstantValueExpressionContext) {}

// ExitConstantValueExpression is called when production constantValueExpression is exited.
func (s *BaseTslParserListener) ExitConstantValueExpression(ctx *ConstantValueExpressionContext) {}

// EnterTimeValue is called when production timeValue is entered.
func (s *BaseTslParserListener) EnterTimeValue(ctx *TimeValueContext) {}

// ExitTimeValue is called when production timeValue is exited.
func (s *BaseTslParserListener) ExitTimeValue(ctx *TimeValueContext) {}

// EnterDateValue is called when production dateValue is entered.
func (s *BaseTslParserListener) EnterDateValue(ctx *DateValueContext) {}

// ExitDateValue is called when production dateValue is exited.
func (s *BaseTslParserListener) ExitDateValue(ctx *DateValueContext) {}

// EnterBoolValue is called when production boolValue is entered.
func (s *BaseTslParserListener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production boolValue is exited.
func (s *BaseTslParserListener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterErrorValue is called when production errorValue is entered.
func (s *BaseTslParserListener) EnterErrorValue(ctx *ErrorValueContext) {}

// ExitErrorValue is called when production errorValue is exited.
func (s *BaseTslParserListener) ExitErrorValue(ctx *ErrorValueContext) {}
