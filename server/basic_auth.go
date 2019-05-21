package server

import (
	"fmt"
	"github.com/tg123/go-htpasswd"
	"log"
	"net/http"
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
