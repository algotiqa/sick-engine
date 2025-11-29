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

package core

import (
	"strings"

	"github.com/tradalia/sick-engine/ast"
	"github.com/tradalia/sick-engine/types"
)

//=============================================================================
//===
//=== Environment
//===
//=============================================================================

type Environment struct {
	constants  map[string]*ast.Constant
	variables  map[string]*ast.Variable
	functions  map[string]*ast.Function
	enums      map[string]*types.EnumType
	classes    map[string]*types.ClassType

	duplicates map[string]bool
	classFunc  []*ast.Function
	callStack *CallStack
}

//=============================================================================

func NewEnvironment() *Environment {
	return &Environment{
		constants  : make(map[string]*ast.Constant),
		variables  : make(map[string]*ast.Variable),
		functions  : make(map[string]*ast.Function),
		enums      : make(map[string]*types.EnumType),
		classes    : make(map[string]*types.ClassType),
		duplicates : make(map[string]bool),
		callStack  : NewCallStack(),
	}
}

//=============================================================================

func (e *Environment) AddScript(s *ast.Script) *ParseErrors {
	pes := NewParseErrors()

	for _,c := range s.Constants {
		e.addConstant(s.PackageName, c, pes)
	}

	for _,v := range s.Variables {
		e.addVariable(s.PackageName, v, pes)
	}

	for _,f := range s.Functions {
		e.addFunction(s.PackageName, f, pes)
	}

	for _,en := range s.Enums {
		e.addEnum(s.PackageName, en, pes)
	}

	for _,c := range s.Classes {
		e.addClass(s.PackageName, c, pes)
	}

	return pes
}

//=============================================================================

func (e *Environment) Build() *ParseError {
	err := e.resolveTypes()
	if err == nil {
		err = e.assignMethodsToClasses()
	}

	return err
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func (e *Environment) addConstant(pack string, c *ast.Constant, pes *ParseErrors) {
	fqn := pack + "." + c.Name

	if e.testDuplicatesAndSet(fqn) {
		pe := NewParseErrorFromInfo(c.Info, "constant name already used: " + fqn)
		pes.AddError(pe)
		return
	}

	e.constants[fqn] = c
}

//=============================================================================

func (e *Environment) addVariable(pack string, v *ast.Variable, pes *ParseErrors) {
	fqn := pack + "." + v.Name

	if e.testDuplicatesAndSet(fqn) {
		pe := NewParseErrorFromInfo(v.Info, "variable name already used: " + fqn)
		pes.AddError(pe)
		return
	}

	e.variables[fqn] = v
}

//=============================================================================

func (e *Environment) addFunction(pack string, f *ast.Function, pes *ParseErrors) {
	if f.Class != "" {
		//--- Function belongs to a class (it's a method)

		f.Class = pack +"."+ f.Class
		e.classFunc = append(e.classFunc, f)
	} else {
		//--- Function is standalone

		fqn := pack + "." + f.Id()

		if e.testDuplicatesAndSet(fqn) {
			pe := NewParseErrorFromInfo(f.Info, "function name already used: " + fqn)
			pes.AddError(pe)
			return
		}

		e.functions[fqn] = f
	}

	for _, p := range f.Params {
		addPackageToType(p.Type, pack)
	}

	for _, r := range f.Returns {
		addPackageToType(r, pack)
	}
}

//=============================================================================

func (e *Environment) addEnum(pack string, en *types.EnumType, pes *ParseErrors) {
	fqn := pack + "." + en.Name

	if e.testDuplicatesAndSet(fqn) {
		pe := NewParseErrorFromInfo(en.Info, "enum name already used: " + fqn)
		pes.AddError(pe)
		return
	}

	e.enums[fqn] = en
}

//=============================================================================

func (e *Environment) addClass(pack string, c *types.ClassType, pes *ParseErrors) {
	fqn := pack + "." + c.Name

	if e.testDuplicatesAndSet(fqn) {
		pe := NewParseErrorFromInfo(c.Info, "class name already used: " + fqn)
		pes.AddError(pe)
		return
	}

	e.classes[fqn] = c

	for _, p := range c.Properties {
		addPackageToType(p.Type, pack)
	}
}

//=============================================================================

func (e *Environment) testDuplicatesAndSet(fqName string) bool {
	if _, ok := e.duplicates[fqName]; ok {
		return true
	}

	e.duplicates[fqName] = true
	return false
}

//=============================================================================

func addPackageToType(t types.Type, pack string) {
	switch t.(type) {
		case *types.ToFindOutType:
			at := t.(*types.ToFindOutType)
			if !strings.Contains(at.Name, ".") {
				at.Name = pack +"."+ at.Name
			}
	}
}

//=============================================================================

func (e *Environment) resolveTypes() *ParseError {
	var err *ParseError

	for _, f := range e.functions {
		for _, p := range f.Params {
			p.Type,err = e.resolveType(p.Type)
			if err != nil {
				return err
			}
		}

		for i, r := range f.Returns {
			r, err = e.resolveType(r)
			if err != nil {
				return err
			}
			f.Returns[i] = r
		}
	}

	for _, c := range e.classes {
		for _, p := range c.Properties {
			p.Type,err = e.resolveType(p.Type)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//=============================================================================

func (e *Environment) resolveType(t types.Type) (types.Type, *ParseError) {
	switch t.(type) {
		case *types.ToFindOutType:
			at := t.(*types.ToFindOutType)
			rt := e.searchType(at.Name)
			if rt == nil {
				return t,NewParseErrorFromInfo(at.Info, "can't resolve type: " + at.Name)
			}
			return rt,nil
	}

	return t,nil
}

//=============================================================================

func (e *Environment) searchType(name string) types.Type {
	if t,ok := e.enums[name]; ok {
		return t
	}

	if t,ok := e.classes[name]; ok {
		return t
	}

	return nil
}

//=============================================================================

func (e *Environment) assignMethodsToClasses() *ParseError{
	for _, f := range e.classFunc {
		if c,ok := e.classes[f.Class]; ok {
			if !c.AddMethod(f) {
				return NewParseErrorFromInfo(f.Info, "method name already used in class: " + f.Name)
			}
		} else {
			return NewParseErrorFromInfo(f.Info, "can't resolve class for function: " + f.Class)
		}
	}

	e.classFunc = nil
	return nil
}

//=============================================================================
