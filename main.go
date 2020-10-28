package main

import (
	"github.com/Jopoleon/AtlantTest/app"
	"github.com/Jopoleon/AtlantTest/config"
	"github.com/Jopoleon/AtlantTest/logger"
)

func main() {

	cfg := config.NewConfig()
	ll := logger.NewLogger(cfg.ProductionStart)

	a, err := app.NewApp(cfg, ll)
	if err != nil {
		ll.Fatal(err)
	}
	a.Run()

}
