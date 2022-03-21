package product

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type ProductService interface {
	GetProduct(ctx context.Context) (dto.ProductsResponse, error)
	GetProductById(ctx context.Context, id int64) (*dto.ProductResponse, error)
	StoreProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error)
}

func NewProductService(Pr impl.ProductRepositoryImpl) ProductService {
	return &impl.ProductServiceImpl{Pr: Pr}
}
