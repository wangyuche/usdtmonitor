package polygon

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wangyuche/goutils/log"
	"github.com/wangyuche/usdtmonitor/sol/erc20"
	"github.com/wangyuche/usdtmonitor/src/structs"
)

type POLYGON struct {
	RPCAddress          string
	USDTContractAddress string
	conn                *ethclient.Client
}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

func (this *POLYGON) Monitor(url string) {

}

func (this *POLYGON) Init() error {
	c, err := ethclient.Dial(this.RPCAddress)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.conn = c
	return nil
}

func (this *POLYGON) GetNowBlockID() (int64, error) {
	header, err := this.conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return header.Number.Int64(), nil
}

func (this *POLYGON) GetUSDTLogByBlockID(blockid int64) ([]structs.USDTLog, error) {
	contractAddress := common.HexToAddress(this.USDTContractAddress)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockid),
		ToBlock:   big.NewInt(blockid),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := this.conn.FilterLogs(context.Background(), query)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(erc20.Erc20ABI)))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	var usdtlogs []structs.USDTLog = make([]structs.USDTLog, 0)
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			var ulog structs.USDTLog
			unpackdata, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				log.Error(err.Error())
				return nil, err
			}
			ulog.From = common.HexToAddress(vLog.Topics[1].Hex()).String()
			ulog.To = common.HexToAddress(vLog.Topics[2].Hex()).String()
			ulog.Tokens = unpackdata[0].(*big.Int).Int64()
			usdtlogs = append(usdtlogs, ulog)
		}
	}
	return usdtlogs, nil
}
