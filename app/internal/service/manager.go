package service

import (
	"go.uber.org/zap"
	"reservation/app/internal/config"
	"reservation/app/internal/repository"
)

type Manager struct {
	Reservation IReservationService
}

func NewManager(repository *repository.Manager, config *config.Config, logger *zap.SugaredLogger) *Manager {
	return &Manager{
		Reservation: NewReservationService(repository, config, logger),
	}
}
