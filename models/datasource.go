package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type DataSource struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrgID       uuid.UUID  `gorm:"type:uuid;not null;index" json:"org_id"`
	TeamID      *uuid.UUID `gorm:"type:uuid;index" json:"team_id,omitempty"`
	Type        string     `gorm:"not null" json:"type"` // google_drive, onedrive, s3, url
	Name        string     `json:"name"`
	Credentials string     `gorm:"type:jsonb" json:"credentials"`
	Visibility  string     `gorm:"type:text;default:'team'" json:"visibility"` // org, team, private
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
}
