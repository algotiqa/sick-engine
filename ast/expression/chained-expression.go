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

package expression

import (
	"errors"
	"strings"

	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Chained expression
//===
//=============================================================================

type ChainedExpression struct {
	HasThis    bool
	HasNew     bool
	Chain      []*ChainItem
	FQClass    *data.FQIdentifier
	InstParams []Expression
	info       *parser.Info
}

//=============================================================================

func NewChainedExpression(hasThis bool, hasNew bool, info *parser.Info) *ChainedExpression {
	return &ChainedExpression{
		HasThis: hasThis,
		HasNew:  hasNew,
		info:    info,
	}
}

//=============================================================================

func (e *ChainedExpression) AddItem(item *ChainItem) {
	e.Chain = append(e.Chain, item)
}

//=============================================================================

func (e *ChainedExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	if e.HasThis {
		if embedder == nil {
			return nil, errors.New("used 'this' outside of a class")
		}

		return e.resolveChainType(embedder, embedder, 0, depth)
	}

	if e.HasNew {
		return e.resolveWithNew(scope, embedder, depth)
	}

	return e.resolvePlain(scope, embedder, depth)
}

//=============================================================================

func (e *ChainedExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func (e *ChainedExpression) resolveWithNew(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	s := scope

	if e.FQClass.Pack != "" {
		sym := s.Resolve(e.FQClass.Pack)
		if sym == nil {
			return nil, errors.New("package referenced by class was not found: " + e.FQClass.Pack)
		}

		if sym.Kind() != interfaces.KindPackage {
			return nil, errors.New("class doesn't reference a package: " + e.FQClass.Pack)
		}

		s = sym.Scope()
	}

	sym := s.Resolve(e.FQClass.Name)

	if sym == nil {
		return nil, errors.New("class was not found in package: " + e.FQClass.String())
	}

	if sym.Kind() != interfaces.KindClass {
		return nil, errors.New("symbol doesn't reference a class: " + e.FQClass.String())
	}

	if len(e.Chain) == 0 {
		return sym.(types.Type), nil
	}

	return e.resolveChainType(sym, embedder, 0, depth)
}

//=============================================================================

func (e *ChainedExpression) resolvePlain(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	ci := e.Chain[0]

	sym, err := e.resolveChainItem(ci, scope, nil, depth)
	if err != nil {
		return nil, err
	}

	if sym == nil {
		return nil, errors.New("symbol was not found: " + ci.Name)
	}

	if sym.Kind() == interfaces.KindPackage {
	}

	return e.resolveChainType(sym, 1, depth)
}

//=============================================================================

func (e *ChainedExpression) resolveChainType(prefix, embedder interfaces.Symbol, index int, depth int) (types.Type, error) {
	for i := index; i < len(e.Chain); i++ {
		item := e.Chain[i]

		sym, err := e.resolveChainItem(item, scope, embedder, depth)

		s := scope.Resolve(item.Name)
		if s == nil {
			return nil, errors.New("'" + item.Name + "' was not found")
		}

	}

	return nil, nil
}

//=============================================================================

func (e *ChainedExpression) resolveChainItem(ci *ChainItem, scope interfaces.Scope, embedder interfaces.Symbol, depth int) (interfaces.Symbol, error) {
	sb := strings.Builder{}
	sb.WriteString(ci.Name)

	if ci.Params != nil {
		//--- Resolve function

		sb.WriteString("|")

		for _, p := range ci.Params {
			sb.WriteString("|")
			t, err := p.ResolveType(scope, embedder, depth)
			if err != nil {
				return nil, err
			}
			sb.WriteString(t.String())
		}
	}

	name := sb.String()
	sym := scope.Resolve(name)

	return sym, nil
}

//=============================================================================
//===
//=== ChainItem
//===
//=============================================================================

type ChainItem struct {
	Name     string
	Accessor Expression
	Params   []Expression
}

//=============================================================================

func NewChainItem(name string, accessor Expression, params []Expression) *ChainItem {
	return &ChainItem{
		Name:     name,
		Accessor: accessor,
		Params:   params,
	}
}

//=============================================================================
