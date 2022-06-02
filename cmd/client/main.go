package main

import (
	"advdiploma/client/pkg"
	"advdiploma/client/provider/http"
	"advdiploma/client/services"
	"advdiploma/client/storage/sqllite"
	"advdiploma/client/tui"
	"context"
	"github.com/rivo/tview"
	"log"
	"time"
)

var app = tview.NewApplication()

func main() {

	cfg, err := pkg.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	provCfg := http.HTTPConfig{
		AuthURL:     "/api/user/login",
		RegisterURL: "/api/user/register",
		SecretURL:   "/api/secret",
		SyncListURL: "/api/sync",
		PingURL:     "/api/ping",
		BaseURL:     cfg.ServerURL,
		Timeout:     time.Millisecond * 500,
	}

	db, err := sqllite.NewStorage(cfg.StorageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	provider := http.NewHTTPProvider(provCfg)

	svcSecret := services.NewSecret(cfg, db)
	svcSync := services.NewSyncService(db, provider, cfg)

	if err := svcSync.Run(context.Background()); err != nil {
		log.Fatal(err)
	}

	tui.SetQ(app, svcSecret)

}
