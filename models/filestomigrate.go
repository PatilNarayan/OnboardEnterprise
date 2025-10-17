package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type FileToMigrate struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrgID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"org_id"`
	TeamID       *uuid.UUID `gorm:"type:uuid;index" json:"team_id,omitempty"`
	DataSourceID *uuid.UUID `gorm:"type:uuid;index" json:"data_source_id,omitempty"`
	FileURL      string     `gorm:"not null" json:"file_url"`
	FileName     string     `json:"file_name"`
	Status       string     `gorm:"type:text;default:'pending'" json:"status"` // pending, processing, done
	Visibility   string     `gorm:"type:text;default:'team'" json:"visibility"`
	CreatedBy    *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type Document struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrgID         uuid.UUID  `gorm:"type:uuid;not null;index" json:"org_id"`
	TeamID        *uuid.UUID `gorm:"type:uuid;index" json:"team_id,omitempty"`
	SourceFileID  *uuid.UUID `gorm:"type:uuid;index" json:"source_file_id,omitempty"`
	Title         string     `json:"title"`
	MimeType      string     `json:"mime_type"`
	SizeBytes     int64      `json:"size_bytes"`
	StoragePath   string     `json:"storage_path"`
	ExtractedText string     `json:"extracted_text"`
	Summary       string     `json:"summary"`
	Keywords      []string   `gorm:"type:text[]" json:"keywords"`
	Entities      string     `gorm:"type:jsonb" json:"entities"`
	Visibility    string     `gorm:"type:text;default:'team'" json:"visibility"`
	AccessControl string     `gorm:"type:jsonb;default:'{}'" json:"access_control"` // allowed_teams, allowed_users
	Indexed       bool       `gorm:"default:false" json:"indexed"`
	CreatedBy     *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type MigrateJobs struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrgID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"org_id"`
	TeamID    *uuid.UUID `gorm:"type:uuid;index" json:"team_id,omitempty"`
	Status    string     `gorm:"type:text;default:'pending'" json:"status"` // pending, processing, done
	CreatedBy *uuid.UUID `gorm:"type:uuid;index" json:"created_by,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
