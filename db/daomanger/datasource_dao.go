package daomanger

import (
	"core/db"
	"core/db/dao"
	"core/models"
	"fmt"

	"github.com/gofrs/uuid"
)

type DataSourceDAO interface {
	GetAll() ([]*models.DataSource, error)
	GetByID(id string) (*models.DataSource, error)
	Create(req models.DataSource) (*models.DataSource, error)
	Update(req models.DataSource) (*models.DataSource, error)
}

type dataSourceDao struct {
	*dao.Query
}

func NewDataSourceDAO() DataSourceDAO {
	return &dataSourceDao{db.DbHandler}
}

func (c *dataSourceDao) GetAll() ([]*models.DataSource, error) {
	return c.DataSource.Where().Find()
}

func (c *dataSourceDao) GetByID(id string) (*models.DataSource, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	return c.DataSource.Where(c.DataSource.ID.Eq(&uid)).First()
}

func (c *dataSourceDao) Create(req models.DataSource) (*models.DataSource, error) {
	err := c.DataSource.Create(&req)
	return &req, err
}

func (c *dataSourceDao) Update(req models.DataSource) (*models.DataSource, error) {
	result, err := c.DataSource.Where(c.DataSource.ID.Eq(&req.ID)).Updates(&req)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("data source %s not found", req.ID)
	}
	return &req, nil
}
