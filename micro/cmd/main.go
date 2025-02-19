package main

import (
	"context"
	"log"
	"log/slog"
	"net"

	"github.com/paulja/go-arch/micro/config"
	"github.com/paulja/go-arch/micro/internal/interceptors"
	"github.com/paulja/go-arch/proto/search/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	vlog := slog.Default()
	slog.SetLogLoggerLevel(slog.LevelDebug)

	vlog.Info("users service", "endpoint", config.GetServePort())
	s := new(server)
	s.logger = vlog
	log.Fatal(s.Run())
}

type server struct {
	search.UnimplementedSearchServiceServer
	logger *slog.Logger
}

func (s *server) Run() error {
	listen, err := net.Listen("tcp", config.GetServePort())
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc.UnaryServerInterceptor(interceptors.CreateLogInterceptor(s.logger)),
		),
	)
	reflection.Register(grpcServer) // do not run in production
	search.RegisterSearchServiceServer(grpcServer, s)
	return grpcServer.Serve(listen)
}

var users []string = []string{"simon", "peter", "jeff", "sarah", "rachel"}

func (s *server) FindUsers(
	ctx context.Context,
	req *search.FindUsersRequest,
) (
	*search.FindUsersResponse,
	error,
) {
	return &search.FindUsersResponse{
		Users: users,
	}, nil
}
