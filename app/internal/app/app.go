package app

import (
	"context"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"reservation/app/internal/config"
	"reservation/app/internal/controller"
	"reservation/app/internal/controller/http"
	"reservation/app/internal/repository"
	"reservation/app/internal/service"
	"reservation/app/pkg/db"
)

type App struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func New(logger *zap.SugaredLogger, config *config.Config) *App {
	return &App{
		logger: logger,
		config: config,
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	gracefullyShutdown(cancel)

	db, err := db.Dial(ctx, a.config.DSN())
	if err != nil {
		log.Fatalf("cannot сonnect to DB '%s:%d': %v", a.config.Database.Host, a.config.Database.Port, err)
	}

	repo := repository.NewManager(db)

	srv := service.NewManager(repo, a.config, a.logger)

	handler := http.NewManager(srv, a.logger)

	HTTPServer := controller.NewServer(a.config, handler)

	return HTTPServer.StartHTTPServer(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
