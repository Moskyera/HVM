package extend

import (
	"fmt"

	"github.com/hacash/HVM/action"
	"github.com/hacash/HVM/trait"
)

// StakingExtendKindMax is the highest HIP-25 external opcode (UNSTAKE_HACD).
const StakingExtendKindMax uint8 = 0x02

/**
 * StakingCallExecutor registers HIP-25 opcodes 0x01 (stake) and 0x02 (unstake).
 * Embed or wrap a StakingHandler from the fullnode Mint integration layer.
 */
type StakingCallExecutor struct {
	Handler trait.StakingHandler
}

func (e *StakingCallExecutor) ExtendKindRange() uint8 {
	return StakingExtendKindMax
}

func (e *StakingCallExecutor) Parse(buf []byte, seek uint32) (trait.VMAction, uint32, error) {
	if seek >= uint32(len(buf)) {
		return nil, 0, fmt.Errorf("empty staking action buffer")
	}
	switch buf[seek] {
	case 0x01:
		act := &action.StakeHacd{}
		sk, err := act.Parse(e, buf, seek)
		return act, sk, err
	case 0x02:
		act := &action.UnstakeHacd{}
		sk, err := act.Parse(e, buf, seek)
		return act, sk, err
	default:
		return nil, 0, fmt.Errorf("unknown staking opcode %d", buf[seek])
	}
}

func (e *StakingCallExecutor) Evaluate(act trait.VMAction, ctx trait.Context) ([]byte, int32, error) {
	res := act.Evaluate(ctx)
	if res.FatalErr() != nil {
		return nil, 0, res.FatalErr()
	}
	return res.RetValue(), 0, nil
}

func (e *StakingCallExecutor) Stake(sender []byte, diamonds []byte) error {
	if e.Handler == nil {
		return fmt.Errorf("staking handler not configured")
	}
	return e.Handler.Stake(sender, diamonds)
}

func (e *StakingCallExecutor) Unstake(sender []byte, diamonds []byte) error {
	if e.Handler == nil {
		return fmt.Errorf("staking handler not configured")
	}
	return e.Handler.Unstake(sender, diamonds)
}