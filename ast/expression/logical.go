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
	"fmt"

	"github.com/tradalia/sick-engine/core/types"
	"github.com/tradalia/sick-engine/core/values"
)

//=============================================================================
//===
//=== And
//===
//=============================================================================

type AndExpression struct {
	expressions []Expression
}

//=============================================================================

func NewAndExpression(expressions []Expression) *AndExpression {
	return &AndExpression{
		expressions: expressions,
	}
}

//=============================================================================

func (e *AndExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *AndExpression) Type() types.Type {
	return nil
}

//=============================================================================
//===
//=== Or
//===
//=============================================================================

type OrExpression struct {
	expressions []Expression
}

//=============================================================================

func NewOrExpression(expressions []Expression) *OrExpression {
	return &OrExpression{
		expressions: expressions,
	}
}

//=============================================================================

func (e *OrExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *OrExpression) Type() types.Type {
	return nil
}

//=============================================================================
//===
//=== Not
//===
//=============================================================================

type NotExpression struct {
	expression Expression
}

//=============================================================================

func NewNotExpression(e Expression) *NotExpression {
	return &NotExpression{
		expression: e,
	}
}

//=============================================================================

func (e *NotExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *NotExpression) Type() types.Type {
	return e.expression.Type()
}

//=============================================================================
//===
//=== Private methods
//===
//=============================================================================

func evalAsBool(e Expression) (bool,error) {
	v,err := e.Eval()
	if err != nil {
		return false,err
	}

	if v.Type() != types.NewBoolType() {
		return false,fmt.Errorf("expected a boolean value. Found %d", v.Data)
	}

	if v.Data == nil {
		return false,nil
	}

	return v.Data().(bool),nil
}

//=============================================================================
