package action

import (
	"bytes"
	"fmt"
)

const (
	DiamondNameWidth  = 6
	MaxDiamondBatch   = 200
	diamondValidChars = "WTYUIAHXVMEKBSZN"
)

func diamondCharValid(b byte) bool {
	return bytes.ContainsByte([]byte(diamondValidChars), b)
}

// parseDiamondNameList reads Uint1 count + concatenated 6-byte literals (DiamondNameListMax200 wire format).
func parseDiamondNameList(buf []byte, seek uint32) ([]byte, uint32, error) {
	if seek >= uint32(len(buf)) {
		return nil, 0, fmt.Errorf("diamond list buffer underflow")
	}
	count := int(buf[seek])
	seek++
	if count == 0 || count > MaxDiamondBatch {
		return nil, 0, fmt.Errorf("diamond count %d invalid", count)
	}
	need := uint32(count * DiamondNameWidth)
	if seek+need > uint32(len(buf)) {
		return nil, 0, fmt.Errorf("diamond list truncated")
	}
	names := make([]byte, need)
	copy(names, buf[seek:seek+need])
	for i := 0; i < count; i++ {
		for j := 0; j < DiamondNameWidth; j++ {
			if !diamondCharValid(names[i*DiamondNameWidth+j]) {
				return nil, 0, fmt.Errorf("invalid diamond name at index %d", i)
			}
		}
	}
	return names, seek + need, nil
}