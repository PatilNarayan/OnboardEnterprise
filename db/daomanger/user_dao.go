package daomanger

import (
	"core/db"
	"core/db/dao"
	"core/models"
	"fmt"

	"github.com/gofrs/uuid"
)

type UserDAO interface {
	GetAll() ([]*models.User, error)
	GetByID(id string) (*models.User, error)
	Create(req models.User) (*models.User, error)
	Update(req models.User) (*models.User, error)
}

type userDao struct {
	*dao.Query
}

func NewUserDAO() UserDAO {
	return &userDao{db.DbHandler}
}

func (c *userDao) GetAll() ([]*models.User, error) {
	return c.User.Where().Find()
}

func (c *userDao) GetByID(id string) (*models.User, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	return c.User.Where(c.User.ID.Eq(&uid)).First()
}

func (c *userDao) Create(req models.User) (*models.User, error) {
	err := c.User.Create(&req)
	return &req, err
}

func (c *userDao) Update(req models.User) (*models.User, error) {
	result, err := c.User.Where(c.User.ID.Eq(&req.ID)).Updates(&req)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user %s not found", req.ID)
	}
	return &req, nil
}
