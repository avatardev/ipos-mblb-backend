package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type ProductCategoryServiceImpl struct {
	Cr ProductCategoryRepositoryImpl
}

func (c *ProductCategoryServiceImpl) GetCategory(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.ProductCategoriesResponse, error) {
	categoryCount, err := c.Cr.Count(ctx, keyword)
	if err != nil {
		return nil, err
	}

	if categoryCount == 0 {
		return nil, errors.ErrInvalidResources
	}

	if limit == 0 {
		limit = categoryCount
	}

	categories, err := c.Cr.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewProductCategoriesResponse(categories, limit, offset, categoryCount), nil
}

func (c *ProductCategoryServiceImpl) GetCategoryById(ctx context.Context, id int64) (*dto.ProductCategoryResponse, error) {
	category, err := c.Cr.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewProductCategoryResponse(*category), nil
}

func (c *ProductCategoryServiceImpl) StoreCategory(ctx context.Context, req *dto.ProductCategoryRequest) (*dto.ProductCategoryResponse, error) {
	category := req.ToEntity()
	data, err := c.Cr.Store(ctx, category)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewProductCategoryResponse(*data), nil
}

func (c *ProductCategoryServiceImpl) UpdateCategory(ctx context.Context, id int64, req *dto.ProductCategoryRequest) (*dto.ProductCategoryResponse, error) {
	category := req.ToEntity()
	category.Id = id

	exists, err := c.Cr.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	data, err := c.Cr.Update(ctx, category)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewProductCategoryResponse(*data), nil
}

func (c *ProductCategoryServiceImpl) DeleteCategory(ctx context.Context, id int64) error {
	exists, err := c.Cr.GetById(ctx, id)
	if err != nil {
		return err
	}

	if exists == nil {
		return errors.ErrNotFound
	}

	if err := c.Cr.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
