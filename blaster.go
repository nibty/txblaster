package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"sync/atomic"
	"time"

	"math/big"
	"sync"
)

type Blaster struct {
	rpcUrl            string
	funderPrivateKey  string
	mnemonic          string
	chainId           *big.Int
	numAccounts       int
	numWorkers        int
	noWait            bool
	gas               uint64
	maxGasFee         uint64
	maxGasPriorityFee uint64
	waitForTx         float64
	xenTorrentAddress common.Address
	xenTorrentTerm    int64
	xenTorrentVMus    int64

	funder        *PrivateKeyAccount
	accounts      []accounts.Account
	wallet        *hdwallet.Wallet
	ctx           context.Context
	client        *ethclient.Client
	gappedService *GappedService
	transaction   *transaction
	xenCrypto     *XenCrypto
}

type AccountJob struct {
	Account    accounts.Account
	Difference *big.Int
	Nonce      uint64
}

type BalanceResult struct {
	Account accounts.Account
	Balance *big.Int
}

// NewBlaster creates a new Blaster instance with the provided parameters.
// rpcUrl: The URL of the RPC server.
func NewBlaster(
	rpcUrl string,
	funderPrivateKey string,
	mnemonic string,
	workers int,
	accounts int,
	noWait bool,
	gas uint64,
	maxGasFee uint64,
	maxGasPriorityFee uint64,
	xenTorrentAddress string,
	xenTorrentTerm int64,
	xenTorrentVMus int64,
	waitForTx float64) *Blaster {
	return &Blaster{
		rpcUrl:            rpcUrl,
		funderPrivateKey:  funderPrivateKey,
		mnemonic:          mnemonic,
		numWorkers:        workers,
		numAccounts:       accounts,
		noWait:            noWait,
		gas:               gas,
		maxGasFee:         maxGasFee,
		maxGasPriorityFee: maxGasPriorityFee,
		waitForTx:         waitForTx,
		xenTorrentAddress: common.HexToAddress(xenTorrentAddress),
		xenTorrentTerm:    xenTorrentTerm,
		xenTorrentVMus:    xenTorrentVMus,
	}
}

func (b *Blaster) connect() {
	err := error(nil)
	b.client, err = ethclient.DialContext(b.ctx, b.rpcUrl)
	if err != nil {
		log.Crit("Failed to connect to the Ethereum client", "err", err)
	}
}

func (b *Blaster) PreparePrivateKeyAccount() PrivateKeyAccount {
	pk, err := crypto.HexToECDSA(b.funderPrivateKey)
	if err != nil {
		log.Crit("Failed to parse private key", "err", err)
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Crit("Failed casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return PrivateKeyAccount{
		PrivateKey: pk,
		Address:    fromAddress,
	}
}

func (b *Blaster) getChainId() *big.Int {
	chainID, err := b.client.NetworkID(b.ctx)
	if err != nil {
		log.Crit("Failed to get chain ID", "err", err)
	}
	return chainID
}

func (b *Blaster) generateAccounts(numAccounts int) (*hdwallet.Wallet, []accounts.Account) {
	log.Debug("Generating accounts", "numAccounts", numAccounts)
	wallet, err := hdwallet.NewFromMnemonic(b.mnemonic)
	if err != nil {
		log.Crit("Failed to parse mnemonic", err)
	}

	accts := make([]accounts.Account, numAccounts)
	for i := 0; i < numAccounts; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
		account, err := wallet.Derive(path, true)
		if err != nil {
			log.Crit("Failed to derive wallet path", "err", err)
		}
		log.Trace("Account", "address", account.Address.Hex())
		accts[i] = account
	}

	return wallet, accts
}

// Prepare sets up the Blaster by connecting to the Ethereum client, preparing the funder account,
// getting the chain ID, and generating the accounts.
func (b *Blaster) Prepare() {
	b.ctx = context.Background()
	b.connect()

	b.funder = NewPrivateKeyAccount(b.funderPrivateKey)
	b.chainId = b.getChainId()
	b.wallet, b.accounts = b.generateAccounts(b.numAccounts)

	b.gappedService = NewGappedService(b.ctx, b.client, b.chainId, b.numWorkers)
	b.gappedService.Update()

	b.transaction = NewTransaction(b.ctx, b.client, b.chainId, b.gas, b.maxGasFee, b.maxGasPriorityFee)

	b.xenCrypto = NewXenCrypto(b.ctx, b.client, b.chainId, b.gas, b.maxGasFee, b.maxGasPriorityFee, b.xenTorrentAddress)
}

// FundAccounts funds the accounts managed by the Blaster. Each account is funded with the specified amount.
// fundValue: The amount to fund each account.
func (b *Blaster) FundAccounts(fundValue int) {
	// check pending transactions
	for {
		if b.gappedService.CheckPendingByAddress(b.funder.Address) {
			log.Info("Found pending TXs. Waiting for 30 seconds...", "address", b.funder.Address.Hex())
			time.Sleep(30 * time.Second)
		} else {
			break
		}

		b.gappedService.Update()
	}

	// Fill gaps in the funder account if needed
	b.gappedService.FillGaps(b.funder.Address, b.funder.PrivateKey)

	// Generate nonces ahead of time
	nonce, err := b.client.PendingNonceAt(b.ctx, b.funder.Address)
	if err != nil {
		log.Crit("Failed to get account nonce", "err", err)
	}
	log.Debug("PrivateKeyAccount Nonce", "address", b.funder.Address, "nonce", nonce)

	// Check balances
	accountsNeedingFunding := b.checkBalances(fundValue, nonce)
	numOfJobs := len(accountsNeedingFunding)

	if numOfJobs == 0 {
		log.Info("No accounts need funding")
		return
	}

	jobs := make(chan *AccountJob, numOfJobs)
	results := make(chan *types.Receipt, numOfJobs)

	// Start workers
	for w := 1; w <= b.numWorkers; w++ {
		go b.fundWorker(w, jobs, results) // Pass the nonces map to the fundWorker
	}

	// Create jobs
	for _, account := range accountsNeedingFunding {
		jobs <- account
	}
	close(jobs)

	log.Info("Need to fund accounts. Starting funder now...", "numAccounts", numOfJobs)

	// Start a goroutine that logs the progress every 30 seconds
	var processedAccounts int32 = 0
	quitProgress := make(chan bool)
	go progressLogger(&processedAccounts, int32(numOfJobs), "Funding Accounts Report", quitProgress)

	// Collect results
	for range numOfJobs {
		txHash := <-results
		if !b.noWait {
			log.Debug("TX confirmed", "from", b.funder.Address.Hex(), "to", txHash.TxHash.Hex())
		}
		atomic.AddInt32(&processedAccounts, 1)
	}
	quitProgress <- true
}

func (b *Blaster) checkBalances(fundValue int, nonce uint64) []*AccountJob {
	desiredBalance := new(big.Int).Mul(big.NewInt(int64(fundValue)), big.NewInt(params.Ether))

	balanceJobs := make(chan accounts.Account, len(b.accounts))
	balanceResults := make(chan *BalanceResult, len(b.accounts))

	for w := 1; w <= b.numWorkers; w++ {
		go b.balanceWorker(w, balanceJobs, balanceResults) // Pass the nonces map to the fundWorker
	}

	// generate array of accounts needing funding
	for _, account := range b.accounts {
		balanceJobs <- account
	}
	close(balanceJobs)

	log.Info("Checking account balances...")
	accountsNeedingFunding := make([]*AccountJob, 0)
	for range len(b.accounts) {
		balanceResult := <-balanceResults
		log.Trace("Account balance", "address", balanceResult.Account.Address.Hex(), "balance", ToDecimal(balanceResult.Balance, 18))

		if balanceResult.Balance.Cmp(desiredBalance) < 0 {
			log.Debug("Account needs funding", "address", balanceResult.Account.Address.Hex(), "balance", ToDecimal(balanceResult.Balance, 18), "desired", ToDecimal(desiredBalance, 18))
			difference := new(big.Int).Sub(desiredBalance, balanceResult.Balance)
			accountsNeedingFunding = append(accountsNeedingFunding, &AccountJob{Account: balanceResult.Account, Difference: difference, Nonce: nonce})
			nonce++
		} else {
			log.Debug("Account does not need funding", "address", balanceResult.Account.Address.Hex(), "balance", ToDecimal(balanceResult.Balance, 18), "desired", ToDecimal(desiredBalance, 18))
		}
	}

	return accountsNeedingFunding
}

func progressLogger(processed *int32, total int32, title string, quit chan bool) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			processed := atomic.LoadInt32(processed)
			printProgressMsg(&processed, total, title)
			return
		case <-ticker.C:
			processed := atomic.LoadInt32(processed)
			printProgressMsg(&processed, total, title)
		}
	}
}

func printProgressMsg(processed *int32, total int32, title string) {
	p := message.NewPrinter(language.English)

	if total > 0 {
		percent := float64(*processed) / float64(total) * 100
		log.Info(title, "processed", p.Sprint(*processed), "total", p.Sprint(total), "percent", percent)
	} else {
		log.Info(title, "processed", p.Sprint(*processed))
	}
}

func (b *Blaster) balanceWorker(id int, balanceJobs <-chan accounts.Account, results chan<- *BalanceResult) {
	for account := range balanceJobs {
		balance, err := b.client.BalanceAt(b.ctx, account.Address, nil)
		if err != nil {
			log.Crit("Failed to get account balance", "err", err)
		}
		results <- &BalanceResult{
			Account: account,
			Balance: balance,
		}
	}
}

func (b *Blaster) fundWorker(id int, jobs <-chan *AccountJob, results chan<- *types.Receipt) {
	for job := range jobs {
		log.Trace("Worker", "id", id, "account", job.Account.Address.Hex())
		txSigned, err := b.transaction.send(b.funder.PrivateKey, b.funder.Address, job.Account.Address, job.Difference, job.Nonce)
		if err != nil {
			log.Crit("Failed to fund account", "err", err)
		}
		txHash, err := bind.WaitMined(b.ctx, b.client, txSigned)
		if err != nil {
			log.Crit("Failed to wait for transaction to be mined", "err", err)
		}
		log.Debug("TX confirmed", "from", b.funder.Address.Hex(), "to", txHash.TxHash.Hex())
		results <- txHash
	}
}

// RunTests runs the tests, which involve sending transactions from each account to the funder account.
// value: The value to send in each transaction (in wei).
func (b *Blaster) RunTests(value uint64) {
	b.gappedService.Update()

	// prepare the accounts in batches per worker, grab their nonce, private key, and fill any nonce gaps
	log.Info("Preparing accounts")
	batches := SplitAccountsIntoBatches(b.accounts, b.numWorkers)

	// prepare the accounts in batches per worker, grab their nonce and private key
	accountChan := make(chan *AccountBatchEntry, b.numWorkers)
	var wg sync.WaitGroup
	for i := 0; i < b.numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for account := range accountChan {
				// get the private key for the account
				pk, err := b.wallet.PrivateKey(account.Account)
				if err != nil {
					log.Crit("Failed to get private key from wallet", "address", account.Account.Address.Hex(), "err", err)
				}
				account.PrivateKey = pk

				// get the nonce for the account
				nonce, err := b.client.PendingNonceAt(b.ctx, account.Account.Address)
				if err != nil {
					log.Crit("Failed to get account nonce", "err", err)
				}
				account.Nonce = nonce
			}
		}()
	}

	for _, batch := range batches {
		for _, account := range batch {
			accountChan <- account
		}
	}

	close(accountChan)
	wg.Wait()

	// fill any nonce gaps if needed
	if b.gappedService.CheckQueued() {
		log.Info("Checking for queued transactions and filling gaps")
		for _, batch := range batches {
			for _, account := range batch {
				b.gappedService.FillGaps(account.Account.Address, account.PrivateKey)
			}
		}
	}

	var processedTxs int32 = 0
	quitProgress := make(chan bool)
	go progressLogger(&processedTxs, 0, "Blaster Report", quitProgress)

	log.Info("Start the blasting!")
	for _, batch := range batches {
		wg.Add(1)
		go func(accs []*AccountBatchEntry) {
			defer wg.Done()

			for {
				for _, acc := range accs {
					var signTx *types.Transaction
					var err error
					if b.xenTorrentAddress != common.HexToAddress("0x0") {
						signTx, err = b.xenCrypto.MintXeNFT(acc.PrivateKey, acc.Nonce, b.xenTorrentTerm, b.xenTorrentVMus)
					} else {
						signTx, err = b.transaction.send(acc.PrivateKey, acc.Account.Address, b.funder.Address, big.NewInt(0).SetUint64(value), acc.Nonce)
					}
					if err != nil {
						log.Crit("Failed to send transaction", "from", acc.Account.Address, "err", err)
					}

					acc.Nonce++

					start := time.Now().UnixMilli()
					time.Sleep(time.Duration(b.waitForTx) * time.Second)
					mined, err := bind.WaitMined(b.ctx, b.client, signTx)
					if err != nil {
						log.Crit("Failed to wait for tx confirmation", "from", acc.Account.Address, "err", err)
					}
					elapsed := float64(time.Now().UnixMilli()-start) / 1000.0
					log.Debug("TX confirmed", "from", acc.Account.Address.Hex(), "to", mined.TxHash.String(), "elapsed", elapsed)

					atomic.AddInt32(&processedTxs, 1)
				}
			}
		}(batch)
	}
	wg.Wait()
	quitProgress <- true
}
