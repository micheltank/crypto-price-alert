package domain

type CreateAlertCommand struct {
	Email     string
	Price     float64
	Direction Direction
	Coin      string
}
