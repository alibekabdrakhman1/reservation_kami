package http

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"reservation/app/internal/model"
	"reservation/app/internal/service"
	"reservation/app/pkg/response"
)

type Reservation struct {
	service *service.Manager
	logger  *zap.SugaredLogger
}

func NewReservationHandler(service *service.Manager, logger *zap.SugaredLogger) *Reservation {
	return &Reservation{
		service: service,
		logger:  logger,
	}
}

func (h *Reservation) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var req *model.InputReservation

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.RespondWithCustomResponse(w, response.ErrorHeader("Invalid request payload", http.StatusBadRequest), nil)
		return
	}

	id, err := h.service.Reservation.CreateReservation(r.Context(), req)
	if err != nil {
		if errors.Is(err, model.ErrReservationConflict) {
			response.RespondWithCustomResponse(w, response.ErrorHeader(err.Error(), http.StatusConflict), nil)
			return
		}

		response.RespondWithCustomResponse(w, response.ErrorHeader("Failed to create reservation", http.StatusInternalServerError), nil)
		return
	}

	response.RespondWithCustomResponse(w, response.SuccessHeader, id)
}

func (h *Reservation) GetReservationsByRoomID(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "room_id")

	reservations, err := h.service.Reservation.GetReservationsByRoomID(r.Context(), roomID)
	if err != nil {
		response.RespondWithCustomResponse(w, response.ErrorHeader("error from get reservations by room ID", http.StatusInternalServerError), nil)
		return
	}

	response.RespondWithCustomResponse(w, response.SuccessHeader, reservations)
}
