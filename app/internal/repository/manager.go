package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"reservation/app/internal/repository/postgre"
)

type Manager struct {
	Reservation IReservationRepository
}

func NewManager(db *pgxpool.Pool) *Manager {
	return &Manager{
		Reservation: postgre.NewReservationRepository(db),
	}
}
