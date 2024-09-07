package http

import (
	"go.uber.org/zap"
	"reservation/app/internal/service"
)

type Manager struct {
	Reservation IReservationHandler
}

func NewManager(service *service.Manager, logger *zap.SugaredLogger) *Manager {
	return &Manager{
		Reservation: NewReservationHandler(service, logger),
	}
}
