
grammar fsm;

// States and events may be interleaved, so a state's transitions can be
// written next to the children they concern.
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

// Two alternatives rather than one, so that a transition carrying neither
// effect nor target is a syntax error rather than something a later stage
// has to catch.
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
    :   'goto' ('.' | 'final' | 'H' | path)
    ;

// The leading '/' of an absolute path is the same token as the one that opens
// an action list, so Prefix covers only the relative forms.
path
    :   (Prefix | '/')? Identifier ('/' Identifier)*
    ;

// ----------------------------------------------------------------------------

Identifier: [a-zA-Z] [a-zA-Z0-9]*;
Prefix: './' | ('../')+;
Preprocessor: '#' ( ~[\r\n] | '\\' [\r\n] | '\\' '\n' )* -> skip;
Blank: [ \t\r\n]+ -> channel(HIDDEN);
