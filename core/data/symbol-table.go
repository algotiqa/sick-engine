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

package data

import (
	"iter"
	"maps"

	"github.com/algotiqa/tiq-engine/core/interfaces"
)

//=============================================================================
//===
//=== Scope implementation
//===
//=============================================================================

type SymbolTable struct {
	previous *SymbolTable
	symbols  map[string]interfaces.Symbol
}

//=============================================================================

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]interfaces.Symbol),
	}
}

//=============================================================================

func (s *SymbolTable) Resolve(name string) interfaces.Symbol {
	currScope := s

	for currScope != nil {
		sym, ok := currScope.symbols[name]
		if ok {
			return sym
		}

		//if sym == nil {
		//	//--- Functions have Ids (not names) that must be processed differently
		//	for k,v := range s.symbols {
		//		if strings.HasPrefix(k,name+"|") {
		//			return v
		//		}
		//	}
		//}
		currScope = currScope.previous
	}

	return nil
}

//=============================================================================

func (s *SymbolTable) Define(symbol interfaces.Symbol) bool {
	_, ok := s.symbols[symbol.Id()]
	if ok {
		return false
	}

	s.symbols[symbol.Id()] = symbol
	return true
}

//=============================================================================

func (s *SymbolTable) Push() interfaces.Scope {
	ns := NewSymbolTable()
	ns.previous = s
	return ns
}

//=============================================================================

func (s *SymbolTable) Pop() interfaces.Scope {
	return s.previous
}

//=============================================================================

func (s *SymbolTable) AllSymbols() iter.Seq[interfaces.Symbol] {
	return maps.Values(s.symbols)
}

//=============================================================================
