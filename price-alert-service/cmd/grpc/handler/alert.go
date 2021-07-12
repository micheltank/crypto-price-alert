package handler

import (
	"context"
	pb "github.com/micheltank/crypto-price-alert/price-alert-service/cmd/grpc/pb"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/application"
	"github.com/pkg/errors"
)

type AlertHandler struct {
	pb.UnimplementedAlertHandlerServer
	service application.IService
}

func NewAlertHandler(service application.IService) *AlertHandler {
	return &AlertHandler{
		service: service,
	}
}

func (h *AlertHandler) GetAlerts(ctx context.Context, in *pb.GetAlertsParams) (*pb.Alerts, error) {
	alerts, err := h.service.GetAlertsAtPrice(in.GetCoin(), in.GetPrice())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get alerts at price")
	}
	var alertsPb []*pb.Alert
	for _, alert := range alerts {
		alertsPb = append(alertsPb, ConvertFromDomain(alert))
	}
	return &pb.Alerts{
		Alert: alertsPb,
	}, nil
}