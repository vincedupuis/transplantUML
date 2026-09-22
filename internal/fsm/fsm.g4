
grammar fsm;

fsm
    :   'fsm' Identifier
        '{'
            (state | event)*
        '}' EOF
    ;

state
    :   'state' Identifier
        '{'
            (state | event)*
        '}'
    ;

event
    :   'on' name=('entry' | 'exit') actions?
    |   'on' name=Identifier guard? actions? goto?
    ;

actions
    :   '/' identifiers
    ;

identifiers
    :   Identifier (',' Identifier)*
    ;

guard
    :  '[' expression ']'
    ;

expression
    :   or_expression
    ;

or_expression
    :   and_expression ('or' and_expression)*
    ;

and_expression
    :   not_expression ('and' not_expression)*
    ;

not_expression
    :   'not'? single_expression
    ;

single_expression
    :   Identifier
    |   '(' expression ')'
    ;

goto
    :   'goto' ('.' | 'final' | 'H' | path)
    ;

path
    :   (Prefix | '/')? Identifier ('/' Identifier)*
    ;

// ----------------------------------------------------------------------------

Identifier: [a-zA-Z] [a-zA-Z0-9]*;
Prefix: './' | ('../')+;
Preprocessor: '#' ( ~[\r\n] | '\\' [\r\n] | '\\' '\n' )* -> skip;
Blank: [ \t\r\n]+ -> channel(HIDDEN);
