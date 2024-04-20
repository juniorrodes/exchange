package main

import (
	"log"
	"net/http"
	"os"

	"github.com/juniorrodes/exchange/internal/api/controller"
	"github.com/juniorrodes/exchange/internal/api/services"
	"github.com/juniorrodes/exchange/internal/infrastructure/server"
)

func main() {
    logger := log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)

	router := server.NewRouter(logger)
    client := http.DefaultClient

    exchngeService := services.NewExchangeService(logger, client)

    converteController := controller.NewConverterController(logger, exchngeService)

    router.Get("/converter", controller.ConvertPage)

	router.Post("/convert", converteController.Convert)

	router.Serve(":8080")
}
