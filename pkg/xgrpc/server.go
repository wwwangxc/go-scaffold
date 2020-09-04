package xgrpc

import (
	"fmt"
	"go-scaffold/pkg/log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func newServer(conf *ServerConfig) *GrpcServer {
	ln, err := net.Listen(conf.Network, conf.Addr)
	if err != nil {
		panic(err.Error())
	}
	srv := grpc.NewServer()
	for _, fn := range conf.fns {
		fn(srv)
	}
	return &GrpcServer{
		ln:   ln,
		srv:  srv,
		conf: conf,
	}
}

// GrpcServer ..
type GrpcServer struct {
	ln   net.Listener
	srv  *grpc.Server
	conf *ServerConfig
}

// Serve ..
func (t *GrpcServer) Serve() {
	serviceKey := fmt.Sprintf("%s:///%s/%s", t.conf.Scheme, t.conf.Name, t.conf.Addr)
	t.conf.register.RegistryService(serviceKey, t.conf.Addr)
	go t.srv.Serve(t.ln)
	fmt.Printf("Listening and serving grpc on %s\n", t.conf.Addr)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	log.Info("Server exiting...")
	t.conf.register.UnRegistryService(serviceKey)
	if i, ok := s.(syscall.Signal); ok {
		os.Exit(int(i))
	} else {
		os.Exit(0)
	}
}

// Close..
func (t *GrpcServer) Close() {
	defer t.ln.Close()
	defer t.srv.GracefulStop()
}
