package impl

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type ProductRepositoryImpl struct {
	db *sql.DB
}

var (
	COUNT_PRODUCT = sq.Select("COUNT(*)")
)

func NewProductRepository(db *database.DatabaseClient) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db.DB}
}
