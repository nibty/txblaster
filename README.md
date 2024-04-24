# TX Blaster

## Install

> Copy the `.env.example` file to `.env` and update the values as needed
```shell
cp .env.example .env
```

## Run

> Fund 10,000 accounts with 1,000 XN each then blast the network with 2,000 concurrent workers
```shell
go run main --rpc ws://localhost:8545 --workers 2000 --accounts 10000 --fund-value 100000
```

> See full list of options
```shell
go run main --help
```
