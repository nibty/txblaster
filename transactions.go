package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type transaction struct {
	client *ethclient.Client
	ctx    context.Context

	chainId           *big.Int
	gas               uint64
	maxGasFee         uint64
	maxGasPriorityFee uint64
}

func NewTransaction(ctx context.Context, client *ethclient.Client, chainId *big.Int, gas uint64, maxGasFee uint64, maxGasPriorityFee uint64) *transaction {
	return &transaction{
		client:            client,
		ctx:               ctx,
		chainId:           chainId,
		gas:               gas,
		maxGasFee:         maxGasFee,
		maxGasPriorityFee: maxGasPriorityFee,
	}
}

func (b *transaction) send(privateKey *ecdsa.PrivateKey, fromAddress common.Address, toAddress common.Address, value *big.Int, nonce uint64) (*types.Transaction, error) {
	gasTipCap := big.NewInt(0)
	if b.maxGasFee != 0 {
		gasTipCap = new(big.Int).SetUint64(b.maxGasPriorityFee)
	} else {
		err := error(nil)
		gasTipCap, err = b.client.SuggestGasTipCap(b.ctx)
		if err != nil {
			return nil, err
		}
	}

	gasPrice := big.NewInt(0)
	if b.maxGasFee != 0 {
		gasPrice = new(big.Int).SetUint64(b.maxGasFee)
	} else {
		err := error(nil)
		gasPrice, err = b.client.SuggestGasPrice(b.ctx)
		if err != nil {
			return nil, err
		}
	}

	err := error(nil)

	log.Debug("Send TX",
		"from", fromAddress.Hex(),
		"to", toAddress,
		"nonce", nonce,
		"value", ToDecimal(value, 18),
		"maxGasFee", gasPrice,
		"maxGasPriorityFee", gasTipCap,
	)

	txParams := &types.DynamicFeeTx{
		ChainID:   b.chainId,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasPrice,
		Gas:       b.gas,
		To:        &toAddress,
		Value:     value,
		Data:      nil,
	}

	tx := types.NewTx(txParams)
	signTx, err := types.SignTx(tx, types.NewLondonSigner(b.chainId), privateKey)
	if err != nil {
		return nil, err
	}

	err = b.client.SendTransaction(b.ctx, signTx)
	if err != nil {
		return nil, err
	}

	return signTx, nil
}
