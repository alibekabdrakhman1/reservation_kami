package service

import (
	"context"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/model"
)

type IReservationService interface {
	CreateReservation(ctx context.Context, reservation *model.InputReservation) (string, error)
	GetReservationsByRoomID(ctx context.Context, roomID string) ([]model.Reservation, error)
}
