package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ElioenaiFerrari/de-fi/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	client, err := ethclient.Dial(os.Getenv("ETHEREUM_URL"))

	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("ETHEREUM_PRIVATE_KEY"))

	if err != nil {
		panic(err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), address)

	if err != nil {
		panic(err)
	}

	gasLimit := uint64(3000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		panic(err)
	}

	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	if err != nil {
		panic(err)
	}

	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)

	_, tx, _, _ := api.DeployVaultContract(auth, client)

	fmt.Println(tx.Hash().Hex())

}
