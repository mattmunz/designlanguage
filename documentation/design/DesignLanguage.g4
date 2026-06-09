
grammar DesignLanguage;

design: preamble? (component NEWLINE)* EOF;

preamble: (author NEWLINE (COMMENT NEWLINE)? | (COMMENT NEWLINE)) NEWLINE;

author: AUTHOR_START AUTHOR_NAME;

component: simpleComponent | simpleComponent (field NEWLINE)+;

simpleComponent: NAME NEWLINE (COMMENT NEWLINE)?;

field: FIELD_START (attribute | method);

attribute:  param (' ' COMMENT)?;

method: NAME ' ' params ' ' (ARROW ' ' params)? (' ' COMMENT)?;

param: NAME ' ' type;

params: '()' | '(' param? (' ' param)* ')';

type: ARRAY? NAME; 

// --- Lexer Rules ---

ARRAY: '[]';

ARROW: '->';

ASTERISK: '*';

FIELD_START: ASTERISK ' ';

AUTHOR_NAME: [A-Z][a-zA-Z ]*;

AUTHOR_START: COMMENT_START 'Author: ';

COMMENT: COMMENT_START COMMENT_TEXT;

COMMENT_START: '-- ';

COMMENT_TEXT: [A-Z]([\p{L}\p{N}\p{Z}]|SPECIAL_CHAR)*;

NAME: [A-Z][a-zA-Z]*;

NEWLINE: '\r'? '\n';

SPECIAL_CHAR : [!@#$%^&*()_\-+={[}\]|:;"'<,>.?~`] ;

// // Whitespace to be ignored.
// WS: [ \t]+ -> skip;