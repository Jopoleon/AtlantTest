package app

import (
	"github.com/Jopoleon/AtlantTest/config"
	"github.com/Jopoleon/AtlantTest/logger"
	"github.com/Jopoleon/AtlantTest/server"
	"github.com/Jopoleon/AtlantTest/storage"
	"github.com/pkg/errors"
)

type App struct {
	Log    *logger.LocalLogger
	Config *config.Config
	Server *server.Server
}

func NewApp(cfg *config.Config, logger *logger.LocalLogger) (*App, error) {

	store, err := storage.NewStorage(cfg, logger)
	if err != nil {
		logger.Error(err)
		return nil, errors.WithStack(err)
	}
	s, err := server.NewServer(cfg, logger, store)
	if err != nil {
		logger.Error(err)
		return nil, errors.WithStack(err)
	}
	return &App{
		Log:    logger,
		Config: cfg,
		Server: s,
	}, nil
}

func (a *App) Run() {
	a.Server.Serve()
}
