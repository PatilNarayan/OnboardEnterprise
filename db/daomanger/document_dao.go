package daomanger

import (
	"core/db"
	"core/db/dao"
	"core/models"
	"fmt"

	"github.com/gofrs/uuid"
)

type DocumentDAO interface {
	GetAll() ([]*models.Document, error)
	GetByID(id string) (*models.Document, error)
	Create(req models.Document) (*models.Document, error)
	Update(req models.Document) (*models.Document, error)
}

type documentDao struct {
	*dao.Query
}

func NewDocumentDAO() DocumentDAO {
	return &documentDao{db.DbHandler}
}

func (c *documentDao) GetAll() ([]*models.Document, error) {
	return c.Document.Where().Find()
}

func (c *documentDao) GetByID(id string) (*models.Document, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	return c.Document.Where(c.Document.ID.Eq(&uid)).First()
}

func (c *documentDao) Create(req models.Document) (*models.Document, error) {
	err := c.Document.Create(&req)
	return &req, err
}

func (c *documentDao) Update(req models.Document) (*models.Document, error) {
	result, err := c.Document.Where(c.Document.ID.Eq(&req.ID)).Updates(&req)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("document %s not found", req.ID)
	}
	return &req, nil
}
