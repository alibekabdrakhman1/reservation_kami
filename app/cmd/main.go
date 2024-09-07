package main

import (
	"fmt"
	"go.uber.org/zap"
	"reservation/app/internal/app"
	"reservation/app/internal/config"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With(zap.String("app", "reservation-service"))

	cfg, err := config.LoadConfig("./")
	fmt.Println(cfg)
	if err != nil {
		l.Error(err)
		l.Fatalf("failed to load configs err: %v", err)
	}

	app := app.New(l, &cfg)

	app.Run()
}
