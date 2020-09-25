package xgin

import (
	"context"
	"fmt"
	"go-scaffold/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

// HTTPServer ..
type HTTPServer struct {
	http.Server
	conf *Config
}

// ListenAndServe ..
func (t *HTTPServer) ListenAndServe() {
	fmt.Printf("Listening and serving HTTP on %s\n", t.Addr)

	go func() {
		if err := t.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("[HTTP Serve Error]", zap.Any("error", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), t.conf.ShutdownTTL)
	defer cancel()
	if err := t.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", zap.Error(err))
	}
	log.Info("Server exited...")
	fmt.Printf("\x1b[32m%s\x1b[0m", `
bye :)
`)
}
