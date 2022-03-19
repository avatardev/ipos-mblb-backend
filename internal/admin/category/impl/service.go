package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type CategoryServiceImpl struct {
	Cr *CategoryRepositoryImpl
}

func (c *CategoryServiceImpl) GetCategory(ctx context.Context) (dto.Categories, error) {
	categories := c.Cr.GetAll(ctx)
	if categories == nil {
		return nil, errors.ErrInvalidResources
	}

	return dto.MapToCategories(categories), nil
}

func (c *CategoryServiceImpl) GetCategoryById(ctx context.Context, id uint64) (*dto.Category, error) {
	category := c.Cr.GetById(ctx, id)
	if category == nil {
		return nil, errors.ErrInvalidResources
	}

	return dto.MapToCategory(*category), nil
}

// TODO make POST, UPDATE, DELETE Routes
