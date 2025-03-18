package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/monsavile/rubik-scrumbler/internal/config"
	scrumblerV1 "github.com/monsavile/rubik-scrumbler/pkg/scrumbler_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServerConfig interface {
	Port() int
}

type server struct {
	scrumblerV1.UnimplementedScrumblerV1Server
}

func (s *server) Scrumble(ctx context.Context, req *scrumblerV1.ScrumbleRequest) (*scrumblerV1.ScrumbleResponse, error) {
	return &scrumblerV1.ScrumbleResponse{
		Cube: "test",
	}, nil
}

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	grpcServerConfig, err := config.NewGRPCServerConfig()
	if err != nil {
		log.Fatalf("failed to get grpc server config: %s", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcServerConfig.Port()))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()
	server := server{}

	reflection.Register(grpcServer)
	scrumblerV1.RegisterScrumblerV1Server(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
