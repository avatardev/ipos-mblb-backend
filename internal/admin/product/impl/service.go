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
