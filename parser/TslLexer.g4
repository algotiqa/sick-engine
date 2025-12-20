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

lexer grammar TslLexer;

//=== Keywords ================================================================

PACKAGE  : 'package';
CONST    : 'const';
VAR      : 'var';
FUNC     : 'func';

IF       : 'if';
ELSE     : 'else';
FOR      : 'for';
TRUE     : 'true';
FALSE    : 'false';
RETURN   : 'return';
AND      : 'and';
OR       : 'or';
NOT      : 'not';
NULL     : 'null';
THIS     : 'this';
NEW      : 'new';

//--- Bar

BAR         : 'bar';
OPEN        : 'open';
HIGH        : 'high';
LOW         : 'low';
CLOSE       : 'close';

//--- Orders

BUY         : 'buy';
SELL        : 'sell';
SELL_SHORT  : 'sellShort';
BUY_TO_COVER: 'buyToCover';
AT          : 'at';
MARKET      : 'market';
STOP        : 'stop';
LIMIT       : 'limit';
CONTRACTS   : 'contracts';

//--- Datatypes

INT        : 'int';
REAL       : 'real';
BOOL       : 'bool';
STRING     : 'string';
TIME       : 'time';
DATE       : 'date';
TIMESERIES : 'timeseries';
ENUM       : 'enum';
CLASS      : 'class';
ERROR      : 'error';
LIST       : 'list';
MAP        : 'map';

//=== Identifiers =============================================================

IDENTIFIER : [a-zA-Z]+ ( '_' | [a-zA-Z] | [0-9] )* ;

//=== Values =================================================================

INT_VALUE : DIGIT+;

REAL_VALUE
    : DIGIT+ '.' DIGIT*
    | '.' DIGIT+
    ;

DIGIT: [0-9];

STRING_VALUE : '"' (~["\\] | ESCAPED_VALUE)* '"' ;

fragment ESCAPED_VALUE: '\\' [bnrt\\"];

//=== Symbols =================================================================

L_PAREN    : '(' ;
R_PAREN    : ')' ;
L_BRACKET  : '[' ;
R_BRACKET  : ']' ;
L_CURLY    : '{' ;
R_CURLY    : '}' ;
EQUAL      : '=' ;

STAR       : '*' ;
SLASH      : '/' ;
PLUS       : '+' ;
MINUS      : '-' ;
COMMA      : ',' ;
DOT        : '.' ;
COLON      : ':' ;
APOS       : '\'';

//=== Relation operators

NOT_EQUAL        : '<>';
LESS_THAN        : '<';
LESS_OR_EQUAL    : '<=';
GREATER_THAN     : '>';
GREATER_OR_EQUAL : '>=';

//=== Comments and white spaces ===============================================

WS           : [ \t]+        -> skip;
NEWLINE      : [\r\n]+       -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

//=============================================================================
