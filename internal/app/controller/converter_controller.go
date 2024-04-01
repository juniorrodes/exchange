package controller

import (
	"log"
	"net/http"

	"github.com/juniorrodes/exchange/internal/app/model"
	"github.com/juniorrodes/exchange/internal/app/views"
)

type ConverterController struct {
    currencyMap model.Currencies
    logger     *log.Logger
}

func NewConverterController(logger *log.Logger) *ConverterController {
    return &ConverterController{
        currencyMap: make(model.Currencies),
        logger: logger,
    }
}

func (c *ConverterController) Convert(w http.ResponseWriter, r *http.Request) {
    fromCurrency := r.PathValue("from")
    toCurrency := r.PathValue("to")
    value := r.PathValue("value")

    c.logger.Printf("Converting %s to %s, value: %s", fromCurrency, toCurrency, value)
}

func ConvertPage(w http.ResponseWriter, r *http.Request) {
    views.ConvertPage(model.CurrenciesCodes).Render(r.Context(), w)
}
