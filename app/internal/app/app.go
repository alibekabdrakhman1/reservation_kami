package app

import (
	"context"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/config"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/controller"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/controller/http"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/repository"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/service"
	"github.com/alibekabdrakhman1/reservation_kami/app/pkg/db"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
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

	srv := service.NewManager(repo, a.logger)

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
