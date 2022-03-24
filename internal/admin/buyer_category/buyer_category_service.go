package buyercategory

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer_category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type BuyerCategoryService interface {
	GetCategory(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.BuyerCategoriesResponse, error)
	GetCategoryById(ctx context.Context, id int64) (*dto.BuyerCategoryResponse, error)
	StoreCategory(ctx context.Context, req *dto.BuyerCategoryRequest) (*dto.BuyerCategoryResponse, error)
	UpdateCategory(ctx context.Context, id int64, req *dto.BuyerCategoryRequest) (*dto.BuyerCategoryResponse, error)
	DeleteCategory(ctx context.Context, id int64) error
}

func NewBuyerCategoryService(Cr impl.BuyerCategoryRepositoryImpl) BuyerCategoryService {
	return &impl.BuyerCategoryServiceImpl{Cr: Cr}
}
