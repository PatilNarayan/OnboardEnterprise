package models

import uuid "github.com/gofrs/uuid"

type DocumentCSVRecord struct {
	OrgID        uuid.UUID
	TeamID       *uuid.UUID
	DataSourceID *uuid.UUID
	FileURL      string
	FileName     string
	Visibility   string
	CreatedBy    *uuid.UUID
}
