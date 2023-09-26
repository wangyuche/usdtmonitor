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
