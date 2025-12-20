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

	"github.com/tradalia/sick-engine/core/interfaces"
	"github.com/tradalia/sick-engine/core/types"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================
//===
//=== Operands
//===
//=============================================================================

const (
	AritOpAdd  = "+"
	AritOpSub  = "-"
	AritOpMult = "*"
	AritOpDiv  = "/"
)

//=============================================================================
//===
//=== Arithmetic
//===
//=============================================================================

type ArithmeticExpression struct {
	operand  string
	left     Expression
	right    Expression
	info     *parser.Info
}

//=============================================================================

func NewArithmeticExpression(operand string, left, right Expression, info *parser.Info) *ArithmeticExpression {
	return &ArithmeticExpression{
		operand : operand,
		left    : left,
		right   : right,
		info    : info,
	}
}

//=============================================================================

func (e *ArithmeticExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	t1, err1 := e.left .ResolveType(scope, embedder, depth)
	t2, err2 := e.right.ResolveType(scope, embedder, depth)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	if supportsStringify(t1, t2) {
		return types.NewStringType(), nil
	}

	t := combineTypes(e.operand, t1, t2)
	if t == nil {
		return nil, errors.New("operand types mismatch: '" +t1.String() + "' and '" + t2.String() +"'")
	}

	return t,nil
}

//=============================================================================

func (e *ArithmeticExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func supportsStringify(t1, t2 types.Type) bool {
	if t1.Code() == types.CodeTimeseries || t2.Code() == types.CodeTimeseries {
		return false
	}

	return t1.Code() == types.CodeString || t2.Code() == types.CodeString
}

//=============================================================================

func combineTypes(operand string, t1, t2 types.Type) types.Type {
	if t1.Code() == types.CodeTimeseries || t2.Code() == types.CodeTimeseries {
		return nil
	}

	if t1.Code() == types.CodeInt && t2.Code() == types.CodeInt {
		return types.NewIntType()
	}

	if (t1.Code() == types.CodeInt || t1.Code() == types.CodeReal) && (t2.Code() == types.CodeInt || t2.Code() == types.CodeReal) {
		return types.NewRealType()
	}

	if operand == AritOpAdd || operand == AritOpSub {
		if (t1.Code() == types.CodeTime || t1.Code() == types.CodeInt) && (t2.Code() == types.CodeTime || t2.Code() == types.CodeInt) {
			return types.NewTimeType()
		}

		if (t1.Code() == types.CodeDate && t2.Code() == types.CodeInt) || (t1.Code() == types.CodeInt && t2.Code() == types.CodeDate) {
			return types.NewDateType()
		}
	}

	return nil
}

//=============================================================================
