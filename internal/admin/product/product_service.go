package product

import (
	"bytes"
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type ProductService interface {
	GetProduct(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.ProductsResponse, error)
	GetProductById(ctx context.Context, id int64) (*dto.ProductResponse, error)
	StoreProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, id int64, req *dto.ProductRequest) (*dto.ProductResponse, error)
	EditProductImage(ctx context.Context, id int64, img *bytes.Buffer, fName string) (res *dto.ProductResponse, err error)
	DeleteProduct(ctx context.Context, id int64) error
}

func NewProductService(Pr impl.ProductRepositoryImpl) ProductService {
	return &impl.ProductServiceImpl{Pr: Pr}
}
