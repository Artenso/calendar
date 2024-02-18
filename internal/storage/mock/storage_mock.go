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

// MockEventsStorage is a mock of EventsStorage interface.
type MockEventsStorage struct {
	ctrl     *gomock.Controller
	recorder *MockEventsStorageMockRecorder
}

// MockEventsStorageMockRecorder is the mock recorder for MockEventsStorage.
type MockEventsStorageMockRecorder struct {
	mock *MockEventsStorage
}

// NewMockEventsStorage creates a new mock instance.
func NewMockEventsStorage(ctrl *gomock.Controller) *MockEventsStorage {
	mock := &MockEventsStorage{ctrl: ctrl}
	mock.recorder = &MockEventsStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventsStorage) EXPECT() *MockEventsStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockEventsStorage) Add(ctx context.Context, info model.EventInfo) (*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, info)
	ret0, _ := ret[0].(*model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockEventsStorageMockRecorder) Add(ctx, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockEventsStorage)(nil).Add), ctx, info)
}

// Edit mocks base method.
func (m *MockEventsStorage) Edit(ctx context.Context, eventID int64, info model.EventInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", ctx, eventID, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit.
func (mr *MockEventsStorageMockRecorder) Edit(ctx, eventID, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockEventsStorage)(nil).Edit), ctx, eventID, info)
}

// GetAllEvents mocks base method.
func (m *MockEventsStorage) GetAllEvents(ctx context.Context) []model.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEvents", ctx)
	ret0, _ := ret[0].([]model.Event)
	return ret0
}

// GetAllEvents indicates an expected call of GetAllEvents.
func (mr *MockEventsStorageMockRecorder) GetAllEvents(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvents", reflect.TypeOf((*MockEventsStorage)(nil).GetAllEvents), ctx)
}

// GetByID mocks base method.
func (m *MockEventsStorage) GetByID(ctx context.Context, eventID int64) (*model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, eventID)
	ret0, _ := ret[0].(model.Event)
	ret1, _ := ret[1].(error)
	return &ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockEventsStorageMockRecorder) GetByID(ctx, eventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockEventsStorage)(nil).GetByID), ctx, eventID)
}

// GetFromToEvents mocks base method.
func (m *MockEventsStorage) GetFromToEvents(ctx context.Context, from, to time.Time) []model.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromToEvents", ctx, from, to)
	ret0, _ := ret[0].([]model.Event)
	return ret0
}

// GetFromToEvents indicates an expected call of GetFromToEvents.
func (mr *MockEventsStorageMockRecorder) GetFromToEvents(ctx, from, to any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromToEvents", reflect.TypeOf((*MockEventsStorage)(nil).GetFromToEvents), ctx, from, to)
}

// Remove mocks base method.
func (m *MockEventsStorage) Remove(ctx context.Context, eventID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, eventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockEventsStorageMockRecorder) Remove(ctx, eventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockEventsStorage)(nil).Remove), ctx, eventID)
}