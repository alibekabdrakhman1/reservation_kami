package repository

import (
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/repository/postgre"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Manager struct {
	Reservation IReservationRepository
}

func NewManager(db *pgxpool.Pool) *Manager {
	return &Manager{
		Reservation: postgre.NewReservationRepository(db),
	}
}
