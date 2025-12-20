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

package runtime

import (
	"github.com/tradalia/sick-engine/ast"
	"github.com/tradalia/sick-engine/core/data"
	"github.com/tradalia/sick-engine/core/interfaces"
	"github.com/tradalia/sick-engine/core/types"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================
//===
//=== Const
//===
//=============================================================================

const MaxResolutionDepth = 16

//=============================================================================
//===
//=== Environment
//===
//=============================================================================

type Environment struct {
	scope interfaces.Scope
}

//=============================================================================

func NewEnvironment() *Environment {
	return &Environment{
		scope: data.NewSymbolTable(),
	}
}

//=============================================================================

func (e *Environment) AddScript(s *ast.Script) *parser.ParseErrors {
	pes := parser.NewParseErrors()
	pac := e.getOrCreatePackage(s.PackageName)

	for _,c := range s.Constants {
		e.addConstant(pac, c, pes)
	}

	for _,v := range s.Variables {
		e.addVariable(pac, v, pes)
	}

	for _,f := range s.Functions {
		e.addFunction(pac, f, pes)
	}

	for _,en := range s.Enums {
		e.addEnum(pac, en, pes)
	}

	for _,c := range s.Classes {
		e.addClass(pac, c, pes)
	}

	return pes
}

//=============================================================================

func (e *Environment) Build() *parser.ParseError {
	err := e.assignMethodsToClasses()
	if err == nil {
		err = e.convertToFindOutTypes()
		if err == nil {
			err = e.initScopes()
			if err == nil {
				err = e.resolveGlobalObjectTypes()
			}
		}
	}

	return err
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func (e *Environment) getOrCreatePackage(name string) *Package {
	pac := e.scope.ResolveLocally(name)
	if pac == nil {
		pac = NewPackage(name, e.scope)
		e.scope.Define(pac)
	}

	return pac.(*Package)
}

//=============================================================================

func (e *Environment) addConstant(pack *Package, c *ast.Constant, pes *parser.ParseErrors) {
	if !pack.scope.Define(c) {
		pe := parser.NewParseErrorFromInfo(c.Info, "constant name already used: " + c.Id())
		pes.AddError(pe)
		return
	}
}

//=============================================================================

func (e *Environment) addVariable(pack *Package, v *ast.Variable, pes *parser.ParseErrors) {
	if !pack.scope.Define(v) {
		pe := parser.NewParseErrorFromInfo(v.Info, "variable name already used: " + v.Id())
		pes.AddError(pe)
		return
	}
}

//=============================================================================

func (e *Environment) addFunction(pack *Package, f *types.Function, pes *parser.ParseErrors) {
	if f.Class != nil {
		//--- Function belongs to a class (it's a method)

		if f.Class.Pack != "" {
			pack = e.getOrCreatePackage(f.Class.Pack)
		}
		pack.AddFunction(f)
	} else {
		//--- Function is standalone

		if !pack.scope.Define(f) {
			pe := parser.NewParseErrorFromInfo(f.Info, "function name already used: " + f.Id())
			pes.AddError(pe)
			return
		}
	}
}

//=============================================================================

func (e *Environment) addEnum(pack *Package, en *types.EnumType, pes *parser.ParseErrors) {
	if !pack.scope.Define(en) {
		pe := parser.NewParseErrorFromInfo(en.Info, "enum name already used: " + en.Id())
		pes.AddError(pe)
		return
	}
}

//=============================================================================

func (e *Environment) addClass(pack *Package, c *types.ClassType, pes *parser.ParseErrors) {
	if !pack.scope.Define(c) {
		pe := parser.NewParseErrorFromInfo(c.Info, "class name already used: " + c.Id())
		pes.AddError(pe)
		return
	}
}

//=============================================================================
//===
//=== Semantic analysis
//===
//=============================================================================

//=============================================================================
//=== Move functions to proper classes if they are methods
//=============================================================================

func (e *Environment) assignMethodsToClasses() *parser.ParseError{
	for p := range e.scope.AllSymbols() {
		pack := p.(*Package)
		err  := pack.AssignMethodsToClasses()
		if err != nil {
			return err
		}
	}

	return nil
}

//=============================================================================
//=== Substitutes ToFindOutType with the real type for
//===  - General functions (parameters and return types)
//===  - Class properties
//===  - Class methods (parameters and return types)
//=============================================================================

func (e *Environment) convertToFindOutTypes() *parser.ParseError {
	var err *parser.ParseError

	for p := range e.scope.AllSymbols() {
		pack := p.(*Package)
		for s := range pack.scope.AllSymbols() {
			if s.Kind() == interfaces.KindFunction {
				err = e.resolveFunctionTypes(pack.scope, s.(*types.Function))
				if err != nil {
					return err
				}
			}

			if s.Kind() == interfaces.KindClass {
				c := s.(*types.ClassType)
				for _, p := range c.Properties {
					p.Type,err = e.resolveType(pack.scope, p.Type)
					if err != nil {
						return err
					}
				}
				for _, f := range c.Functions {
					err = e.resolveFunctionTypes(pack.scope, f)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

//=============================================================================

func (e *Environment) resolveFunctionTypes(scope interfaces.Scope, f *types.Function) *parser.ParseError {
	var err *parser.ParseError

	for _, p := range f.Params {
		p.Type,err = e.resolveType(scope, p.Type)
		if err != nil {
			return err
		}
	}

	for i, r := range f.Returns {
		r, err = e.resolveType(scope, r)
		if err != nil {
			return err
		}
		f.Returns[i] = r
	}

	return nil
}

//=============================================================================

func (e *Environment) resolveType(scope interfaces.Scope, t types.Type) (types.Type, *parser.ParseError) {
	switch t.(type) {
		case *types.ToFindOutType:
			at := t.(*types.ToFindOutType)
			rt := e.searchType(scope, at.Name)
			if rt == nil {
				return t,parser.NewParseErrorFromInfo(at.Info, "can't resolve type: " + at.Name.String())
			}
			return rt,nil
	}

	return t,nil
}

//=============================================================================

func (e *Environment) searchType(scope interfaces.Scope, name *data.FQIdentifier) types.Type {
	if name.Pack != "" {
		s := scope.ResolveGlobally(name.Pack)
		if s == nil || s.Kind() != interfaces.KindPackage {
			return nil
		}

		scope = s.(*Package).scope
	}

	t := scope.ResolveLocally(name.Name)
	if t != nil {
		if t.Specie() == interfaces.SpecieType {
			return t.(types.Type)
		}
	}

	return nil
}

//=============================================================================
//=== Scope initialization
//===  - Adds all entities to containers, checking for duplicates
//=============================================================================

func (e *Environment) initScopes() *parser.ParseError {
	for s := range e.scope.AllSymbols() {
		err := s.InitScope(e.scope)
		if err != nil {
			return err
		}
	}
	return nil
}

//=============================================================================
//=== Type resolution for global Constants and variables
//=============================================================================

func (e *Environment) resolveGlobalObjectTypes() *parser.ParseError {
	//--- First, resolve type for global constants

	for s := range e.scope.AllSymbols() {
		if s.Kind() == interfaces.KindConst {
			c   := s.(*ast.Constant)
			err := c.SetupType(e.scope, MaxResolutionDepth)

			if err != nil {
				return parser.NewParseErrorFromInfo(c.Info, "cannot resolve type for constant '"+ c.Id() +"' : "+ err.Error())
			}
		}
	}

	//--- Once all constants are fine, resolve type for global variables

	for s := range e.scope.AllSymbols() {
		if s.Kind() == interfaces.KindVar {
			v   := s.(*ast.Variable)
			err := v.SetupType(e.scope, MaxResolutionDepth)

			if err != nil {
				return parser.NewParseErrorFromInfo(v.Info, "cannot resolve type for constant '"+ v.Id() +"' : "+ err.Error())
			}
		}
	}

	return nil
}

//=============================================================================
//=============================================================================

//func (e *Environment) validateFunctions() *parser.ParseError {
//	for s := range e.globalScope.AllSymbols() {
//		if s.Kind() == interfaces.KindFunction {
//			err := e.validateFunction(s.(*ast.Function))
//			if err != nil {
//				return err
//			}
//		}
//
//		if s.Kind() == interfaces.KindClass {
//			c := s.(*types.ClassType)
//			for _, f := range c.Methods {
//				err := e.validateFunction(f.(*ast.Function))
//				if err != nil {
//					return err
//				}
//			}
//		}
//	}
//
//	return nil
//}

//=============================================================================

//func (e *Environment) validateFunction(f *ast.Function) *parser.ParseError {
//	return nil
//}

//=============================================================================
