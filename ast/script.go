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
	"github.com/tradalia/sick-engine/data"
	"github.com/tradalia/sick-engine/types"
)

//=============================================================================

type Script struct {
	Filename    string
	PackageName string
	Constants   []*Constant
	Variables   []*Variable
	Functions   []*Function
	Enums       []*types.EnumType
	Classes     []*types.ClassType

	identifiers map[string]bool
}

//=============================================================================

func NewScript(filename string) *Script {
	return &Script{
		Filename   : filename,
		identifiers: make(map[string]bool),
	}
}

//=============================================================================

func (s *Script) AddConstant(c *Constant) bool {
	if s.testAndSet(c.Name) {
		return false
	}

	s.Constants = append(s.Constants, c)
	return true
}

//=============================================================================

func (s *Script) AddVariable(v *Variable) bool {
	if s.testAndSet(v.Name) {
		return false
	}

	s.Variables = append(s.Variables, v)
	return true
}

//=============================================================================

func (s *Script) AddFunction(f *Function) bool {
	if s.testAndSet(f.Id()) {
		return false
	}

	s.Functions = append(s.Functions, f)
	return true
}

//=============================================================================

func (s *Script) AddEnum(e *types.EnumType) bool {
	if s.testAndSet(e.Name) {
		return false
	}

	s.Enums = append(s.Enums, e)
	return true
}

//=============================================================================

func (s *Script) AddClass(c *types.ClassType) bool {
	if s.testAndSet(c.Name) {
		return false
	}

	s.Classes = append(s.Classes, c)
	return true
}

//=============================================================================

func (s *Script) testAndSet(name string) bool {
	if _, ok := s.identifiers[name]; ok {
		return true
	}

	s.identifiers[name] = true
	return false
}

//=============================================================================
//===
//=== Constant
//===
//=============================================================================

type Constant struct {
	Name  string
	Value expression.Expression
	Info  *data.Info
}

//=============================================================================

func NewConstant(name string, e expression.Expression, info *data.Info) *Constant {
	return &Constant{
		Name : name,
		Value: e,
		Info : info,
	}
}

//=============================================================================
//===
//=== Variable
//===
//=============================================================================

type Variable struct {
	Name  string
	Value expression.Expression
	Info  *data.Info
}

//=============================================================================

func NewVariable(name string, e expression.Expression, info *data.Info) *Variable {
	return &Variable{
		Name : name,
		Value: e,
		Info : info,
	}
}

//=============================================================================
