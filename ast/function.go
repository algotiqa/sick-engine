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
	"strings"

	"github.com/tradalia/sick-engine/ast/statement"
	"github.com/tradalia/sick-engine/core/types"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================
//===
//=== Function
//===
//=============================================================================

type Function struct {
	Name     string
	Class    string
	Params   []*Param
	Returns  []types.Type
	Block    *statement.Block
	Info     *parser.Info
}

//=============================================================================

func NewFunction(name string, info *parser.Info) *Function {
	return &Function{
		Name : name,
		Info : info,
	}
}
//=============================================================================

func (f *Function) Id() string {
	sb := strings.Builder{}
	sb.WriteString(f.Name)
	sb.WriteString("|")
	sb.WriteString(f.Class)
	sb.WriteString("|")

	for _,p := range f.Params {
		sb.WriteString("|")
		sb.WriteString(p.Type.String())
	}

	return sb.String()
}

//=============================================================================

func (f *Function) AddParam(p *Param) {
	f.Params = append(f.Params, p)
}

//=============================================================================

func (f *Function) AddReturnType(t types.Type) {
	f.Returns = append(f.Returns, t)
}

//=============================================================================
//===
//=== Param
//===
//=============================================================================

type Param struct {
	Name string
	Type types.Type
	Info *parser.Info
}

//=============================================================================

func NewParam(name string, t types.Type, info *parser.Info) *Param {
	return &Param{
		Name: name,
		Type: t,
		Info: info,
	}
}

//=============================================================================
