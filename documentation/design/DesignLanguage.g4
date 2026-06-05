
grammar DesignLanguage;

design: (component NEWLINE?)* EOF;

component: NAME NEWLINE field*;

field: attribute | method;

attribute: ASTERISK ' ' param NEWLINE;

method: ASTERISK ' ' NAME ' ' params ' ' ARROW ' ' params NEWLINE;

param: NAME ' ' type;

params: '()' | '(' param? (' ' param)* ')';

type: ARRAY? NAME; 

// --- Lexer Rules ---

ARRAY: '[]';

NAME: [A-Z][a-zA-Z]*;

ASTERISK: '*';

ARROW: '->';

NEWLINE: '\r'? '\n';

// // Whitespace to be ignored.
// WS: [ \t]+ -> skip;