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

package values

import "github.com/tradalia/sick-engine/types"

//=============================================================================
//===
//=== Map value
//===
//=============================================================================

type MapValue struct {
	values    map[Value]Value
	keyType   types.Type
	valueType types.Type
}

//=============================================================================

func NewMapValue(keyType, valueType types.Type) *MapValue {
	return &MapValue{
		keyType  : keyType,
		valueType: valueType,
		values   : map[Value]Value{},
	}
}

//=============================================================================

func (v *MapValue) Set(key Value, value Value) {
	v.values[key] = value
}

//=============================================================================

func (v *MapValue) Data() any {
	return v.values
}

//=============================================================================

func (v *MapValue) Type() types.Type {
	return v.keyType
}

//=============================================================================

func (v *MapValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *MapValue) LessThan(other Value) bool {
	return false
}

//=============================================================================
