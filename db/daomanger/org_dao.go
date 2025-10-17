package daomanger

import (
	"core/db"
	"core/db/dao"
	"core/models"
	"fmt"

	"github.com/gofrs/uuid"
)

type OrgDAO interface {
	GetAll() ([]*models.Organization, error)
	GetByID(id string) (*models.Organization, error)
	Create(req models.Organization) (*models.Organization, error)
	Update(req models.Organization) (*models.Organization, error)
}

type orgDao struct {
	*dao.Query
}

func NewOrgDAO() OrgDAO {
	return &orgDao{db.DbHandler}
}

func (c *orgDao) GetAll() ([]*models.Organization, error) {
	orgs, err := c.Organization.Where().Find()
	if err != nil {
		return nil, err
	}
	if orgs == nil {
		return nil, fmt.Errorf("no orgs found")
	}
	return orgs, nil
}

func (c *orgDao) GetByID(id string) (*models.Organization, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	return c.Organization.Where(c.Organization.ID.Eq(&uid)).First()
}

func (c *orgDao) Create(req models.Organization) (*models.Organization, error) {
	err := c.Organization.Create(&req)
	return &req, err
}

func (c *orgDao) Update(req models.Organization) (*models.Organization, error) {
	result, err := c.Organization.Where(c.Organization.ID.Eq(&req.ID)).Updates(&req)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("organization %s not found", req.ID)
	}
	return &req, nil
}
