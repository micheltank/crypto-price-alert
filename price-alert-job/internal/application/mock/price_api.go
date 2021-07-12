// Code generated by MockGen. DO NOT EDIT.
// Source: price_api.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
reflect "reflect"

gomock "github.com/golang/mock/gomock"
domain "github.com/micheltank/crypto-price-alert/price-alert-job/internal/domain"
)

// MockPriceApi is a mock of PriceApi interface.
type MockPriceApi struct {
	ctrl     *gomock.Controller
	recorder *MockPriceApiMockRecorder
}

// MockPriceApiMockRecorder is the mock recorder for MockPriceApi.
type MockPriceApiMockRecorder struct {
	mock *MockPriceApi
}

// NewMockPriceApi creates a new mock instance.
func NewMockPriceApi(ctrl *gomock.Controller) *MockPriceApi {
	mock := &MockPriceApi{ctrl: ctrl}
	mock.recorder = &MockPriceApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceApi) EXPECT() *MockPriceApiMockRecorder {
	return m.recorder
}

// GetPrice mocks base method.
func (m *MockPriceApi) GetPrice(coin string) (domain.Price, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice", coin)
	ret0, _ := ret[0].(domain.Price)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockPriceApiMockRecorder) GetPrice(coin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockPriceApi)(nil).GetPrice), coin)
}
