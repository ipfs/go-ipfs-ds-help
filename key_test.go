package dshelp

import (
	"testing"

	cid "github.com/ipfs/go-cid"
)

func TestKey(t *testing.T) {
	c, _ := cid.Decode("QmP63DkAFEnDYNjDYBpyNDfttu1fvUw99x1brscPzpqmmq")
	k := c.Hash()
	dsKey1 := CidToDsKey(c)
	dsKey2 := MultihashToDsKey(k)
	if dsKey1.String() != dsKey2.String() {
		t.Fatal("CidToDsKey and MultihashToDsKey should produce same result")
	}
	k2, err := DsKeyToMultihash(dsKey2)
	if err != nil {
		t.Fatal(err)
	}
	if string(k) != string(k2) {
		t.Fatal("should have parsed the same key")
	}
}
