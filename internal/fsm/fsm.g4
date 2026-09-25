
grammar fsm;

fsm
    :   Note? FSM Identifier
        '{'
            (start | state | parallel | submachine | final | terminate | choice | junction | fork | join | point | event)*
        '}' EOF
    ;

state
    :   Note? INITIAL? STATE Identifier stereotype?
        '{'
            (start | state | parallel | submachine | final | terminate | choice | junction | fork | join | point | history | event)*
        '}'
    ;

parallel
    :   Note? INITIAL? PARALLEL STATE Identifier stereotype?
        '{'
            (region | point | event)*
        '}'
    ;

submachine
    :   Note? INITIAL? SUBMACHINE Identifier stereotype?
        '{'
            (point | reference | event)*
        '}'
    ;

final
    :   Note? FINAL STATE Identifier? stereotype?
    ;

terminate
    :   Note? TERMINATE STATE Identifier? stereotype?
    ;

region
    :   Note? REGION Identifier stereotype?
        '{'
            (start | state | parallel | submachine | final | terminate | choice | junction | fork | join | history)*
        '}'
    ;

choice
    :   Note? INITIAL? CHOICE Identifier stereotype?
        (   '{'
                branch*
            '}'
        |   branch
        )
    ;

junction
    :   Note? INITIAL? JUNCTION Identifier stereotype?
        (   '{'
                branch*
            '}'
        |   branch
        )
    ;

fork
    :   Note? FORK Identifier stereotype?
        '{'
            (Note? actions? goto)*
        '}'
    ;

join
    :   Note? JOIN Identifier stereotype? actions? goto
    ;

point
    :   Note? kind=(ENTRY | EXIT) POINT Identifier stereotype? actions? goto
    ;

reference
    :   Note? ENTRY POINT Identifier stereotype?
    ;

history
    :   Note? kind=(H | H_DEEP) stereotype? (actions? goto)?
    ;

start
    :   INITIAL actions? GOTO Identifier
    ;

stereotype
    :   '<<' Identifier '>>'
    ;

branch
    :   Note? ('[' ELSE ']' | guard)? actions? goto
    ;

event
    :   Note?
        (   name=(ENTRY | EXIT | DO) actions
        |   ON name=Identifier '/' DEFER
        |   INVARIANT guard
        |   trigger guard? actions goto?
        |   trigger guard? goto
        |   guard? actions? goto
        )
    ;

trigger
    :   ON name=Identifier
    |   AFTER '(' delay=(Duration | Identifier) ')'
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
    :   and_expression (OR and_expression)*
    ;

and_expression
    :   not_expression (AND not_expression)*
    ;

not_expression
    :   NOT? single_expression
    ;

single_expression
    :   Identifier
    |   '(' expression ')'
    ;

goto
    :   GOTO (FINAL | TERMINATE)
    |   GOTO LOCAL? Identifier
    |   GOTO LOCAL? (Identifier '.')? (H | H_DEEP)
    ;

FSM: 'fsm';
STATE: 'state';
PARALLEL: 'parallel';
SUBMACHINE: 'submachine';
REGION: 'region';
FINAL: 'final';
TERMINATE: 'terminate';
CHOICE: 'choice';
JUNCTION: 'junction';
FORK: 'fork';
JOIN: 'join';
ENTRY: 'entry';
EXIT: 'exit';
POINT: 'point';
DO: 'do';
ON: 'on';
AFTER: 'after';
ELSE: 'else';
GOTO: 'goto';
LOCAL: 'local';
INITIAL: 'initial';
DEFER: 'defer';
INVARIANT: 'invariant';
AND: 'and';
OR: 'or';
NOT: 'not';
H: 'H';
H_DEEP: 'H*';
Duration: [0-9]+ ('.' [0-9]+)? ('ms' | 's');
Identifier: [a-zA-Z_] [a-zA-Z0-9_]*;
Note: '|' ( '\\' . | ~[|\\] )* '|';
Comment: '#' ( ~[\r\n] )* -> skip;
Blank: [ \t\r\n]+ -> channel(HIDDEN);
