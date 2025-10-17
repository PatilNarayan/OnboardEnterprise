package daomanger

import (
	"core/db"
	"core/db/dao"
	"core/models"
	"fmt"

	"github.com/gofrs/uuid"
)

type TeamMemberDAO interface {
	GetAll() ([]*models.TeamMember, error)
	GetByID(id string) (*models.TeamMember, error)
	Create(req models.TeamMember) (*models.TeamMember, error)
	Update(req models.TeamMember) (*models.TeamMember, error)
}

type teamMemberDao struct {
	*dao.Query
}

func NewTeamMemberDAO() TeamMemberDAO {
	return &teamMemberDao{db.DbHandler}
}

func (c *teamMemberDao) GetAll() ([]*models.TeamMember, error) {
	return c.TeamMember.Where().Find()
}

func (c *teamMemberDao) GetByID(id string) (*models.TeamMember, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	return c.TeamMember.Where(c.TeamMember.ID.Eq(&uid)).First()
}

func (c *teamMemberDao) Create(req models.TeamMember) (*models.TeamMember, error) {
	err := c.TeamMember.Create(&req)
	return &req, err
}

func (c *teamMemberDao) Update(req models.TeamMember) (*models.TeamMember, error) {
	result, err := c.TeamMember.Where(c.TeamMember.ID.Eq(&req.ID)).Updates(&req)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("team member %s not found", req.ID)
	}
	return &req, nil
}
