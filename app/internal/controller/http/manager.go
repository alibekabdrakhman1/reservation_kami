package http

import (
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/service"
	"go.uber.org/zap"
)

type Manager struct {
	Reservation IReservationHandler
}

func NewManager(service *service.Manager, logger *zap.SugaredLogger) *Manager {
	return &Manager{
		Reservation: NewReservationHandler(service, logger),
	}
}
