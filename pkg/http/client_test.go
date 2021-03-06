//Copyright 2019 Expedia, Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestSuccess(t *testing.T) {
	path := "/path"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if want, have := path, r.URL.Path; want != have {
			t.Errorf("unexpected path, want: %q, have %q", want, have)
		}
	}))
	defer server.Close()

	c := NewClient(server.URL, false)
	reqBody := ""
	resp := c.SendRequest("GET", path, map[string]string{}, &reqBody)
	assert.Nil(t, resp.Err)
}

func TestHttpError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(400)
	}))
	defer server.Close()

	c := NewClient(server.URL, false)
	reqBody := ""
	resp := c.SendRequest("GET", "/", map[string]string{}, &reqBody)
	assert.Nil(t, resp.Err)
	assert.Equal(t, resp.StatusCode, 400)
}

func TestConnectionError(t *testing.T) {
	c := NewClient("http://localhost:9999", false)
	reqBody := ""
	resp := c.SendRequest("GET", "/potato", map[string]string{}, &reqBody)
	assert.NotNil(t, resp.Err)
}
