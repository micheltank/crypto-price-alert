package domain

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidPriceAbove = errors.New("above price cannot be equal or less than zero")
	ErrInvalidPriceBelow = errors.New("below price cannot be equal or less than zero")
	ErrUnsupportedCoin = errors.New(fmt.Sprintf("this coin isn't supported. Available: %s", strings.Join(SupportedCoins, ", ")))
)