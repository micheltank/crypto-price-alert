package domain

type Price struct {
	USD float64
}

func NewPrice(usdPrice float64) Price {
	return Price{
		USD: usdPrice,
	}
}
