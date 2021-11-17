package monitor

import "github.com/google/uuid"

type Device struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	User     string    `json:"user"`
	Password string    `json:"password"`
}

type MapDevice map[uuid.UUID]*Device
