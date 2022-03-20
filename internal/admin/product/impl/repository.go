package impl

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/product/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *database.DatabaseClient) ProductRepositoryImpl {
	return ProductRepositoryImpl{DB: db.DB}
}

var (
	SELECT_PRODUCT = sq.Select("id", "id_kategori", "nama_produk", "harga_std_m3", "keterangan", "status", "deleted_at", "created_at", "updated_at").From("produks")
	COUNT_PRODUCT  = sq.Select("COUNT(*)")
)

func (pr ProductRepositoryImpl) GetAll(ctx context.Context) (entity.Products, error) {
	stmt, params, err := SELECT_PRODUCT.Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[Product.GetAll] error: %v\n", err)
		return nil, err
	}

	prpd, err := pr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Product.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Product.GetAll] error:%v\n", err)
		return nil, err
	}

	defer rows.Close()
	data := entity.Products{}

	for rows.Next() {
		temp := &entity.Product{}
		err := rows.Scan(&temp.Id, &temp.CategoryId, &temp.Name, &temp.Price, &temp.Description, &temp.Status, &temp.Deleted, &temp.Created, &temp.Updated)
		if err != nil {
			log.Printf("[Product.GetAll] error: %v\n", err)
			return nil, err
		}

		data = append(data, temp)
	}

	return data, err
}
