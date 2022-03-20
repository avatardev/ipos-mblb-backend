package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type CategoryServiceImpl struct {
	Cr *CategoryRepositoryImpl
}

func (c *CategoryServiceImpl) GetCategory(ctx context.Context) (dto.CategoriesResponse, error) {
	categories := c.Cr.GetAll(ctx)
	if categories == nil {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewCategoriesResponse(categories), nil
}

func (c *CategoryServiceImpl) GetCategoryById(ctx context.Context, id int64) (*dto.CategoryResponse, error) {
	category := c.Cr.GetById(ctx, id)
	if category == nil {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewCategoryReponse(*category), nil
}

func (c *CategoryServiceImpl) StoreCategory(ctx context.Context, req *dto.CategoryRequest) (*dto.CategoryResponse, error) {
	category := req.ToEntity()
	data := c.Cr.Store(ctx, category)
	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewCategoryReponse(*data), nil
}

// TODO make POST, UPDATE, DELETE Routes
