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
	"github.com/tradalia/sick-engine/core/types"
	"github.com/tradalia/sick-engine/core/values"
)

//=============================================================================
//===
//=== Chained expression
//===
//=============================================================================

type ChainedExpression struct {
	HasThis bool
	Items   []*ChainItem
}

//=============================================================================

func NewChainedExpression(hasThis bool) *ChainedExpression {
	return &ChainedExpression{
		HasThis : hasThis,
	}
}

//=============================================================================

func (e *ChainedExpression) AddItem(item *ChainItem) {
	e.Items = append(e.Items, item)
}

//=============================================================================

func (e *ChainedExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *ChainedExpression) Type() types.Type {
	return nil
}

//=============================================================================
//===
//=== ChainItem
//===
//=============================================================================

const (
	CITypeIdentifier   = 0
	CITypeFunctionCall = 1
)

//=============================================================================

type ChainItem struct {
	Type     int
	Name     string
	Accessor Expression
	Params []Expression
}

//=============================================================================

func NewChainItem(type_ int, name string, accessor Expression, params []Expression) *ChainItem {
	return &ChainItem{
		Type    : type_,
		Name    : name,
		Accessor: accessor,
		Params  : params,
	}
}

//=============================================================================
