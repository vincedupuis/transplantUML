
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
    |   trigger guard? actions goto?
    |   trigger guard? goto
    ;

trigger
    :   'on' name=Identifier
    |   'after' '(' delay=(Duration | Identifier) ')'
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
Duration: [0-9]+ ('.' [0-9]+)? ('ms' | 's');
Identifier: [a-zA-Z_] [a-zA-Z0-9_.\-]*;
Comment: '#' ( ~[\r\n] )* -> skip;
Blank: [ \t\r\n]+ -> channel(HIDDEN);
