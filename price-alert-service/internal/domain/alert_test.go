package domain

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestAlert(t *testing.T) {
	t.Run("Regular alert creation", func(t *testing.T) {
		g := NewGomegaWithT(t)

		_, err := NewAlert("john@gmail.com", 100, "BTC", DirectionAbove)
		g.Expect(err).Should(
			Not(HaveOccurred()))
	})
	t.Run("Invalid email", func(t *testing.T) {
		g := NewGomegaWithT(t)

		_, err := NewAlert("invalid.email", 100, "BTC", DirectionAbove)
		g.Expect(err).Should(
			Equal(ErrInvalidEmail))
	})
	t.Run("Invalid price above", func(t *testing.T) {
		g := NewGomegaWithT(t)

		_, err := NewAlert("john@gmail.com", -100, "BTC", DirectionAbove)
		g.Expect(err).Should(
			Equal(ErrInvalidPriceAbove))
	})
	t.Run("Invalid price below", func(t *testing.T) {
		g := NewGomegaWithT(t)

		_, err := NewAlert("john@gmail.com", -100, "BTC", DirectionBelow)
		g.Expect(err).Should(
			Equal(ErrInvalidPriceBelow))
	})
	t.Run("Unsupported coin", func(t *testing.T) {
		g := NewGomegaWithT(t)

		_, err := NewAlert("john@gmail.com", 100, "ASDADASDA", DirectionBelow)
		g.Expect(err).Should(
			Equal(ErrUnsupportedCoin))
	})
}