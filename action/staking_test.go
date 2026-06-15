package action

import (
	"testing"
)

func TestParseDiamondNameList(t *testing.T) {
	buf := []byte{0x02, 'W', 'T', 'Y', 'U', 'I', 'A', 'H', 'X', 'V', 'M', 'E'}
	names, next, err := parseDiamondNameList(buf, 0)
	if err != nil {
		t.Fatal(err)
	}
	if next != 13 {
		t.Fatalf("expected seek 13, got %d", next)
	}
	if len(names) != 12 {
		t.Fatalf("expected 12 name bytes, got %d", len(names))
	}
}

func TestParseDiamondNameListRejectsZeroCount(t *testing.T) {
	_, _, err := parseDiamondNameList([]byte{0x00}, 0)
	if err == nil {
		t.Fatal("expected error for zero count")
	}
}

func TestStakeHacdParse(t *testing.T) {
	buf := []byte{0x01, 0x01, 'W', 'T', 'Y', 'U', 'I', 'A'}
	act := &StakeHacd{}
	sk, err := act.Parse(nil, buf, 0)
	if err != nil {
		t.Fatal(err)
	}
	if sk != 8 {
		t.Fatalf("expected seek 8, got %d", sk)
	}
	if len(act.Diamonds) != 6 {
		t.Fatalf("expected 6 diamond bytes, got %d", len(act.Diamonds))
	}
}