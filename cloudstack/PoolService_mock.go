//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/PoolService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPoolServiceIface is a mock of PoolServiceIface interface.
type MockPoolServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockPoolServiceIfaceMockRecorder
}

// MockPoolServiceIfaceMockRecorder is the mock recorder for MockPoolServiceIface.
type MockPoolServiceIfaceMockRecorder struct {
	mock *MockPoolServiceIface
}

// NewMockPoolServiceIface creates a new mock instance.
func NewMockPoolServiceIface(ctrl *gomock.Controller) *MockPoolServiceIface {
	mock := &MockPoolServiceIface{ctrl: ctrl}
	mock.recorder = &MockPoolServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoolServiceIface) EXPECT() *MockPoolServiceIfaceMockRecorder {
	return m.recorder
}

// CreateStoragePool mocks base method.
func (m *MockPoolServiceIface) CreateStoragePool(p *CreateStoragePoolParams) (*CreateStoragePoolResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStoragePool", p)
	ret0, _ := ret[0].(*CreateStoragePoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStoragePool indicates an expected call of CreateStoragePool.
func (mr *MockPoolServiceIfaceMockRecorder) CreateStoragePool(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStoragePool", reflect.TypeOf((*MockPoolServiceIface)(nil).CreateStoragePool), p)
}

// DeleteStoragePool mocks base method.
func (m *MockPoolServiceIface) DeleteStoragePool(p *DeleteStoragePoolParams) (*DeleteStoragePoolResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStoragePool", p)
	ret0, _ := ret[0].(*DeleteStoragePoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStoragePool indicates an expected call of DeleteStoragePool.
func (mr *MockPoolServiceIfaceMockRecorder) DeleteStoragePool(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStoragePool", reflect.TypeOf((*MockPoolServiceIface)(nil).DeleteStoragePool), p)
}

// FindStoragePoolsForMigration mocks base method.
func (m *MockPoolServiceIface) FindStoragePoolsForMigration(p *FindStoragePoolsForMigrationParams) (*FindStoragePoolsForMigrationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStoragePoolsForMigration", p)
	ret0, _ := ret[0].(*FindStoragePoolsForMigrationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStoragePoolsForMigration indicates an expected call of FindStoragePoolsForMigration.
func (mr *MockPoolServiceIfaceMockRecorder) FindStoragePoolsForMigration(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStoragePoolsForMigration", reflect.TypeOf((*MockPoolServiceIface)(nil).FindStoragePoolsForMigration), p)
}

// GetStoragePoolByID mocks base method.
func (m *MockPoolServiceIface) GetStoragePoolByID(id string, opts ...OptionFunc) (*StoragePool, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStoragePoolByID", varargs...)
	ret0, _ := ret[0].(*StoragePool)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePoolByID indicates an expected call of GetStoragePoolByID.
func (mr *MockPoolServiceIfaceMockRecorder) GetStoragePoolByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolByID", reflect.TypeOf((*MockPoolServiceIface)(nil).GetStoragePoolByID), varargs...)
}

// GetStoragePoolByName mocks base method.
func (m *MockPoolServiceIface) GetStoragePoolByName(name string, opts ...OptionFunc) (*StoragePool, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStoragePoolByName", varargs...)
	ret0, _ := ret[0].(*StoragePool)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePoolByName indicates an expected call of GetStoragePoolByName.
func (mr *MockPoolServiceIfaceMockRecorder) GetStoragePoolByName(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolByName", reflect.TypeOf((*MockPoolServiceIface)(nil).GetStoragePoolByName), varargs...)
}

// GetStoragePoolID mocks base method.
func (m *MockPoolServiceIface) GetStoragePoolID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStoragePoolID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePoolID indicates an expected call of GetStoragePoolID.
func (mr *MockPoolServiceIfaceMockRecorder) GetStoragePoolID(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolID", reflect.TypeOf((*MockPoolServiceIface)(nil).GetStoragePoolID), varargs...)
}

// ListStoragePools mocks base method.
func (m *MockPoolServiceIface) ListStoragePools(p *ListStoragePoolsParams) (*ListStoragePoolsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStoragePools", p)
	ret0, _ := ret[0].(*ListStoragePoolsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStoragePools indicates an expected call of ListStoragePools.
func (mr *MockPoolServiceIfaceMockRecorder) ListStoragePools(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStoragePools", reflect.TypeOf((*MockPoolServiceIface)(nil).ListStoragePools), p)
}

// NewCreateStoragePoolParams mocks base method.
func (m *MockPoolServiceIface) NewCreateStoragePoolParams(name, url, zoneid string) *CreateStoragePoolParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateStoragePoolParams", name, url, zoneid)
	ret0, _ := ret[0].(*CreateStoragePoolParams)
	return ret0
}

// NewCreateStoragePoolParams indicates an expected call of NewCreateStoragePoolParams.
func (mr *MockPoolServiceIfaceMockRecorder) NewCreateStoragePoolParams(name, url, zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateStoragePoolParams", reflect.TypeOf((*MockPoolServiceIface)(nil).NewCreateStoragePoolParams), name, url, zoneid)
}

// NewDeleteStoragePoolParams mocks base method.
func (m *MockPoolServiceIface) NewDeleteStoragePoolParams(id string) *DeleteStoragePoolParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteStoragePoolParams", id)
	ret0, _ := ret[0].(*DeleteStoragePoolParams)
	return ret0
}

// NewDeleteStoragePoolParams indicates an expected call of NewDeleteStoragePoolParams.
func (mr *MockPoolServiceIfaceMockRecorder) NewDeleteStoragePoolParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteStoragePoolParams", reflect.TypeOf((*MockPoolServiceIface)(nil).NewDeleteStoragePoolParams), id)
}

// NewFindStoragePoolsForMigrationParams mocks base method.
func (m *MockPoolServiceIface) NewFindStoragePoolsForMigrationParams(id string) *FindStoragePoolsForMigrationParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFindStoragePoolsForMigrationParams", id)
	ret0, _ := ret[0].(*FindStoragePoolsForMigrationParams)
	return ret0
}

// NewFindStoragePoolsForMigrationParams indicates an expected call of NewFindStoragePoolsForMigrationParams.
func (mr *MockPoolServiceIfaceMockRecorder) NewFindStoragePoolsForMigrationParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFindStoragePoolsForMigrationParams", reflect.TypeOf((*MockPoolServiceIface)(nil).NewFindStoragePoolsForMigrationParams), id)
}

// NewListStoragePoolsParams mocks base method.
func (m *MockPoolServiceIface) NewListStoragePoolsParams() *ListStoragePoolsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListStoragePoolsParams")
	ret0, _ := ret[0].(*ListStoragePoolsParams)
	return ret0
}

// NewListStoragePoolsParams indicates an expected call of NewListStoragePoolsParams.
func (mr *MockPoolServiceIfaceMockRecorder) NewListStoragePoolsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListStoragePoolsParams", reflect.TypeOf((*MockPoolServiceIface)(nil).NewListStoragePoolsParams))
}

// NewUpdateStoragePoolParams mocks base method.
func (m *MockPoolServiceIface) NewUpdateStoragePoolParams(id string) *UpdateStoragePoolParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateStoragePoolParams", id)
	ret0, _ := ret[0].(*UpdateStoragePoolParams)
	return ret0
}

// NewUpdateStoragePoolParams indicates an expected call of NewUpdateStoragePoolParams.
func (mr *MockPoolServiceIfaceMockRecorder) NewUpdateStoragePoolParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateStoragePoolParams", reflect.TypeOf((*MockPoolServiceIface)(nil).NewUpdateStoragePoolParams), id)
}

// UpdateStoragePool mocks base method.
func (m *MockPoolServiceIface) UpdateStoragePool(p *UpdateStoragePoolParams) (*UpdateStoragePoolResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStoragePool", p)
	ret0, _ := ret[0].(*UpdateStoragePoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStoragePool indicates an expected call of UpdateStoragePool.
func (mr *MockPoolServiceIfaceMockRecorder) UpdateStoragePool(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStoragePool", reflect.TypeOf((*MockPoolServiceIface)(nil).UpdateStoragePool), p)
}
