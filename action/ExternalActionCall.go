package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

/********** Static **********/

type StaticExternalActionCall struct {
	ExternalAction trait.VMAction
}

func (elm StaticExternalActionCall) VMKind() uint8 {
	return 0
}

func (elm StaticExternalActionCall) IsBurning90PersentTxFees() bool {
	return false
}

func (s *StaticExternalActionCall) Parse(extca trait.ExtendCallExecutor, buf []byte, seek uint32) (uint32, error) {
	panic(any("Never call StaticExternalActionCall.Parse"))
}

func (elm StaticExternalActionCall) ChildActions() []trait.VMAction {
	return []trait.VMAction{elm.ExternalAction}
}

func (elm StaticExternalActionCall) Childs() []trait.ASTNode {
	return []trait.ASTNode{elm.ExternalAction}
}

func (elm StaticExternalActionCall) Evaluate(ctx trait.Context) trait.EvalResult {
	// check
	var exec = ctx.GetExtendCallExecutor()
	var evmk = elm.ExternalAction.VMKind()
	if evmk > exec.ExtendKindRange() {
		return eval.ResultFatalErr(fmt.Errorf("ExternalAction VMKind %d error", evmk))
	}
	// evaluate
	var resv, _, err = exec.Evaluate(elm.ExternalAction, ctx)
	if err != nil {
		return eval.ResultFatalErr(err) // error
	}
	// ok
	return eval.ResultValue(resv)
}

/********** Dynamic **********/

type DynamicExternalActionCall struct {
	ActionOverallData trait.VMAction
}

func (elm DynamicExternalActionCall) VMKind() uint8 {
	return 254
}

func (elm DynamicExternalActionCall) IsBurning90PersentTxFees() bool {
	return true
}

func (elm DynamicExternalActionCall) ChildActions() []trait.VMAction {
	return []trait.VMAction{elm.ActionOverallData}
}

/* func (elm DynamicExternalActionCall) Childs() []trait.ASTNode {
	return []trait.ASTNode{elm.ActionOverallData}
} */

func (s *DynamicExternalActionCall) Parse(extca trait.ExtendCallExecutor, buf []byte, seek uint32) (uint32, error) {
	return 0, nil
}

func (elm DynamicExternalActionCall) Evaluate(ctx trait.Context) trait.EvalResult {
	var exec = ctx.GetExtendCallExecutor()
	var actdts = elm.ActionOverallData.Evaluate(ctx)
	if actdts.CheckInterrupt() {
		return actdts
	}
	var data = actdts.RetValue()
	if len(data) < 2 {
		return eval.ResultFatalErr(fmt.Errorf("Action data length %d less than kind length %d",
			len(data), 2))
	}
	// build action
	var actobj, _, e = exec.Parse(data, 0)
	if e != nil {
		return eval.ResultFatalErr(fmt.Errorf("Action build error: %s ", e))
	}
	// evaluate
	var resv, _, err = ctx.GetExtendCallExecutor().Evaluate(actobj, ctx)
	if err != nil {
		return eval.ResultFatalErr(err) // error
	}
	// ok
	return eval.ResultValue(resv)

}
