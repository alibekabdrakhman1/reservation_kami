package service

import (
	"context"
	"reservation/app/internal/model"
)

type IReservationService interface {
	CreateReservation(ctx context.Context, reservation *model.InputReservation) (string, error)
	GetReservationsByRoomID(ctx context.Context, roomID string) ([]model.Reservation, error)
}
