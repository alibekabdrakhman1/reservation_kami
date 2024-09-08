package service

import (
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/repository"
	"go.uber.org/zap"
)

type Manager struct {
	Reservation IReservationService
}

func NewManager(repository *repository.Manager, logger *zap.SugaredLogger) *Manager {
	return &Manager{
		Reservation: NewReservationService(repository, logger),
	}
}
