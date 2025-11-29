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
	"strings"

	"github.com/tradalia/sick-engine/types"
	"github.com/tradalia/sick-engine/values"
)

//=============================================================================
//===
//=== Identifier
//===
//=============================================================================

type IdentifierExpression struct {
	name     *FQIdentifier
	accessor  Expression
}

//=============================================================================

func (e *IdentifierExpression) Eval() (values.Value,error) {
	return nil,nil
}

//=============================================================================

func (e *IdentifierExpression) Type() types.Type {
	return nil
}

//=============================================================================

func NewIdentifierExpression(name *FQIdentifier, accessor Expression) *IdentifierExpression {
	return &IdentifierExpression{
		name    : name,
		accessor: accessor,
	}
}

//=============================================================================
//===
//=== FQIdentifier
//===
//=============================================================================

type FQIdentifier struct {
	Scopes []string
}

//=============================================================================

func NewFQIdentifier() *FQIdentifier {
	return &FQIdentifier{}
}

//=============================================================================

func (i *FQIdentifier) AddScope(name string) {
	i.Scopes = append(i.Scopes, name)
}

//=============================================================================

func (i *FQIdentifier) String() string {
	var sb strings.Builder

	for index, name := range i.Scopes {
		if index != 0 {
			sb.WriteString(".")
		}
		sb.WriteString(name)
	}

	return sb.String()
}

//=============================================================================
