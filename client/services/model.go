package services

import "github.com/google/uuid"

var (
	SyncActions = map[string]int{
		"UPLOAD_NEW":     1,
		"UPLOAD":         2,
		"DOWNLOAD":       3,
		"DELETE_LOCALLY": 4,
		"SEND_DELETE":    5,
		"COLLISION":      6,
	}
)

type SyncTask struct {
	LocID    int64
	SecretId uuid.UUID
	Ver      int
	ActionID int
}
