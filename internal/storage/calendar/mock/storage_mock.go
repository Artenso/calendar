// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go
//
// Generated by this command:
//
//	mockgen -source=storage.go -destination=mock/storage_mock.go
//

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/Artenso/calendar/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockIEventsStorage is a mock of IEventsStorage interface.
type MockIEventsStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIEventsStorageMockRecorder
}

// MockIEventsStorageMockRecorder is the mock recorder for MockIEventsStorage.
type MockIEventsStorageMockRecorder struct {
	mock *MockIEventsStorage
}

// NewMockIEventsStorage creates a new mock instance.
func NewMockIEventsStorage(ctrl *gomock.Controller) *MockIEventsStorage {
	mock := &MockIEventsStorage{ctrl: ctrl}
	mock.recorder = &MockIEventsStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEventsStorage) EXPECT() *MockIEventsStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIEventsStorage) Add(ctx context.Context, info *model.EventInfo) (*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, info)
	ret0, _ := ret[0].(*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIEventsStorageMockRecorder) Add(ctx, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIEventsStorage)(nil).Add), ctx, info)
}

// Edit mocks base method.
func (m *MockIEventsStorage) Edit(ctx context.Context, eventID int64, info *model.UpdateEventInfo) (*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", ctx, eventID, info)
	ret0, _ := ret[0].(*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Edit indicates an expected call of Edit.
func (mr *MockIEventsStorageMockRecorder) Edit(ctx, eventID, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockIEventsStorage)(nil).Edit), ctx, eventID, info)
}

// GetAllEvents mocks base method.
func (m *MockIEventsStorage) GetAllEvents(ctx context.Context) ([]*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEvents", ctx)
	ret0, _ := ret[0].([]*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEvents indicates an expected call of GetAllEvents.
func (mr *MockIEventsStorageMockRecorder) GetAllEvents(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvents", reflect.TypeOf((*MockIEventsStorage)(nil).GetAllEvents), ctx)
}

// GetByID mocks base method.
func (m *MockIEventsStorage) GetByID(ctx context.Context, eventID int64) (*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, eventID)
	ret0, _ := ret[0].(*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIEventsStorageMockRecorder) GetByID(ctx, eventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIEventsStorage)(nil).GetByID), ctx, eventID)
}

// GetFromToEvents mocks base method.
func (m *MockIEventsStorage) GetFromToEvents(ctx context.Context, from, to time.Time) ([]*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromToEvents", ctx, from, to)
	ret0, _ := ret[0].([]*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromToEvents indicates an expected call of GetFromToEvents.
func (mr *MockIEventsStorageMockRecorder) GetFromToEvents(ctx, from, to any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromToEvents", reflect.TypeOf((*MockIEventsStorage)(nil).GetFromToEvents), ctx, from, to)
}

// Remove mocks base method.
func (m *MockIEventsStorage) Remove(ctx context.Context, eventID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, eventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockIEventsStorageMockRecorder) Remove(ctx, eventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockIEventsStorage)(nil).Remove), ctx, eventID)
}