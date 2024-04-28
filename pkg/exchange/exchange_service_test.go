package exchange

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetConvertion(t *testing.T) {
    t.Run("Should not call exchange when cache is populated", func(t *testing.T) {
        service := Service {
            logger: log.New(os.Stdout, "TEST:", log.Ltime | log.Lshortfile),
            cache: cacheMap{
                "USD": &exchangeData{
                    Rates: rates{"BRL": 5.2},
                    expiration: time.Now().Add(1 * time.Second),
                },
            },
            client: nil,
        }

        ratesGetter = func(service *Service, body *exchangeData, fromCurrency string) error {
            t.Fatalf("rateGetter should have not beem called")
            return nil
        }

        value, err := service.GetConvertion("BRL", "USD", 10)

        expected := 10 * 5.2
        assert.Nil(t, err)
        assert.Equalf(t, expected, value, "expected: %f, found %f", expected, value)
    })

    t.Run("Should call exchange when cache is empty", func(t *testing.T) {
        service := Service {
            logger: log.New(os.Stdout, "TEST:", log.Ltime | log.Lshortfile),
            cache: cacheMap{},
            client: nil,
        }

        calledTimes := 0
        ratesGetter = func(_ *Service, body *exchangeData, _ string) error {
            calledTimes++
            body.Rates = rates{"BRL": 5.2}
            body.NextUpdate = time.Now().Add(1 * time.Hour).Format(time.RFC1123Z) 
            
            return nil
        }
         
        value, err := service.GetConvertion("BRL", "USD", 10)

        expected := 10 * 5.2
        assert.Nil(t, err)
        assert.Equalf(t, 
            1, calledTimes,
            "expected ratesGetter to be called once, but it was called: %d", calledTimes,
        )
        assert.Equalf(t, expected, value, "expected: %f, found %f", expected, value)
    })
}
