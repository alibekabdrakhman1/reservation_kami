package postgre

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"reservation/app/internal/model"
)

type ReservationRepository struct {
	DB *pgxpool.Pool
}

func NewReservationRepository(db *pgxpool.Pool) *ReservationRepository {
	return &ReservationRepository{
		DB: db,
	}
}

func (r *ReservationRepository) CreateReservation(ctx context.Context, reservation *model.Reservation) (string, error) {
	var count int
	query := `
        SELECT COUNT(*)
        FROM reservations
        WHERE room_id = $1 AND (
            (start_time < $3 AND end_time > $2) OR
            (start_time < $2 AND end_time > $2) OR
            (start_time < $3 AND end_time > $3)
        )
    `

	err := r.DB.QueryRow(ctx, query, reservation.RoomID, reservation.StartTime, reservation.EndTime).Scan(&count)
	if err != nil {
		return "", err
	}

	if count > 0 {
		return "", model.ErrReservationConflict
	}

	var id string
	query = `
		INSERT INTO reservations (room_id, start_time, end_time)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err = r.DB.QueryRow(ctx, query, reservation.RoomID, reservation.StartTime, reservation.EndTime).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *ReservationRepository) GetReservationsByRoomID(ctx context.Context, roomID string) ([]model.Reservation, error) {
	query := `
		SELECT * FROM reservations
		WHERE room_id = $1
		ORDER BY start_time
	`

	rows, err := r.DB.Query(ctx, query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []model.Reservation
	for rows.Next() {
		var res model.Reservation
		if err := rows.Scan(&res.ID, &res.RoomID, &res.StartTime, &res.EndTime); err != nil {
			return nil, err
		}
		reservations = append(reservations, res)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}
