package extend

import (
	"fmt"

	"github.com/hacash/HVM/action"
)

// MintStakingCallbacks is implemented by the fullnode Mint layer (Rust FFI or in-process Go mint).
type MintStakingCallbacks interface {
	StakeHacd(sender []byte, diamonds []byte) error
	UnstakeHacd(sender []byte, diamonds []byte) error
}

// MintStakingHandler bridges HVM opcodes 0x01/0x02 to Mint staking_apply_* logic.
type MintStakingHandler struct {
	Mint MintStakingCallbacks
}

func (h *MintStakingHandler) Stake(sender []byte, diamonds []byte) error {
	if h.Mint == nil {
		return fmt.Errorf("mint staking callbacks not configured")
	}
	if err := action.ValidateDiamondWire(diamonds); err != nil {
		return err
	}
	return h.Mint.StakeHacd(sender, diamonds)
}

func (h *MintStakingHandler) Unstake(sender []byte, diamonds []byte) error {
	if h.Mint == nil {
		return fmt.Errorf("mint staking callbacks not configured")
	}
	if err := action.ValidateDiamondWire(diamonds); err != nil {
		return err
	}
	return h.Mint.UnstakeHacd(sender, diamonds)
}

// NewStakingExecutor returns an ExtendCallExecutor wired for HIP-25.
func NewStakingExecutor(mint MintStakingCallbacks) *StakingCallExecutor {
	return &StakingCallExecutor{
		Handler: &MintStakingHandler{Mint: mint},
	}
}