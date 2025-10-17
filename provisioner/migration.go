package provisioner

import (
	"core/db/daomanger"
	"core/models"
	migration "core/worker"

	uuid "github.com/gofrs/uuid"
)

type MigrationProvisioner struct{}
type MigrationProvisionerInterface interface {
	Migrate(orgName string, teamName string, records []models.DocumentCSVRecord)
}

func NewMigrationProvisioner() *MigrationProvisioner {
	return &MigrationProvisioner{}
}

func (p *MigrationProvisioner) Migrate(orgName string, teamName string, records []models.DocumentCSVRecord) (*uuid.UUID, error) {

	// 1️⃣ Read org and team names present or not than insert
	org, err := daomanger.NewOrgDAO().Create(models.Organization{Name: orgName})
	if err != nil {
		return nil, err
	}
	team, err := daomanger.NewTeamDAO().Create(models.Team{OrgID: org.ID, Name: teamName})
	if err != nil {
		return nil, err
	}

	migrationData, err := daomanger.NewMigrateJobssDAO().Create(models.MigrateJobs{
		OrgID:  org.ID,
		TeamID: &team.ID,
		Status: "initiated",
	})
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		record.TeamID = &team.ID
		record.OrgID = org.ID
	}

	go migration.Migrate(orgName, teamName, records)

	return &migrationData.ID, nil

}
