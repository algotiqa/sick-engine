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
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================
//===
//=== Enum
//===
//=============================================================================

type EnumType struct {
	Name     string
	IsInt    bool
	Items    []*EnumItem
	ItemMap  map[string]*EnumItem
	Info    *parser.Info
}

//=============================================================================

func NewEnumType(name string, isInt bool, info *parser.Info) *EnumType {
	return &EnumType{
		Name   : name,
		IsInt  : isInt,
		ItemMap: make(map[string]*EnumItem),
		Info   : info,
	}
}

//=============================================================================

func (e *EnumType) AddItem(i *EnumItem) bool {
	if _,ok := e.ItemMap[i.Name]; ok {
		return false
	}

	e.Items = append(e.Items, i)
	e.ItemMap[i.Name] = i
	return true
}

//=============================================================================

func (e *EnumType) Size() int {
	return len(e.Items)
}

//=============================================================================

func (e *EnumType) AssignCodes() {
	for i, item := range e.Items {
		item.Code = i+1
	}
}

//=============================================================================

func (e *EnumType) Id() int8 {
	return IdEnum
}

//=============================================================================

func (e *EnumType) String() string {
	return "enum="+ e.Name
}

//=============================================================================
//===
//=== EnumItem
//===
//=============================================================================

type EnumItem struct {
	Name  string
	Code  int
	Value string
}

//=============================================================================

func NewEnumItem(name string, code int, value string) *EnumItem {
	return &EnumItem{name, code, value}
}

//=============================================================================
