// package goproject
package Go_project

import (
	"context"
	"net/http"
	"time"
)

// абстракция над http.Server
type Server struct {
	httpServer *http.Server
}

// Запуск сервера
func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Остановка сервера
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
