package model

import "github.com/google/uuid"

type LoginRequest struct {
	Login      string    `json:"login"`
	MasterHash string    `json:"master_hash"`
	Password   string    `json:"password"`
	DeviceID   uuid.UUID `json:"device_id"`
}
