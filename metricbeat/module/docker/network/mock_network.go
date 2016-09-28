package network

import (
	"github.com/stretchr/testify/mock"
)
// Mock of NetworkCalculator interface
type MockNetworkCalculator struct {
	mock.Mock
}

func (_m *MockNetworkCalculator) getRxBytesPerSecond(newStats *NETRaw, oldStats *NETRaw) float64 {
	ret := _m.Called()
	ret0, _ := ret[0].(float64)
	return ret0
}
func (_m *MockNetworkCalculator) getRxDroppedPerSecond(newStats *NETRaw, oldStats *NETRaw) float64 {
	ret := _m.Called()
	ret0, _ := ret[0].(float64)
	return ret0
}
func (_m *MockNetworkCalculator) getRxErrorsPerSecond(newStats *NETRaw, oldStats *NETRaw) float64 {
	ret := _m.Called()
	ret0, _ := ret[0].(float64)
	return ret0
}
func (_m *MockNetworkCalculator) getRxPacketsPerSecond(newStats *NETRaw, oldStats *NETRaw) float64 {
	ret := _m.Called()
	ret0, _ := ret[0].(float64)
	return ret0
}
func (_m *MockNetworkCalculator) getTxBytesPerSecond(newStats *NETRaw, oldStats *NETRaw) float64 {
	ret := _m.Called()
	ret0, _ := ret[0].(float64)
	return ret0
}
func (_m *MockNetworkCalculator) getTxDroppedPerSecond(newStats *NETRaw, oldStats *NETRaw) float64 {
	ret := _m.Called()
	ret0, _ := ret[0].(float64)
	return ret0
}
func (_m *MockNetworkCalculator) getTxErrorsPerSecond(newStats *NETRaw, oldStats *NETRaw) float64 {
	ret := _m.Called()
	ret0, _ := ret[0].(float64)
	return ret0
}
func (_m *MockNetworkCalculator) getTxPacketsPerSecond(newStats *NETRaw, oldStats *NETRaw) float64 {
	ret := _m.Called()
	ret0, _ := ret[0].(float64)
	return ret0
}
