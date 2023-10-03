package main

import (
	"os"
	"strconv"

	"github.com/wangyuche/goutils/log"
	"github.com/wangyuche/usdtmonitor/src/eth"
	"github.com/wangyuche/usdtmonitor/src/polygon"
	"github.com/wangyuche/usdtmonitor/src/structs"
	"github.com/wangyuche/usdtmonitor/src/tron"
)

type iweb3 interface {
	Init() error
	GetNowBlockID() (int64, error)
	GetUSDTLogByBlockID(blockid int64) ([]structs.USDTLog, error)
	Monitor(url string)
}

type Web3Type string

const (
	ethType     Web3Type = "eth"
	polygonType Web3Type = "polygon"
	tronType    Web3Type = "tron"
)

func New(t Web3Type) iweb3 {
	switch t {
	case ethType:
		return &eth.ETH{
			RPCAddress:          os.Getenv("ETHRPCAddress"),
			USDTContractAddress: os.Getenv("ETHContractAddress"),
		}
	case polygonType:
		return &polygon.POLYGON{
			RPCAddress:          os.Getenv("POLYGONRPCAddress"),
			USDTContractAddress: os.Getenv("POLYGONContractAddress"),
		}
	case tronType:
		return &tron.TRON{
			RPCAddress:          os.Getenv("TRONRPCAddress"),
			USDTContractAddress: os.Getenv("TRONContractAddress"),
		}
	}
	return nil
}

func main() {
	log.New(log.LogType(os.Getenv("LogType")))
	chain := New(tronType)
	err := chain.Init()
	if err != nil {
		log.Fail(err.Error())
	}
	/*
		for {
			blockid, err := chain.GetNowBlockID()
			if err != nil {
				log.Error(err.Error())
				continue
			}
			log.Debug(strconv.FormatInt(blockid, 10))
			time.Sleep(1 * time.Second)

		}
	*/
	datas, err := chain.GetUSDTLogByBlockID(55228207)
	if err != nil {
		log.Fail(err.Error())
	}
	for _, data := range datas {
		log.Info(data.From)
		log.Info(data.To)
		log.Info(strconv.FormatInt(data.Tokens, 10))
	}
	//a.Monitor(os.Getenv("TRONRPCAddress"))
}
