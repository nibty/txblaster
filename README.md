# TX Blaster

## Install

> Copy the `.env.example` file to `.env` and update the values as needed
```shell
cp .env.example .env
```

## Run

> Fund 5,000 accounts with 1,000 XN each then blast the network with 1,000 concurrent workers
```shell
go run main --rpc ws://localhost:8545 --workers 1000 --accounts 5000 --fund-value 10000
```

> Enable verbose logging
```shell
go run main --rpc ws://localhost:8545 --workers 1000 --accounts 5000 --fund-value 10000 --verbose 4
```

> See full list of options
```shell
go run main --help
```
