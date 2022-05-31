package controllers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/ElioenaiFerrari/de-fi/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
)

type VaultController struct {
	client *ethclient.Client
	conn   *api.VaultContract
	auth   *bind.TransactOpts
}

func NewVaultController(client *ethclient.Client, conn *api.VaultContract, auth *bind.TransactOpts) *VaultController {
	return &VaultController{
		client: client,
		conn:   conn,
		auth:   auth,
	}
}

func (vaultController VaultController) Create(c echo.Context) error {
	type Request struct {
		From   string `json:"from"`
		Amount int64  `json:"amount"`
	}

	var request Request

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	from := common.HexToAddress(request.From)
	amount := big.NewInt(request.Amount)

	vaultController.auth.Nonce.Add(vaultController.auth.Nonce, big.NewInt(1))

	reply, err := vaultController.conn.CreateAccount(vaultController.auth, from, amount)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, reply)

}

func (vaultController VaultController) Update(c echo.Context) error {

	fromStr := c.Param("from")

	if fromStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "from is required")
	}

	type Request struct {
		Type   string `json:"type"`
		To     string `json:"to"`
		Amount int64  `json:"amount"`
	}

	var request Request

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	from := common.HexToAddress(fromStr)
	to := common.HexToAddress(request.To)
	amount := big.NewInt(request.Amount)

	fmt.Println(from)

	switch request.Type {
	case "deposit":
		vaultController.auth.Nonce.Add(vaultController.auth.Nonce, big.NewInt(1))
		reply, err := vaultController.conn.Deposit(vaultController.auth, from, amount)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, reply)

	case "withdraw":
		vaultController.auth.Nonce.Add(vaultController.auth.Nonce, big.NewInt(1))
		reply, err := vaultController.conn.Withdraw(vaultController.auth, from, amount)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, reply)

	case "transfer":
		vaultController.auth.Nonce.Add(vaultController.auth.Nonce, big.NewInt(1))
		reply, err := vaultController.conn.Transfer(vaultController.auth, from, to, amount)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, reply)

	default:
		return c.JSON(http.StatusBadRequest, "invalid type")
	}
}

func (vaultController VaultController) Get(c echo.Context) error {
	fromStr := c.Param("from")

	if fromStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "from is required")
	}

	from := common.HexToAddress(fromStr)

	balance, err := vaultController.conn.BalanceOf(&bind.CallOpts{}, from)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalSupply, err := vaultController.conn.GetTotalSupply(&bind.CallOpts{})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"balance":     balance,
		"totalSupply": totalSupply,
	})
}
