package handler

import (
	pb "github.com/micheltank/crypto-price-alert/price-alert-service/cmd/grpc/pb"
	"github.com/micheltank/crypto-price-alert/price-alert-service/internal/domain"
)

func ConvertFromDomain(alert domain.Alert) *pb.Alert {
	return &pb.Alert{
		Id:        alert.GetId(),
		Email:     alert.GetEmail(),
		Price:     alert.GetPrice(),
		Coin:      alert.GetCoin(),
		Direction: ConvertDirectionFromDomain(alert.GetDirection()),
	}
}

func ConvertDirectionFromDomain(direction domain.Direction) pb.Alert_Direction {
	switch direction {
	case domain.DirectionAbove:
		return pb.Alert_ABOVE
	case domain.DirectionBelow:
		return pb.Alert_BELOW
	}
	return 0
}
