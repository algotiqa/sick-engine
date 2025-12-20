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

package types

import (
	"github.com/tradalia/sick-engine/core/data"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================
//===
//=== Types
//===
//=============================================================================

func ConvertType(tree parser.ITypeContext) Type {
	if tree.ListType() != nil {
		lt      := tree.ListType()
		subType := ConvertType(lt.Type_())
		return NewListType(subType)
	}

	if tree.MapType() != nil {
		mt      := tree.MapType()
		keyType := ConvertKeyType(mt.KeyType())
		valType := ConvertType(mt.Type_())
		return NewMapType(keyType, valType)
	}

	if tree.INT() != nil {
		return NewIntType()
	}

	if tree.REAL() != nil {
		return NewRealType()
	}

	if tree.BOOL() != nil {
		return NewBoolType()
	}

	if tree.STRING() != nil {
		return NewStringType()
	}

	if tree.TIME() != nil {
		return NewTimeType()
	}

	if tree.DATE() != nil {
		return NewDateType()
	}

	if tree.TimeSeriesType() != nil {
		tst := tree.TimeSeriesType()
		var subType Type = NewRealType()
		if tst.Type_() != nil {
			subType = ConvertType(tst.Type_())
		}
		return NewTimeSeriesType(subType)
	}

	if tree.ERROR() != nil {
		return NewErrorType()
	}

	fqi := data.NewFQIdentifier(tree.FqIdentifier())

	return NewToFindOutType(fqi, parser.NewInfo(tree))
}

//=============================================================================

func ConvertKeyType(tree parser.IKeyTypeContext) Type {
	if tree.INT() != nil {
		return NewIntType()
	}

	if tree.STRING() != nil {
		return NewStringType()
	}

	if tree.TIME() != nil {
		return NewTimeType()
	}

	if tree.DATE() != nil {
		return NewDateType()
	}

	panic("Unknown key type : " + tree.GetText())
	return nil
}

//=============================================================================
