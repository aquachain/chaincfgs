package chaincfgs_test

import (
	"fmt"
	"log"

	"github.com/aquachain/chaincfgs"
	"github.com/aquachain/hdwallet"
)

func ExampleCoin() {
	coincfg, err := chaincfgs.Coin("NMC")
	if err != nil {
		log.Fatalln(err)
	}
	wallet, err := hdwallet.NewFromMnemonic("tag volcano eight thank tide danger coast health above argue embrace heavy", "")
	if err != nil {
		log.Fatalln(err)
	}

	// bitcoin-style address
	addr, err := wallet.MasterKey().Address(&coincfg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(addr)
}
