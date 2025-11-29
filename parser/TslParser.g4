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

parser grammar TslParser;

options {
    tokenVocab = TslLexer;
}

//=============================================================================

script
    : packageDef? ( constantsDef | variablesDef | functionDef | enumDef | classDef )*
    ;

//=============================================================================
//===
//=== Top entities
//===
//=============================================================================

packageDef
    : PACKAGE IDENTIFIER
    ;

//=============================================================================

constantsDef
    : CONST ( L_CURLY ( constantDef )* R_CURLY | constantDef )
    ;

constantDef
    : IDENTIFIER EQUAL simplifiedExpression
    ;

//=============================================================================

variablesDef
    : VAR ( L_CURLY ( variableDef )* R_CURLY | variableDef )
    ;

variableDef
    : IDENTIFIER EQUAL simplifiedExpression
    ;

//=============================================================================

functionDef
    : FUNC (class)? IDENTIFIER parameters results? block
    ;

class
    : L_PAREN IDENTIFIER R_PAREN
    ;

parameters
    : L_PAREN (parameterDecl (COMMA parameterDecl)* )? COMMA? R_PAREN
    ;

parameterDecl
    : IDENTIFIER type
    ;

results
    : type (COMMA type)*
    ;

//=============================================================================

enumDef
    : ENUM IDENTIFIER L_CURLY (enumItem)+ R_CURLY;

enumItem
    : IDENTIFIER ( L_PAREN enumValue R_PAREN )?
    ;

enumValue
    : INT_VALUE | STRING_VALUE;

//=============================================================================

classDef
    : CLASS IDENTIFIER L_CURLY (property)* R_CURLY;

property
    : IDENTIFIER type
    ;

//=============================================================================
//===
//=== Types
//===
//=============================================================================

type
    : INT
    | REAL
    | BOOL
    | STRING
    | TIME
    | DATE
    | TIMESERIES
    | fqIdentifier
    | ERROR
    | listType
    | mapType
    ;

listType
    : LIST OF L_PAREN type R_PAREN
    ;

mapType
    : MAP OF L_PAREN keyType COMMA type R_PAREN
    ;

keyType
    : INT
    | STRING
    | TIME
    | DATE
    ;

//=============================================================================
//===
//=== Statements
//===
//=============================================================================

block
    : L_CURLY (statement)* R_CURLY;

//=============================================================================

statement
    : varDeclaration
    | ifStatement
    | functionCall
    | returnStatement
    ;

//=============================================================================

varDeclaration
    : fqIdentifier ( COMMA fqIdentifier)* EQUAL expression ( COMMA expression)*
    ;

//=============================================================================

ifStatement
    : IF expression block (elseIfBlock)* (elseBlock)?
    ;

elseIfBlock : ELSE IF expression block;

elseBlock : ELSE block;

//=============================================================================

functionCall
    : fqIdentifier paramsExpression
    ;

//=============================================================================

returnStatement
    : RETURN
    | RETURN expression ( COMMA expression)*
    ;

//=============================================================================
//===
//=== Expressions
//===
//=============================================================================

expression
    : unaryExpression
    | expression (STAR | SLASH) expression
    | expression (PLUS | MINUS) expression
    | expression ( EQUAL | NOT_EQUAL | LESS_THAN | LESS_OR_EQUAL | GREATER_THAN | GREATER_OR_EQUAL ) expression
    | NOT expression
    | expression ( AND expression )+
    | expression ( OR  expression )+
    | listExpression
    | mapExpression
    ;

//=============================================================================

unaryExpression
    : (MINUS|PLUS) unaryExpression
    | identifierExpression
    | functionCallExpression
    | expressionInParenthesis
    | constantValueExpression
    ;

//=============================================================================

identifierExpression
    : fqIdentifier ( accessorExpression )?
    ;

fqIdentifier
    : IDENTIFIER ( DOT IDENTIFIER )*
    ;

accessorExpression
    : L_BRACKET expression R_BRACKET
    ;

//=============================================================================

functionCallExpression
    : fqIdentifier paramsExpression
    ;

paramsExpression
    : L_PAREN (expression (COMMA expression)* )? R_PAREN
    ;

//=============================================================================

expressionInParenthesis
    : L_PAREN expression R_PAREN;

//=============================================================================

listExpression
    : listType initialListValues?
    ;

initialListValues
    : L_CURLY expression ( COMMA expression)* R_CURLY
    ;

//=============================================================================

mapExpression
    : mapType initialMapValues?
    ;

initialMapValues
    : L_CURLY keyValueCouple ( COMMA keyValueCouple)* R_CURLY
    ;

keyValueCouple
    : keyValue COLON expression
    ;

keyValue
    : INT_VALUE
    | STRING_VALUE
    | timeValue
    | dateValue
    ;

//=============================================================================

simplifiedExpression
    : (MINUS|PLUS) simplifiedExpression
    | identifierExpression
    | constantValueExpression
    | simplifiedExpression (STAR | SLASH) simplifiedExpression
    | simplifiedExpression (PLUS | MINUS) simplifiedExpression
    ;

//=============================================================================

constantValueExpression
    : INT_VALUE
    | REAL_VALUE
    | boolValue
    | STRING_VALUE
    | timeValue
    | dateValue
    | errorValue
    ;

timeValue
    : APOS INT_VALUE COLON INT_VALUE APOS;

dateValue
    : APOS INT_VALUE MINUS INT_VALUE MINUS INT_VALUE APOS;

boolValue
    : TRUE | FALSE ;

errorValue
    : ERROR L_PAREN STRING_VALUE R_PAREN;

//=============================================================================
