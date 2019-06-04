package rules

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stackrox/default-authz-plugin/pkg/httperr"
	"github.com/stackrox/default-authz-plugin/pkg/payload"
)

// handler wraps a rule engine in a HTTP request handler.
type handler struct {
	engine Engine
}

// NewHandler creates and returns an http.Handler that processes authorization requests by dispatching them to the
// given rules engine.
func NewHandler(engine Engine) http.Handler {
	return &handler{
		engine: engine,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respBytes, err := h.handleHTTPRequest(r)
	if err != nil {
		httperr.Write(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(respBytes); err != nil {
		log.Println("Could not send HTTP response:", err)
		return
	}
	// Print a newline character to make output from, e.g., curl more readable.
	_, _ = fmt.Fprintln(w)
}

func (h *handler) handleHTTPRequest(r *http.Request) ([]byte, error) {
	if r.Method != http.MethodPost {
		return nil, httperr.New(http.StatusMethodNotAllowed, "only POST requests are allowed")
	}

	req, err := payload.ParseAndValidateRequest(r.Body)
	if err != nil {
		return nil, httperr.Newf(http.StatusBadRequest, "reading authorization request: %v", err)
	}

	resp, err := h.processAuthorizationRequest(req)
	if err != nil {
		return nil, httperr.Newf(http.StatusInternalServerError, "could not process authorization request: %v", err)
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		return nil, httperr.Newf(http.StatusInternalServerError, "could not marshal authorization response: %v", err)
	}
	return respBytes, nil
}

func (h *handler) processAuthorizationRequest(req *payload.AuthorizationRequest) (*payload.AuthorizationResponse, error) {
	resp := &payload.AuthorizationResponse{}

	for _, scope := range req.RequestedScopes {
		if ok, err := h.engine.Authorized(&req.Principal, &scope); err != nil {
			return nil, err
		} else if ok {
			resp.AuthorizedScopes = append(resp.AuthorizedScopes, scope)
		}
	}

	return resp, nil
}
