package categories

import (
	"github.com/dhucsik/bookers/internal/models"
	"github.com/samber/lo"
)

type categoryModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type categoryModels []*categoryModel

func (m *categoryModel) convert() *models.Category {
	return &models.Category{
		ID:   m.ID,
		Name: m.Name,
	}
}

func convertCategory(category *models.Category) *categoryModel {
	return &categoryModel{
		ID:   category.ID,
		Name: category.Name,
	}
}

func (m categoryModels) convert() []*models.Category {
	return lo.Map(m, func(model *categoryModel, _ int) *models.Category {
		return model.convert()
	})
}
