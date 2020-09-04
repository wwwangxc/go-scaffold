package xgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func newServer(conf *Config) *HTTPServer {
	return &HTTPServer{
		conf: conf,
		Server: http.Server{
			Addr:    fmt.Sprintf(":%d", conf.Port),
			Handler: newEngine(conf),
		},
	}
}

func newEngine(conf *Config) *gin.Engine {
	gin.SetMode(conf.Mode)
	engine := gin.New()
	for _, fn := range conf.fns {
		fn(engine)
	}
	return engine
}
