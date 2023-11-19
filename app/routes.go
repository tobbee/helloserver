package app

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	_ "net/http/pprof" // For profiling at /debug

	"github.com/go-chi/chi/middleware"
)

// Routes defines dispatches for all routes.
func (s *Server) Routes() error {
	s.Router.Mount("/debug", middleware.Profiler())
	s.Router.MethodFunc("GET", "/healthz", s.healthzHandlerFunc)
	s.Router.MethodFunc("GET", "/static/*", s.staticHandlerFunc)
	s.Router.MethodFunc("HEAD", "/static/*", s.staticHandlerFunc)
	s.Router.MethodFunc("GET", "/", s.indexHandlerFunc)
	return nil
}

// optionsHandlerFunc provides the allowed methods.
func (s *Server) optionsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "OPTIONS, GET, HEAD")
}

func (s *Server) healthzHandlerFunc(w http.ResponseWriter, r *http.Request) {
	s.jsonResponse(w, true, http.StatusOK)
}

// jsonResponse marshals message and give response with code
//
// Don't add any more content after this since Content-Length is set
func (s *Server) jsonResponse(w http.ResponseWriter, message interface{}, code int) {
	raw, err := json.Marshal(message)
	if err != nil {
		http.Error(w, fmt.Sprintf("{message: \"%s\"}", err), http.StatusInternalServerError)
		slog.Error(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Header().Set("Content-Length", strconv.Itoa(len(raw)))
	_, err = w.Write(raw)
	if err != nil {
		slog.Error("could not write HTTP response", "err", err)
	}
}
