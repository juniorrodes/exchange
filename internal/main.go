package main

import (
	"log"
	"os"

	"github.com/juniorrodes/exchange/internal/app/controller"
	"github.com/juniorrodes/exchange/internal/infrastructure/server"
)

func main() {
    logger := log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)

	router := server.NewRouter(logger)
    converteController := controller.NewConverterController(logger)

    router.Get("/convert", controller.ConvertPage)

	router.Get("/convert/{from}/{to}/{value}", converteController.Convert)

	router.Serve(":8080")
}
