package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/juniorrodes/exchange/components"
	"github.com/juniorrodes/exchange/pkg/exchange"
)

type ConverterController struct {
    logger     *log.Logger
    exchangeService *exchange.Service
}

func NewConverterController(logger *log.Logger, service *exchange.Service) *ConverterController {
    return &ConverterController{
        logger: logger,
        exchangeService: service,
    }
}

func (c *ConverterController) Convert(w http.ResponseWriter, r *http.Request) {
    fromCurrency := r.FormValue("from")
    toCurrency := r.FormValue("to")
    value, err := strconv.ParseFloat(r.FormValue("value"), 32)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("field value is malformed"))

    }
    
    convertedValue, err := c.exchangeService.GetConvertion(toCurrency, fromCurrency, value)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("<div>%s %f, is worth %s %f</div>", fromCurrency, value, toCurrency, convertedValue)))
}

func ConvertPage(w http.ResponseWriter, r *http.Request) {
    components.ConvertPage(exchange.SupportedCurrencies).Render(r.Context(), w)
}
