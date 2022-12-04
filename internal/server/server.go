package server

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/internal/controller/v1/user"
	"github.com/gaulzhw/go-server/internal/store"
	"github.com/gaulzhw/go-server/pkg/core"
	"github.com/gaulzhw/go-server/pkg/errno"
)

type Server struct {
	factory store.Factory

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

func (s *Server) InjectStoreFactory(factory store.Factory) {
	s.factory = factory
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

	r.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ParseCoder(errno.ErrUnknown.Error()), nil)
	})

	// pprof
	{
		r.GET("/debug/pprof/", gin.WrapF(pprof.Index))
		r.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
		r.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
		r.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
		r.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))
	}

	// middleware
	v1 := r.Group("/api/v1")
	{
		// user RESTful resources
		userv1 := v1.Group("/users")
		{
			userController := user.NewController(s.factory)
			userv1.GET(":name", userController.Get)
			userv1.POST("", userController.Create)
		}
	}

	return r
}
