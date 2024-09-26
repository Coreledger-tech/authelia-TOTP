// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-crypt/crypt/algorithm (interfaces: Hash)
//
// Generated by this command:
//
//	mockgen -package authentication -destination file_user_provider_hash_mock_test.go -mock_names Hash=MockHash github.com/go-crypt/crypt/algorithm Hash
//

// Package authentication is a generated GoMock package.
package authentication

import (
	reflect "reflect"

	algorithm "github.com/go-crypt/crypt/algorithm"
	gomock "go.uber.org/mock/gomock"
)

// MockHash is a mock of Hash interface.
type MockHash struct {
	ctrl     *gomock.Controller
	recorder *MockHashMockRecorder
}

// MockHashMockRecorder is the mock recorder for MockHash.
type MockHashMockRecorder struct {
	mock *MockHash
}

// NewMockHash creates a new mock instance.
func NewMockHash(ctrl *gomock.Controller) *MockHash {
	mock := &MockHash{ctrl: ctrl}
	mock.recorder = &MockHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHash) EXPECT() *MockHashMockRecorder {
	return m.recorder
}

// Hash mocks base method.
func (m *MockHash) Hash(arg0 string) (algorithm.Digest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", arg0)
	ret0, _ := ret[0].(algorithm.Digest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash.
func (mr *MockHashMockRecorder) Hash(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockHash)(nil).Hash), arg0)
}

// HashWithSalt mocks base method.
func (m *MockHash) HashWithSalt(arg0 string, arg1 []byte) (algorithm.Digest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashWithSalt", arg0, arg1)
	ret0, _ := ret[0].(algorithm.Digest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashWithSalt indicates an expected call of HashWithSalt.
func (mr *MockHashMockRecorder) HashWithSalt(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashWithSalt", reflect.TypeOf((*MockHash)(nil).HashWithSalt), arg0, arg1)
}

// MustHash mocks base method.
func (m *MockHash) MustHash(arg0 string) algorithm.Digest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustHash", arg0)
	ret0, _ := ret[0].(algorithm.Digest)
	return ret0
}

// MustHash indicates an expected call of MustHash.
func (mr *MockHashMockRecorder) MustHash(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustHash", reflect.TypeOf((*MockHash)(nil).MustHash), arg0)
}

// Validate mocks base method.
func (m *MockHash) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockHashMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockHash)(nil).Validate))
}
