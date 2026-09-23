
grammar fsm;

fsm
    :   'fsm' Identifier
        '{'
            (state | event)*
        '}' EOF
    ;

state
    :   Initial? 'state' Identifier
        '{'
            (state | event)*
        '}'
    ;

event
    :   'on' name=('entry' | 'exit') actions?
    |   'on' name=Identifier guard? actions goto?
    |   'on' name=Identifier guard? goto
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
    :   'goto' ('.' | 'final' | 'H' | Identifier)
    ;

Initial: 'initial';
Identifier: [a-zA-Z_] [a-zA-Z0-9_.\-]*;
Comment: '#' ( ~[\r\n] )* -> skip;
Blank: [ \t\r\n]+ -> channel(HIDDEN);
