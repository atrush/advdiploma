package main

import (
	"advdiploma/server/api"
	"advdiploma/server/pkg"
	"advdiploma/server/services/auth"
	"advdiploma/server/services/secret"
	"advdiploma/server/storage/psql"
	"advdiploma/server/storage/psql/migrations"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	cfg, err := pkg.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	if cfg.Migrate {

		log.Println("starting migrations")
		if err := migrations.RunMigrations(cfg.DatabaseDSN, cfg.TableName); err != nil {
			log.Fatal(err.Error())
		}
		log.Println("migrations ended")
		return
	}

	db, err := psql.NewStorage(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err.Error())
	}

	jwtAuth, err := auth.NewAuth(db)
	if err != nil {
		log.Fatalf("error starting auth service:%v", err.Error())
	}

	svcSecret, err := secret.NewSecret(db)
	if err != nil {
		log.Fatalf("error starting secret service:%v", err.Error())
	}

	server, err := api.NewServer(cfg, jwtAuth, svcSecret)
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		if err := server.Run(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatal(err)
			}
		}
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	<-sigc

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("error shutdown server: %s\n", err.Error())
	}

}
