package service

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/model"
	"github.com/alibekabdrakhman1/reservation_kami/app/internal/repository"
	"github.com/alibekabdrakhman1/reservation_kami/app/pkg/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "qwerty"
	dbName     = "reservation"
)

var service *ReservationService

func TestMain(m *testing.M) {
	DB, err := setupDB()
	if err != nil {
		fmt.Printf("failed to connect to the database: %v", err)
		os.Exit(1)
	}
	defer DB.Close()

	repo := repository.NewManager(DB)
	l := zap.NewNop().Sugar()

	service = NewReservationService(repo, l)

	code := m.Run()

	os.Exit(code)
}

func setupDB() (*pgxpool.Pool, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	ctx := context.Background()
	DB, err := db.Dial(ctx, psqlInfo)
	if err != nil {
		return nil, err
	}

	err = DB.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return DB, nil
}

func TestCreateReservationSuccess(t *testing.T) {
	input := &model.InputReservation{
		RoomID:    "101",
		StartTime: time.Now().UTC().Add(1 * time.Hour).Format(time.RFC3339),
		EndTime:   time.Now().UTC().Add(2 * time.Hour).Format(time.RFC3339),
	}

	id, err := service.CreateReservation(context.Background(), input)
	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestCreateReservationConflict(t *testing.T) {
	input1 := &model.InputReservation{
		RoomID:    "101",
		StartTime: time.Now().UTC().Add(3 * time.Hour).Format(time.RFC3339),
		EndTime:   time.Now().UTC().Add(4 * time.Hour).Format(time.RFC3339),
	}
	_, err := service.CreateReservation(context.Background(), input1)
	assert.NoError(t, err)

	input2 := &model.InputReservation{
		RoomID:    "101",
		StartTime: time.Now().UTC().Add(3 * time.Hour).Format(time.RFC3339),
		EndTime:   time.Now().UTC().Add(5 * time.Hour).Format(time.RFC3339),
	}
	_, err = service.CreateReservation(context.Background(), input2)
	require.Error(t, err)
	assert.Equal(t, "reservation conflicts with an existing reservation", err.Error())
}

func TestConcurrentReservations(t *testing.T) {
	input := &model.InputReservation{
		RoomID:    "101",
		StartTime: time.Now().UTC().Add(6 * time.Hour).Format(time.RFC3339),
		EndTime:   time.Now().UTC().Add(7 * time.Hour).Format(time.RFC3339),
	}

	errChan := make(chan error, 2)

	go func() {
		_, err := service.CreateReservation(context.Background(), input)
		errChan <- err
	}()

	go func() {
		_, err := service.CreateReservation(context.Background(), input)
		errChan <- err
	}()

	for i := 0; i < 2; i++ {
		err := <-errChan
		if err != nil {
			assert.Contains(t, err.Error(), "reservation conflicts with an existing reservation")
		}
	}
}
