package main

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//----------------------------Benchmarks----------------------------

// func BenchmarkIPC(b *testing.B) {

// 	ipcPath := ""
// 	client, err := ethclient.Dial(ipcPath)
// 	if err != nil {
// 		b.Fatal("Error when initializing client", err)
// 	}

// 	targetAddress := common.HexToAddress("0x29F06Dc85D64B9748465735A2C61a524Ffd28BDe")

// 	for n := 0; n < b.N; n++ {
// 		client.BalanceAt(context.Background(), targetAddress, nil)
// 	}
// }

func BenchmarkRPC(b *testing.B) {

	rpcPort := "https://eth-mainnet.g.alchemy.com/v2/AXIyCY50RSPKbuzwVYONrQJ-7R9UFGvm"
	client, err := ethclient.Dial(rpcPort)
	if err != nil {
		b.Fatal("Error when initializing client", err)
	}

	targetAddress := common.HexToAddress("0x29F06Dc85D64B9748465735A2C61a524Ffd28BDe")

	for n := 0; n < b.N; n++ {
		client.BalanceAt(context.Background(), targetAddress, nil)
	}
}