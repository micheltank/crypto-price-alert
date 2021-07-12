package repository

import (
	"database/sql"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/domain"
)

type AlertPostgreSql struct {
	db *sql.DB
}

func NewAlertPostgreSql(db *sql.DB) *AlertPostgreSql {
	return &AlertPostgreSql{
		db: db,
	}
}

func (r *AlertPostgreSql) Create(alert domain.Alert) (domain.Alert, error) {
	var lastInsertId int64

	query := `INSERT INTO alert (
                email,
				price,
			    coin,
				direction) 
			  VALUES($1,$2,$3,$4)
			  RETURNING id`
	err := r.db.QueryRow(query,
		alert.GetEmail(),
		alert.GetPrice(),
		alert.GetCoin(),
		alert.GetDirection()).Scan(&lastInsertId)
	if err != nil {
		return domain.Alert{}, err
	}
	if err != nil {
		return domain.Alert{}, err
	}
	alert = domain.NewAlertFromRepository(lastInsertId, alert.GetEmail(), alert.GetPrice(), alert.GetCoin(), alert.GetDirection())
	return alert, nil
}

func (r *AlertPostgreSql) List(email string) (domain.Alerts, error) {
	rows, err := r.db.Query(`SELECT
							id,
							email,
							price,
							direction
						FROM alert
						WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var alerts domain.Alerts
	for rows.Next() {
		var alert Alert
		err = rows.Scan(&alert.Id,
			&alert.Email,
			&alert.Price,
			&alert.Coin,
			&alert.Direction)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, alert.ToDomain())
	}
	return alerts, nil
}

func (r *AlertPostgreSql) GetAlertsAtPrice(coin string, price float64) (domain.Alerts, error) {
	rows, err := r.db.Query(`SELECT
							id,
							email,
							price,
       						coin,
							direction
						FROM alert
						WHERE coin=$1 
						  AND (($2 > price AND direction='ABOVE')
						    OR ($2 < price AND direction='BELOW'))`, coin, price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var alerts domain.Alerts
	for rows.Next() {
		var alert Alert
		err = rows.Scan(&alert.Id,
			&alert.Email,
			&alert.Price,
			&alert.Coin,
			&alert.Direction)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, alert.ToDomain())
	}
	return alerts, nil
}
