// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lucas-clemente/quic-go (interfaces: StreamFrameSource)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "gx/ipfs/QmU44KWVkSHno7sNDTeUcL4FBgxgoidkFuTUyTXWJPXXFJ/quic-go/internal/protocol"
	wire "gx/ipfs/QmU44KWVkSHno7sNDTeUcL4FBgxgoidkFuTUyTXWJPXXFJ/quic-go/internal/wire"
)

// MockStreamFrameSource is a mock of StreamFrameSource interface
type MockStreamFrameSource struct {
	ctrl     *gomock.Controller
	recorder *MockStreamFrameSourceMockRecorder
}

// MockStreamFrameSourceMockRecorder is the mock recorder for MockStreamFrameSource
type MockStreamFrameSourceMockRecorder struct {
	mock *MockStreamFrameSource
}

// NewMockStreamFrameSource creates a new mock instance
func NewMockStreamFrameSource(ctrl *gomock.Controller) *MockStreamFrameSource {
	mock := &MockStreamFrameSource{ctrl: ctrl}
	mock.recorder = &MockStreamFrameSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamFrameSource) EXPECT() *MockStreamFrameSourceMockRecorder {
	return m.recorder
}

// HasCryptoStreamData mocks base method
func (m *MockStreamFrameSource) HasCryptoStreamData() bool {
	ret := m.ctrl.Call(m, "HasCryptoStreamData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasCryptoStreamData indicates an expected call of HasCryptoStreamData
func (mr *MockStreamFrameSourceMockRecorder) HasCryptoStreamData() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasCryptoStreamData", reflect.TypeOf((*MockStreamFrameSource)(nil).HasCryptoStreamData))
}

// PopCryptoStreamFrame mocks base method
func (m *MockStreamFrameSource) PopCryptoStreamFrame(arg0 protocol.ByteCount) *wire.StreamFrame {
	ret := m.ctrl.Call(m, "PopCryptoStreamFrame", arg0)
	ret0, _ := ret[0].(*wire.StreamFrame)
	return ret0
}

// PopCryptoStreamFrame indicates an expected call of PopCryptoStreamFrame
func (mr *MockStreamFrameSourceMockRecorder) PopCryptoStreamFrame(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopCryptoStreamFrame", reflect.TypeOf((*MockStreamFrameSource)(nil).PopCryptoStreamFrame), arg0)
}

// PopStreamFrames mocks base method
func (m *MockStreamFrameSource) PopStreamFrames(arg0 protocol.ByteCount) []*wire.StreamFrame {
	ret := m.ctrl.Call(m, "PopStreamFrames", arg0)
	ret0, _ := ret[0].([]*wire.StreamFrame)
	return ret0
}

// PopStreamFrames indicates an expected call of PopStreamFrames
func (mr *MockStreamFrameSourceMockRecorder) PopStreamFrames(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopStreamFrames", reflect.TypeOf((*MockStreamFrameSource)(nil).PopStreamFrames), arg0)
}
