package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type Organization struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Teams []Team `gorm:"foreignKey:OrgID" json:"teams,omitempty"`
}
