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

func testCoin(t *testing.T, cfg *Params, expectAddr string) {
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
	if addr.String() != expectAddr {
		t.Logf("expected:\n%s\ngot:\n%s\n", expectAddr, addr.String())
		t.FailNow()
	}
	fmt.Println("Address:", addr)

}

func TestFakecoin(t *testing.T) {
	testCoin(t, &FAKEMainnet, "FnPmeiMjLLYURLTnhKfRqfAZUy8HPL3HGt")
}
func TestBitcoin(t *testing.T) {
}

func TestLitecoin(t *testing.T) {
}
