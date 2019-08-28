// Copyright 2019 The chaincfgs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package chaincfgs is a collection of chaincfgs for bitcoin and altcoins
//
// With this package, you can interact with multiple chains in one application
//
// ..such as signing transactions, creating addresses, and using addresses.
//
// Adding a new altcoin
//
// all coins must have at least COINMainnet, and optional alternate networks.
// Testnets can be named anything but must start with capital letter.
// Examples: BTCMainnet, BTCSimnet, ALTMainnet, COINMainnet, COINTestnet10
//
// That is the extent of this package, and anything else is entirely out of scope.
//
// To add a coin, edit this 'chains.go' file, using a similar coin as template.
package chaincfgs

import (
	"github.com/btcsuite/btcd/wire"

	btcchaincfg "github.com/btcsuite/btcd/chaincfg"
	dcrchaincfg "github.com/decred/dcrd/chaincfg"
	ltcchaincfg "github.com/ltcsuite/ltcd/chaincfg"
)

type Params = btcchaincfg.Params

// Alt type is the minimal fields needed for what we need
type Alt struct {
	Name string
	Net  uint32 // wire.BitcoinNet
	// Human-readable part for Bech32 encoded segwit addresses, as defined
	// in BIP 173.
	Bech32HRPSegwit string

	// Address encoding magics
	PubKeyHashAddrID        byte // First byte of a P2PKH address
	ScriptHashAddrID        byte // First byte of a P2SH address
	PrivateKeyID            byte // First byte of a WIF private key
	WitnessPubKeyHashAddrID byte // First byte of a P2WPKH address
	WitnessScriptHashAddrID byte // First byte of a P2WSH address

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID [4]byte
	HDPublicKeyID  [4]byte

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType uint32
}

// Basic returns a *btcchaincfg.Params from the given Alt fields
func Basic(alt Alt) btcchaincfg.Params {
	return btcchaincfg.Params{
		Name:                    alt.Name,
		Net:                     wire.BitcoinNet(alt.Net),
		Bech32HRPSegwit:         alt.Bech32HRPSegwit,
		PubKeyHashAddrID:        alt.PubKeyHashAddrID,
		ScriptHashAddrID:        alt.ScriptHashAddrID,
		PrivateKeyID:            alt.PrivateKeyID,
		WitnessPubKeyHashAddrID: alt.WitnessPubKeyHashAddrID,
		WitnessScriptHashAddrID: alt.WitnessScriptHashAddrID,
		HDPrivateKeyID:          alt.HDPrivateKeyID,
		HDPublicKeyID:           alt.HDPublicKeyID,
		HDCoinType:              alt.HDCoinType,
	}
}

// these coins have implemented their own chaincfg package.
// so, let's use them if we can. otherwise, create custom Params
var (
	BTCMainnet  = btcchaincfg.MainNetParams
	BTCTestnet3 = btcchaincfg.TestNet3Params
	BTCSimnet   = btcchaincfg.SimNetParams
	BTCRegnet   = btcchaincfg.RegressionNetParams
	LTCMainnet  = ltcchaincfg.MainNetParams
	DCRMainnet  = dcrchaincfg.MainNetParams
	LTCTestnet4 = (ltcchaincfg.TestNet4Params)
	LTCSimnet   = (ltcchaincfg.SimNetParams)
	LTCRegnet   = (ltcchaincfg.RegressionNetParams)
	DCRTestnet3 = (dcrchaincfg.TestNet3Params)
	DCRSimnet   = (dcrchaincfg.SimNetParams)
	DCRRegnet   = (dcrchaincfg.RegNetParams)
)

// if the altcoin doesn't have a chaincfg package, we can make a basic one containing only relevant fields
var (
	FAKEMainnet = Basic(Alt{"FAKE", 0x12141c17, "fake", 0x24, 0x26, 0x80, 0x06, 0x0A, [4]byte{0x04, 0x88, 0xad, 0xe4}, [4]byte{0x04, 0x88, 0xb2, 0x1e}, 0})
)
