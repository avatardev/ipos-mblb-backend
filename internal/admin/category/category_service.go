package category

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type CategoryService interface {
	GetCategory(ctx context.Context) (dto.Categories, error)
	GetCategoryById(ctx context.Context, id uint64) (*dto.Category, error)
}

func NewCategoryService(Cr *impl.CategoryRepositoryImpl) CategoryService {
	return &impl.CategoryServiceImpl{Cr: Cr}
}
