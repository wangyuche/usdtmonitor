package tron

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/wangyuche/goutils/log"
	"github.com/wangyuche/usdtmonitor/proto/tronproto/api"
	"github.com/wangyuche/usdtmonitor/src/common"
	"github.com/wangyuche/usdtmonitor/src/structs"
	"google.golang.org/grpc"
)

type TRON struct {
	RPCAddress          string
	USDTContractAddress string
	conn                api.WalletClient
}

func (this *TRON) Monitor(url string) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	g := api.NewWalletClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {

		result, err := g.GetNowBlock2(ctx, new(api.EmptyMessage))
		if err != nil {
			panic(err)
		}

		print(result.BlockHeader.RawData.Number)
		time.Sleep(1 * time.Second)
	}
}

func (this *TRON) Init() error {
	conn, err := grpc.Dial(this.RPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.conn = api.NewWalletClient(conn)
	return nil
}

func (this *TRON) GetNowBlockID() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := this.conn.GetNowBlock(ctx, new(api.EmptyMessage))
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return result.BlockHeader.RawData.Number, nil
}

func (this *TRON) GetUSDTLogByBlockID(blockid int64) ([]structs.USDTLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	var blockidreq *api.BlockReq = &api.BlockReq{}
	blockidreq.IdOrNum = strconv.FormatInt(blockid, 10)
	blockidreq.Detail = true
	ex, err := this.conn.GetBlock(ctx, blockidreq)
	if err != nil {
		log.Error(err.Error())
		cancel()
		return nil, err
	}
	cancel()
	var usdtlogs []structs.USDTLog = make([]structs.USDTLog, 0)
	for _, data := range ex.Transactions {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		var bytemessage *api.BytesMessage = &api.BytesMessage{}
		bytemessage.Value = data.Txid
		transactionsinfo, err := this.conn.GetTransactionInfoById(ctx, bytemessage)
		if err != nil {
			log.Error(err.Error())
			cancel()
			return nil, err
		}
		cancel()
		for _, event := range transactionsinfo.Log {
			if len(event.Topics) == 3 {
				contractaddress := common.EncodeCheck(append([]byte{0x41}, event.GetAddress()...))
				if contractaddress == this.USDTContractAddress {
					var ulog structs.USDTLog
					event.Topics[1][11] = 0x41
					event.Topics[2][11] = 0x41
					amount := new(big.Int).SetBytes(common.TrimLeftZeroes(event.Data)).Int64()
					ulog.From = common.EncodeCheck(event.Topics[1][11:])
					ulog.To = common.EncodeCheck(event.Topics[1][11:])
					ulog.Tokens = amount
					usdtlogs = append(usdtlogs, ulog)
				}
			}
		}
	}
	return usdtlogs, nil
}
