package products

import (
	"context"

	"github.com/Jopoleon/AtlantTest/models"

	"github.com/Jopoleon/AtlantTest/client"
	"github.com/Jopoleon/AtlantTest/config"
	"github.com/Jopoleon/AtlantTest/logger"
	"github.com/Jopoleon/AtlantTest/server/proto_server"
	"github.com/Jopoleon/AtlantTest/storage"
)

type Service struct {
	Log     *logger.LocalLogger
	Config  *config.Config
	fetcher client.CSVFetcher
	repo    *storage.Storage
}

func (s Service) List(ctx context.Context, request *proto_server.ListRequest) (*proto_server.ListReply, error) {
	ps, err := s.repo.DB.GetProducts(ctx, request)
	if err != nil {
		s.Log.Error(err)
		return nil, err
	}
	return &proto_server.ListReply{
		Products: models.Products(ps).ToGRPCProducts(),
		SortType: request.Sort.String(),
	}, nil
}

func (s Service) Fetch(ctx context.Context, request *proto_server.FetchRequest) (*proto_server.BasicReply, error) {
	ps, err := s.fetcher.GetCSV(request.Url)
	if err != nil {
		s.Log.Error(err)
		return nil, err
	}
	err = s.repo.DB.SaveProducts(ctx, ps)
	if err != nil {
		s.Log.Error(err)
		return nil, err
	}
	return &proto_server.BasicReply{
		Message: "all products saved",
		Code:    200,
	}, nil
}

func NewProductsService(log *logger.LocalLogger, config *config.Config,
	repo *storage.Storage, fetcher client.CSVFetcher) *Service {
	return &Service{
		Log:     log,
		Config:  config,
		repo:    repo,
		fetcher: fetcher,
	}
}
