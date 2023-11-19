package app

import (
	"embed"
	"fmt"
	htmpl "html/template"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	port          int
	Router        *chi.Mux
	htmlTemplates *htmpl.Template
}

//go:embed static/* templates/*

var embeddedContent embed.FS

// staticHandlerFunc handles static files in tree starting at static
func (s *Server) staticHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.FS(embeddedContent))
	fs.ServeHTTP(w, r)
}

// indexHandlerFunc handles access to /.
func (s *Server) indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	wi := welcomeInfo{Host: fmt.Sprintf(":%d", s.port), Version: s.getVersion(),
		Scouts: []string{"Kalle", "Britta"}}
	err := s.htmlTemplates.ExecuteTemplate(w, "welcome.html", wi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) compileTemplates() error {
	var err error
	s.htmlTemplates, err = compileHTMLTemplates(embeddedContent, "templates")
	if err != nil {
		return fmt.Errorf("compileHTMLTemplates: %w", err)
	}
	slog.Info("html templates", "defined", s.htmlTemplates.DefinedTemplates())
	return nil
}

func NewServer(port int) (*Server, error) {
	logger := slog.Default()
	s := Server{
		port:   port,
		Router: chi.NewRouter(),
	}
	s.Router.Use(middleware.Recoverer)
	err := s.compileTemplates()
	if err != nil {
		return nil, err
	}
	err = s.Routes()
	if err != nil {
		return nil, fmt.Errorf("routes: %w", err)
	}
	logger.Info("helloserver starting", "version", s.getVersion(), "port", port)
	return &s, nil
}

func (s *Server) getVersion() string {
	return "0.1"
}

// Template information
type welcomeInfo struct {
	Host    string
	Version string
	Scouts  []string
}
