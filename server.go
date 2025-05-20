package WhyAi

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 28,
		ReadTimeout:    40 * time.Second,
		WriteTimeout:   40 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Close(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
