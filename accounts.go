package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type PrivateKeyAccount struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

func NewPrivateKeyAccount(privateKey string) *PrivateKeyAccount {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Crit("Failed to parse private key", "err", err)
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Crit("Failed casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &PrivateKeyAccount{
		Address:    address,
		PrivateKey: pk,
	}
}
