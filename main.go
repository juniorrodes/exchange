package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	controller "github.com/juniorrodes/exchange/pkg/controllers"
	"github.com/juniorrodes/exchange/pkg/exchange"
	"github.com/juniorrodes/exchange/pkg/server"
)

var logger = log.New(os.Stdout, "[exchange-server]", log.Ldate | log.Ltime | log.Lshortfile)
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
    filePath := r.URL.Path[len("/static/"):]
    
    logger.Println(filePath)

    fullPath := filepath.Join(".", "static", filePath)
    http.ServeFile(w, r, fullPath)
}

func main() {

	router := server.NewRouter(logger)
    client := http.DefaultClient

    exchngeService := exchange.NewExchangeService(logger, client)

    converteController := controller.NewConverterController(logger, exchngeService)

    router.Get("/converter", controller.ConvertPage)
    router.Get("/static/", serveStaticFiles)

	router.Post("/convert", converteController.Convert)

	router.Serve(":8080")
}
