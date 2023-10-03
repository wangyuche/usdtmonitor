# USDT Monitor

### ETH Guide
Using [go-ethereum](<https://geth.ethereum.org/docs>) for develop eth.<br>
[golang api example](<https://goethereumbook.org/zh/>)

Build sol to go file
```sh
docker pull ethereum/solc:0.8.4
docker run --rm -v $(pwd)/sol/erc20:/root ethereum/solc:0.4.24 --abi --bin /root/erc20.sol -o /root/build

docker pull ethereum/client-go:alltools-latest
docker run --rm -v $(pwd)/sol/erc20:/root ethereum/client-go:alltools-latest abigen --bin=/root/build/ERC20.bin --abi=/root/build/ERC20.abi --pkg=erc20 --out=/root/ERC20.go
```

### Polygon Guide
Because Polygon is EVM serise.<br>
So you can reference ETH Guide.<br>


### Tron Guide
Add [tron proto](<https://github.com/tronprotocol/protocol/tree/master>) to submodules<br>
Install protoc-gen-grpc-gateway and copy google proto from golang folder.<br>
[Tron RPC List](<https://developers.tron.network/v3.7/docs/official-public-node>)<br>

```sh
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```
