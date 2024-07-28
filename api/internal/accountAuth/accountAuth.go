package accountAuth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type AccountAuth struct {
	ethClient *ethclient.Client
}

func NewAccountAuth(ethClient *ethclient.Client) *AccountAuth {
	return &AccountAuth{
		ethClient: ethClient,
	}
}

// GetAccountAuth function to create accountAuth for any account from its private key
func (l *AccountAuth) GetAccountAuth(privateKeyAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := l.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := l.ethClient.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(6000000)
	gasPrice, err := l.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	auth.GasPrice = gasPrice

	return auth
}
