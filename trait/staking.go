package trait

// StakingFeeSharePercent is the HIP-25 fee redirect share (inscription + transfer fees).
const StakingFeeSharePercent = 22

/**
 * HIP-25 staking hooks invoked from HVM external actions 0x01 / 0x02.
 * Fullnode wiring provides a concrete implementation that updates Mint state.
 */
type StakingHandler interface {
	Stake(sender []byte, diamonds []byte) error
	Unstake(sender []byte, diamonds []byte) error
}

/**
 * Optional transaction context used by staking actions to read the fee payer / signer.
 */
type TransactionContext interface {
	Context
	Sender() []byte
}