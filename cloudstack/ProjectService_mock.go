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
// Source: ./cloudstack/ProjectService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProjectServiceIface is a mock of ProjectServiceIface interface.
type MockProjectServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockProjectServiceIfaceMockRecorder
}

// MockProjectServiceIfaceMockRecorder is the mock recorder for MockProjectServiceIface.
type MockProjectServiceIfaceMockRecorder struct {
	mock *MockProjectServiceIface
}

// NewMockProjectServiceIface creates a new mock instance.
func NewMockProjectServiceIface(ctrl *gomock.Controller) *MockProjectServiceIface {
	mock := &MockProjectServiceIface{ctrl: ctrl}
	mock.recorder = &MockProjectServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectServiceIface) EXPECT() *MockProjectServiceIfaceMockRecorder {
	return m.recorder
}

// ActivateProject mocks base method.
func (m *MockProjectServiceIface) ActivateProject(p *ActivateProjectParams) (*ActivateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateProject", p)
	ret0, _ := ret[0].(*ActivateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateProject indicates an expected call of ActivateProject.
func (mr *MockProjectServiceIfaceMockRecorder) ActivateProject(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateProject", reflect.TypeOf((*MockProjectServiceIface)(nil).ActivateProject), p)
}

// CreateProject mocks base method.
func (m *MockProjectServiceIface) CreateProject(p *CreateProjectParams) (*CreateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", p)
	ret0, _ := ret[0].(*CreateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectServiceIfaceMockRecorder) CreateProject(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectServiceIface)(nil).CreateProject), p)
}

// DeleteProject mocks base method.
func (m *MockProjectServiceIface) DeleteProject(p *DeleteProjectParams) (*DeleteProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", p)
	ret0, _ := ret[0].(*DeleteProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockProjectServiceIfaceMockRecorder) DeleteProject(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectServiceIface)(nil).DeleteProject), p)
}

// DeleteProjectInvitation mocks base method.
func (m *MockProjectServiceIface) DeleteProjectInvitation(p *DeleteProjectInvitationParams) (*DeleteProjectInvitationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectInvitation", p)
	ret0, _ := ret[0].(*DeleteProjectInvitationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectInvitation indicates an expected call of DeleteProjectInvitation.
func (mr *MockProjectServiceIfaceMockRecorder) DeleteProjectInvitation(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectInvitation", reflect.TypeOf((*MockProjectServiceIface)(nil).DeleteProjectInvitation), p)
}

// GetProjectByID mocks base method.
func (m *MockProjectServiceIface) GetProjectByID(id string, opts ...OptionFunc) (*Project, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectByID", varargs...)
	ret0, _ := ret[0].(*Project)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectByID indicates an expected call of GetProjectByID.
func (mr *MockProjectServiceIfaceMockRecorder) GetProjectByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectByID", reflect.TypeOf((*MockProjectServiceIface)(nil).GetProjectByID), varargs...)
}

// GetProjectByName mocks base method.
func (m *MockProjectServiceIface) GetProjectByName(name string, opts ...OptionFunc) (*Project, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectByName", varargs...)
	ret0, _ := ret[0].(*Project)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectByName indicates an expected call of GetProjectByName.
func (mr *MockProjectServiceIfaceMockRecorder) GetProjectByName(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectByName", reflect.TypeOf((*MockProjectServiceIface)(nil).GetProjectByName), varargs...)
}

// GetProjectID mocks base method.
func (m *MockProjectServiceIface) GetProjectID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectID indicates an expected call of GetProjectID.
func (mr *MockProjectServiceIfaceMockRecorder) GetProjectID(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectID", reflect.TypeOf((*MockProjectServiceIface)(nil).GetProjectID), varargs...)
}

// GetProjectInvitationByID mocks base method.
func (m *MockProjectServiceIface) GetProjectInvitationByID(id string, opts ...OptionFunc) (*ProjectInvitation, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectInvitationByID", varargs...)
	ret0, _ := ret[0].(*ProjectInvitation)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectInvitationByID indicates an expected call of GetProjectInvitationByID.
func (mr *MockProjectServiceIfaceMockRecorder) GetProjectInvitationByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectInvitationByID", reflect.TypeOf((*MockProjectServiceIface)(nil).GetProjectInvitationByID), varargs...)
}

// ListProjectInvitations mocks base method.
func (m *MockProjectServiceIface) ListProjectInvitations(p *ListProjectInvitationsParams) (*ListProjectInvitationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectInvitations", p)
	ret0, _ := ret[0].(*ListProjectInvitationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectInvitations indicates an expected call of ListProjectInvitations.
func (mr *MockProjectServiceIfaceMockRecorder) ListProjectInvitations(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectInvitations", reflect.TypeOf((*MockProjectServiceIface)(nil).ListProjectInvitations), p)
}

// ListProjects mocks base method.
func (m *MockProjectServiceIface) ListProjects(p *ListProjectsParams) (*ListProjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", p)
	ret0, _ := ret[0].(*ListProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockProjectServiceIfaceMockRecorder) ListProjects(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockProjectServiceIface)(nil).ListProjects), p)
}

// NewActivateProjectParams mocks base method.
func (m *MockProjectServiceIface) NewActivateProjectParams(id string) *ActivateProjectParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewActivateProjectParams", id)
	ret0, _ := ret[0].(*ActivateProjectParams)
	return ret0
}

// NewActivateProjectParams indicates an expected call of NewActivateProjectParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewActivateProjectParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewActivateProjectParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewActivateProjectParams), id)
}

// NewCreateProjectParams mocks base method.
func (m *MockProjectServiceIface) NewCreateProjectParams(displaytext, name string) *CreateProjectParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateProjectParams", displaytext, name)
	ret0, _ := ret[0].(*CreateProjectParams)
	return ret0
}

// NewCreateProjectParams indicates an expected call of NewCreateProjectParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewCreateProjectParams(displaytext, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateProjectParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewCreateProjectParams), displaytext, name)
}

// NewDeleteProjectInvitationParams mocks base method.
func (m *MockProjectServiceIface) NewDeleteProjectInvitationParams(id string) *DeleteProjectInvitationParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteProjectInvitationParams", id)
	ret0, _ := ret[0].(*DeleteProjectInvitationParams)
	return ret0
}

// NewDeleteProjectInvitationParams indicates an expected call of NewDeleteProjectInvitationParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewDeleteProjectInvitationParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteProjectInvitationParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewDeleteProjectInvitationParams), id)
}

// NewDeleteProjectParams mocks base method.
func (m *MockProjectServiceIface) NewDeleteProjectParams(id string) *DeleteProjectParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteProjectParams", id)
	ret0, _ := ret[0].(*DeleteProjectParams)
	return ret0
}

// NewDeleteProjectParams indicates an expected call of NewDeleteProjectParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewDeleteProjectParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteProjectParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewDeleteProjectParams), id)
}

// NewListProjectInvitationsParams mocks base method.
func (m *MockProjectServiceIface) NewListProjectInvitationsParams() *ListProjectInvitationsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListProjectInvitationsParams")
	ret0, _ := ret[0].(*ListProjectInvitationsParams)
	return ret0
}

// NewListProjectInvitationsParams indicates an expected call of NewListProjectInvitationsParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewListProjectInvitationsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListProjectInvitationsParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewListProjectInvitationsParams))
}

// NewListProjectsParams mocks base method.
func (m *MockProjectServiceIface) NewListProjectsParams() *ListProjectsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListProjectsParams")
	ret0, _ := ret[0].(*ListProjectsParams)
	return ret0
}

// NewListProjectsParams indicates an expected call of NewListProjectsParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewListProjectsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListProjectsParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewListProjectsParams))
}

// NewSuspendProjectParams mocks base method.
func (m *MockProjectServiceIface) NewSuspendProjectParams(id string) *SuspendProjectParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSuspendProjectParams", id)
	ret0, _ := ret[0].(*SuspendProjectParams)
	return ret0
}

// NewSuspendProjectParams indicates an expected call of NewSuspendProjectParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewSuspendProjectParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSuspendProjectParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewSuspendProjectParams), id)
}

// NewUpdateProjectInvitationParams mocks base method.
func (m *MockProjectServiceIface) NewUpdateProjectInvitationParams(projectid string) *UpdateProjectInvitationParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateProjectInvitationParams", projectid)
	ret0, _ := ret[0].(*UpdateProjectInvitationParams)
	return ret0
}

// NewUpdateProjectInvitationParams indicates an expected call of NewUpdateProjectInvitationParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewUpdateProjectInvitationParams(projectid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateProjectInvitationParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewUpdateProjectInvitationParams), projectid)
}

// NewUpdateProjectParams mocks base method.
func (m *MockProjectServiceIface) NewUpdateProjectParams(id string) *UpdateProjectParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateProjectParams", id)
	ret0, _ := ret[0].(*UpdateProjectParams)
	return ret0
}

// NewUpdateProjectParams indicates an expected call of NewUpdateProjectParams.
func (mr *MockProjectServiceIfaceMockRecorder) NewUpdateProjectParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateProjectParams", reflect.TypeOf((*MockProjectServiceIface)(nil).NewUpdateProjectParams), id)
}

// SuspendProject mocks base method.
func (m *MockProjectServiceIface) SuspendProject(p *SuspendProjectParams) (*SuspendProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuspendProject", p)
	ret0, _ := ret[0].(*SuspendProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuspendProject indicates an expected call of SuspendProject.
func (mr *MockProjectServiceIfaceMockRecorder) SuspendProject(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuspendProject", reflect.TypeOf((*MockProjectServiceIface)(nil).SuspendProject), p)
}

// UpdateProject mocks base method.
func (m *MockProjectServiceIface) UpdateProject(p *UpdateProjectParams) (*UpdateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", p)
	ret0, _ := ret[0].(*UpdateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockProjectServiceIfaceMockRecorder) UpdateProject(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectServiceIface)(nil).UpdateProject), p)
}

// UpdateProjectInvitation mocks base method.
func (m *MockProjectServiceIface) UpdateProjectInvitation(p *UpdateProjectInvitationParams) (*UpdateProjectInvitationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectInvitation", p)
	ret0, _ := ret[0].(*UpdateProjectInvitationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectInvitation indicates an expected call of UpdateProjectInvitation.
func (mr *MockProjectServiceIfaceMockRecorder) UpdateProjectInvitation(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectInvitation", reflect.TypeOf((*MockProjectServiceIface)(nil).UpdateProjectInvitation), p)
}
