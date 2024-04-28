package exchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Service struct {
    client *http.Client
    logger *log.Logger
    cache cacheMap
}

var ratesGetter = getRates

func getRates(service *Service, body *exchangeData, fromCurrency string) error {
    service.logger.Print("Calling exchange API...")
    exchangeURL := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", fromCurrency)

    req, err := http.NewRequest(http.MethodGet, exchangeURL, nil)
    if err != nil {
        return err
    }
    resp, err := service.client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    if err = json.Unmarshal(bodyBytes, body); err != nil {
        return err
    }

    return nil
}

func NewExchangeService(logger *log.Logger, client *http.Client) *Service {
    return &Service{
        logger: logger,
        client: client,
        cache: make(cacheMap),
    }
}


func (s *Service) GetConvertion(
    toCurrency, fromCurrency string, 
    value float64) (float64, error) {
    if s.cache[fromCurrency] != nil && (!s.cache[fromCurrency].expiration.IsZero() && time.Now().Before(s.cache[fromCurrency].expiration)) {
        rate, ok := s.cache[fromCurrency].Rates[toCurrency]
        if ok {
            s.logger.Print("Using cached rates")
            return value * rate, nil
        }

        s.logger.Printf("Did not found rate for %s", fromCurrency)

        return 0.0, errors.New("rate does not exist") 
    }
    
    body := &exchangeData{}
    err := ratesGetter(s, body, fromCurrency)
    if err != nil {
        return 0.0, err
    }

    s.cache[fromCurrency] = body
    s.logger.Printf("NextUpdate: %s", body.NextUpdate)
    s.cache[fromCurrency].expiration, err = time.Parse(time.RFC1123Z, body.NextUpdate)
    if err != nil {
        s.cache[fromCurrency].expiration = time.Now()
        return 0.0, err
    }

    cachedRate := s.cache[fromCurrency] 

    s.logger.Printf("Got rates %v", cachedRate.Rates)
    s.logger.Printf("Cache expiration is: %d %v, %d:%d", 
        cachedRate.expiration.Day(),
        cachedRate.expiration.Month(),
        cachedRate.expiration.Hour(),
        cachedRate.expiration.Minute())

    rate := cachedRate.Rates[toCurrency]

    return value * rate, nil
}
