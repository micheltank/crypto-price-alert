package infra

import (
	"encoding/json"
	"fmt"
	"github.com/micheltank/crypto-price-alert/price-alert-job/internal/domain"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type CryptoComparePriceApi struct {
	host   string
	apiKey string
	client *http.Client
}

func NewCryptoComparePriceApi(host string, apiKey string) CryptoComparePriceApi {
	return CryptoComparePriceApi{
		host:   host,
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c CryptoComparePriceApi) GetPrice(coin string) (domain.Price, error) {
	url := fmt.Sprintf("%s/data/price?fsym=%s&tsyms=USD&ApiKey=%s", c.host, coin, c.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return domain.Price{}, errors.Wrap(err, "failed to get price from crypto compare api")
	}
	if resp.StatusCode != http.StatusOK {
		return domain.Price{}, errors.New(fmt.Sprintf("crypto compare api returned %d status code error", resp.StatusCode))
	}
	var price Price
	err = json.NewDecoder(resp.Body).Decode(&price)
	if err != nil {
		return domain.Price{}, errors.Wrap(err, "failed to decode price from crypto compare api")
	}
	return domain.NewPrice(price.USD), nil
}
