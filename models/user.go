package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrgID       uuid.UUID `gorm:"type:uuid;not null;index" json:"org_id"`
	Email       string    `gorm:"unique;not null" json:"email"`
	DisplayName string    `json:"display_name"`
	Role        string    `gorm:"type:text;default:'member'" json:"role"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type TeamMember struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TeamID    uuid.UUID `gorm:"type:uuid;not null;index" json:"team_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type AccessControl struct {
	AllowedTeams []uuid.UUID `json:"allowed_teams"`
	AllowedUsers []uuid.UUID `json:"allowed_users"`
	DeniedUsers  []uuid.UUID `json:"denied_users"`
}
