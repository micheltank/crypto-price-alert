package domain

var (
	SupportedCoins = []string{"BTC", "ETH", "BNB"}
)

type Alert struct {
	id        int64
	email     string
	price     float64
	coin      string
	direction Direction
}

func NewAlert(id int64, email string, price float64, coin string, direction Direction) Alert {
	return Alert{
		id:        id,
		email:     email,
		price:     price,
		direction: direction,
		coin:      coin,
	}
}

type Alerts []Alert

func (a *Alert) GetId() int64 {
	return a.id
}

func (a *Alert) GetEmail() string {
	return a.email
}

func (a *Alert) GetPrice() float64 {
	return a.price
}

func (a *Alert) SetPrice(price float64) {
	a.price = price
}

func (a *Alert) SetCoin(coin string) {
	a.coin = coin
}

func (a *Alert) GetCoin() string {
	return a.coin
}

func (a *Alert) GetDirection() Direction {
	return a.direction
}
