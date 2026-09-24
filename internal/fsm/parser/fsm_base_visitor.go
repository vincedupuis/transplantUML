// Code generated from fsm.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // fsm
import "github.com/antlr4-go/antlr/v4"

type BasefsmVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasefsmVisitor) VisitFsm(ctx *FsmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitState(ctx *StateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitParallel(ctx *ParallelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitSubmachine(ctx *SubmachineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitFinal(ctx *FinalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitTerminate(ctx *TerminateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitRegion(ctx *RegionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitChoice(ctx *ChoiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitJunction(ctx *JunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitFork(ctx *ForkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitJoin(ctx *JoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitPoint(ctx *PointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitHistory(ctx *HistoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitStereotype(ctx *StereotypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitBranch(ctx *BranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitEvent(ctx *EventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitTrigger(ctx *TriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitActions(ctx *ActionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitIdentifiers(ctx *IdentifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitGuard(ctx *GuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitOr_expression(ctx *Or_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitAnd_expression(ctx *And_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitNot_expression(ctx *Not_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitSingle_expression(ctx *Single_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefsmVisitor) VisitGoto(ctx *GotoContext) interface{} {
	return v.VisitChildren(ctx)
}
