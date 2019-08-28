package chaincfgs

import (
	"fmt"
	"testing"

	"github.com/aquachain/hdwallet"
	"github.com/btcsuite/btcutil/hdkeychain"
)

var (
	testPhrase = "tag volcano eight thank tide danger coast health above argue embrace heavy"
	testPW     = ""
)

func TestFakecoin(t *testing.T) {
	cfg := &FAKEMainnet
	seed, err := hdwallet.NewSeedFromMnemonic(testPhrase, testPW)
	if err != nil {
		t.Error(err)
	}
	key, err := hdkeychain.NewMaster(seed, cfg)
	if err != nil {
		t.Error(err)
	}
	addr, err := key.Address(cfg)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "FnPmeiMjLLYURLTnhKfRqfAZUy8HPL3HGt" {
		t.Logf("expected:\n%s\ngot:\n%s\n", "FnPmeiMjLLYURLTnhKfRqfAZUy8HPL3HGt", addr.String())
		t.FailNow()
	}
	fmt.Println("Address:", addr)

}

func TestBitcoin(t *testing.T) {
	t.Skip()
}

func TestLitecoin(t *testing.T) {
	t.Skip()
}
