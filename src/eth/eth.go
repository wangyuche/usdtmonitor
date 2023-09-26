package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wangyuche/usdtmonitor/sol/erc20"
)

type ETH struct {
}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

func (this *ETH) Monitor(url string) {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("ETHContractAddress"))
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(18218382),
		ToBlock:   big.NewInt(18218382),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(erc20.Erc20ABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		fmt.Printf("Log TxHash: %s\n", vLog.TxHash.String())
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent LogTransfer
			aaa, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			transferEvent.Tokens = aaa[0].(*big.Int)
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())
		default:
		}
	}

}
