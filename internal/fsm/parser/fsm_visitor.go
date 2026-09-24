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

	// Visit a parse tree produced by fsmParser#parallel.
	VisitParallel(ctx *ParallelContext) interface{}

	// Visit a parse tree produced by fsmParser#submachine.
	VisitSubmachine(ctx *SubmachineContext) interface{}

	// Visit a parse tree produced by fsmParser#final.
	VisitFinal(ctx *FinalContext) interface{}

	// Visit a parse tree produced by fsmParser#terminate.
	VisitTerminate(ctx *TerminateContext) interface{}

	// Visit a parse tree produced by fsmParser#region.
	VisitRegion(ctx *RegionContext) interface{}

	// Visit a parse tree produced by fsmParser#choice.
	VisitChoice(ctx *ChoiceContext) interface{}

	// Visit a parse tree produced by fsmParser#junction.
	VisitJunction(ctx *JunctionContext) interface{}

	// Visit a parse tree produced by fsmParser#fork.
	VisitFork(ctx *ForkContext) interface{}

	// Visit a parse tree produced by fsmParser#join.
	VisitJoin(ctx *JoinContext) interface{}

	// Visit a parse tree produced by fsmParser#point.
	VisitPoint(ctx *PointContext) interface{}

	// Visit a parse tree produced by fsmParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by fsmParser#history.
	VisitHistory(ctx *HistoryContext) interface{}

	// Visit a parse tree produced by fsmParser#stereotype.
	VisitStereotype(ctx *StereotypeContext) interface{}

	// Visit a parse tree produced by fsmParser#branch.
	VisitBranch(ctx *BranchContext) interface{}

	// Visit a parse tree produced by fsmParser#event.
	VisitEvent(ctx *EventContext) interface{}

	// Visit a parse tree produced by fsmParser#trigger.
	VisitTrigger(ctx *TriggerContext) interface{}

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
}
