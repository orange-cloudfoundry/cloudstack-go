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

package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectService_CreateProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createProject"
		response, err := ParseAsyncResponse(apiName, "ProjectService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Project.NewCreateProjectParams("testProject", "testProject")
	resp, err := client.Project.CreateProject(params)
	if err != nil {
		t.Errorf("Failed to create Project due to: %v", err)
		return
	}
	if resp == nil || resp.Name != "testProject" {
		t.Errorf("Failed to create project")
	}
}

func TestProjectService_ActivateProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "activateProject"
		response, err := ParseAsyncResponse(apiName, "ProjectService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Project.NewActivateProjectParams("99a842a4-e50f-4265-8ca7-249959506c13")
	resp, err := client.Project.ActivateProject(params)
	if err != nil {
		t.Errorf("Failed to activate Project due to: %v", err)
		return
	}
	if resp == nil || resp.Id != "99a842a4-e50f-4265-8ca7-249959506c13" || resp.State != "Active" {
		t.Errorf("Failed to activate project")
	}
}

func TestProjectService_SuspendProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "suspendProject"
		response, err := ParseAsyncResponse(apiName, "ProjectService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Project.NewSuspendProjectParams("99a842a4-e50f-4265-8ca7-249959506c13")
	resp, err := client.Project.SuspendProject(params)
	if err != nil {
		t.Errorf("Failed to suspend Project due to: %v", err)
		return
	}
	if resp == nil || resp.Id != "99a842a4-e50f-4265-8ca7-249959506c13" || resp.State != "Suspended" {
		t.Errorf("Failed to suspend project")
	}
}

func TestProjectService_UpdateProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "updateProject"
		response, err := ParseAsyncResponse(apiName, "ProjectService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Project.NewUpdateProjectParams("99a842a4-e50f-4265-8ca7-249959506c13")
	params.SetDisplaytext("testProjectUpdated")
	resp, err := client.Project.UpdateProject(params)
	if err != nil {
		t.Errorf("Failed to update Project details due to: %v", err)
		return
	}

	if resp == nil || resp.Id != "69646881-8d7f-4800-987d-106698a42608" || resp.Displaytext != "testProjectUpdate" {
		t.Errorf("Failed to update project name")
	}
}
