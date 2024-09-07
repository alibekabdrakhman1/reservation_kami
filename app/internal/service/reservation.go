package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"reservation/app/internal/config"
	"reservation/app/internal/model"
	"reservation/app/internal/repository"
	"time"
)

type ReservationService struct {
	repository *repository.Manager
	config     *config.Config
	logger     *zap.SugaredLogger
}

func NewReservationService(repository *repository.Manager, config *config.Config, logger *zap.SugaredLogger) *ReservationService {
	return &ReservationService{repository: repository, config: config, logger: logger}
}

func (s *ReservationService) CreateReservation(ctx context.Context, reservation *model.InputReservation) (string, error) {
	newReservation := &model.Reservation{
		ID:     uuid.New(),
		RoomID: reservation.RoomID,
	}
	var err error
	newReservation.StartTime, newReservation.EndTime, err = parseTimes(reservation.StartTime, reservation.EndTime)
	if err != nil {
		return "", err
	}

	return s.repository.Reservation.CreateReservation(ctx, newReservation)
}

func (s *ReservationService) GetReservationsByRoomID(ctx context.Context, roomID string) ([]model.Reservation, error) {
	return s.repository.Reservation.GetReservationsByRoomID(ctx, roomID)
}

func parseTimes(startTimeStr, endTimeStr string) (time.Time, time.Time, error) {
	const layout = "2006-01-02T15:04:05Z"
	startTime, err := time.Parse(layout, startTimeStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	endTime, err := time.Parse(layout, endTimeStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	if startTime.After(endTime) {
		return time.Time{}, time.Time{}, errors.New("start time must be before end time")
	}

	return startTime, endTime, nil
}
