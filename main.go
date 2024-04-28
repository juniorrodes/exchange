package main

import (
	"log"
	"net/http"
	"os"

	controller "github.com/juniorrodes/exchange/pkg/controllers"
	"github.com/juniorrodes/exchange/pkg/exchange"
	"github.com/juniorrodes/exchange/pkg/server"
)

func main() {
    logger := log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)

	router := server.NewRouter(logger)
    client := http.DefaultClient

    exchngeService := exchange.NewExchangeService(logger, client)

    converteController := controller.NewConverterController(logger, exchngeService)

    router.Get("/converter", controller.ConvertPage)

	router.Post("/convert", converteController.Convert)

	router.Serve(":8080")
}
