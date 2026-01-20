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

	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================

const (
	BarOpen  = 1
	BarHigh  = 2
	BarLow   = 3
	BarClose = 4
)

//=============================================================================
//===
//=== Bar Access
//===
//=============================================================================

type BarAccessExpression struct {
	Bar      int
	Accessor Expression
	info     *parser.Info
}

//=============================================================================

func NewBarAccessExpression(bar int, accessor Expression, info *parser.Info) *BarAccessExpression {
	return &BarAccessExpression{
		Bar:      bar,
		Accessor: accessor,
		info:     info,
	}
}

//=============================================================================

func (e *BarAccessExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	t, err := e.Accessor.ResolveType(scope, embedder, depth)
	if err != nil {
		return nil, err
	}

	if t.Code() != types.CodeInt {
		return nil, errors.New("accessor's type must be int: '" + t.String() + "'")
	}

	return types.NewRealType(), nil
}

//=============================================================================

func (e *BarAccessExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
