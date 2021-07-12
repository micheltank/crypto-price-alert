package infra

import (
	"github.com/micheltank/crypto-price-alert/price-alert-job/internal/domain"
	pb "github.com/micheltank/crypto-price-alert/price-alert-service/cmd/grpc/pb"
)

func ConvertAlertToDomain(alert pb.Alert) domain.Alert {
	return domain.NewAlert(
		alert.GetId(),
		alert.GetEmail(),
		alert.GetPrice(),
		alert.GetCoin(),
		ConvertAlertDirection(alert.GetDirection()),
	)
}

func ConvertAlertDirection(direction pb.Alert_Direction) domain.Direction {
	switch direction {
	case pb.Alert_ABOVE:
		return domain.DirectionAbove
	case pb.Alert_BELOW:
		return domain.DirectionBelow
	}
	return ""
}
