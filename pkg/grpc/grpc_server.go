package grpc

import (
	"github.com/ololko/simple-HTTP-server/pkg/events/apis"
	gw "github.com/ololko/simple-HTTP-server/pkg/events/models"
	"log"
	"net"
	"os"
	"os/signal"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func RunGrpcServer(ctx context.Context, svc *apis.Service, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	gw.RegisterEventsServer(server, svc)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
