package daomanger

import (
	"core/db"
	"core/db/dao"
	"core/models"
	"fmt"

	"github.com/gofrs/uuid"
)

type TeamDAO interface {
	GetAll() ([]*models.Team, error)
	GetByID(id string) (*models.Team, error)
	Create(req models.Team) (*models.Team, error)
	Update(req models.Team) (*models.Team, error)
}

type teamDao struct {
	*dao.Query
}

func NewTeamDAO() TeamDAO {
	return &teamDao{db.DbHandler}
}

func (c *teamDao) GetAll() ([]*models.Team, error) {
	return c.Team.Where().Find()
}

func (c *teamDao) GetByID(id string) (*models.Team, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	return c.Team.Where(c.Team.ID.Eq(&uid)).First()
}

func (c *teamDao) Create(req models.Team) (*models.Team, error) {
	err := c.Team.Create(&req)
	return &req, err
}

func (c *teamDao) Update(req models.Team) (*models.Team, error) {
	result, err := c.Team.Where(c.Team.ID.Eq(&req.ID)).Updates(&req)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("team %s not found", req.ID)
	}
	return &req, nil
}
