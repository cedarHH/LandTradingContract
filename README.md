### How to run
- Run server:
```shell
docker-compose up --build -d
```
- Remove server:
```shell
docker-compose down --rmi all
```

- To deply contract
```shell
solc --optimize --abi ./smartContract/LandTransaction.sol -o build
solc --optimize --bin ./smartContract/LandTransaction.sol -o build
abigen --abi=./build/LandTransaction.abi --bin=./build/LandTransaction.bin --pkg=contract --out=./api/contract/LandTransaction.go
```