package storage

import (
	"context"

	"github.com/Jopoleon/AtlantTest/config"
	"github.com/Jopoleon/AtlantTest/logger"
	"github.com/Jopoleon/AtlantTest/models"
	"github.com/Jopoleon/AtlantTest/server/proto_server"
	"github.com/Jopoleon/AtlantTest/storage/mongo"
	"github.com/pkg/errors"
)

type Storage struct {
	Logger *logger.LocalLogger
	DB     ProductManager
}

type ProductManager interface {
	GetProducts(ctx context.Context, filter *proto_server.ListRequest) ([]models.Product, error)
	SaveProducts(ctx context.Context, ps []models.Product) error
}

func NewStorage(cfg *config.Config, log *logger.LocalLogger) (*Storage, error) {
	res := &Storage{}
	db, err := mongo.NewMongoDB(cfg, log)
	if err != nil {
		log.Error(err)
		return res, errors.WithStack(err)
	}

	res.DB = db
	res.Logger = log
	return res, nil
}
