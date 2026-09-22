// Code generated from fsm.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // fsm
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by fsmParser.
type fsmVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by fsmParser#fsm.
	VisitFsm(ctx *FsmContext) interface{}

	// Visit a parse tree produced by fsmParser#state.
	VisitState(ctx *StateContext) interface{}

	// Visit a parse tree produced by fsmParser#event.
	VisitEvent(ctx *EventContext) interface{}

	// Visit a parse tree produced by fsmParser#actions.
	VisitActions(ctx *ActionsContext) interface{}

	// Visit a parse tree produced by fsmParser#identifiers.
	VisitIdentifiers(ctx *IdentifiersContext) interface{}

	// Visit a parse tree produced by fsmParser#guard.
	VisitGuard(ctx *GuardContext) interface{}

	// Visit a parse tree produced by fsmParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by fsmParser#or_expression.
	VisitOr_expression(ctx *Or_expressionContext) interface{}

	// Visit a parse tree produced by fsmParser#and_expression.
	VisitAnd_expression(ctx *And_expressionContext) interface{}

	// Visit a parse tree produced by fsmParser#not_expression.
	VisitNot_expression(ctx *Not_expressionContext) interface{}

	// Visit a parse tree produced by fsmParser#single_expression.
	VisitSingle_expression(ctx *Single_expressionContext) interface{}

	// Visit a parse tree produced by fsmParser#goto.
	VisitGoto(ctx *GotoContext) interface{}

	// Visit a parse tree produced by fsmParser#path.
	VisitPath(ctx *PathContext) interface{}
}
