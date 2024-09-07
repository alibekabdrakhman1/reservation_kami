package http

import "net/http"

type IReservationHandler interface {
	CreateReservation(w http.ResponseWriter, r *http.Request)
	GetReservationsByRoomID(w http.ResponseWriter, r *http.Request)
}
