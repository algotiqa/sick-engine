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
	"strings"

	"github.com/tradalia/sick-engine/core/data"
	"github.com/tradalia/sick-engine/core/interfaces"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================
//===
//=== Function
//===
//=============================================================================

type Function struct {
	name     string
	Class    *data.FQIdentifier
	Params   []*Param
	Returns  []Type
	Block    *data.Block
	Info     *parser.Info
	scope    interfaces.Scope
}

//=============================================================================

func NewFunction(name string, info *parser.Info) *Function {
	return &Function{
		name : name,
		Info : info,
	}
}
//=============================================================================

func (f *Function) AddParam(p *Param) {
	f.Params = append(f.Params, p)
}

//=============================================================================

func (f *Function) AddReturnType(t Type) {
	f.Returns = append(f.Returns, t)
}

//=============================================================================

func (f *Function) SetEmbedder(embedder interfaces.Symbol) {
	f.Block.SetEmbedder(embedder)
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (f *Function) Id() string {
	sb := strings.Builder{}
	sb.WriteString(f.name)
	sb.WriteString("|")

	for _,p := range f.Params {
		sb.WriteString("|")
		sb.WriteString(p.Type.String())
	}

	return sb.String()
}

//=============================================================================

func (f *Function) Kind() interfaces.Kind {
	return interfaces.KindFunction
}

//=============================================================================

func (f *Function) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================

func (f *Function) InitScope(parent interfaces.Scope) *parser.ParseError {
	f.scope = parent.Push()

	for _, param := range f.Params {
		if !f.scope.Define(param) {
			return parser.NewParseErrorFromInfo(param.Info, "parameter duplicated in function: "+ param.Id())
		}
	}

	return f.Block.InitScope(f.scope)
}

//=============================================================================

func (f *Function) Scope() interfaces.Scope {
	return f.scope
}

//=============================================================================
//===
//=== Param
//===
//=============================================================================

type Param struct {
	Name string
	Type Type
	Info *parser.Info
}

//=============================================================================

func NewParam(name string, t Type, info *parser.Info) *Param {
	return &Param{
		Name: name,
		Type: t,
		Info: info,
	}
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (p *Param) Id() string {
	return p.Name
}

//=============================================================================

func (p *Param) Kind() interfaces.Kind {
	return interfaces.KindParameter
}

//=============================================================================

func (p *Param) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================

func (p *Param) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (p *Param) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
