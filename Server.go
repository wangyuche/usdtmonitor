package main

import (
	"os"

	"github.com/wangyuche/usdtmonitor/src/eth"
)

type iweb3 interface {
	Monitor(url string)
}

type Web3Type string

const (
	ethType  Web3Type = "eth"
	tronType Web3Type = "tron"
)

func New(t Web3Type) iweb3 {
	switch t {
	case ethType:
		return &eth.ETH{}
	}
	return nil
}

func main() {
	a := New(ethType)
	a.Monitor(os.Getenv("ETHRPCAddress"))
}
