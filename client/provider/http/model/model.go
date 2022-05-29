package model

import "github.com/google/uuid"

type LoginRequest struct {
	Login      string    `json:"login"`
	MasterHash string    `json:"master_hash"`
	Password   string    `json:"password"`
	DeviceID   uuid.UUID `json:"device_id"`
}

type SecretRequest struct {
	Data string    `json:"data,omitempty"`
	ID   uuid.UUID `json:"id,omitempty"`
	Ver  int       `json:"ver,omitempty"`
}

func (s *SecretRequest) IsValidResponseUpload() bool {
	return s.ID != uuid.Nil && s.Ver > 0
}

func (s *SecretRequest) IsValidResponseDownload() bool {
	return s.ID != uuid.Nil && s.Ver > 0 && len(s.Data) > 0
}

func (s *SecretRequest) IsValidUpload() bool {
	if s.ID == uuid.Nil && s.Ver > 1 {
		return false
	}
	if s.Ver == 0 {
		return false
	}
	if len(s.Data) == 0 {
		return false
	}
	return true
}
func (s *SecretRequest) IsValidDelete() bool {
	if s.ID == uuid.Nil {
		return false
	}
	if len(s.Data) != 0 || s.Ver > 0 {
		return false
	}
	return true
}

func (s *SecretRequest) IsValidDownload() bool {
	if s.ID == uuid.Nil {
		return false
	}
	if len(s.Data) != 0 || s.Ver > 0 {
		return false
	}
	return true
}

type SyncResponse struct {
	List map[uuid.UUID]int `json:"list"`
}
