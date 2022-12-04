package server

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpSrv *http.Server
}

func NewServer(addr string, tls *tls.Config) *Server {
	s := &http.Server{
		Addr: addr,
	}

	if tls != nil {
		s.TLSConfig = tls
	}

	return &Server{
		httpSrv: s,
	}
}

func (s *Server) Start() error {
	// route
	s.httpSrv.Handler = s.router()

	if s.httpSrv.TLSConfig == nil {
		return s.httpSrv.ListenAndServe()
	}
	return s.httpSrv.ListenAndServeTLS("", "")
}

func (s *Server) Stop() {
	if s.httpSrv != nil {
		s.httpSrv.Shutdown(context.TODO())
	}
}

func (s *Server) router() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// pprof
	{
		r.GET("/debug/pprof/", gin.WrapF(pprof.Index))
		r.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
		r.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
		r.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
		r.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))
	}

	return r
}
