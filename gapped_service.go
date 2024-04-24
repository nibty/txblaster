package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sort"
	"strconv"
	"sync"
)

var (
	txPoolContent = TxPoolContent{}
)

type GappedService struct {
	chainId    *big.Int
	numWorkers int
	client     *ethclient.Client
	ctx        context.Context
}

type TXStatus struct {
	Pending *hexutil.Big `json:"pending"`
	Queued  *hexutil.Big `json:"queued"`
}

type TxPoolContent struct {
	Pending TxAccountMap `json:"pending"`
	Queued  TxAccountMap `json:"queued"`
}

type TxAccountMap map[string]TxNonceMap

type TxNonceMap map[string]Transaction

type Transaction struct {
	Hash common.Hash `json:"hash"`
}

func NewGappedService(ctx context.Context, client *ethclient.Client, chainId *big.Int, numWorkers int) *GappedService {
	return &GappedService{
		chainId:    chainId,
		numWorkers: numWorkers,
		client:     client,
		ctx:        ctx,
	}
}

func (g *GappedService) checkStatus() *TXStatus {
	txStatus := TXStatus{}
	//var raw json.RawMessage

	// check the rpc txpool.Status for gapped transactions
	err := g.client.Client().CallContext(g.ctx, &txStatus, "txpool_status")
	if err != nil {
		log.Crit("Failed to get txpool status", "err", err)
	}

	return &txStatus
}

func (g *GappedService) CheckQueued() bool {
	txStatus := g.checkStatus()
	return txStatus.Queued.ToInt().Cmp(big.NewInt(0)) > 0
}

func (g *GappedService) CheckQueuedByAddress(address common.Address) bool {
	queued := g.GetQueuedByAddress(address)
	return len(*queued) > 0
}

func (g *GappedService) CheckPending() bool {
	txStatus := g.checkStatus()
	return txStatus.Pending.ToInt().Cmp(big.NewInt(0)) > 0
}

func (g *GappedService) Update() {
	txPoolContent = TxPoolContent{}
	// check the rpc txpool.Content for gapped transactions
	err := g.client.Client().CallContext(g.ctx, &txPoolContent, "txpool_content")
	if err != nil {
		log.Crit("Failed to get txpool content", "err", err)
	}
}

func (g *GappedService) GetQueuedByAddress(address common.Address) *TxAccountMap {
	// create copy of txPoolContent.Queued
	queued := make(TxAccountMap)
	for key, _ := range txPoolContent.Queued {
		if key == address.Hex() {
			queued[key] = txPoolContent.Queued[key]
		}
	}
	return &queued
}

func (g *GappedService) GetPendingTxByAddress(address common.Address) []common.Hash {
	var pendingTxs []common.Hash
	for key, _ := range txPoolContent.Pending {
		log.Trace("Pending TX", "key", key, "address", address.Hex())
		if key == address.Hex() {
			for _, tx := range txPoolContent.Pending[key] {
				pendingTxs = append(pendingTxs, tx.Hash)
			}
		}
	}
	return pendingTxs
}

func (g *GappedService) CheckPendingByAddress(address common.Address) bool {
	for key, _ := range txPoolContent.Pending {
		if key == address.Hex() {
			return true
		}
	}
	return false
}

func (g *GappedService) GetGappedNonces(currentNonce uint64, address common.Address) *[]uint64 {
	value := txPoolContent.Queued[address.Hex()]

	keys := make([]uint64, 0, len(value))
	for k := range value {
		i, err := strconv.ParseUint(k, 10, 64)
		if err != nil {
			log.Crit("Failed to parse nonce", "err", err)
		}

		keys = append(keys, i)
	}

	if len(keys) == 0 {
		return &keys
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	firstGaps := keys[0] - currentNonce
	gappedNonces := make([]uint64, firstGaps)
	gappedNonces[0] = currentNonce
	// loop though range of firstgaps and add to gappedNonces
	for i := uint64(1); i < firstGaps; i++ {
		gappedNonces[i] = currentNonce + i
	}

	// loop through keys and add and gaps to gappedNonces
	for i := 0; i < len(keys)-1; i++ {
		gap := keys[i+1] - keys[i]
		for j := uint64(1); j < gap; j++ {
			gappedNonces = append(gappedNonces, keys[i]+j)
		}
	}

	return &gappedNonces
}

func (g *GappedService) FillGaps(address common.Address, privateKey *ecdsa.PrivateKey) {
	if !g.CheckQueuedByAddress(address) {
		log.Trace("No queued transactions")
		return
	}

	currentNonce, err := g.client.PendingNonceAt(g.ctx, address)
	if err != nil {
		log.Crit("Failed to get account nonce", "err", err)
	}

	gappedNonces := g.GetGappedNonces(currentNonce, address)
	if len(*gappedNonces) == 0 {
		log.Debug("No gaps to fill", "address", address.Hex())
		return
	}

	log.Info("Filling gaps", "address", address.Hex(), "gappedNonces", gappedNonces)

	nonceChan := make(chan uint64, 10)
	var wg sync.WaitGroup

	for i := 0; i < g.numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for nonce := range nonceChan {
				// create a transaction
				tx := NewTransaction(g.ctx, g.client, g.chainId, 21000, 0, 0)

				// send the transaction
				txSigned, err := tx.send(privateKey, address, address, big.NewInt(0), nonce)
				if err != nil {
					log.Crit("Failed to send transaction", "err", err)
				}

				// wait for the transaction to be mined
				txHash, err := bind.WaitMined(g.ctx, g.client, txSigned)
				if err != nil {
					log.Crit("Failed to wait for transaction to be mined", "err", err)
				}

				log.Info("TX confirmed", "from", address.Hex(), "to", txHash.TxHash.Hex())
			}
		}()
	}

	for _, nonce := range *gappedNonces {
		nonceChan <- nonce
	}

	close(nonceChan)
	wg.Wait()
}

func (g *GappedService) ReplacePending(address common.Address, privateKey *ecdsa.PrivateKey) {
	if !g.CheckPending() {
		log.Trace("No pending transactions")
		return
	}

	// get the pending transactions
	pendingTxs := g.GetPendingTxByAddress(address)
	log.Info("Replacing Pending transactions", "address", address.Hex(), "num_txs", len(pendingTxs))

	for _, pendingTx := range pendingTxs {
		// get transaction details. such as nonce and gas price
		tx, _, err := g.client.TransactionByHash(g.ctx, pendingTx)
		if err != nil {
			log.Crit("Failed to get transaction details", "tx_hash", pendingTx.Hex(), "err", err)
		}

		// create a replacement transaction with the same nonce and increased gas price 2x uint64
		newGasPrice := new(big.Int).Mul(tx.GasPrice(), big.NewInt(2)).Uint64()
		newGasTipPrice := new(big.Int).Mul(tx.GasTipCap(), big.NewInt(2)).Uint64()

		log.Debug("Replacing pending transaction",
			"tx_hash", pendingTx,
			"nonce", tx.Nonce(),
			"gasPrice", tx.GasPrice(),
			"gasTipPrice", tx.GasTipCap(),
			"replacement_gasPrice", newGasPrice,
			"replacement_gasTipPrice", newGasTipPrice)

		currentGasPrice, err := g.client.SuggestGasPrice(g.ctx)
		if err != nil {
			log.Crit("Failed to get current gas price", "err", err)
		}

		currentGasTipPrice, err := g.client.SuggestGasTipCap(g.ctx)
		if err != nil {
			log.Crit("Failed to get current gas tip price", "err", err)
		}

		if newGasPrice < currentGasPrice.Uint64() {
			newGasPrice = currentGasPrice.Uint64()
		}

		if newGasTipPrice < currentGasTipPrice.Uint64() {
			newGasTipPrice = currentGasTipPrice.Uint64()
		}

		transaction := NewTransaction(g.ctx, g.client, g.chainId, tx.Gas(), newGasPrice, newGasTipPrice)
		send, err := transaction.send(privateKey, address, address, big.NewInt(0), tx.Nonce())
		if err != nil {
			log.Crit("Failed to send replacement transaction", "err", err)
		}

		// wait for the transaction to be mined
		txHash, err := bind.WaitMined(g.ctx, g.client, send)
		if err != nil {
			log.Crit("Failed to wait for transaction to be mined", "err", err)
		}

		log.Info("TX confirmed", "from", address.Hex(), "to", txHash.TxHash.Hex())
	}

}
