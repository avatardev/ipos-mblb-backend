package category

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/category/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type CategoryRepository interface {
	GetAll(ctx context.Context) entity.Categories
	GetById(ctx context.Context, id int64) *entity.Category
	Store(ctx context.Context, category entity.Category) *entity.Category
	Update(ctx context.Context, category entity.Category) *entity.Category
	Delete(ctx context.Context, id int64) bool
}

func NewCategoryRepository(db *database.DatabaseClient) impl.CategoryRepositoryImpl {
	return impl.CategoryRepositoryImpl{DB: db.DB}
}
