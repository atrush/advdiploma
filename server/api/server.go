package api

import (
	"advdiploma/server/api/handler"
	"advdiploma/server/pkg"
	"advdiploma/server/services/auth"
	"advdiploma/server/services/secret"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type Server struct {
	httpServer http.Server
}

func NewServer(cfg *pkg.Config, a auth.Authenticator, secret secret.SecretManager) (*Server, error) {

	h, err := handler.NewHandler(a, secret)
	if err != nil {
		return nil, fmt.Errorf("ошибка запуска server:%w", err)
	}

	return &Server{
		httpServer: http.Server{
			Addr:    cfg.ServerPort,
			Handler: handler.GetRouter(h),
		},
	}, nil
}

// Run Start server
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown server
func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.ListenAndServe(); err == http.ErrServerClosed {
		return errors.New("http server not runned")
	}

	return s.httpServer.Shutdown(ctx)
}
