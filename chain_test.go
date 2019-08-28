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

	gold = map[string]string{
		"DOGED": "DNN9k4a15FmDWjdGYofWvuXqhwad68TvAm",
		"MAZA":  "MR8DSEXmGs2jsQR13CKtdQyaJ2jVM2bTPH",
		"BLK":   "BMg8pX5ZXMSqRZvqRhzvWHAuZRHuY3xvBm",
		"GCR":   "Gb4ycvxJkhUE4CjxkAL4oui8jyeAot5z37",
		"JBS":   "JbkzYUSkJbnc9NSPsFzfEY54tVvtKQSJpR",
		"MONA":  "MR8DSEXmGs2jsQR13CKtdQyaJ2jVM2bTPH",
		"NBT":   "2L9nGGYZPPJNHBX5DVMHq16MW3A4VZ3SG1k",
		"NEOS":  "SeX4EeQWWD48W3E8Mef2w3WoUb5kaW9sY7",
		"CCN":   "CZgwmqyRetqTssM6VxzsxezGSw4jjGhg9A",
		"DOGE":  "DNN9k4a15FmDWjdGYofWvuXqhwad68TvAm",
		"FTC":   "6vxVzKoPiNMCRoPtA6LRAuBFdsTXkBegGN",
		"LTC":   "LcT1U1wBrW6zEY8pzMfFfAS132DbwMbiKy",
		"BTC":   "1JE4CodMmqrvyjSfpDfxP9NEporKrqDeJs",
		"CDN":   "CZgwmqyRetqTssM6VxzsxezGSw4jjGhg9A",
		"NSR":   "2L9nGGYZPPJNHBX5DVMHq16MW3A4VZ3SG1k",
		"NVC":   "4Wvt5g1fTHZwXCZN1aLWG9YXrqusTVVM2S",
		"PPC":   "PRpEMn2CpmM7xa7SAHzV43LWSZ2CwjrA95",
		"BTCD":  "RSWFHKWeNffW3josHPf5UfhSb5JvQYdKxn",
		"GRS":   "FnPmeiMjLLYURLTnhKfRqfAZUy8HPL3HGt",
		"TEST":  "mxk1VriLasJBkqvHXneLD4aZgoT2gfaYJs",
		"AC":    "AYzvrJUz6zX5nhefNsLHY2dLJQn2AEaCEJ",
		"ETH":   "1JE4CodMmqrvyjSfpDfxP9NEporKrqDeJs",
		"PKB":   "PRpEMn2CpmM7xa7SAHzV43LWSZ2CwjrA95",
		"POT":   "PRpEMn2CpmM7xa7SAHzV43LWSZ2CwjrA95",
		"DASH":  "Xsuu34HFjZ5X8g3Fg6zBEg42f9S1rA4FJJ",
		"DGB":   "DNN9k4a15FmDWjdGYofWvuXqhwad68TvAm",
		"DGC":   "DNN9k4a15FmDWjdGYofWvuXqhwad68TvAm",
		"???":   "AYzvrJUz6zX5nhefNsLHY2dLJQn2AEaCEJ",
		"CLAM":  "xRXh6gTq3ZLy27YYAhJcpKf23v9ZDvJ3NN",
		"DOPE":  "4Wvt5g1fTHZwXCZN1aLWG9YXrqusTVVM2S",
		"EFL":   "LcT1U1wBrW6zEY8pzMfFfAS132DbwMbiKy",
		"NMC":   "NDoRQT8LhDxVWGhB62zXbfX9Z3FNic1RMP",
		"OK":    "PRpEMn2CpmM7xa7SAHzV43LWSZ2CwjrA95",
		"RBY":   "RqqrGRow5r8NsAwxJozPxnyEDaZsBzJYDH",
	}
)

func testCoin(t *testing.T, cfg Params, expectAddr string, testName string) {
	seed, err := hdwallet.NewSeedFromMnemonic(testPhrase, testPW)
	if err != nil {
		t.Error(err)
	}
	key, err := hdkeychain.NewMaster(seed, &cfg)
	if err != nil {
		t.Error(err)
	}
	addr, err := key.Address(&cfg)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != expectAddr {
		t.Logf("[%s] expected:\n%s\ngot:\n%s\n", testName, expectAddr, addr.String())
		t.Error("not expect")
	}
	fmt.Printf("\"%s\": \"%s\",\n", testName, addr)

}

func TestFakecoin(t *testing.T) {
	testCoin(t, FAKEMainnet, "FnPmeiMjLLYURLTnhKfRqfAZUy8HPL3HGt", "FAKE")
}

func TestAlts(t *testing.T) {
	for i := range altcoins {
		testCoin(t, Basic(Alt(altcoins[i].toAlt(i))), gold[i], i)
	}
}

func TestBitcoin(t *testing.T) {
}

func TestLitecoin(t *testing.T) {
}
