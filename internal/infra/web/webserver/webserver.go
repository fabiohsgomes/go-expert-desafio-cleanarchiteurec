package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string][]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string][]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers[method] = append(s.Handlers[method], map[string]http.HandlerFunc{path: handler})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	
	for method, handlers := range s.Handlers {
		for _, handler := range handlers {
			for path, fn := range handler {
				switch method {
				case http.MethodGet:
					s.Router.Get(path, fn)
				case http.MethodPost:
					s.Router.Post(path, fn)
				case http.MethodPut:
					s.Router.Put(path, fn)
				case http.MethodDelete:
					s.Router.Delete(path, fn)
				default:
					s.Router.Method(method, path, fn)
				}
			}
		}
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
