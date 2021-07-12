package infra

import (
	"context"
	"github.com/micheltank/crypto-price-alert/price-alert-job/internal/domain"
	pb "github.com/micheltank/crypto-price-alert/price-alert-service/cmd/grpc/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type AlertApi struct {
	host   string
	client pb.AlertHandlerClient
}

func NewAlertApi(host string) (AlertApi, error) {
	connection, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return AlertApi{}, errors.Wrap(err, "failed to connect on alert grpc api")
	}
	client := pb.NewAlertHandlerClient(connection)
	return AlertApi{
		client: client,
	}, nil
}

func (a AlertApi) GetAlerts(coin string, price float64) (domain.Alerts, error) {
	req := &pb.GetAlertsParams{
		Coin:  coin,
		Price: price,
	}
	res, err := a.client.GetAlerts(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get alerts from alert api")
	}
	var alerts domain.Alerts
	for _, alert := range res.GetAlert() {
		alerts = append(alerts, ConvertAlertToDomain(*alert))
	}
	return alerts, nil
}
