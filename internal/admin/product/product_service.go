package product

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type ProductService interface {
	GetProduct(ctx context.Context) (dto.ProductsResponse, error)
}

func NewProductService(Pr impl.ProductRepositoryImpl) ProductService {
	return &impl.ProductServiceImpl{Pr: Pr}
}
