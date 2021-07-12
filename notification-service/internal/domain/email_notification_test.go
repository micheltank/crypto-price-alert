package domain_test

import (
	"github.com/micheltank/crypto-price-alert/notification-service/internal/domain"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAlertService(t *testing.T) {
	t.Run("Test email content template above", func(t *testing.T) {
		g := NewGomegaWithT(t)

		fields := domain.NewEmailNotificationFields(1000, 2000, "USD")
		emailNotification := domain.NewEmailNotification("1", "john@gmail.com", domain.TemplateAbove, fields)
		subject, body := emailNotification.BuildContent()

		g.Expect(subject).Should(
			Equal("The price went up!"))
		g.Expect(body).Should(
			Equal("Your alert for price above USD 1,000.00\n\t\t\tThe price is now USD 2,000.00"))
	})
	t.Run("Test email content template below", func(t *testing.T) {
		g := NewGomegaWithT(t)

		fields := domain.NewEmailNotificationFields(2000, 1000, "USD")
		emailNotification := domain.NewEmailNotification("1", "john@gmail.com", domain.TemplateBelow, fields)
		subject, body := emailNotification.BuildContent()

		g.Expect(subject).Should(
			Equal("The price has gone down!"))
		g.Expect(body).Should(
			Equal("Your alert for price below USD 2,000.00\n\t\t\tThe price is now USD 1,000.00"))
	})
}
