package model

import "github.com/google/uuid"

type Anggota struct {
	UserID uuid.UUID `json:"user_id"`
	RuangID uuid.UUID `json:"ruang_id"`
}