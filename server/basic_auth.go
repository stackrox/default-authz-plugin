/*
Copyright 2019 StackRox Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tg123/go-htpasswd"
)

var (
	// only support bcrypt for security reasons
	htpasswdParsers = []htpasswd.PasswdParser{
		htpasswd.AcceptBcrypt,
	}
)

func handleWithBasicAuth(w http.ResponseWriter, r *http.Request, htpasswdFile *htpasswd.File, handler http.Handler) {
	username, password, ok := r.BasicAuth()
	log.Println("Handle with basic auth", username, password, ok)
	if !ok {
		w.Header().Set(`WWW-Authenticate`, `Basic realm="StackRox Authorization Plugin"`)
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("not authorized"))
		return
	}

	if !htpasswdFile.Match(username, password) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("invalid credentials"))
		return
	}

	handler.ServeHTTP(w, r)
}

func createBasicAuthHandler(htpasswdFilePath string, handler http.Handler) (http.Handler, error) {
	badLines := 0
	htpasswdFile, err := htpasswd.New(htpasswdFilePath, htpasswdParsers, func(error) { badLines++ })
	if err != nil {
		return nil, err
	}
	if badLines > 0 {
		return nil, fmt.Errorf("htpasswd file %s contains %d bad lines (note that only bcrypt is allowed)", htpasswdFilePath, badLines)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleWithBasicAuth(w, r, htpasswdFile, handler)
	}), nil
}
