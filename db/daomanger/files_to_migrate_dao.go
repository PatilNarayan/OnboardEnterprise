package daomanger

import (
	"core/db"
	"core/db/dao"
	"core/models"
	"fmt"

	"github.com/gofrs/uuid"
)

type FileToMigrateDAO interface {
	GetAll() ([]*models.FileToMigrate, error)
	GetByID(id string) (*models.FileToMigrate, error)
	Create(req models.FileToMigrate) (*models.FileToMigrate, error)
	Update(req models.FileToMigrate) (*models.FileToMigrate, error)
}

type fileToMigrateDao struct {
	*dao.Query
}

func NewFileToMigrateDAO() FileToMigrateDAO {
	return &fileToMigrateDao{db.DbHandler}
}

func (c *fileToMigrateDao) GetAll() ([]*models.FileToMigrate, error) {
	return c.FileToMigrate.Where().Find()
}

func (c *fileToMigrateDao) GetByID(id string) (*models.FileToMigrate, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	return c.FileToMigrate.Where(c.FileToMigrate.ID.Eq(&uid)).First()
}

func (c *fileToMigrateDao) Create(req models.FileToMigrate) (*models.FileToMigrate, error) {
	err := c.FileToMigrate.Create(&req)
	return &req, err
}

func (c *fileToMigrateDao) Update(req models.FileToMigrate) (*models.FileToMigrate, error) {
	result, err := c.FileToMigrate.Where(c.FileToMigrate.ID.Eq(&req.ID)).Updates(&req)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("file to migrate %s not found", req.ID)
	}
	return &req, nil
}
