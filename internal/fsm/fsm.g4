
grammar fsm;

fsm
    :   Note? 'fsm' Identifier
        '{'
            (state | parallel | submachine | final | terminate | choice | junction | fork | join | point | event)*
        '}' EOF
    ;

state
    :   Note? Initial? 'state' Identifier stereotype?
        '{'
            (state | parallel | submachine | final | terminate | choice | junction | fork | join | point | history | event)*
        '}'
    ;

parallel
    :   Note? Initial? 'parallel' 'state' Identifier stereotype?
        '{'
            (region | point | event)*
        '}'
    ;

submachine
    :   Note? Initial? 'submachine' Identifier stereotype?
        '{'
            event*
        '}'
    ;

final
    :   Note? 'final' 'state' Identifier? stereotype?
    ;

terminate
    :   Note? 'terminate' 'state' Identifier? stereotype?
    ;

region
    :   Note? 'region' Identifier stereotype?
        '{'
            (state | parallel | submachine | final | terminate | choice | junction | fork | join | history)*
        '}'
    ;

choice
    :   Note? 'choice' Identifier stereotype?
        (   '{'
                branch*
            '}'
        |   branch
        )
    ;

junction
    :   Note? 'junction' Identifier stereotype?
        (   '{'
                branch*
            '}'
        |   branch
        )
    ;

fork
    :   Note? 'fork' Identifier stereotype?
        '{'
            (Note? actions? goto)*
        '}'
    ;

join
    :   Note? 'join' Identifier stereotype? actions? goto
    ;

point
    :   Note? kind=('entry' | 'exit') 'point' Identifier stereotype? actions? goto
    ;

history
    :   Note? kind=('H' | 'H*') stereotype? (actions? goto)?
    ;

stereotype
    :   '<<' Identifier '>>'
    ;

branch
    :   Note? ('[' 'else' ']' | guard)? actions? goto
    ;

event
    :   Note?
        (   name=('entry' | 'exit' | 'do') actions
        |   'on' name=Identifier '/' Defer
        |   Invariant guard
        |   trigger guard? actions goto?
        |   trigger guard? goto
        |   guard? actions? goto
        )
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
Invariant: 'invariant';
Defer: 'defer';
Local: 'local';
Duration: [0-9]+ ('.' [0-9]+)? ('ms' | 's');
Identifier: [a-zA-Z_] [a-zA-Z0-9_]*;
Note: '|' ( '\\' . | ~[|\\] )* '|';
Comment: '#' ( ~[\r\n] )* -> skip;
Blank: [ \t\r\n]+ -> channel(HIDDEN);
