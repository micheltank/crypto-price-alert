package domain

type PriceApi interface {
	GetPrice(coin string) (Price, error)
}
