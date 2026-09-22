// Package fsm is tpuml's own state machine language. The grammar lives in
// fsm.g4; the lexer, parser and visitor in the parser subpackage are generated
// from it by ANTLR (`make generate`) and committed, so building tpuml needs no
// Java.
package fsm

//go:generate java -jar $ANTLR_JAR -Dlanguage=Go -visitor -no-listener -o parser -package parser -Xexact-output-dir fsm.g4
