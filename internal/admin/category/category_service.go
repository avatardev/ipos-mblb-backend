package category

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type CategoryService interface {
	GetCategory(ctx context.Context) (dto.CategoriesResponse, error)
	GetCategoryById(ctx context.Context, id int64) (*dto.CategoryResponse, error)
	StoreCategory(ctx context.Context, req *dto.CategoryRequest) (*dto.CategoryResponse, error)
	UpdateCategory(ctx context.Context, id int64, req *dto.CategoryRequest) (*dto.CategoryResponse, error)
	DeleteCategory(ctx context.Context, id int64) error
}

func NewCategoryService(Cr *impl.CategoryRepositoryImpl) CategoryService {
	return &impl.CategoryServiceImpl{Cr: Cr}
}
