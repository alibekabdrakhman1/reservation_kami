package repository

import (
	"context"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/model"
)

type IReservationRepository interface {
	CreateReservation(ctx context.Context, reservation *model.Reservation) (string, error)
	GetReservationsByRoomID(ctx context.Context, roomID string) ([]model.Reservation, error)
}
