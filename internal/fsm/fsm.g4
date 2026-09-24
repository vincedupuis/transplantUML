
grammar fsm;

fsm
    :   'fsm' Identifier
        '{'
            (state | parallel | submachine | choice | junction | fork | join | point | event)*
        '}' EOF
    ;

state
    :   Initial? 'state' Identifier
        '{'
            (state | parallel | submachine | choice | junction | fork | join | point | event)*
        '}'
    ;

parallel
    :   Initial? 'parallel' 'state' Identifier
        '{'
            (region | point | event)*
        '}'
    ;

submachine
    :   Initial? 'submachine' Identifier
        '{'
            event*
        '}'
    ;

region
    :   'region' Identifier
        '{'
            (state | parallel | submachine | choice | junction | fork | join)*
        '}'
    ;

choice
    :   'choice' Identifier
        (   '{'
                branch*
            '}'
        |   branch
        )
    ;

junction
    :   'junction' Identifier
        (   '{'
                branch*
            '}'
        |   branch
        )
    ;

fork
    :   'fork' Identifier
        '{'
            (actions? goto)*
        '}'
    ;

join
    :   'join' Identifier actions? goto
    ;

point
    :   kind=('entry' | 'exit') 'point' Identifier actions? goto
    ;

branch
    :   ('[' 'else' ']' | guard)? actions? goto
    ;

event
    :   name=('entry' | 'exit' | 'do') actions
    |   'on' name=Identifier '/' Defer
    |   trigger guard? actions goto?
    |   trigger guard? goto
    |   guard? actions? goto
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
    :   'goto' ('final' | 'terminate')
    |   'goto' Local? Identifier
    |   'goto' Local? (Identifier '.')? ('H' | 'H*')
    ;

Initial: 'initial';
Defer: 'defer';
Local: 'local';
Duration: [0-9]+ ('.' [0-9]+)? ('ms' | 's');
Identifier: [a-zA-Z_] [a-zA-Z0-9_]*;
Comment: '#' ( ~[\r\n] )* -> skip;
Blank: [ \t\r\n]+ -> channel(HIDDEN);
