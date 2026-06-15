package action

// ValidateDiamondWire checks HIP-25 / DiamondNameListMax200 payload without constructing an action.
func ValidateDiamondWire(raw []byte) error {
	_, _, err := parseDiamondNameList(raw, 0)
	return err
}