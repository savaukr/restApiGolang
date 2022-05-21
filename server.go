package todo

import "net/http"

type Server struct {
	httpServer *http.Server
}

func(s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		MaxHeaderBytes: 1<< 28, 
		ReaderTimeout: 10*time.Second,
		WriteTimeout: 10 *time.Second,
	}

	return s.httpServer.ListenAndServer()
} 

func (s*Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}