//=============================================================================
/*
Copyright © 2025 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package ast

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tradalia/sick-engine/ast/expression"
	"github.com/tradalia/sick-engine/ast/statement"
	"github.com/tradalia/sick-engine/data"
	"github.com/tradalia/sick-engine/parser"
	"github.com/tradalia/sick-engine/tool"
	"github.com/tradalia/sick-engine/types"
	"github.com/tradalia/sick-engine/values"
)

//=============================================================================

const (
	DefaultPackage = "main"
)

//=============================================================================
//=== We cannot use the parser to match these as we want some flexibility, like
//=== matching close, bar.close and any other bar.X that can emerge in the future
//=============================================================================

var ReservedIdentifiers = map[string]bool{
	"bar"   : true,
	"trade" : true,
	"open"  : true,
	"close" : true,
	"high"  : true,
	"low"   : true,
}

//=============================================================================
//===
//=== AST builder
//===
//=============================================================================

func Build(tree parser.IScriptContext) *Script {
	script := NewScript(tree.GetParser().GetInputStream().GetSourceName())

	script.PackageName = convertPackage(tree.PackageDef())

	convertConstants(script, tree.AllConstantsDef())
	convertVariables(script, tree.AllVariablesDef())
	convertFunctions(script, tree.AllFunctionDef())
	convertEnums    (script, tree.AllEnumDef())
	convertClasses  (script, tree.AllClassDef())

	return script
}

//=============================================================================
//===
//=== Top entities
//===
//=============================================================================

func convertPackage(tree parser.IPackageDefContext) string {
	if tree != nil {
		id := tree.IDENTIFIER().GetText()

		if _,ok := ReservedIdentifiers[id]; ok {
			raiseError(tree.GetParser(), "package's name is reserved: "+ id)
		}

		if !tool.StartsWithLowerCase(id) {
			raiseError(tree.GetParser(), "package's name must start with a lowercase character: "+ id)
		}

		return id
	}

	return DefaultPackage
}

//=============================================================================
//=== Constants
//=============================================================================

func convertConstants(script *Script, tree []parser.IConstantsDefContext) {
	for _, cds := range tree {
		for _,cd := range cds.AllConstantDef() {
			c := convertConstantDef(cd)
			if !script.AddConstant(c) {
				raiseError(cds.GetParser(), "constant's name already used: "+ c.Name)
			}
		}
	}
}

//=============================================================================

func convertConstantDef(tree parser.IConstantDefContext) *Constant {
	name  := tree.IDENTIFIER().GetText()
	value := tree.SimplifiedExpression()

	if !tool.StartsWithUpperCase(name) {
		raiseError(tree.GetParser(), "constant's name must start with an uppercase character: "+ name)
	}

	return NewConstant(name, convertSimplifiedExpression(value), data.NewInfo(tree))
}

//=============================================================================
//=== Variables
//=============================================================================

func convertVariables(script *Script, tree []parser.IVariablesDefContext) {
	for _, vds := range tree {
		for _,vd := range vds.AllVariableDef() {
			v := convertVariableDef(vd)
			if !script.AddVariable(v) {
				raiseError(vd.GetParser(), "variable's name already used: "+ v.Name)
			}
		}
	}
}

//=============================================================================

func convertVariableDef(tree parser.IVariableDefContext) *Variable {
	name  := tree.IDENTIFIER().GetText()
	value := tree.SimplifiedExpression()

	if _,ok := ReservedIdentifiers[name]; ok {
		raiseError(tree.GetParser(), "variable's name is reserved: "+ name)
	}

	if !tool.StartsWithLowerCase(name) {
		raiseError(tree.GetParser(), "variable's name must start with a lowercase character: "+ name)
	}

	return NewVariable(name, convertSimplifiedExpression(value), data.NewInfo(tree))
}

//=============================================================================
//=== Functions
//=============================================================================

func convertFunctions(script *Script, tree []parser.IFunctionDefContext) {
	for _, fd := range tree {
		f := convertFunctionDef(fd)
		if !script.AddFunction(f) {
			raiseError(fd.GetParser(), "function's name already used: "+ f.Name)
		}
	}
}

//=============================================================================

func convertFunctionDef(tree parser.IFunctionDefContext) *Function {
	name := tree.IDENTIFIER().GetText()

	if _,ok := ReservedIdentifiers[name]; ok {
		raiseError(tree.GetParser(), "function's name is reserved: "+ name)
	}

	if !tool.StartsWithLowerCase(name) {
		raiseError(tree.GetParser(), "function's name must start with a lowercase character: "+ name)
	}

	f := NewFunction(name, data.NewInfo(tree))

	if tree.Class() != nil {
		f.Class = tree.Class().IDENTIFIER().GetText()
		if !tool.StartsWithUpperCase(f.Class) {
			raiseError(tree.GetParser(), "function's class name must start with an uppercase character: "+ f.Class)
		}
	}

	convertFunctionParams (f, tree.Parameters())
	convertFunctionResults(f, tree.Results())
	f.Block = convertStatementBlock(tree.Block())

	return f
}

//=============================================================================

func convertFunctionParams(f *Function, params parser.IParametersContext) {
	names := map[string]bool{}

	for _, pd := range params.AllParameterDecl() {
		name  := pd.IDENTIFIER().GetText()
		type_ := pd.Type_()

		if _,ok := ReservedIdentifiers[name]; ok {
			raiseError(params.GetParser(), "function's parameter is reserved: "+ name)
		}

		if !tool.StartsWithLowerCase(name) {
			raiseError(params.GetParser(), "function's parameter must start with a lowercase character: "+ name)
		}

		if _,ok := names[name]; ok {
			raiseError(params.GetParser(), "function's parameter name already used: "+ name)
		}

		names[name] = true
		t := convertType(type_)
		p := NewParam(name,t, data.NewInfo(pd))
		f.AddParam(p)
	}
}

//=============================================================================

func convertFunctionResults(f *Function, results parser.IResultsContext) {
	if results != nil {
		for _,r := range results.AllType_() {
			t := convertType(r)
			f.AddReturnType(t)
		}
	}
}

//=============================================================================
//=== Enums
//=============================================================================

func convertEnums(script *Script, tree []parser.IEnumDefContext) {
	for _, ed := range tree {
		e := convertEnumDef(ed)
		if !script.AddEnum(e) {
			raiseError(ed.GetParser(), "enum's name already used: "+ e.Name)
		}
	}
}

//=============================================================================

func convertEnumDef(tree parser.IEnumDefContext) *types.EnumType {
	name := tree.IDENTIFIER().GetText()
	e := types.NewEnumType(name, true, data.NewInfo(tree))

	if !tool.StartsWithUpperCase(name) {
		raiseError(tree.GetParser(), "enum's name must start with an uppercase character: "+ name)
	}

	countInts := 0
	countStr  := 0

	for _, ei := range tree.AllEnumItem() {
		i := convertEnumItem(ei)
		if !e.AddItem(i) {
			raiseError(tree.GetParser(), "enum's item already used: "+ i.Name)
		}

		if i.Code != -1 {
			countInts++
		}

		if i.Value != "" {
			countStr++
		}
	}

	badInts := countInts > 0 && countInts != e.Size()
	badStr  := countStr  > 0 && countStr  != e.Size()

	if badInts && badStr {
		raiseError(tree.GetParser(), "mixing integers and strings in enum: "+ name)
	}

	if countInts == 0 && countStr == 0 {
		e.AssignCodes()
	} else if countStr != 0 {
		e.IsInt = false
	}

	return e
}

//=============================================================================

func convertEnumItem(tree parser.IEnumItemContext) *types.EnumItem {
	name  := tree.IDENTIFIER().GetText()
	code  := -1
	value := ""

	if !tool.StartsWithUpperCase(name) {
		raiseError(tree.GetParser(), "enum item's name must start with an uppercase character: "+ name)
	}

	enumVal := tree.EnumValue()
	if enumVal != nil {
		intValue := enumVal.INT_VALUE()
		strValue := enumVal.STRING_VALUE()

		if intValue != nil {
			iValue,err := strconv.Atoi(intValue.GetText())
			if err != nil {
				raiseError(tree.GetParser(), "Invalid integer value for enum : " + intValue.GetText())
			}

			code = iValue
		}
		if strValue != nil {
			text := strValue.GetText()
			value = text[1:len(text)-1]
		}
	}
	return types.NewEnumItem(name, code, value)
}

//=============================================================================
//=== Classes
//=============================================================================

func convertClasses(script *Script, tree []parser.IClassDefContext) {
	for _, cd := range tree {
		c := convertClassDef(cd)
		if !script.AddClass(c) {
			raiseError(cd.GetParser(), "class's name already used: "+ c.Name)
		}
	}
}

//=============================================================================

func convertClassDef(tree parser.IClassDefContext) *types.ClassType {
	name := tree.IDENTIFIER().GetText()
	c := types.NewClassType(name, data.NewInfo(tree))

	if !tool.StartsWithUpperCase(name) {
		raiseError(tree.GetParser(), "class's name must start with an uppercase character: "+ name)
	}

	for _, p := range tree.AllProperty() {
		if !c.AddProperty(convertProperty(p)) {
			raiseError(p.GetParser(), "class's property already used: "+ c.Name)
		}
	}

	return c
}

//=============================================================================

func convertProperty(tree parser.IPropertyContext) *types.Property {
	name  := tree.IDENTIFIER().GetText()
	type_ := tree.Type_()

	if _,ok := ReservedIdentifiers[name]; ok {
		raiseError(tree.GetParser(), "class's property is reserved: "+ name)
	}

	if !tool.StartsWithLowerCase(name) {
		raiseError(tree.GetParser(), "class's property must start with a lowercase character: "+ name)
	}

	t := convertType(type_)

	return types.NewProperty(name,t, data.NewInfo(tree))
}

//=============================================================================
//===
//=== Statements
//===
//=============================================================================

func convertStatementBlock(tree parser.IBlockContext) *statement.Block {
	block := &statement.Block{}

	for _,stmt := range tree.AllStatement() {
		varDecl := stmt.VarDeclaration()
		ifStmt  := stmt.IfStatement()
		callStmt:= stmt.FunctionCall()
		retStmt := stmt.ReturnStatement()

		if varDecl != nil {
			block.Add(convertVarDeclaration(varDecl))
		} else if ifStmt != nil {
			block.Add(convertIfStatement(ifStmt))
		} else if callStmt != nil {
			block.Add(convertFunctionCallStatement(callStmt))
		} else if retStmt != nil {
			block.Add(convertReturnStatement(retStmt))
		} else {
			panic("Unknown statement type : " + stmt.GetText())
		}
	}

	return block
}

//=============================================================================

func convertVarDeclaration(tree parser.IVarDeclarationContext) *statement.VarDeclaration {
	vd := statement.NewVarDeclaration()

	for _, fqId := range tree.AllFqIdentifier() {
		name := convertFQIdentifier(fqId)
		if ! tool.IsLowerCase(name) {
			raiseError(tree.GetParser(), "Fully qualified identifier must be in lowercase: "+ name.String())
		}

		var acc expression.Expression

		vd.AddIdentifier(expression.NewIdentifierExpression(name, acc))
	}

	for _, e := range tree.AllExpression() {
		expr := convertExpression(e)
		vd.AddExpression(expr)
	}

	return vd
}

//=============================================================================

func convertIfStatement(tree parser.IIfStatementContext) *statement.IfStatement {
	is    := statement.NewIfStatement()
	cond  := convertExpression(tree.Expression())
	block := convertStatementBlock(tree.Block())
	is.AddConditionalBlock(statement.NewConditionalBlock(cond, block))

	for _, eib := range tree.AllElseIfBlock() {
		cond  = convertExpression(eib.Expression())
		block = convertStatementBlock(eib.Block())

		is.AddConditionalBlock(statement.NewConditionalBlock(cond, block))
	}

	if tree.ElseBlock() != nil {
		cond  = expression.NewTrueExpression()
		block = convertStatementBlock(tree.ElseBlock().Block())

		is.AddConditionalBlock(statement.NewConditionalBlock(cond, block))
	}

	return is
}

//=============================================================================

func convertReturnStatement(tree parser.IReturnStatementContext) *statement.ReturnStatement {
	rs := statement.NewReturnStatement()

	for _, e := range tree.AllExpression() {
		expr := convertExpression(e)
		rs.AddExpression(expr)
	}

	return rs
}

//=============================================================================

func convertFunctionCallStatement(tree parser.IFunctionCallContext) *statement.FunctionCallStatement {
	fqid := convertFQIdentifier(tree.FqIdentifier())
	if ! tool.IsLowerCase(fqid) {
		raiseError(tree.GetParser(), "Fully qualified function call must be in lowercase: "+ fqid.String())
	}

	s := statement.NewFunctionCallStatement(fqid)

	for _,e := range tree.ParamsExpression().AllExpression() {
		expr := convertExpression(e)
		s.AddExpression(expr)
	}

	return s
}

//=============================================================================
//===
//=== Expressions
//===
//=============================================================================

func convertExpression(tree parser.IExpressionContext) expression.Expression {
	if tree == nil {
		return nil
	}

	//--- First, try with an arithmetic expression

	ae := convertArithmeticExpression(tree)
	if ae != nil {
		return ae
	}

	//--- Then try with a relational expression

	re := convertRelationalExpression(tree)
	if re != nil {
		return re
	}

	//--- Now, try with AND/OR

	be := convertBooleanExpression(tree)
	if be != nil {
		return be
	}

	//--- Lastly, other possibilities

	unaryExpr := tree.UnaryExpression()
	listExpr  := tree.ListExpression()
	mapExpr   := tree.MapExpression()

	if unaryExpr != nil {
		return convertUnaryExpression(unaryExpr)
	} else if listExpr != nil {
		return convertListExpression(listExpr)
	} else if mapExpr != nil {
		return convertMapExpression(mapExpr)
	} else {
		panic("Unknown expression type : " + tree.GetText())
	}
}

//=============================================================================

func convertArithmeticExpression(tree parser.IExpressionContext) expression.Expression {
	if tree.STAR() != nil {
		return expression.NewMultExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.SLASH() != nil {
		return expression.NewDivExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.PLUS() != nil {
		return expression.NewAddExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}
	if tree.MINUS() != nil {
		return expression.NewSubExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	return nil
}

//=============================================================================

func convertRelationalExpression(tree parser.IExpressionContext) expression.Expression {
	if tree.EQUAL() != nil {
		return expression.NewEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.LESS_OR_EQUAL() != nil {
		return expression.NewLessOrEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.GREATER_OR_EQUAL() != nil {
		return expression.NewGreaterOrEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.LESS_THAN() != nil {
		return expression.NewLessThanExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.GREATER_THAN() != nil {
		return expression.NewGreaterThanExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.NOT_EQUAL() != nil {
		return expression.NewNotEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	return nil
}

//=============================================================================

func convertBooleanExpression(tree parser.IExpressionContext) expression.Expression {
	if tree.AllAND() != nil {
		expr := buildExpressionList(tree.AllExpression())
		return expression.NewAndExpression(expr)
	}

	if tree.AllOR() != nil {
		expr := buildExpressionList(tree.AllExpression())
		return expression.NewOrExpression(expr)
	}

	if tree.NOT() != nil {
		expr := convertExpression(tree.Expression(0))
		return expression.NewNotExpression(expr)
	}

	return nil
}

//=============================================================================

func buildExpressionList(list []parser.IExpressionContext) []expression.Expression {
	var res []expression.Expression

	for _, expr := range list {
		res = append(res, convertExpression(expr))
	}

	return res
}

//=============================================================================

func convertUnaryExpression(tree parser.IUnaryExpressionContext) expression.Expression {
	identExpr := tree.IdentifierExpression()
	exprParen := tree.ExpressionInParenthesis()
	constValue:= tree.ConstantValueExpression()
	funcCall  := tree.FunctionCallExpression()

	if identExpr != nil {
		return convertIdentifierExpression(identExpr)
	} else if exprParen != nil {
		return convertExpression(exprParen.Expression())
	} else if constValue != nil {
		return convertConstantValueExpression(constValue)
	} else if funcCall != nil {
		return convertFunctionCallExpression(funcCall)
	}

	if tree.PLUS() != nil {
		return convertUnaryExpression(tree.UnaryExpression())
	}

	if tree.MINUS() != nil {
		value := values.NewIntValue(int64(-1))
		return expression.NewMultExpression(expression.NewConstantValueExpression(value), convertUnaryExpression(tree.UnaryExpression()))
	}

	panic("Unknown unary expression type : " + tree.GetText())
}

//=============================================================================

func convertListExpression(tree parser.IListExpressionContext) expression.Expression {
	t  := convertType(tree.ListType().Type_())
	lv := values.NewListValue(t)
	le := expression.NewListExpression(lv)

	if tree.InitialListValues() != nil {
		for _,e := range tree.InitialListValues().AllExpression() {
			le.AddExpression(convertExpression(e))
		}
	}

	return le
}

//=============================================================================

func convertMapExpression(tree parser.IMapExpressionContext) expression.Expression {
	kt  := convertKeyType(tree.MapType().KeyType())
	vt  := convertType   (tree.MapType().Type_())
	mv  := values.NewMapValue(kt, vt)
	mex := expression.NewMapExpression(mv)

	if tree.InitialMapValues() != nil {
		for _,me := range tree.InitialMapValues().AllKeyValueCouple() {
			k := convertKeyValue  (me.KeyValue())
			e := convertExpression(me.Expression())
			mex.Set(k, e)
		}
	}

	return mex
}

//=============================================================================

func convertIdentifierExpression(tree parser.IIdentifierExpressionContext) expression.Expression {
	name     := convertFQIdentifier(tree.FqIdentifier())
	accessor := tree.AccessorExpression()

	var acc expression.Expression

	if accessor != nil {
		acc = convertExpression(accessor.Expression())
	}

	return expression.NewIdentifierExpression(name, acc)
}

//=============================================================================

func convertFunctionCallExpression(tree parser.IFunctionCallExpressionContext) expression.Expression {
	name   := convertFQIdentifier(tree.FqIdentifier())
	params := tree.ParamsExpression()
	list   := buildExpressionList(params.AllExpression())

	return expression.NewFunctionCallExpression(name, list)
}

//=============================================================================

func convertConstantValueExpression(tree parser.IConstantValueExpressionContext) expression.Expression {
	value := convertConstantValue(tree)
	return expression.NewConstantValueExpression(value)
}

//=============================================================================

func convertConstantValue(tree parser.IConstantValueExpressionContext) values.Value {
	//--- Integers

	intVal := tree.INT_VALUE()
	if intVal != nil {
		return convertConstantInt(tree.GetParser(), intVal)
	}

	//--- Reals

	realVal := tree.REAL_VALUE()
	if realVal != nil {
		return convertConstantReal(tree.GetParser(), realVal)
	}

	//--- Time objects

	timeVal := tree.TimeValue()
	if timeVal != nil {
		return convertConstantTime(tree.GetParser(), timeVal)
	}

	//--- Date objects

	dateVal := tree.DateValue()
	if dateVal != nil {
		return convertConstantDate(tree.GetParser(), dateVal)
	}

	//--- Strings

	strVal  := tree.STRING_VALUE()
	if strVal != nil {
		return convertConstantString(strVal)
	}

	//--- Booleans

	boolVal := tree.BoolValue()
	if boolVal != nil {
		v := boolVal.GetText() == "true"

		return values.NewBoolValue(v)
	}

	//--- Errors

	errVal := tree.ErrorValue()
	if errVal != nil {
		text := errVal.STRING_VALUE().GetText()
		text  = text[1:len(text)-1]

		return values.NewErrorValue(text)
	}

	//--- Unknown

	panic("Unknown constant expression type : " + tree.GetText())
	return nil
}

//=============================================================================

func convertConstantInt(parser antlr.Parser, val antlr.TerminalNode) values.Value {
	i,err := strconv.Atoi(val.GetText())
	if err != nil {
		raiseError(parser, "invalid integer value : " + val.GetText())
	}

	return values.NewIntValue(int64(i))
}

//=============================================================================

func convertConstantReal(parser antlr.Parser, val antlr.TerminalNode) values.Value {
	f,err := strconv.ParseFloat(val.GetText(), 64)
	if err != nil {
		raiseError(parser, "Invalid real value : " + val.GetText())
	}

	return values.NewRealValue(f)
}

//=============================================================================

func convertConstantTime(parser antlr.Parser, val parser.ITimeValueContext) values.Value {
	hh,_ := strconv.Atoi(val.INT_VALUE(0).GetText())
	mm,_ := strconv.Atoi(val.INT_VALUE(1).GetText())

	v := data.NewTime(hh, mm)

	if !v.IsValid() {
		raiseError(parser, "Invalid time value : " + val.GetText())
	}

	return values.NewTimeValue(v)
}

//=============================================================================

func convertConstantDate(parser antlr.Parser, val parser.IDateValueContext) values.Value {
	y,_ := strconv.Atoi(val.INT_VALUE(0).GetText())
	m,_ := strconv.Atoi(val.INT_VALUE(1).GetText())
	d,_ := strconv.Atoi(val.INT_VALUE(2).GetText())

	v := data.NewDate(y, m, d)

	if !v.IsValid() {
		raiseError(parser, "Invalid date value : " + val.GetText())
	}

	return values.NewDateValue(v)
}

//=============================================================================

func convertConstantString(val antlr.TerminalNode) values.Value {
	text := val.GetText()
	text  = text[1:len(text)-1]

	return values.NewStringValue(text)
}

//=============================================================================

func convertKeyValue(tree parser.IKeyValueContext) values.Value{
	//--- Integers

	intVal  := tree.INT_VALUE()
	if intVal != nil {
		return convertConstantInt(tree.GetParser(), intVal)
	}

	//--- Time objects

	timeVal := tree.TimeValue()
	if timeVal != nil {
		return convertConstantTime(tree.GetParser(), timeVal)
	}

	//--- Date objects

	dateVal := tree.DateValue()
	if dateVal != nil {
		return convertConstantDate(tree.GetParser(), dateVal)
	}

	//--- Strings

	strVal  := tree.STRING_VALUE()
	if strVal != nil {
		return convertConstantString(strVal)
	}

	panic("Unknown constant expression type : " + tree.GetText())
}

//=============================================================================

func convertType(tree parser.ITypeContext) types.Type {
	if tree.ListType() != nil {
		lt := tree.ListType()
		subType := convertType(lt.Type_())
		return types.NewListType(subType)
	}

	if tree.MapType() != nil {
		mt := tree.MapType()
		keyType := convertKeyType(mt.KeyType())
		valType := convertType(mt.Type_())
		return types.NewMapType(keyType, valType)
	}

	if tree.INT() != nil {
		return types.NewIntType()
	}

	if tree.REAL() != nil {
		return types.NewRealType()
	}

	if tree.BOOL() != nil {
		return types.NewBoolType()
	}

	if tree.STRING() != nil {
		return types.NewStringType()
	}

	if tree.TIME() != nil {
		return types.NewTimeType()
	}

	if tree.DATE() != nil {
		return types.NewDateType()
	}

	if tree.TIMESERIES() != nil {
		return types.NewTimeseriesType()
	}

	if tree.ERROR() != nil {
		return types.NewErrorType()
	}

	return types.NewToFindOutType(tree.GetText(), data.NewInfo(tree))
}

//=============================================================================

func convertKeyType(tree parser.IKeyTypeContext) types.Type {
	if tree.INT() != nil {
		return types.NewIntType()
	}

	if tree.STRING() != nil {
		return types.NewStringType()
	}

	if tree.TIME() != nil {
		return types.NewTimeType()
	}

	if tree.DATE() != nil {
		return types.NewDateType()
	}

	panic("Unknown key type : " + tree.GetText())
	return nil
}

//=============================================================================

func convertFQIdentifier(tree parser.IFqIdentifierContext) *expression.FQIdentifier {
	fqi := expression.NewFQIdentifier()

	for _, s := range tree.AllIDENTIFIER() {
		name := s.GetText()
		fqi.AddScope(name)
	}

	return fqi
}

//=============================================================================

func convertSimplifiedExpression(tree parser.ISimplifiedExpressionContext) expression.Expression {
	if tree.STAR() != nil {
		return expression.NewMultExpression(
					convertSimplifiedExpression(tree.SimplifiedExpression(0)),
					convertSimplifiedExpression(tree.SimplifiedExpression(1)))
	}

	if tree.SLASH() != nil {
		return expression.NewDivExpression(
					convertSimplifiedExpression(tree.SimplifiedExpression(0)),
					convertSimplifiedExpression(tree.SimplifiedExpression(1)))
	}

	if tree.PLUS() != nil {
		return expression.NewAddExpression(
					convertSimplifiedExpression(tree.SimplifiedExpression(0)),
					convertSimplifiedExpression(tree.SimplifiedExpression(1)))
	}
	if tree.MINUS() != nil {
		return expression.NewSubExpression(
					convertSimplifiedExpression(tree.SimplifiedExpression(0)),
					convertSimplifiedExpression(tree.SimplifiedExpression(1)))
	}

	//--- Lastly, other possibilities

	identExpr := tree.IdentifierExpression()
	constValue:= tree.ConstantValueExpression()

	if identExpr != nil {
		return convertIdentifierExpression(identExpr)
	} else if constValue != nil {
		return convertConstantValueExpression(constValue)
	}

	if tree.PLUS() != nil {
		return convertSimplifiedExpression(tree.SimplifiedExpression(0))
	}

	if tree.MINUS() != nil {
		value := values.NewIntValue(int64(-1))
		return expression.NewMultExpression(
					expression.NewConstantValueExpression(value),
					convertSimplifiedExpression(tree.SimplifiedExpression(0)))
	}

	panic("Unknown simplified expression type : " + tree.GetText())
}

//=============================================================================
//===
//=== Other private functions
//===
//=============================================================================

func raiseError(p antlr.Parser, message string) {
	p.NotifyErrorListeners(message, nil, nil)
}

//=============================================================================
