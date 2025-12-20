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

package types

import (
	"github.com/tradalia/sick-engine/core/data"
	"github.com/tradalia/sick-engine/core/interfaces"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================
//===
//=== Class
//===
//=============================================================================

type ClassType struct {
	name        string
	Properties  []*Property
	Functions   []*Function
	Info        *parser.Info
	classScope  interfaces.Scope
}

//=============================================================================

func NewClassType(name string, info *parser.Info) *ClassType {
	return &ClassType{
		name : name,
		Info : info,
	}
}

//=============================================================================

func (t *ClassType) AddProperty(p *Property) {
	t.Properties = append(t.Properties, p)
}

//=============================================================================

func (t *ClassType) AddFunction(f *Function) {
	t.Functions = append(t.Functions, f)
}

//=============================================================================

func (t *ClassType) Code() int8 {
	return CodeClass
}

//=============================================================================

func (t *ClassType) String() string {
	return "class=("+ t.name +")"
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (t *ClassType) Id() string {
	return t.name
}

//=============================================================================

func (t *ClassType) Kind() interfaces.Kind {
	return interfaces.KindClass
}

//=============================================================================

func (t *ClassType) Specie() interfaces.Specie {
	return interfaces.SpecieType
}

//=============================================================================

func (t *ClassType) InitScope(parent interfaces.Scope) *parser.ParseError {
	t.classScope = data.NewSymbolTable()

	for _, prop := range t.Properties {
		if !t.classScope.Define(prop) {
			return parser.NewParseErrorFromInfo(prop.Info, "property duplicated in class: "+ prop.Id())
		}
	}

	for _, f := range t.Functions {
		if !t.classScope.Define(f) {
			return parser.NewParseErrorFromInfo(f.Info, "function duplicated in class: "+ f.Id())
		}

		err := f.InitScope(parent)
		if err != nil {
			return err
		}

		f.SetEmbedder(t)
	}

	return nil
}

//=============================================================================

func (t *ClassType) Scope() interfaces.Scope {
	return t.classScope
}

//=============================================================================
//===
//=== Property
//===
//=============================================================================

type Property struct {
	name  string
	Type  Type
	Info *parser.Info
}

//=============================================================================

func NewProperty(name string, type_ Type, info *parser.Info) *Property {
	return &Property{
		name: name,
		Type: type_,
		Info: info,
	}
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (p *Property) Id() string {
	return p.name
}

//=============================================================================

func (p *Property) Kind() interfaces.Kind {
	return interfaces.KindProperty
}

//=============================================================================

func (p *Property) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================

func (p *Property) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (p *Property) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
