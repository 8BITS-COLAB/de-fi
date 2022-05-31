package main

import (
	"os"

	"github.com/ElioenaiFerrari/de-fi/api"
	"github.com/ElioenaiFerrari/de-fi/controllers"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	godotenv.Load()
}

func main() {
	client, err := ethclient.Dial(os.Getenv("ETHEREUM_URL"))

	if err != nil {
		panic(err)
	}

	auth, err := api.Auth(client)

	if err != nil {
		panic(err)
	}

	addr, _, _, _ := api.DeployVaultContract(auth, client)
	conn, err := api.NewVaultContract(addr, client)

	if err != nil {
		panic(err)
	}

	server := echo.New()
	router := server.Group("/api")

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.Use(middleware.Secure())
	server.Use(middleware.Gzip())
	server.Use(middleware.BodyLimit("1M"))

	vaultController := controllers.NewVaultController(client, conn, auth)

	router.POST("/vaults", vaultController.Create)
	router.PUT("/vaults/:from", vaultController.Update)
	router.GET("/vaults/:from", vaultController.Get)

	server.Logger.Fatal(server.Start(":3000"))

}
