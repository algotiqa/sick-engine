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

package core

import (
	"fmt"
	"strings"

	"github.com/tradalia/sick-engine/data"
)

//=============================================================================
//===
//=== Parse error
//===
//=============================================================================

type ParseErrors struct {
	Errors []*ParseError
}

//=============================================================================

func NewParseErrors() *ParseErrors {
	return &ParseErrors{
	}
}

//=============================================================================

func (pe *ParseErrors) AddError(p *ParseError) {
	pe.Errors = append(pe.Errors, p)
}

//=============================================================================

func (pe *ParseErrors) IsEmpty() bool {
	return len(pe.Errors) == 0
}

//=============================================================================

func (pe *ParseErrors) String() string {
	var sb strings.Builder

	for _, err := range pe.Errors {
		sb.WriteString(err.String() + "\n")
	}

	return sb.String()
}

//=============================================================================
//===
//=== Parse error
//===
//=============================================================================

type ParseError struct {
	File   string
	Line   int
	Column int
	Error  string
}

//=============================================================================

func NewParseError(file string, line int, column int, err string) *ParseError {
	return &ParseError{file, line, column, err}
}

//=============================================================================

func NewParseErrorFromInfo(info *data.Info, err string) *ParseError {
	return &ParseError{info.Filename, info.Line, info.Column, err}
}

//=============================================================================

func (pe *ParseError) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("File: %s, Line: %d, Col: %d, Error: %s", pe.File, pe.Line, pe.Column, pe.Error))
	return sb.String()
}

//=============================================================================
