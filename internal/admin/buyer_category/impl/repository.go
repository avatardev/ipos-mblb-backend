package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer_category/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type BuyerCategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewBuyerCategoryRepository(db *database.DatabaseClient) BuyerCategoryRepositoryImpl {
	return BuyerCategoryRepositoryImpl{DB: db.DB}
}

var (
	COUNT_CATEGORY  = sq.Select("COUNT(*)").From("kategori_kendaraans c")
	SELECT_CATEGORY = sq.Select("c.id", "c.nama_kategori", "c.multi_produk").From("kategori_kendaraans c")
	INSERT_CATEGORY = sq.Insert("kategori_kendaraans").Columns("nama_kategori", "multi_produk", "created_at", "updated_at")
	UPDATE_CATEGORY = sq.Update("kategori_kendaraans")
)

func (cr BuyerCategoryRepositoryImpl) Count(ctx context.Context, keyword string) (uint64, error) {
	stmt, params, err := COUNT_CATEGORY.Where(sq.And{sq.Eq{"c.deleted_at": nil}, sq.Like{"c.nama_kategori": fmt.Sprintf("%%%s%%", keyword)}}).ToSql()
	if err != nil {
		log.Printf("[BuyerCategory.Count] error: %v\n", err)
		return 0, err
	}

	var categoryCount uint64
	row := cr.DB.QueryRowContext(ctx, stmt, params...)
	queryErr := row.Scan(&categoryCount)
	if queryErr != nil {
		log.Printf("[BuyerCategory.Count] error: %v\n", err)
		return 0, err
	}

	return categoryCount, nil
}

func (cr BuyerCategoryRepositoryImpl) GetAll(ctx context.Context, keyword string, limit uint64, offset uint64) (entity.BuyersCategories, error) {
	stmt, params, err := SELECT_CATEGORY.Where(sq.And{sq.Eq{"deleted_at": nil}, sq.Like{"c.nama_kategori": fmt.Sprintf("%%%s%%", keyword)}}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[BuyerCategory.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := cr.DB.QueryContext(ctx, stmt, params...)
	if err != nil {
		log.Printf("[BuyerCategory.GetAll] error: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	data := entity.BuyersCategories{}

	for rows.Next() {
		temp := &entity.BuyerCategory{}
		err := rows.Scan(&temp.Id, &temp.Name, &temp.IsMultiProduct)
		if err != nil {
			log.Printf("[BuyerCategory.GetAll] error: %v\n", err)
			return nil, err
		}

		data = append(data, temp)
	}

	return data, nil
}

func (cr BuyerCategoryRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.BuyerCategory, error) {
	stmt, params, err := SELECT_CATEGORY.Where(sq.And{sq.Eq{"c.deleted_at": nil}, sq.Eq{"c.id": id}}).ToSql()
	if err != nil {
		log.Printf("[BuyerCategory.GetById] error: %v\n", err)
		return nil, err
	}

	rows := cr.DB.QueryRowContext(ctx, stmt, params...)

	category := &entity.BuyerCategory{}
	queryErr := rows.Scan(&category.Id, &category.Name, &category.IsMultiProduct)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[BuyerCategory.GetById] error: %v\n", queryErr)
		return nil, queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[BuyerCategory.GetById] error: %v\n", queryErr)
		return nil, nil
	}

	return category, nil
}

func (cr BuyerCategoryRepositoryImpl) Store(ctx context.Context, category entity.BuyerCategory) (*entity.BuyerCategory, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_CATEGORY.Values(category.Name, category.IsMultiProduct, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[BuyerCategory.Store] error: %v\n", err)
		return nil, err
	}

	res, err := cr.DB.ExecContext(ctx, stmt, params...)
	if err != nil {
		log.Printf("[BuyerCategory.Store] error: %v\n", err)
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[BuyerCategory.Store] error: %v\n", err)
		return nil, err
	}

	return cr.GetById(ctx, lid)
}

func (cr BuyerCategoryRepositoryImpl) Update(ctx context.Context, category entity.BuyerCategory) (*entity.BuyerCategory, error) {
	updateMap := map[string]interface{}{
		"id":            category.Id,
		"nama_kategori": category.Name,
		"multi_produk":  category.IsMultiProduct,
		"updated_at":    time.Now(),
	}

	stmt, params, err := UPDATE_CATEGORY.SetMap(updateMap).Where(sq.Eq{"id": category.Id}).ToSql()
	if err != nil {
		log.Printf("[BuyerCategory.Update] error: %v\n", err)
		return nil, err
	}

	if _, err := cr.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[BuyerCategory.Update] error: %v\n", err)
		return nil, err
	}

	return cr.GetById(ctx, category.Id)
}

func (cr BuyerCategoryRepositoryImpl) Delete(ctx context.Context, id int64) error {
	updateMap := map[string]interface{}{
		"deleted_at": time.Now(),
		"updated_at": time.Now(),
	}

	stmt, params, err := UPDATE_CATEGORY.SetMap(updateMap).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Printf("[BuyerCategory.Update] error: %v\n", err)
		return err
	}

	if _, err := cr.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[BuyerCategory.Update] error: %v\n", err)
		return err
	}

	return nil
}
