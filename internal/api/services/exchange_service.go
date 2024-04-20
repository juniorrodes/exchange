package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type rates map[string]float64

type exchangeData struct {
    Rates rates `json:"rates"`
    NextUpdate string `json:"time_next_update_utc"`
    expiration time.Time
}

type ExchangeService struct {
    client *http.Client
    logger *log.Logger
    cache exchangeData
}

func NewExchangeService(logger *log.Logger, client *http.Client) *ExchangeService {
    return &ExchangeService{
        logger: logger,
        client: client,
    }
}


func (s *ExchangeService) GetConvertion(
    toCurrency, fromCurrency string, 
    value float64) (float64, error) {
    if !s.cache.expiration.IsZero() && time.Now().Before(s.cache.expiration) {
        rate, ok := s.cache.Rates[toCurrency]
        if ok {
            s.logger.Print("Using cached rates")
            return value * rate, nil
        }

        s.logger.Printf("Did not found rate for %s", fromCurrency)

        return 0.0, errors.New("rate does not exist") 
    }

    s.logger.Print("Calling exchange API...")
    exchangeURL := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", fromCurrency)

    req, err := http.NewRequest(http.MethodGet, exchangeURL, nil)
    if err != nil {
        return 0.0, err
    }
    resp, err := s.client.Do(req)
    if err != nil {
        return 0.0, err
    }
    defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0.0, err
    }

    if err = json.Unmarshal(bodyBytes, &s.cache); err != nil {
        return 0.0, err
    }

    s.cache.expiration, err = time.Parse(time.RFC1123Z, s.cache.NextUpdate)
    if err != nil {
        s.cache.expiration = time.Now()
        return 0.0, err
    }
    s.logger.Printf("Got rates %v", s.cache.Rates)

    rate := s.cache.Rates[toCurrency]

    return value * rate, nil
}
