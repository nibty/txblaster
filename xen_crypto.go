package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"main/contracts/xentorrent"
	"math/big"
)

type XenCrypto struct {
	chainId           *big.Int
	account           *PrivateKeyAccount
	gas               uint64
	maxGasFee         uint64
	maxGasPriorityFee uint64

	xenTorrentTransactor *xentorrent.XentorrentTransactor

	client *ethclient.Client
	ctx    context.Context
}

func NewXenCrypto(
	ctx context.Context,
	client *ethclient.Client,
	chainId *big.Int,
	gas uint64,
	maxGasFee uint64,
	maxGasPriorityFee uint64,
	xenTorrentAddress common.Address) *XenCrypto {
	xenTorrentTransactor, err := xentorrent.NewXentorrentTransactor(xenTorrentAddress, client)
	if err != nil {
		log.Crit("Failed to create transactor", "err", err)
	}

	return &XenCrypto{
		chainId:           chainId,
		gas:               gas,
		maxGasFee:         maxGasFee,
		maxGasPriorityFee: maxGasPriorityFee,

		xenTorrentTransactor: xenTorrentTransactor,
		client:               client,
		ctx:                  ctx,
	}
}

func (x XenCrypto) MintXeNFT(privateKey *ecdsa.PrivateKey, nonce uint64, vmus int64, term int64) (*types.Transaction, error) {
	gasTipCap := big.NewInt(0)
	if x.maxGasFee != 0 {
		gasTipCap = new(big.Int).SetUint64(x.maxGasPriorityFee)
	} else {
		err := error(nil)
		gasTipCap, err = x.client.SuggestGasTipCap(x.ctx)
		if err != nil {
			return nil, err
		}
	}

	gasPrice := big.NewInt(0)
	if x.maxGasFee != 0 {
		gasPrice = new(big.Int).SetUint64(x.maxGasFee)
	} else {
		err := error(nil)
		gasPrice, err = x.client.SuggestGasPrice(x.ctx)
		if err != nil {
			return nil, err
		}
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, x.chainId)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = x.gas
	auth.GasFeeCap = gasPrice
	auth.GasTipCap = gasTipCap

	tx, err := x.xenTorrentTransactor.BulkClaimRank(auth, big.NewInt(vmus), big.NewInt(term))
	if err != nil {
		return nil, err
	}

	return tx, nil
}
