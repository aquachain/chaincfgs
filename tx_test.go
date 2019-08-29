package chaincfgs

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"

	"github.com/btcsuite/btcutil"
)

func TestTx(t *testing.T) {
	// Ordinarily the private key would come from whatever storage mechanism
	// is being used, but for this example just hard code it.
	for _, cfg := range AllCoins {

		fmt.Printf("\n[txtest] %s\n", cfg.Name)

		fromAddr := getAddress(cfg, 0)
		privKey, err := getKey(cfg, 0).ECPrivKey()
		if err != nil {
			panic(err)
		}

		toAddr := getAddress(cfg, 1)
		fmt.Println("From: ", fromAddr)
		fmt.Println("To:   ", toAddr)

		// For this example, create a fake transaction that represents what
		// would ordinarily be the real transaction that is being spent.  It
		// contains a single output that pays to address in the amount of 1 BTC.
		originTx := wire.NewMsgTx(wire.TxVersion)
		prevOut := wire.NewOutPoint(&chainhash.Hash{}, ^uint32(0))
		txIn := wire.NewTxIn(prevOut, []byte{txscript.OP_0, txscript.OP_0}, nil)
		originTx.AddTxIn(txIn)
		pkScript, err := txscript.PayToAddrScript(toAddr)
		if err != nil {
			t.Error(err)
			fmt.Println(err)
			return
		}
		txOut := wire.NewTxOut(100000000, pkScript)
		originTx.AddTxOut(txOut)
		originTxHash := originTx.TxHash()

		// Create the transaction to redeem the fake transaction.
		redeemTx := wire.NewMsgTx(wire.TxVersion)

		// Add the input(s) the redeeming transaction will spend.  There is no
		// signature script at this point since it hasn't been created or signed
		// yet, hence nil is provided for it.
		prevOut = wire.NewOutPoint(&originTxHash, 0)
		txIn = wire.NewTxIn(prevOut, nil, nil)
		redeemTx.AddTxIn(txIn)

		// Ordinarily this would contain that actual destination of the funds,
		// but for this example don't bother.
		txOut = wire.NewTxOut(0, nil)
		redeemTx.AddTxOut(txOut)

		// Sign the redeeming transaction.
		lookupKey := func(a btcutil.Address) (*btcec.PrivateKey, bool, error) {
			// Ordinarily this function would involve looking up the private
			// key for the provided address, but since the only thing being
			// signed in this example uses the address associated with the
			// private key from above, simply return it with the compressed
			// flag set since the address is using the associated compressed
			// public key.
			//
			// NOTE: If you want to prove the code is actually signing the
			// transaction properly, uncomment the following line which
			// intentionally returns an invalid key to sign with, which in
			// turn will result in a failure during the script execution
			// when verifying the signature.
			//
			// privKey.D.SetInt64(12345)
			//
			return privKey, true, nil
		}
		// Notice that the script database parameter is nil here since it isn't
		// used.  It must be specified when pay-to-script-hash transactions are
		// being signed.
		sigScript, err := txscript.SignTxOutput(&cfg,
			redeemTx, 0, originTx.TxOut[0].PkScript, txscript.SigHashAll,
			txscript.KeyClosure(lookupKey), nil, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		redeemTx.TxIn[0].SignatureScript = sigScript

		// Prove that the transaction has been validly signed by executing the
		// script pair.
		flags := txscript.ScriptBip16 | txscript.ScriptVerifyDERSignatures |
			txscript.ScriptStrictMultiSig |
			txscript.ScriptDiscourageUpgradableNops
		vm, err := txscript.NewEngine(originTx.TxOut[0].PkScript, redeemTx, 0,
			flags, nil, nil, -1)
		if err != nil {
			t.Error(err)
			fmt.Println(err)
			return
		}
		if err := vm.Execute(); err != nil {
			t.Log("will fail:")
			t.Error(err)
			fmt.Println(err)
			return
		}

		fmt.Printf("%x\n", sigScript)
		fmt.Println("Transaction successfully signed")
	}
}
