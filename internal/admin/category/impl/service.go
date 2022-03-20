package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type CategoryServiceImpl struct {
	Cr CategoryRepositoryImpl
}

func (c *CategoryServiceImpl) GetCategory(ctx context.Context) (dto.CategoriesResponse, error) {
	categories, err := c.Cr.GetAll(ctx)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if categories == nil {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewCategoriesResponse(categories), nil
}

func (c *CategoryServiceImpl) GetCategoryById(ctx context.Context, id int64) (*dto.CategoryResponse, error) {
	category, err := c.Cr.GetById(ctx, id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if category == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewCategoryReponse(*category), nil
}

func (c *CategoryServiceImpl) StoreCategory(ctx context.Context, req *dto.CategoryRequest) (*dto.CategoryResponse, error) {
	category := req.ToEntity()
	data, err := c.Cr.Store(ctx, category)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewCategoryReponse(*data), nil
}

func (c *CategoryServiceImpl) UpdateCategory(ctx context.Context, id int64, req *dto.CategoryRequest) (*dto.CategoryResponse, error) {
	category := req.ToEntity()
	category.Id = id

	exists, err := c.Cr.GetById(ctx, id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	data, err := c.Cr.Update(ctx, category)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewCategoryReponse(*data), nil
}

func (c *CategoryServiceImpl) DeleteCategory(ctx context.Context, id int64) error {
	exists, err := c.Cr.GetById(ctx, id)
	if err != nil {
		return errors.ErrUnknown
	}

	if exists == nil {
		return errors.ErrNotFound
	}

	if err := c.Cr.Delete(ctx, id); err != nil {
		return errors.ErrUnknown
	}

	return nil
}
