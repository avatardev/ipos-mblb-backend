package productcategory

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/product_category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type ProductCategoryService interface {
	GetCategory(ctx context.Context, limit uint64, offset uint64) (*dto.ProductCategoriesResponse, error)
	GetCategoryById(ctx context.Context, id int64) (*dto.ProductCategoryResponse, error)
	StoreCategory(ctx context.Context, req *dto.ProductCategoryRequest) (*dto.ProductCategoryResponse, error)
	UpdateCategory(ctx context.Context, id int64, req *dto.ProductCategoryRequest) (*dto.ProductCategoryResponse, error)
	DeleteCategory(ctx context.Context, id int64) error
}

func NewProductCategoryService(Cr impl.ProductCategoryRepositoryImpl) ProductCategoryService {
	return &impl.ProductCategoryServiceImpl{Cr: Cr}
}
