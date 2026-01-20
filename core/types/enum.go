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
	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Enum
//===
//=============================================================================

type EnumType struct {
	name  string
	IsInt bool
	Items []*EnumItem
	Info  *parser.Info
	scope interfaces.Scope
}

//=============================================================================

func NewEnumType(name string, isInt bool, info *parser.Info) *EnumType {
	return &EnumType{
		name:  name,
		IsInt: isInt,
		Info:  info,
	}
}

//=============================================================================

func (e *EnumType) AddItem(i *EnumItem) {
	e.Items = append(e.Items, i)
}

//=============================================================================

func (e *EnumType) Size() int {
	return len(e.Items)
}

//=============================================================================

func (e *EnumType) AssignCodes() {
	for i, item := range e.Items {
		item.Code = i + 1
	}
}

//=============================================================================

func (e *EnumType) Code() int8 {
	return CodeEnum
}

//=============================================================================

func (e *EnumType) String() string {
	return "enum=" + e.name
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (e *EnumType) Id() string {
	return e.name
}

//=============================================================================

func (e *EnumType) Kind() interfaces.Kind {
	return interfaces.KindEnum
}

//=============================================================================

func (e *EnumType) Specie() interfaces.Specie {
	return interfaces.SpecieType
}

//=============================================================================

func (e *EnumType) InitScope(parent interfaces.Scope) *parser.ParseError {
	e.scope = data.NewSymbolTable()

	for _, item := range e.Items {
		if !e.scope.Define(item) {
			return parser.NewParseErrorFromInfo(e.Info, "item duplicated in enum: "+item.Id())
		}
	}
	return nil
}

//=============================================================================

func (e *EnumType) Scope() interfaces.Scope {
	return e.scope
}

//=============================================================================
//===
//=== EnumItem
//===
//=============================================================================

type EnumItem struct {
	name  string
	Code  int
	Value string
}

//=============================================================================

func NewEnumItem(name string, code int, value string) *EnumItem {
	return &EnumItem{name, code, value}
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (e *EnumItem) Id() string {
	return e.name
}

//=============================================================================

func (e *EnumItem) Kind() interfaces.Kind {
	return interfaces.KindEnumItem
}

//=============================================================================

func (e *EnumItem) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================

func (e *EnumItem) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (e *EnumItem) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
