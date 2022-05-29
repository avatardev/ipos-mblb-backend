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
	SELECT_PRODUCT = sq.Select("p.id", "p.id_kategori", "k.nama_kategori", "p.nama_produk", "p.harga_std_m3", "k.pajak", "p.keterangan", "p.status", "p.img").
			From("produks p").LeftJoin("kategoris k ON p.id_kategori = k.id")
	INSERT_PRODUCT = sq.Insert("produks").Columns("id_kategori", "nama_produk", "harga_std_m3", "keterangan", "status", "created_at", "updated_at")
	UPDATE_PRODUCT = sq.Update("produks")
)

var (
	SELECT_SELLER        = sq.Select("id").From("sellers")
	INSERT_MERCHANT_ITEM = sq.Insert("produk_sellers").Columns("id_produk", "harga", "status", "id_seller", "updated_at", "created_at")
)

func (pr ProductRepositoryImpl) Count(ctx context.Context, keyword string) (uint64, error) {
	stmt, params, err := COUNT_PRODUCT.Where(sq.And{sq.Eq{"produks.deleted_at": nil}, sq.Like{"produks.nama_produk": fmt.Sprintf("%%%s%%", keyword)}}).ToSql()
	if err != nil {
		log.Printf("[Product.Count] error: %v\n", err)
		return 0, err
	}

	var productCount uint64
	row := pr.DB.QueryRowContext(ctx, stmt, params...)
	queryErr := row.Scan(&productCount)
	if queryErr != nil {
		log.Printf("[Product.Count] error: %v\n", err)
		return 0, err
	}

	return productCount, nil
}

func (pr ProductRepositoryImpl) GetAll(ctx context.Context, keyword string, limit uint64, offset uint64) (entity.Products, error) {
	stmt, params, err := SELECT_PRODUCT.Where(sq.And{sq.Eq{"p.deleted_at": nil}, sq.Like{"p.nama_produk": fmt.Sprintf("%%%s%%", keyword)}}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[Product.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := pr.DB.QueryContext(ctx, stmt, params...)
	if err != nil {
		log.Printf("[Product.GetAll] error:%v\n", err)
		return nil, err
	}

	defer rows.Close()
	data := entity.Products{}

	for rows.Next() {
		temp := &entity.Product{}
		err := rows.Scan(&temp.Id, &temp.CategoryId, &temp.CategoryName, &temp.Name, &temp.Price, &temp.Tax, &temp.Description, &temp.Status, &temp.Img)
		if err != nil {
			log.Printf("[Product.GetAll] error: %v\n", err)
			return nil, err
		}

		data = append(data, temp)
	}

	return data, nil
}

func (pr ProductRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Product, error) {
	stmt, params, err := SELECT_PRODUCT.Where(sq.And{sq.Eq{"p.id": id}, sq.Eq{"p.deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Product.GetById] id: %v, error: %v\n", id, err)
		return nil, err
	}

	rows := pr.DB.QueryRowContext(ctx, stmt, params...)

	product := &entity.Product{}
	queryErr := rows.Scan(&product.Id, &product.CategoryId, &product.CategoryName, &product.Name, &product.Price, &product.Tax, &product.Description, &product.Status, &product.Img)
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

	res, err := pr.DB.ExecContext(ctx, stmt, params...)
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

	if _, err := pr.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[Product.Update] categoryId: %v, name: %v, desc: %v, status: %v, error: %v\n", product.CategoryId, product.Name, product.Description, product.Status, err)
		return nil, err
	}

	return pr.GetById(ctx, product.Id)
}

func (pr ProductRepositoryImpl) UpdateImage(ctx context.Context, id int64, img string) (*entity.Product, error) {
	updateMap := map[string]interface{}{
		"img":        img,
		"updated_at": time.Now(),
	}

	stmt, params, err := UPDATE_PRODUCT.SetMap(updateMap).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Printf("[Product.UpdateImage] categoryId: %v, error: %v\n", id, err)
		return nil, err
	}

	if _, err := pr.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[Product.UpdateImage] categoryId: %v, error: %v\n", id, err)
		return nil, err
	}

	return pr.GetById(ctx, id)
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

	if _, err := pr.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[Product.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	return nil
}

// SELECT_SELLER        = sq.Select("id").From("sellers")
// INSERT_MERCHANT_ITEM = sq.Insert("produk_sellers").Columns("id_produk", "harga", "status", "id_seller", "updated_at", "created_at")

func (pr ProductRepositoryImpl) FindActiveSeller(ctx context.Context) (sellers []int64, err error) {
	stmt, args, err := SELECT_SELLER.Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[Product.FindActiveSeller] error: %v\n", err)
		return
	}

	rows, err := pr.DB.QueryContext(ctx, stmt, args...)
	if err != nil {
		log.Printf("[Product.FindActiveSeller] error: %v\n", err)
		return
	}

	sellers = []int64{}
	for rows.Next() {
		var id int64
		if err = rows.Scan(&id); err != nil {
			sellers = nil
			return
		}
		sellers = append(sellers, id)
	}

	return
}

func (pr ProductRepositoryImpl) StoreNewMerchantItem(ctx context.Context, id int64, product entity.Product) (err error) {
	currTime := time.Now()
	stmt, args, err := INSERT_MERCHANT_ITEM.Values(product.Id, product.Price, 0, id, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Product.StoreNewMerchantItem] error: %v\n", err)
		return
	}

	_, err = pr.DB.ExecContext(ctx, stmt, args...)
	if err != nil {
		log.Printf("[Product.StoreNewMerchantItem] error: %v\n", err)
		return
	}

	return
}
