// Code generated from TslParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TslParser
import "github.com/antlr4-go/antlr/v4"

// TslParserListener is a complete listener for a parse tree produced by TslParser.
type TslParserListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterPackageDef is called when entering the packageDef production.
	EnterPackageDef(c *PackageDefContext)

	// EnterConstantsDef is called when entering the constantsDef production.
	EnterConstantsDef(c *ConstantsDefContext)

	// EnterConstantDef is called when entering the constantDef production.
	EnterConstantDef(c *ConstantDefContext)

	// EnterVariablesDef is called when entering the variablesDef production.
	EnterVariablesDef(c *VariablesDefContext)

	// EnterVariableDef is called when entering the variableDef production.
	EnterVariableDef(c *VariableDefContext)

	// EnterFunctionDef is called when entering the functionDef production.
	EnterFunctionDef(c *FunctionDefContext)

	// EnterClass is called when entering the class production.
	EnterClass(c *ClassContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParameterDecl is called when entering the parameterDecl production.
	EnterParameterDecl(c *ParameterDeclContext)

	// EnterResults is called when entering the results production.
	EnterResults(c *ResultsContext)

	// EnterEnumDef is called when entering the enumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterEnumItem is called when entering the enumItem production.
	EnterEnumItem(c *EnumItemContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// EnterClassDef is called when entering the classDef production.
	EnterClassDef(c *ClassDefContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTimeSeriesType is called when entering the timeSeriesType production.
	EnterTimeSeriesType(c *TimeSeriesTypeContext)

	// EnterListType is called when entering the listType production.
	EnterListType(c *ListTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterKeyType is called when entering the keyType production.
	EnterKeyType(c *KeyTypeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVarsDeclaration is called when entering the varsDeclaration production.
	EnterVarsDeclaration(c *VarsDeclarationContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterVarsAssignmentOrFunctionCall is called when entering the varsAssignmentOrFunctionCall production.
	EnterVarsAssignmentOrFunctionCall(c *VarsAssignmentOrFunctionCallContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterElseIfBlock is called when entering the elseIfBlock production.
	EnterElseIfBlock(c *ElseIfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterExpressionInParenthesis is called when entering the expressionInParenthesis production.
	EnterExpressionInParenthesis(c *ExpressionInParenthesisContext)

	// EnterChainedExpression is called when entering the chainedExpression production.
	EnterChainedExpression(c *ChainedExpressionContext)

	// EnterChainItem is called when entering the chainItem production.
	EnterChainItem(c *ChainItemContext)

	// EnterParamsExpression is called when entering the paramsExpression production.
	EnterParamsExpression(c *ParamsExpressionContext)

	// EnterAccessorExpression is called when entering the accessorExpression production.
	EnterAccessorExpression(c *AccessorExpressionContext)

	// EnterBarExpression is called when entering the barExpression production.
	EnterBarExpression(c *BarExpressionContext)

	// EnterListExpression is called when entering the listExpression production.
	EnterListExpression(c *ListExpressionContext)

	// EnterMapExpression is called when entering the mapExpression production.
	EnterMapExpression(c *MapExpressionContext)

	// EnterKeyValueCouple is called when entering the keyValueCouple production.
	EnterKeyValueCouple(c *KeyValueCoupleContext)

	// EnterKeyValue is called when entering the keyValue production.
	EnterKeyValue(c *KeyValueContext)

	// EnterConstantValueExpression is called when entering the constantValueExpression production.
	EnterConstantValueExpression(c *ConstantValueExpressionContext)

	// EnterTimeValue is called when entering the timeValue production.
	EnterTimeValue(c *TimeValueContext)

	// EnterDateValue is called when entering the dateValue production.
	EnterDateValue(c *DateValueContext)

	// EnterBoolValue is called when entering the boolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterErrorValue is called when entering the errorValue production.
	EnterErrorValue(c *ErrorValueContext)

	// EnterFqIdentifier is called when entering the fqIdentifier production.
	EnterFqIdentifier(c *FqIdentifierContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitPackageDef is called when exiting the packageDef production.
	ExitPackageDef(c *PackageDefContext)

	// ExitConstantsDef is called when exiting the constantsDef production.
	ExitConstantsDef(c *ConstantsDefContext)

	// ExitConstantDef is called when exiting the constantDef production.
	ExitConstantDef(c *ConstantDefContext)

	// ExitVariablesDef is called when exiting the variablesDef production.
	ExitVariablesDef(c *VariablesDefContext)

	// ExitVariableDef is called when exiting the variableDef production.
	ExitVariableDef(c *VariableDefContext)

	// ExitFunctionDef is called when exiting the functionDef production.
	ExitFunctionDef(c *FunctionDefContext)

	// ExitClass is called when exiting the class production.
	ExitClass(c *ClassContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameterDecl is called when exiting the parameterDecl production.
	ExitParameterDecl(c *ParameterDeclContext)

	// ExitResults is called when exiting the results production.
	ExitResults(c *ResultsContext)

	// ExitEnumDef is called when exiting the enumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitEnumItem is called when exiting the enumItem production.
	ExitEnumItem(c *EnumItemContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTimeSeriesType is called when exiting the timeSeriesType production.
	ExitTimeSeriesType(c *TimeSeriesTypeContext)

	// ExitListType is called when exiting the listType production.
	ExitListType(c *ListTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitKeyType is called when exiting the keyType production.
	ExitKeyType(c *KeyTypeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVarsDeclaration is called when exiting the varsDeclaration production.
	ExitVarsDeclaration(c *VarsDeclarationContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitVarsAssignmentOrFunctionCall is called when exiting the varsAssignmentOrFunctionCall production.
	ExitVarsAssignmentOrFunctionCall(c *VarsAssignmentOrFunctionCallContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitElseIfBlock is called when exiting the elseIfBlock production.
	ExitElseIfBlock(c *ElseIfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitExpressionInParenthesis is called when exiting the expressionInParenthesis production.
	ExitExpressionInParenthesis(c *ExpressionInParenthesisContext)

	// ExitChainedExpression is called when exiting the chainedExpression production.
	ExitChainedExpression(c *ChainedExpressionContext)

	// ExitChainItem is called when exiting the chainItem production.
	ExitChainItem(c *ChainItemContext)

	// ExitParamsExpression is called when exiting the paramsExpression production.
	ExitParamsExpression(c *ParamsExpressionContext)

	// ExitAccessorExpression is called when exiting the accessorExpression production.
	ExitAccessorExpression(c *AccessorExpressionContext)

	// ExitBarExpression is called when exiting the barExpression production.
	ExitBarExpression(c *BarExpressionContext)

	// ExitListExpression is called when exiting the listExpression production.
	ExitListExpression(c *ListExpressionContext)

	// ExitMapExpression is called when exiting the mapExpression production.
	ExitMapExpression(c *MapExpressionContext)

	// ExitKeyValueCouple is called when exiting the keyValueCouple production.
	ExitKeyValueCouple(c *KeyValueCoupleContext)

	// ExitKeyValue is called when exiting the keyValue production.
	ExitKeyValue(c *KeyValueContext)

	// ExitConstantValueExpression is called when exiting the constantValueExpression production.
	ExitConstantValueExpression(c *ConstantValueExpressionContext)

	// ExitTimeValue is called when exiting the timeValue production.
	ExitTimeValue(c *TimeValueContext)

	// ExitDateValue is called when exiting the dateValue production.
	ExitDateValue(c *DateValueContext)

	// ExitBoolValue is called when exiting the boolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitErrorValue is called when exiting the errorValue production.
	ExitErrorValue(c *ErrorValueContext)

	// ExitFqIdentifier is called when exiting the fqIdentifier production.
	ExitFqIdentifier(c *FqIdentifierContext)
}
