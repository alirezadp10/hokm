// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/database/sqlite.go

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// AutoMigrate mocks base method.
func (m *MockDB) AutoMigrate(dst ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range dst {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AutoMigrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoMigrate indicates an expected call of AutoMigrate.
func (mr *MockDBMockRecorder) AutoMigrate(dst ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockDB)(nil).AutoMigrate), dst...)
}

// Begin mocks base method.
func (m *MockDB) Begin(opts ...*sql.TxOptions) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Begin", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockDBMockRecorder) Begin(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockDB)(nil).Begin), opts...)
}

// Commit mocks base method.
func (m *MockDB) Commit() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockDBMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockDB)(nil).Commit))
}

// Count mocks base method.
func (m *MockDB) Count(count *int64) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", count)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockDBMockRecorder) Count(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDB)(nil).Count), count)
}

// Create mocks base method.
func (m *MockDB) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDBMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDB)(nil).Create), value)
}

// CreateInBatches mocks base method.
func (m *MockDB) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInBatches", value, batchSize)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// CreateInBatches indicates an expected call of CreateInBatches.
func (mr *MockDBMockRecorder) CreateInBatches(value, batchSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInBatches", reflect.TypeOf((*MockDB)(nil).CreateInBatches), value, batchSize)
}

// Delete mocks base method.
func (m *MockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{value}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDBMockRecorder) Delete(value interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{value}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDB)(nil).Delete), varargs...)
}

// Exec mocks base method.
func (m *MockDB) Exec(sql string, values ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{sql}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockDBMockRecorder) Exec(sql interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDB)(nil).Exec), varargs...)
}

// Find mocks base method.
func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockDBMockRecorder) Find(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDB)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockDBMockRecorder) First(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockDB)(nil).First), varargs...)
}

// FirstOrCreate mocks base method.
func (m *MockDB) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FirstOrCreate", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// FirstOrCreate indicates an expected call of FirstOrCreate.
func (mr *MockDBMockRecorder) FirstOrCreate(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrCreate", reflect.TypeOf((*MockDB)(nil).FirstOrCreate), varargs...)
}

// FirstOrInit mocks base method.
func (m *MockDB) FirstOrInit(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FirstOrInit", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// FirstOrInit indicates an expected call of FirstOrInit.
func (mr *MockDBMockRecorder) FirstOrInit(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrInit", reflect.TypeOf((*MockDB)(nil).FirstOrInit), varargs...)
}

// Last mocks base method.
func (m *MockDB) Last(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Last", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *MockDBMockRecorder) Last(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockDB)(nil).Last), varargs...)
}

// Pluck mocks base method.
func (m *MockDB) Pluck(column string, dest interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pluck", column, dest)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Pluck indicates an expected call of Pluck.
func (mr *MockDBMockRecorder) Pluck(column, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pluck", reflect.TypeOf((*MockDB)(nil).Pluck), column, dest)
}

// Rollback mocks base method.
func (m *MockDB) Rollback() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockDBMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockDB)(nil).Rollback))
}

// RollbackTo mocks base method.
func (m *MockDB) RollbackTo(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTo", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// RollbackTo indicates an expected call of RollbackTo.
func (mr *MockDBMockRecorder) RollbackTo(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTo", reflect.TypeOf((*MockDB)(nil).RollbackTo), name)
}

// Row mocks base method.
func (m *MockDB) Row() *sql.Row {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Row")
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// Row indicates an expected call of Row.
func (mr *MockDBMockRecorder) Row() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Row", reflect.TypeOf((*MockDB)(nil).Row))
}

// Rows mocks base method.
func (m *MockDB) Rows() (*sql.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rows")
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rows indicates an expected call of Rows.
func (mr *MockDBMockRecorder) Rows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rows", reflect.TypeOf((*MockDB)(nil).Rows))
}

// Save mocks base method.
func (m *MockDB) Save(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockDBMockRecorder) Save(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDB)(nil).Save), value)
}

// SavePoint mocks base method.
func (m *MockDB) SavePoint(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePoint", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// SavePoint indicates an expected call of SavePoint.
func (mr *MockDBMockRecorder) SavePoint(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePoint", reflect.TypeOf((*MockDB)(nil).SavePoint), name)
}

// Scan mocks base method.
func (m *MockDB) Scan(dest interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", dest)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockDBMockRecorder) Scan(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockDB)(nil).Scan), dest)
}

// ScanRows mocks base method.
func (m *MockDB) ScanRows(rows *sql.Rows, dest interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanRows", rows, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanRows indicates an expected call of ScanRows.
func (mr *MockDBMockRecorder) ScanRows(rows, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanRows", reflect.TypeOf((*MockDB)(nil).ScanRows), rows, dest)
}

// Take mocks base method.
func (m *MockDB) Take(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Take", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Take indicates an expected call of Take.
func (mr *MockDBMockRecorder) Take(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*MockDB)(nil).Take), varargs...)
}

// Update mocks base method.
func (m *MockDB) Update(column string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", column, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDBMockRecorder) Update(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDB)(nil).Update), column, value)
}

// UpdateColumn mocks base method.
func (m *MockDB) UpdateColumn(column string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", column, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// UpdateColumn indicates an expected call of UpdateColumn.
func (mr *MockDBMockRecorder) UpdateColumn(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*MockDB)(nil).UpdateColumn), column, value)
}

// UpdateColumns mocks base method.
func (m *MockDB) UpdateColumns(values interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumns", values)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// UpdateColumns indicates an expected call of UpdateColumns.
func (mr *MockDBMockRecorder) UpdateColumns(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumns", reflect.TypeOf((*MockDB)(nil).UpdateColumns), values)
}

// Updates mocks base method.
func (m *MockDB) Updates(values interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", values)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockDBMockRecorder) Updates(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockDB)(nil).Updates), values)
}

// Where mocks base method.
func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockDBMockRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockDB)(nil).Where), varargs...)
}
