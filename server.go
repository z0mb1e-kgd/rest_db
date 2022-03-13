package main

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port ...string) error {
	addrPort := serverDefaultPort
	if len(port) > 0 {
		addrPort = port[0]
	}
	addrPort = ":" + addrPort

	s.httpServer = &http.Server{
		Addr:           addrPort,
		MaxHeaderBytes: serverMaxHeaderBytes,
		ReadTimeout:    serverReadTimeout,
		WriteTimeout:   serverWriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
