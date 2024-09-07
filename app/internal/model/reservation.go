package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Reservation struct {
	ID        uuid.UUID `json:"id"`
	RoomID    string    `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type InputReservation struct {
	RoomID    string `json:"room_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

var (
	ErrReservationConflict = errors.New("reservation conflicts with an existing reservation")
)
