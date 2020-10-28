package mongo

import (
	"context"
	"fmt"

	"github.com/Jopoleon/AtlantTest/config"
	"github.com/Jopoleon/AtlantTest/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDB struct {
	mc        *mongo.Client
	cfg       *config.Config
	productDB *mongo.Database
	Logger    *logger.LocalLogger
}

func NewMongoDB(cfg *config.Config, log *logger.LocalLogger) (*MongoDB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s", cfg.DBConfig.DBHost, cfg.DBConfig.DBPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	return &MongoDB{
		mc:        client,
		Logger:    log,
		cfg:       cfg,
		productDB: client.Database(cfg.DBConfig.DBName),
	}, client.Ping(ctx, &readpref.ReadPref{})
}
