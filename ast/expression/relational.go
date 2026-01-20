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
//===
//=== Operands
//===
//=============================================================================

const (
	RelOpEqual          = "="
	RelOpNotEqual       = "<>"
	RelOpLessThan       = "<"
	RelOpGreaterThan    = ">"
	RelOpLessOrEqual    = "<="
	RelOpGreaterOrEqual = ">="
)

//=============================================================================
//===
//=== Relational
//===
//=============================================================================

type RelationalExpression struct {
	operand string
	left    Expression
	right   Expression
	info    *parser.Info
}

//=============================================================================

func NewRelationalExpression(operand string, left, right Expression, info *parser.Info) *RelationalExpression {
	return &RelationalExpression{
		operand: operand,
		left:    left,
		right:   right,
		info:    info,
	}
}

//=============================================================================

func (e *RelationalExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	t1, err1 := e.left.ResolveType(scope, embedder, depth)
	t2, err2 := e.right.ResolveType(scope, embedder, depth)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	ok := false

	switch e.operand {
	case RelOpEqual:
		ok = canEquateTo(t1, t2)
	case RelOpNotEqual:
		ok = canEquateTo(t1, t2)
	case RelOpLessThan:
		ok = canCompareTo(t1, t2)
	case RelOpLessOrEqual:
		ok = canCompareTo(t1, t2)
	case RelOpGreaterThan:
		ok = canCompareTo(t1, t2)
	case RelOpGreaterOrEqual:
		ok = canCompareTo(t1, t2)
	}

	if !ok {
		return nil, errors.New("operand '" + e.operand + "'s not usable with '" + t1.String() + "' and '" + t2.String() + "'")
	}

	return types.NewBoolType(), nil
}

//=============================================================================

func (e *RelationalExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func canEquateTo(t1, t2 types.Type) bool {
	if t1.Code() == types.CodeTimeseries || t2.Code() == types.CodeTimeseries {
		return false
	}

	if t1.Code() == t2.Code() {
		return true
	}

	if t1.Code() == types.CodeInt && t2.Code() == types.CodeReal {
		return true
	}

	if t1.Code() == types.CodeReal && t2.Code() == types.CodeInt {
		return true
	}

	return false
}

//=============================================================================

func canCompareTo(t1, t2 types.Type) bool {
	if t1.Code() == types.CodeTimeseries || t2.Code() == types.CodeTimeseries {
		return false
	}

	if t1.Code() == types.CodeDate && t2.Code() == types.CodeDate {
		return true
	}

	if t1.Code() == types.CodeTime && t2.Code() == types.CodeTime {
		return true
	}

	if t1.Code() == types.CodeString && t2.Code() == types.CodeString {
		return true
	}

	if (t1.Code() == types.CodeInt || t1.Code() == types.CodeReal) && (t2.Code() == types.CodeInt || t2.Code() == types.CodeReal) {
		return true
	}

	return false
}

//=============================================================================
