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

	"github.com/tradalia/sick-engine/ast/expression"
	"github.com/tradalia/sick-engine/ast/statement"
	"github.com/tradalia/sick-engine/core/types"
	"github.com/tradalia/sick-engine/parser"
	"github.com/tradalia/sick-engine/tool"
)

//=============================================================================

const (
	DefaultPackage = "main"
)

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

		if !tool.StartsWithLowerCase(id) {
			parser.RaiseError(tree.GetParser(), "package's name must start with a lowercase character: "+ id)
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
				parser.RaiseError(cds.GetParser(), "constant's name already used: "+ c.Name)
			}
		}
	}
}

//=============================================================================

func convertConstantDef(tree parser.IConstantDefContext) *Constant {
	name  := tree.IDENTIFIER().GetText()
	value := tree.Expression()

	if !tool.StartsWithUpperCase(name) {
		parser.RaiseError(tree.GetParser(), "constant's name must start with an uppercase character: "+ name)
	}

	return NewConstant(name, expression.ConvertExpression(value), parser.NewInfo(tree))
}

//=============================================================================
//=== Variables
//=============================================================================

func convertVariables(script *Script, tree []parser.IVariablesDefContext) {
	for _, vds := range tree {
		for _,vd := range vds.AllVariableDef() {
			v := convertVariableDef(vd)
			if !script.AddVariable(v) {
				parser.RaiseError(vd.GetParser(), "variable's name already used: "+ v.Name)
			}
		}
	}
}

//=============================================================================

func convertVariableDef(tree parser.IVariableDefContext) *Variable {
	name  := tree.IDENTIFIER().GetText()
	value := tree.Expression()

	if !tool.StartsWithLowerCase(name) {
		parser.RaiseError(tree.GetParser(), "variable's name must start with a lowercase character: "+ name)
	}

	return NewVariable(name, expression.ConvertExpression(value), parser.NewInfo(tree))
}

//=============================================================================
//=== Functions
//=============================================================================

func convertFunctions(script *Script, tree []parser.IFunctionDefContext) {
	for _, fd := range tree {
		f := convertFunctionDef(fd)
		if !script.AddFunction(f) {
			parser.RaiseError(fd.GetParser(), "function's name already used: "+ f.Name)
		}
	}
}

//=============================================================================

func convertFunctionDef(tree parser.IFunctionDefContext) *Function {
	name := tree.IDENTIFIER().GetText()

	if !tool.StartsWithLowerCase(name) {
		parser.RaiseError(tree.GetParser(), "function's name must start with a lowercase character: "+ name)
	}

	f := NewFunction(name, parser.NewInfo(tree))

	if tree.Class() != nil {
		fqi := NewFQIdentifier(tree.Class().FqIdentifier())
		if !tool.StartsWithUpperCase(fqi.Name) {
			parser.RaiseError(tree.GetParser(), "function's class name must start with an uppercase character: "+ fqi.Name)
		}

		f.Class = tree.Class().FqIdentifier().GetText()
	}

	convertFunctionParams (f, tree.Parameters())
	convertFunctionResults(f, tree.Results())
	f.Block = statement.ConvertBlock(tree.Block())

	return f
}

//=============================================================================

func convertFunctionParams(f *Function, params parser.IParametersContext) {
	names := map[string]bool{}

	for _, pd := range params.AllParameterDecl() {
		name  := pd.IDENTIFIER().GetText()
		type_ := pd.Type_()

		if !tool.StartsWithLowerCase(name) {
			parser.RaiseError(params.GetParser(), "function's parameter must start with a lowercase character: "+ name)
		}

		if _,ok := names[name]; ok {
			parser.RaiseError(params.GetParser(), "function's parameter name already used: "+ name)
		}

		names[name] = true
		t := types.ConvertType(type_)
		p := NewParam(name,t, parser.NewInfo(pd))
		f.AddParam(p)
	}
}

//=============================================================================

func convertFunctionResults(f *Function, results parser.IResultsContext) {
	if results != nil {
		for _,r := range results.AllType_() {
			t := types.ConvertType(r)
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
			parser.RaiseError(ed.GetParser(), "enum's name already used: "+ e.Name)
		}
	}
}

//=============================================================================

func convertEnumDef(tree parser.IEnumDefContext) *types.EnumType {
	name := tree.IDENTIFIER().GetText()
	e := types.NewEnumType(name, true, parser.NewInfo(tree))

	if !tool.StartsWithUpperCase(name) {
		parser.RaiseError(tree.GetParser(), "enum's name must start with an uppercase character: "+ name)
	}

	countInts := 0
	countStr  := 0

	for _, ei := range tree.AllEnumItem() {
		i := convertEnumItem(ei)
		if !e.AddItem(i) {
			parser.RaiseError(tree.GetParser(), "enum's item already used: "+ i.Name)
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
		parser.RaiseError(tree.GetParser(), "mixing integers and strings in enum: "+ name)
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
		parser.RaiseError(tree.GetParser(), "enum item's name must start with an uppercase character: "+ name)
	}

	enumVal := tree.EnumValue()
	if enumVal != nil {
		intValue := enumVal.INT_VALUE()
		strValue := enumVal.STRING_VALUE()

		if intValue != nil {
			iValue,err := strconv.Atoi(intValue.GetText())
			if err != nil {
				parser.RaiseError(tree.GetParser(), "Invalid integer value for enum : " + intValue.GetText())
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
			parser.RaiseError(cd.GetParser(), "class's name already used: "+ c.Name)
		}
	}
}

//=============================================================================

func convertClassDef(tree parser.IClassDefContext) *types.ClassType {
	name := tree.IDENTIFIER().GetText()
	c := types.NewClassType(name, parser.NewInfo(tree))

	if !tool.StartsWithUpperCase(name) {
		parser.RaiseError(tree.GetParser(), "class's name must start with an uppercase character: "+ name)
	}

	for _, p := range tree.AllProperty() {
		if !c.AddProperty(convertProperty(p)) {
			parser.RaiseError(p.GetParser(), "class's property already used: "+ c.Name)
		}
	}

	return c
}

//=============================================================================

func convertProperty(tree parser.IPropertyContext) *types.Property {
	name  := tree.IDENTIFIER().GetText()
	type_ := tree.Type_()

	if !tool.StartsWithLowerCase(name) {
		parser.RaiseError(tree.GetParser(), "class's property must start with a lowercase character: "+ name)
	}

	t := types.ConvertType(type_)

	return types.NewProperty(name,t, parser.NewInfo(tree))
}

//=============================================================================
