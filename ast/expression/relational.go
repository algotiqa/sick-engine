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
	"github.com/tradalia/sick-engine/types"
	"github.com/tradalia/sick-engine/values"
)

//=============================================================================
//===
//=== Equal
//===
//=============================================================================

type EqualExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func NewEqualExpression(left, right Expression) *EqualExpression {
	return &EqualExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================

func (e *EqualExpression) Eval() (values.Value,error) {
	left, err1 := e.left .Eval()
	right,err2 := e.right.Eval()

	if err1 != nil {
		return nil,err1
	}

	if err2 != nil {
		return nil,err2
	}

	_ = left == right
	return nil,nil
}

//=============================================================================

func (e *EqualExpression) Type() types.Type {
	return nil
}

//=============================================================================
//===
//=== Not equal
//===
//=============================================================================

type NotEqualExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func NewNotEqualExpression(left, right Expression) *NotEqualExpression {
	return &NotEqualExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================

func (e *NotEqualExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *NotEqualExpression) Type() types.Type {
	return nil
}

//=============================================================================
//===
//=== Less or equal
//===
//=============================================================================

type LessOrEqualExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *LessOrEqualExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *LessOrEqualExpression) Type() types.Type {
	return nil
}

//=============================================================================

func NewLessOrEqualExpression(left, right Expression) *LessOrEqualExpression {
	return &LessOrEqualExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Greater or equal
//===
//=============================================================================

type GreaterOrEqualExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *GreaterOrEqualExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *GreaterOrEqualExpression) Type() types.Type {
	return nil
}

//=============================================================================

func NewGreaterOrEqualExpression(left, right Expression) *GreaterOrEqualExpression {
	return &GreaterOrEqualExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Less than
//===
//=============================================================================

type LessThanExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *LessThanExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *LessThanExpression) Type() types.Type {
	return nil
}

//=============================================================================

func NewLessThanExpression(left, right Expression) *LessThanExpression {
	return &LessThanExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Greater than
//===
//=============================================================================

type GreaterThanExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *GreaterThanExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *GreaterThanExpression) Type() types.Type {
	return nil
}

//=============================================================================

func NewGreaterThanExpression(left, right Expression) *GreaterThanExpression {
	return &GreaterThanExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
