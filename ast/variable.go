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
	"github.com/tradalia/sick-engine/ast/expression"
	"github.com/tradalia/sick-engine/core/interfaces"
	"github.com/tradalia/sick-engine/core/types"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================
//===
//=== Variable
//===
//=============================================================================

type Variable struct {
	name  string
	Value expression.Expression
	Type  types.Type
	Info  *parser.Info
}

//=============================================================================

func NewVariable(name string, e expression.Expression, info *parser.Info) *Variable {
	return &Variable{
		name : name,
		Value: e,
		Info : info,
	}
}

//=============================================================================

func (v *Variable) SetupType(s interfaces.Scope, depth int) error {
	t,err := v.Value.ResolveType(s, nil, depth)
	if err != nil {
		return err
	}

	v.Type = t
	return nil
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (v *Variable) Id() string {
	return v.name
}

//=============================================================================

func (v *Variable) Kind() interfaces.Kind {
	return interfaces.KindConst
}

//=============================================================================

func (v *Variable) Specie() interfaces.Specie {
	return interfaces.SpecieObject
}

//=============================================================================

func (v *Variable) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (v *Variable) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
