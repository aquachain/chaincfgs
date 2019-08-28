package chaincfgs

import (
	"testing"

	"github.com/aquachain/hdwallet"
)

var (
	testPhrase = "tag volcano eight thank tide danger coast health above argue embrace heavy"
	testPW     = ""
)

func TestFakecoin(t *testing.T) {
	_, err := hdwallet.NewFromMnemonicCfg(&FAKEMainnet, testPhrase, testPW)
	if err != nil {
		t.Error(err)
	}
}

func TestBitcoin(t *testing.T) {
	t.Skip()
}

func TestLitecoin(t *testing.T) {
	t.Skip()
}
