package authors

import (
	"github.com/dhucsik/bookers/internal/models"
	"github.com/samber/lo"
)

type authorModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type authorModels []*authorModel

func (m *authorModel) convert() *models.Author {
	return &models.Author{
		ID:   m.ID,
		Name: m.Name,
	}
}

func convertAuthor(author *models.Author) *authorModel {
	return &authorModel{
		ID:   author.ID,
		Name: author.Name,
	}
}

func (m authorModels) convert() []*models.Author {
	return lo.Map(m, func(model *authorModel, _ int) *models.Author {
		return model.convert()
	})
}
