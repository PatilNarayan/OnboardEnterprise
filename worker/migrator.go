package migration

import (
	"context"
	"fmt"
	"log"
	"time"

	"core/db/daomanger"
	"core/models"

	"github.com/gofrs/uuid"
)

// Migrate handles file migration from CSV data
func Migrate(orgName string, teamName string, records []models.DocumentCSVRecord) error {
	if orgName == "" {
		return fmt.Errorf("org name cannot be empty")
	}
	if len(records) == 0 {
		return fmt.Errorf("no records to migrate")
	}

	orgDAO := daomanger.NewOrgDAO()
	fileDAO := daomanger.NewFileToMigrateDAO()

	// 1️⃣ Get the organization by name
	orgs, err := orgDAO.GetAll()
	if err != nil {
		return fmt.Errorf("failed to get orgs: %v", err)
	}

	var orgID uuid.UUID
	found := false
	for _, org := range orgs {
		if org.Name == orgName {
			orgID = org.ID
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("organization %q not found", orgName)
	}

	// 2️⃣ Create a migration record in FileToMigrate
	for _, record := range records {

		migrationID := uuid.Must(uuid.NewV4())
		fileToMigrate := models.FileToMigrate{
			ID:           migrationID,
			OrgID:        orgID,
			TeamID:       record.TeamID,
			DataSourceID: record.DataSourceID,
			FileURL:      record.FileURL,
			FileName:     record.FileName,
			Visibility:   record.Visibility,
			CreatedBy:    record.CreatedBy,
			Status:       "Pending",
		}

		if _, err := fileDAO.Create(fileToMigrate); err != nil {
			return fmt.Errorf("failed to create migration record: %v", err)
		}

		log.Printf("✅ Migration %s created for org %s (records: %d)", migrationID, orgName, len(records))

		// 3️⃣ Start background worker
		go startWorker(context.Background(), migrationID, records)
	}
	return nil
}

// startWorker processes records asynchronously
func startWorker(ctx context.Context, migrationID uuid.UUID, records []models.DocumentCSVRecord) {
	fileDAO := daomanger.NewFileToMigrateDAO()

	log.Printf("🚀 Starting migration worker for ID: %s", migrationID)

	for i, row := range records {
		select {
		case <-ctx.Done():
			log.Printf("⚠️ Migration %s cancelled", migrationID)
			return
		default:
			// Simulate processing; here you could insert into Document table
			log.Printf("Processing row %d: %v", i+1, row)
			time.Sleep(100 * time.Millisecond) // simulate work
		}
	}

	// 4️⃣ Update migration status to COMPLETED
	update := models.FileToMigrate{
		ID:        migrationID,
		Status:    "COMPLETED",
		UpdatedAt: time.Now(),
	}
	if _, err := fileDAO.Update(update); err != nil {
		log.Printf("❌ Failed to update migration status: %v", err)
		return
	}

	log.Printf("✅ Migration %s completed successfully", migrationID)
}
