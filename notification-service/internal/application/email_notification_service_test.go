package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/micheltank/crypto-price-alert/notification-service/internal/application"
	mock_application "github.com/micheltank/crypto-price-alert/notification-service/internal/application/mock"
	"github.com/micheltank/crypto-price-alert/notification-service/internal/domain"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAlertService(t *testing.T) {
	ctrl := gomock.NewController(t)
	emailService := mock_application.NewMockEmailService(ctrl)

	t.Run("Send email", func(t *testing.T) {
		g := NewGomegaWithT(t)
		email := "john@gmail.com"
		fields := domain.NewEmailNotificationFields(1000, 2000, "USD")
		emailNotification := domain.NewEmailNotification("1", email, domain.TemplateAbove, fields)
		subject, body := emailNotification.BuildContent()

		emailService.EXPECT().
			Send(email, subject, body).
			DoAndReturn(func(to, subject, body string) error {
				return nil
			}).
			Times(1)

		service := application.NewEmailNotificationService(emailService)
		err := service.Send(emailNotification)
		g.Expect(err).Should(
			Not(HaveOccurred()))
	})
}
