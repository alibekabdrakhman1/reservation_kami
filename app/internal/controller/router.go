package controller

import "github.com/go-chi/chi"

func (s *Server) SetupRoutes() {
	s.App.Route("/reservations", func(r chi.Router) {
		r.Post("/", s.handler.Reservation.CreateReservation)
		r.Get("/{room_id}", s.handler.Reservation.GetReservationsByRoomID)
	})
}
