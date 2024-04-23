package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/log"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

const (
	FakeValidatorPrivateKey = "163f5f0f9a621d72fedd85ffca3d08d131ab4e812181e0d30ffd1c885d20aac7"
	FakeMnemonic            = "tag volcano eight thank tide danger coast health above argue embrace heavy"
	defaultNumWorkers       = 10
	defaultNumAccounts      = 10
	defaultRpcURL           = "http://127.0.0.1:8545"
)

func main() {
	logger := log.NewLogger(log.NewTerminalHandler(os.Stdout, true))
	log.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file")
	}

	privateKey, _ := strings.CutPrefix(os.Getenv("PRIVATE_KEY"), "0x")
	if privateKey == "" {
		log.Warn("No private key provided, using fake private key")
		privateKey = FakeValidatorPrivateKey
	}

	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		log.Warn("No mnemonic provided, using fake mnemonic")
		mnemonic = FakeMnemonic
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Warn("No RPC URL provided, using default RPC URL", "defaultRpcURL", defaultRpcURL)
		rpcURL = defaultRpcURL
	}

	//numTxs := flag.Int("num_txs", 100000, "Limit the number of transactions to send")
	noWait := flag.Bool("no-wait", false, "Don't wait for tx receipt")
	gas := flag.Uint64("gas", 21000, "The gas limit")
	maxGasFee := flag.Uint64("max-gas-fee", 0, "The max gas fee (gwei)")
	maxGasPriorityFee := flag.Uint64("max-gas-priority-fee", 0, "The max gas priority fee (gwei)")
	value := flag.Uint64("value", 1, "The value to send (wei)")
	rpc := flag.String("rpc", rpcURL, "The RPC URL to use")
	workers := flag.Int("workers", defaultNumWorkers, "The number of workers to use")
	accounts := flag.Int("accounts", defaultNumAccounts, "The number of accounts to use")
	fundValue := flag.Int("fund-value", 5, "The amount to fund each account.")
	shouldFundAccounts := flag.Bool("fund-accounts", true, "Fund the accounts")
	shouldFunTests := flag.Bool("run-tests", true, "Run the tests")
	verbosity := flag.Int("verbosity", 3, "The log verbosity")
	waitForTx := flag.Float64("confirm-wait", 1.2, "The time to wait before checking the tx confirmation (in seconds)")

	flag.Parse()
	if flag.Arg(0) == "-h" || flag.Arg(0) == "--help" {
		flag.Usage()
		os.Exit(0)
	}

	logger = log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.FromLegacyLevel(*verbosity), true))
	log.SetDefault(logger)

	blaster := NewBlaster(*rpc, privateKey, mnemonic, *workers, *accounts, *noWait, *gas, *maxGasFee, *maxGasPriorityFee, *waitForTx)
	blaster.Prepare()

	if *shouldFundAccounts {
		blaster.FundAccounts(*fundValue)
	}

	if *shouldFunTests {
		blaster.RunTests(*value)
	}
}
