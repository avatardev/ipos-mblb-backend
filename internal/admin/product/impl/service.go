package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type ProductServiceImpl struct {
	Pr ProductRepositoryImpl
}

func (p *ProductServiceImpl) GetProduct(ctx context.Context, query string, limit uint64, offset uint64) (*dto.ProductsResponse, error) {
	productCount, err := p.Pr.Count(ctx, query)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if productCount == 0 {
		return nil, errors.ErrInvalidResources
	}

	products, err := p.Pr.GetAll(ctx, query, limit, offset)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if len(products) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewProductsResponse(products, limit, offset, productCount), nil
}

func (p *ProductServiceImpl) GetProductById(ctx context.Context, id int64) (*dto.ProductResponse, error) {
	product, err := p.Pr.GetById(ctx, id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if product == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewProductResponse(*product), nil
}

func (p *ProductServiceImpl) StoreProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error) {
	product := req.ToEntity()

	data, err := p.Pr.Store(ctx, product)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewProductResponse(*data), nil
}

func (p *ProductServiceImpl) UpdateProduct(ctx context.Context, id int64, req *dto.ProductRequest) (*dto.ProductResponse, error) {
	product := req.ToEntity()
	product.Id = id

	exists, err := p.Pr.GetById(ctx, id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	data, err := p.Pr.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewProductResponse(*data), nil
}

func (p *ProductServiceImpl) DeleteProduct(ctx context.Context, id int64) error {
	exists, err := p.Pr.GetById(ctx, id)
	if err != nil {
		return errors.ErrUnknown
	}

	if exists == nil {
		return errors.ErrNotFound
	}

	if err := p.Pr.Delete(ctx, id); err != nil {
		return errors.ErrUnknown
	}

	return nil
}
