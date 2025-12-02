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
//=== Class
//===
//=============================================================================

type ClassType struct {
	Name        string
	Properties  map[string]*Property
	Methods     map[string]Method
	Info       *parser.Info
}

//=============================================================================

func NewClassType(name string, info *parser.Info) *ClassType {
	return &ClassType{
		Name      : name,
		Properties: map[string]*Property{},
		Methods   : map[string]Method{},
		Info      : info,
	}
}

//=============================================================================

func (t *ClassType) AddProperty(p *Property) bool {
	if _,ok := t.Properties[p.Name]; ok {
		return false
	}

	t.Properties[p.Name] = p
	return true
}

//=============================================================================

func (t *ClassType) AddMethod(m Method) bool {
	if _,ok := t.Methods[m.Id()]; ok {
		return false
	}

	t.Methods[m.Id()] = m
	return true
}

//=============================================================================

func (t *ClassType) Id() int8 {
	return IdClass
}

//=============================================================================

func (t *ClassType) String() string {
	return "class=("+ t.Name +")"
}

//=============================================================================
//===
//=== Property
//===
//=============================================================================

type Property struct {
	Name  string
	Type  Type
	Info *parser.Info
}

//=============================================================================

func NewProperty(name string, type_ Type, info *parser.Info) *Property {
	return &Property{
		Name: name,
		Type: type_,
		Info: info,
	}
}

//=============================================================================
//===
//=== Method
//===
//=============================================================================

type Method interface {
	Id() string
}

//=============================================================================
