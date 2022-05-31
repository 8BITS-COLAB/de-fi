package api

import (
	"context"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Auth(client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("ETHEREUM_PRIVATE_KEY"))

	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), address)

	if err != nil {
		return nil, err
	}

	gasLimit := uint64(3000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		return nil, err

	}

	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	if err != nil {
		return nil, err
	}

	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)

	return auth, nil
}
