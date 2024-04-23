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

type TxNonceMap map[string]string

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

func (g *GappedService) updateQueued() {
	if !g.CheckQueued() {
		return
	}

	// check the rpc txpool.Content for gapped transactions
	err := g.client.Client().CallContext(g.ctx, &txPoolContent, "txpool_inspect")
	if err != nil {
		log.Crit("Failed to get txpool content", "err", err)
	}
}

func (g *GappedService) GetQueuedByAddress(address common.Address) *TxAccountMap {
	for key, _ := range txPoolContent.Queued {
		if key != address.Hex() {
			delete(txPoolContent.Queued, key)
		}
	}
	return &txPoolContent.Queued
}

func (g *GappedService) GetGappedNonces(currentNonce uint64, address common.Address) *[]uint64 {
	txPoolContent := TxPoolContent{}

	// check the rpc txpool.Content for gapped transactions
	err := g.client.Client().CallContext(g.ctx, &txPoolContent, "txpool_inspect")
	if err != nil {
		log.Crit("Failed to get txpool content", "err", err)
	}

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
	if !g.CheckQueued() {
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
