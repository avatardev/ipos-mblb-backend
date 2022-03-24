package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type BuyerCategoryServiceImpl struct {
	Cr BuyerCategoryRepositoryImpl
}

func (c *BuyerCategoryServiceImpl) GetCategory(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.BuyerCategoriesResponse, error) {
	categoryCount, err := c.Cr.Count(ctx, keyword)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if categoryCount == 0 {
		return nil, errors.ErrInvalidResources
	}

	categories, err := c.Cr.GetAll(ctx, keyword, limit, offset)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if len(categories) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewBuyerCategoriesResponse(categories, limit, offset, categoryCount), nil
}

func (c *BuyerCategoryServiceImpl) GetCategoryById(ctx context.Context, id int64) (*dto.BuyerCategoryResponse, error) {
	category, err := c.Cr.GetById(ctx, id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if category == nil {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewBuyerCategoryResponse(category), nil
}

func (c *BuyerCategoryServiceImpl) StoreCategory(ctx context.Context, req *dto.BuyerCategoryRequest) (*dto.BuyerCategoryResponse, error) {
	category := req.ToEntity()

	data, err := c.Cr.Store(ctx, category)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewBuyerCategoryResponse(data), nil
}

func (c *BuyerCategoryServiceImpl) UpdateCategory(ctx context.Context, id int64, req *dto.BuyerCategoryRequest) (*dto.BuyerCategoryResponse, error) {
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
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewBuyerCategoryResponse(data), nil
}

func (c *BuyerCategoryServiceImpl) DeleteCategory(ctx context.Context, id int64) error {
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
