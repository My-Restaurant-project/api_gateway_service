package models

type CheckReservationFilter struct {
	ReservationID string `json:"reservation_id string,omitempty"`
	UserID        string `json:"user_id string,omitempty"`
}
