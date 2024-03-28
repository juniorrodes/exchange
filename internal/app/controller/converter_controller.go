package controller

import "github.com/juniorrodes/exchange/internal/app/model"

type ConverterController struct {
    currencyMap model.Currencies
}

func NewConverterController() *ConverterController {
    return &ConverterController{
        currencyMap: make(model.Currencies),
    }
}

func (c *ConverterController) Convert(w http.ResponseWriter, r *http.Request) {
    
}
