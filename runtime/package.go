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
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Package
//===
//=============================================================================

type Package struct {
	name      string
	scope     interfaces.Scope
	classFunc []*types.Function
}

//=============================================================================

func NewPackage(name string, parent interfaces.Scope) *Package {
	return &Package{
		name:  name,
		scope: parent.Push(),
	}
}

//=============================================================================

func (p *Package) AddFunction(f *types.Function) {
	f.Class.Pack = p.name
	p.classFunc = append(p.classFunc, f)
}

//=============================================================================

func (p *Package) AssignMethodsToClasses() *parser.ParseError {
	for _, f := range p.classFunc {
		s := p.scope.Resolve(f.Class.Name)
		if s != nil {
			if s.Kind() != interfaces.KindClass {
				return parser.NewParseErrorFromInfo(f.Info, "function doesn't reference a class: "+f.Class.Name)
			}

			c := s.(*types.ClassType)
			c.AddFunction(f)
		} else {
			return parser.NewParseErrorFromInfo(f.Info, "can't resolve class for function: "+f.Class.String())
		}
	}

	p.classFunc = nil
	return nil
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (p *Package) Id() string {
	return p.name
}

//=============================================================================

func (p *Package) Kind() interfaces.Kind {
	return interfaces.KindPackage
}

//=============================================================================

func (p *Package) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================
//--- A package already has a scope populated with data
//--- We just need to initialize the scopes of the children

func (p *Package) InitScope(parent interfaces.Scope) *parser.ParseError {
	for s := range p.scope.AllSymbols() {
		err := s.InitScope(p.scope)
		if err != nil {
			return err
		}
	}

	return nil
}

//=============================================================================

func (p *Package) Scope() interfaces.Scope {
	return p.scope
}

//=============================================================================
