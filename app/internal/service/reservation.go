package service

import (
	"context"
	"errors"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/model"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type ReservationService struct {
	repository *repository.Manager
	logger     *zap.SugaredLogger
}

func NewReservationService(repository *repository.Manager, logger *zap.SugaredLogger) *ReservationService {
	return &ReservationService{repository: repository, logger: logger}
}

func (s *ReservationService) CreateReservation(ctx context.Context, reservation *model.InputReservation) (string, error) {
	newReservation := &model.Reservation{
		ID:     uuid.New(),
		RoomID: reservation.RoomID,
	}
	var err error
	newReservation.StartTime, newReservation.EndTime, err = parseTimes(reservation.StartTime, reservation.EndTime)
	if err != nil {
		s.logger.Error(err)
		return "", err
	}

	id, err := s.repository.Reservation.CreateReservation(ctx, newReservation)
	if err != nil {
		s.logger.Error(err)
		return "", err
	}
	return id, nil
}

func (s *ReservationService) GetReservationsByRoomID(ctx context.Context, roomID string) ([]model.Reservation, error) {
	reservations, err := s.repository.Reservation.GetReservationsByRoomID(ctx, roomID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return reservations, nil
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

	now := time.Now().UTC()
	if startTime.Before(now) {
		return time.Time{}, time.Time{}, errors.New("start time must be in the future")
	}

	return startTime, endTime, nil
}
