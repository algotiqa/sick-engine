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
//=== Mult
//===
//=============================================================================

type MultExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *MultExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *MultExpression) Type() types.Type {
	return nil
}

//=============================================================================

func NewMultExpression(left, right Expression) *MultExpression {
	return &MultExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Div
//===
//=============================================================================

type DivExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func NewDivExpression(left, right Expression) *DivExpression {
	return &DivExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================

func (e *DivExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *DivExpression) Type() types.Type {
	return nil
}

//=============================================================================
//===
//=== Add
//===
//=============================================================================

type AddExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *AddExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *AddExpression) Type() types.Type {
	return nil
}

//=============================================================================

func NewAddExpression(left, right Expression) *AddExpression {
	return &AddExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Sub
//===
//=============================================================================

type SubExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *SubExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *SubExpression) Type() types.Type {
	return nil
}

//=============================================================================

func NewSubExpression(left, right Expression) *SubExpression {
	return &SubExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
