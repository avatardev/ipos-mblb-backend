package impl

import (
	"context"
	"database/sql"
	"log"
	"time"

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
	COUNT_PRODUCT  = sq.Select("COUNT(*)")
	SELECT_PRODUCT = sq.Select("id", "id_kategori", "nama_produk", "harga_std_m3", "keterangan", "status", "deleted_at", "created_at", "updated_at").From("produks")
	INSERT_PRODUCT = sq.Insert("produks").Columns("id_kategori", "nama_produk", "harga_std_m3", "keterangan", "status", "created_at", "updated_at")
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

func (pr ProductRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Product, error) {
	stmt, params, err := SELECT_PRODUCT.Where(sq.And{sq.Eq{"id": id}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Product.GetById] id: %v, error: %v\n", id, err)
		return nil, err
	}

	prpd, err := pr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Product.GetById] id: %v, error: %v\n", id, err)
		return nil, err
	}

	rows := prpd.QueryRowContext(ctx, params...)

	product := &entity.Product{}
	queryErr := rows.Scan(&product.Id, &product.CategoryId, &product.Name, &product.Price, &product.Description, &product.Status, &product.Deleted, &product.Created, &product.Updated)
	if queryErr != nil {
		log.Printf("[Product.GetById] id: %v, error: %v\n", id, queryErr)
		return nil, queryErr
	}

	return product, nil
}

func (pr ProductRepositoryImpl) Store(ctx context.Context, product entity.Product) (*entity.Product, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_PRODUCT.Values(product.CategoryId, product.Name, product.Price, product.Description, product.Status, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Product.Store] error: %v\n", err)
		return nil, err
	}

	prpd, err := pr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Product.Store] error: %v\n", err)
		return nil, err
	}

	res, err := prpd.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("[Product.Store] error: %v\n", err)
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[Product.Store] error: %v\n", err)
	}

	return pr.GetById(ctx, lid)
}
