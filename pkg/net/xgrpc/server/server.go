package server

import (
	"fmt"
	"go-scaffold/pkg/log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

// GrpcServer ..
type GrpcServer struct {
	conf *Config

	ln  net.Listener
	srv *grpc.Server

	closed bool
}

// create grpc server instance.
func newServer(conf *Config) *GrpcServer {
	ln, err := net.Listen(conf.Network, fmt.Sprintf("127.0.0.1:%d", conf.Port))
	if err != nil {
		panic(err.Error())
	}
	srv := grpc.NewServer()
	for _, service := range conf.services {
		service(srv)
	}
	return &GrpcServer{
		ln:     ln,
		srv:    srv,
		conf:   conf,
		closed: false,
	}
}

// Serve ..
func (t *GrpcServer) Serve() {
	serviceKey := fmt.Sprintf("%s:///%s/%s", t.conf.Scheme, t.conf.Name, t.ln.Addr().String())
	t.conf.register.RegistryService(serviceKey, t.ln.Addr().String())
	go func() {
		if err := t.srv.Serve(t.ln); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Listening and serving grpc on %s\n", t.ln.Addr().String())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutdown Server ...")
	t.conf.register.UnRegistryService(serviceKey)
	log.Info("Server exited...")
	fmt.Printf("\x1b[32m%s\x1b[0m", `
bye :)
`)
	t.srv.GracefulStop()
}
