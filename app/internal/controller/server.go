package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	http2 "net/http"
	"reservation/app/internal/config"
	"reservation/app/internal/controller/http"
	"time"
)

type Server struct {
	cfg     *config.Config
	handler *http.Manager
	App     *chi.Mux
}

func NewServer(cfg *config.Config, handler *http.Manager) *Server {
	return &Server{
		cfg:     cfg,
		handler: handler,
	}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()

	server := &http2.Server{
		Addr:    fmt.Sprintf(":%v", s.cfg.HttpServer.Port),
		Handler: s.App,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http2.ErrServerClosed) {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("controller Shutdown Failed:%v", err)
	}

	log.Print("controller exited properly")
	return nil
}

func (s *Server) BuildEngine() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "*"))

	return r
}
