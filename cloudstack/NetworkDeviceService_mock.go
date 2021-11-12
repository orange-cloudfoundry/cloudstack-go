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
// Source: ./cloudstack/NetworkDeviceService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNetworkDeviceServiceIface is a mock of NetworkDeviceServiceIface interface.
type MockNetworkDeviceServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkDeviceServiceIfaceMockRecorder
}

// MockNetworkDeviceServiceIfaceMockRecorder is the mock recorder for MockNetworkDeviceServiceIface.
type MockNetworkDeviceServiceIfaceMockRecorder struct {
	mock *MockNetworkDeviceServiceIface
}

// NewMockNetworkDeviceServiceIface creates a new mock instance.
func NewMockNetworkDeviceServiceIface(ctrl *gomock.Controller) *MockNetworkDeviceServiceIface {
	mock := &MockNetworkDeviceServiceIface{ctrl: ctrl}
	mock.recorder = &MockNetworkDeviceServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkDeviceServiceIface) EXPECT() *MockNetworkDeviceServiceIfaceMockRecorder {
	return m.recorder
}

// AddNetworkDevice mocks base method.
func (m *MockNetworkDeviceServiceIface) AddNetworkDevice(p *AddNetworkDeviceParams) (*AddNetworkDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNetworkDevice", p)
	ret0, _ := ret[0].(*AddNetworkDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNetworkDevice indicates an expected call of AddNetworkDevice.
func (mr *MockNetworkDeviceServiceIfaceMockRecorder) AddNetworkDevice(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNetworkDevice", reflect.TypeOf((*MockNetworkDeviceServiceIface)(nil).AddNetworkDevice), p)
}

// DeleteNetworkDevice mocks base method.
func (m *MockNetworkDeviceServiceIface) DeleteNetworkDevice(p *DeleteNetworkDeviceParams) (*DeleteNetworkDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetworkDevice", p)
	ret0, _ := ret[0].(*DeleteNetworkDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNetworkDevice indicates an expected call of DeleteNetworkDevice.
func (mr *MockNetworkDeviceServiceIfaceMockRecorder) DeleteNetworkDevice(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkDevice", reflect.TypeOf((*MockNetworkDeviceServiceIface)(nil).DeleteNetworkDevice), p)
}

// ListNetworkDevice mocks base method.
func (m *MockNetworkDeviceServiceIface) ListNetworkDevice(p *ListNetworkDeviceParams) (*ListNetworkDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetworkDevice", p)
	ret0, _ := ret[0].(*ListNetworkDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetworkDevice indicates an expected call of ListNetworkDevice.
func (mr *MockNetworkDeviceServiceIfaceMockRecorder) ListNetworkDevice(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetworkDevice", reflect.TypeOf((*MockNetworkDeviceServiceIface)(nil).ListNetworkDevice), p)
}

// NewAddNetworkDeviceParams mocks base method.
func (m *MockNetworkDeviceServiceIface) NewAddNetworkDeviceParams() *AddNetworkDeviceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddNetworkDeviceParams")
	ret0, _ := ret[0].(*AddNetworkDeviceParams)
	return ret0
}

// NewAddNetworkDeviceParams indicates an expected call of NewAddNetworkDeviceParams.
func (mr *MockNetworkDeviceServiceIfaceMockRecorder) NewAddNetworkDeviceParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddNetworkDeviceParams", reflect.TypeOf((*MockNetworkDeviceServiceIface)(nil).NewAddNetworkDeviceParams))
}

// NewDeleteNetworkDeviceParams mocks base method.
func (m *MockNetworkDeviceServiceIface) NewDeleteNetworkDeviceParams(id string) *DeleteNetworkDeviceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteNetworkDeviceParams", id)
	ret0, _ := ret[0].(*DeleteNetworkDeviceParams)
	return ret0
}

// NewDeleteNetworkDeviceParams indicates an expected call of NewDeleteNetworkDeviceParams.
func (mr *MockNetworkDeviceServiceIfaceMockRecorder) NewDeleteNetworkDeviceParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteNetworkDeviceParams", reflect.TypeOf((*MockNetworkDeviceServiceIface)(nil).NewDeleteNetworkDeviceParams), id)
}

// NewListNetworkDeviceParams mocks base method.
func (m *MockNetworkDeviceServiceIface) NewListNetworkDeviceParams() *ListNetworkDeviceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListNetworkDeviceParams")
	ret0, _ := ret[0].(*ListNetworkDeviceParams)
	return ret0
}

// NewListNetworkDeviceParams indicates an expected call of NewListNetworkDeviceParams.
func (mr *MockNetworkDeviceServiceIfaceMockRecorder) NewListNetworkDeviceParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListNetworkDeviceParams", reflect.TypeOf((*MockNetworkDeviceServiceIface)(nil).NewListNetworkDeviceParams))
}
