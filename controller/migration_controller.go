package controller

import (
	"bytes"
	"core/generated"
	"core/models"
	"core/provisioner"
	"encoding/csv"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

var (
	_ generated.MigrationAPIAPI = &migrationApiController{}
)

type migrationApiController struct{}

func NewMigrationApiController() *migrationApiController { return &migrationApiController{} }

// MigrationStartPost reads org/team info and CSV data
func (a *migrationApiController) MigrationStartPost(c *gin.Context) {
	// 1️⃣ Read org and team names
	orgName := c.PostForm("org_name")
	teamName := c.PostForm("team_name")

	if orgName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "org_name is required"})
		return
	}

	// 2️⃣ Read uploaded CSV file
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CSV file is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open uploaded CSV"})
		return
	}
	defer file.Close()

	// 3️⃣ Parse CSV
	records, err := parseCSV(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse CSV"})
		return
	}

	if len(records) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CSV file is empty"})
		return
	}

	// 4️⃣ Start migration process
	migProvisioner := provisioner.NewMigrationProvisioner()
	migrationId, err := migProvisioner.Migrate(orgName, teamName, records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"migration_id": migrationId, "message": "Migration started successfully. CSV parsing in progress.", "status": "initiated"})
}

// parseCSV reads all CSV rows
func parseCSV(file multipart.File) ([]models.DocumentCSVRecord, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(buf.Bytes()))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []models.DocumentCSVRecord
	for _, row := range rows {
		if len(row) < 4 {
			// require at least org_id, file_url
			continue
		}

		orgID, err := uuid.FromString(row[0])
		if err != nil {
			continue // skip invalid UUIDs
		}

		var teamID *uuid.UUID
		if len(row) > 1 && row[1] != "" {
			tid, err := uuid.FromString(row[1])
			if err == nil {
				teamID = &tid
			}
		}

		var dataSourceID *uuid.UUID
		if len(row) > 2 && row[2] != "" {
			dsID, err := uuid.FromString(row[2])
			if err == nil {
				dataSourceID = &dsID
			}
		}

		fileURL := row[3]

		fileName := ""
		if len(row) > 4 {
			fileName = row[4]
		}

		visibility := "team"
		if len(row) > 5 && row[5] != "" {
			visibility = row[5]
		}

		var createdBy *uuid.UUID
		if len(row) > 6 && row[6] != "" {
			cid, err := uuid.FromString(row[6])
			if err == nil {
				createdBy = &cid
			}
		}

		record := models.DocumentCSVRecord{
			OrgID:        orgID,
			TeamID:       teamID,
			DataSourceID: dataSourceID,
			FileURL:      fileURL,
			FileName:     fileName,
			Visibility:   visibility,
			CreatedBy:    createdBy,
		}

		records = append(records, record)
	}

	return records, nil
}
