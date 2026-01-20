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
	"io"

	"github.com/algotiqa/tiq-engine/ast"
	"github.com/algotiqa/tiq-engine/parser"
	"github.com/algotiqa/tiq-engine/runtime"
	"github.com/algotiqa/tiq-engine/tool"
	"github.com/antlr4-go/antlr/v4"
)

//=============================================================================
//===
//=== Environment creation
//===
//=============================================================================

func CreateEnvironment(path string) (*runtime.Environment, *parser.ParseErrors) {
	pr := parser.NewParseErrors()

	files, err := tool.FindFiles(path, ".tsl")
	if err != nil {
		pe := parser.NewParseError(path, -1, -1, err.Error())
		pr.AddError(pe)
		return nil, pr
	}

	e := runtime.NewEnvironment()

	for _, file := range files {
		script, errs := ParseFile(file)
		if !errs.IsEmpty() {
			return nil, errs
		}

		errs = e.AddScript(script)
		if !errs.IsEmpty() {
			return nil, errs
		}
	}

	return e, pr
}

//=============================================================================
//===
//=== Parsing functions
//===
//=============================================================================

func ParseFile(filename string) (*ast.Script, *parser.ParseErrors) {
	stream, err := antlr.NewFileStream(filename)
	if err != nil {
		pr := parser.NewParseErrors()
		pe := parser.NewParseError(filename, -1, -1, err.Error())
		pr.AddError(pe)
		return nil, pr
	}

	return parse(stream)
}

//=============================================================================

func ParseString(input string) (*ast.Script, *parser.ParseErrors) {
	return parse(antlr.NewInputStream(input))
}

//=============================================================================

func ParseReader(r io.Reader) (*ast.Script, *parser.ParseErrors) {
	return parse(antlr.NewIoStream(r))
}

//=============================================================================

func parse(input antlr.CharStream) (*ast.Script, *parser.ParseErrors) {
	lexer := parser.NewTslLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewTslParser(stream)

	res := parser.NewParseErrors()
	lis := parser.NewParseErrorListener(input.GetSourceName(), res)

	//--- Add error collection listener
	p.RemoveErrorListeners()
	p.AddErrorListener(lis)

	//--- Ask the parser to report all ambiguities
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
	tree := p.Script()
	if !res.IsEmpty() {
		return nil, res
	}

	script := ast.Build(tree)

	return script, res
}

//=============================================================================
