package impl

import (
	"context"
	"database/sql"
	"fmt"
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
	COUNT_PRODUCT  = sq.Select("COUNT(*)").From("produks")
	SELECT_PRODUCT = sq.Select("produks.id", "k.nama_kategori", "produks.nama_produk", "produks.harga_std_m3", "k.pajak", "produks.keterangan", "produks.status").
			From("produks").LeftJoin("kategoris AS k ON produks.id_kategori = k.id")
	INSERT_PRODUCT = sq.Insert("produks").Columns("id_kategori", "nama_produk", "harga_std_m3", "keterangan", "status", "created_at", "updated_at")
	UPDATE_PRODUCT = sq.Update("produks")
)

func (pr ProductRepositoryImpl) Count(ctx context.Context) (uint64, error) {
	stmt, params, err := COUNT_PRODUCT.Where(sq.Eq{"produks.deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[Product.Count] error: %v\n", err)
		return 0, err
	}

	prpd, err := pr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Product.Count] error: %v\n", err)
		return 0, err
	}

	var productCount uint64
	row := prpd.QueryRowContext(ctx, params...)
	queryErr := row.Scan(&productCount)
	if queryErr != nil {
		log.Printf("[Product.Count] error: %v\n", err)
		return 0, err
	}

	return productCount, nil
}

func (pr ProductRepositoryImpl) GetAll(ctx context.Context, keyword string, limit uint64, offset uint64) (entity.Products, error) {
	stmt, params, err := SELECT_PRODUCT.Where(sq.And{sq.Eq{"produks.deleted_at": nil}, sq.Like{"produks.nama_produk": fmt.Sprintf("%%%s%%", keyword)}}).ToSql()
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
		err := rows.Scan(&temp.Id, &temp.CategoryName, &temp.Name, &temp.Price, &temp.Tax, &temp.Description, &temp.Status)
		if err != nil {
			log.Printf("[Product.GetAll] error: %v\n", err)
			return nil, err
		}

		data = append(data, temp)
	}

	return data, err
}

func (pr ProductRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Product, error) {
	stmt, params, err := SELECT_PRODUCT.Where(sq.And{sq.Eq{"produks.id": id}, sq.Eq{"produks.deleted_at": nil}}).ToSql()
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
	queryErr := rows.Scan(&product.Id, &product.CategoryName, &product.Name, &product.Price, &product.Description, &product.Status)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[Product.GetById] id: %v, error: %v\n", id, queryErr)
		return nil, queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Product.GetById] id: %v, error: %v\n", id, queryErr)
		return nil, nil
	}

	return product, nil
}

func (pr ProductRepositoryImpl) Store(ctx context.Context, product entity.Product) (*entity.Product, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_PRODUCT.Values(product.CategoryId, product.Name, product.Price, product.Description, product.Status, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Product.Store] categoryId: %v, name: %v, desc: %v, status: %v, error: %v\n", product.CategoryId, product.Name, product.Description, product.Status, err)
		return nil, err
	}

	prpd, err := pr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Product.Store] categoryId: %v, name: %v, desc: %v, status: %v, error: %v\n", product.CategoryId, product.Name, product.Description, product.Status, err)
		return nil, err
	}

	res, err := prpd.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("[Product.Store] categoryId: %v, name: %v, desc: %v, status: %v, error: %v\n", product.CategoryId, product.Name, product.Description, product.Status, err)
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[Product.Store] categoryId: %v, name: %v, desc: %v, status: %v, error: %v\n", product.CategoryId, product.Name, product.Description, product.Status, err)
		return nil, err
	}

	return pr.GetById(ctx, lid)
}

func (pr ProductRepositoryImpl) Update(ctx context.Context, product entity.Product) (*entity.Product, error) {
	updateMap := map[string]interface{}{
		"id_kategori":  product.CategoryId,
		"nama_produk":  product.Name,
		"harga_std_m3": product.Price,
		"keterangan":   product.Description,
		"status":       product.Status,
		"updated_at":   time.Now(),
	}

	stmt, params, err := UPDATE_PRODUCT.SetMap(updateMap).Where(sq.Eq{"id": product.Id}).ToSql()
	if err != nil {
		log.Printf("[Product.Update] categoryId: %v, name: %v, desc: %v, status: %v, error: %v\n", product.CategoryId, product.Name, product.Description, product.Status, err)
		return nil, err
	}

	prpd, err := pr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Product.Update] categoryId: %v, name: %v, desc: %v, status: %v, error: %v\n", product.CategoryId, product.Name, product.Description, product.Status, err)
		return nil, err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Product.Update] categoryId: %v, name: %v, desc: %v, status: %v, error: %v\n", product.CategoryId, product.Name, product.Description, product.Status, err)
		return nil, err
	}

	return pr.GetById(ctx, product.Id)
}

func (pr ProductRepositoryImpl) Delete(ctx context.Context, id int64) error {
	currTime := time.Now()

	updateMap := map[string]interface{}{
		"status":     false,
		"updated_at": currTime,
		"deleted_at": currTime,
	}

	stmt, params, err := UPDATE_PRODUCT.SetMap(updateMap).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Printf("[Product.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	prpd, err := pr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Product.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Product.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	return nil
}
