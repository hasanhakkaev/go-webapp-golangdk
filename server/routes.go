package server

import (
	"go-web-app/handlers"
)

func (s *Server) setupRoutes() {
	handlers.Home(s.mux)
	handlers.Team(s.mux)
	handlers.Projects(s.mux)
	handlers.Calendar(s.mux)
}
