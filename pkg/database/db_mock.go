// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/database/db.go

// Package database is a generated GoMock package.
package database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDatabase) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDatabaseMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDatabase)(nil).Create), value)
}

// Delete mocks base method.
func (m *MockDatabase) Delete(arg0 interface{}, arg1 ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDatabaseMockRecorder) Delete(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatabase)(nil).Delete), varargs...)
}

// Error mocks base method.
func (m *MockDatabase) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockDatabaseMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockDatabase)(nil).Error))
}

// Find mocks base method.
func (m *MockDatabase) Find(arg0 interface{}, arg1 ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockDatabaseMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDatabase)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockDatabase) First(dest interface{}, conds ...interface{}) Database {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(Database)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockDatabaseMockRecorder) First(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockDatabase)(nil).First), varargs...)
}

// Limit mocks base method.
func (m *MockDatabase) Limit(limit int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit", limit)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Limit indicates an expected call of Limit.
func (mr *MockDatabaseMockRecorder) Limit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockDatabase)(nil).Limit), limit)
}

// Model mocks base method.
func (m *MockDatabase) Model(model interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", model)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockDatabaseMockRecorder) Model(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockDatabase)(nil).Model), model)
}

// Offset mocks base method.
func (m *MockDatabase) Offset(offset int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", offset)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Offset indicates an expected call of Offset.
func (mr *MockDatabaseMockRecorder) Offset(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockDatabase)(nil).Offset), offset)
}

// Order mocks base method.
func (m *MockDatabase) Order(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Order indicates an expected call of Order.
func (mr *MockDatabaseMockRecorder) Order(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockDatabase)(nil).Order), value)
}

// Updates mocks base method.
func (m *MockDatabase) Updates(arg0 interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", arg0)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockDatabaseMockRecorder) Updates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockDatabase)(nil).Updates), arg0)
}

// Where mocks base method.
func (m *MockDatabase) Where(query interface{}, args ...interface{}) Database {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(Database)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockDatabaseMockRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockDatabase)(nil).Where), varargs...)
}
