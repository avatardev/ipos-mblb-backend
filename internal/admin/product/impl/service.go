package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type ProductServiceImpl struct {
	Pr ProductRepositoryImpl
}

func (p *ProductServiceImpl) GetProduct(ctx context.Context) (dto.ProductsResponse, error) {
	products, err := p.Pr.GetAll(ctx)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if products == nil {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewProductsResponse(products), nil
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
