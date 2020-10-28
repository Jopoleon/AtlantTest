package server

import (
	"net"
	"time"

	"github.com/Jopoleon/AtlantTest/client"

	"github.com/Jopoleon/AtlantTest/server/proto_server"

	"github.com/Jopoleon/AtlantTest/server/products"

	"github.com/Jopoleon/AtlantTest/config"
	"github.com/Jopoleon/AtlantTest/logger"
	"github.com/Jopoleon/AtlantTest/storage"

	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Log        *logger.LocalLogger
	Config     *config.Config
	Repository *storage.Storage
	Fetcher    client.CSVFetcher
}

func NewServer(cfg *config.Config,
	log *logger.LocalLogger,
	storage *storage.Storage,
) (*Server, error) {
	return &Server{
		Fetcher:    client.NewCSVFetchClient(log),
		Config:     cfg,
		Log:        log,
		Repository: storage,
	}, nil
}

func (s *Server) Serve() {

	listen, err := net.Listen("tcp", ":"+s.Config.HttpPort)
	if err != nil {
		s.Log.Fatalf("failed to listen: %v\n", err)
	}
	ps := products.NewProductsService(s.Log, s.Config, s.Repository, s.Fetcher)

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: (time.Duration(10) * time.Second),
			Time:              (time.Duration(10) * time.Second),
			Timeout:           (time.Duration(7) * time.Second),
		}),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             (time.Duration(10) * time.Second),
				PermitWithoutStream: true,
			},
		),
		grpc.MaxRecvMsgSize(1000000000),
	)

	reflection.Register(grpcServer)
	proto_server.RegisterProductServiceServer(grpcServer, ps)

	s.Log.Info("grpc server started on port: ", s.Config.HttpPort)
	s.Log.Fatal(grpcServer.Serve(listen))
}
