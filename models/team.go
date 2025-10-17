package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type Team struct {
	ID           uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrgID        uuid.UUID    `gorm:"type:uuid;not null;index" json:"org_id"`
	Name         string       `gorm:"not null" json:"name"`
	Description  string       `json:"description"`
	CreatedAt    time.Time    `gorm:"autoCreateTime" json:"created_at"`
	Organization Organization `gorm:"foreignKey:OrgID" json:"organization,omitempty"`
}
