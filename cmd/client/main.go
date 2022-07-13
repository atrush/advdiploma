package main

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"advdiploma/client/provider/http"
	"advdiploma/client/services"
	"advdiploma/client/storage/sqllite"
	"advdiploma/client/tui"
	"context"
	"fmt"
	"github.com/rivo/tview"
	"log"
	"os"
	"os/signal"
	"time"
)

var app = tview.NewApplication()

func main() {
	fmt.Printf("Build version:%v\n", model.BuildVersion)
	fmt.Printf("Build date:%v\n", model.BuildDate)
	fmt.Printf("Build commit:%v\n", model.BuildCommit)

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

	//  run tui
	tui.SetQ(app, svcSecret)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	<-sigc
}
