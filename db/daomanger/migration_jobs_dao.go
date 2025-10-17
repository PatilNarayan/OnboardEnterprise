package daomanger

import (
	"core/db"
	"core/db/dao"
	"core/models"
	"fmt"

	"github.com/gofrs/uuid"
)

// MigrateJobssDAO defines CRUD operations for MigrateJobss table
type MigrateJobssDAO interface {
	GetAll() ([]*models.MigrateJobs, error)
	GetByID(id uuid.UUID) (*models.MigrateJobs, error)
	Create(job models.MigrateJobs) (*models.MigrateJobs, error)
	Update(job models.MigrateJobs) (*models.MigrateJobs, error)
	Delete(id uuid.UUID) error
}

// MigrateJobsDao implements MigrateJobssDAO
type MigrateJobsDao struct {
	*dao.Query
}

// NewMigrateJobssDAO returns a new DAO instance
func NewMigrateJobssDAO() MigrateJobssDAO {
	return &MigrateJobsDao{db.DbHandler}
}

// GetAll returns all migrate jobs
func (c *MigrateJobsDao) GetAll() ([]*models.MigrateJobs, error) {
	jobs, err := c.MigrateJobs.Where().Find()
	if err != nil {
		return nil, err
	}
	if jobs == nil {
		return nil, fmt.Errorf("no migrate jobs found")
	}
	return jobs, nil
}

// GetByID returns a migrate job by ID
func (c *MigrateJobsDao) GetByID(id uuid.UUID) (*models.MigrateJobs, error) {
	return c.MigrateJobs.Where(c.MigrateJobs.ID.Eq(&id)).First()
}

// Create inserts a new migrate job
func (c *MigrateJobsDao) Create(job models.MigrateJobs) (*models.MigrateJobs, error) {
	err := c.MigrateJobs.Create(&job)
	return &job, err
}

// Update updates an existing migrate job
func (c *MigrateJobsDao) Update(job models.MigrateJobs) (*models.MigrateJobs, error) {
	result, err := c.MigrateJobs.Where(c.MigrateJobs.ID.Eq(&job.ID)).Updates(&job)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("migrate job %s not found", job.ID)
	}
	return &job, nil
}

// Delete deletes a migrate job by ID
func (c *MigrateJobsDao) Delete(id uuid.UUID) error {
	result, err := c.MigrateJobs.Where(c.MigrateJobs.ID.Eq(&id)).Delete()
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("migrate job %s not found", id)
	}
	return nil
}
