package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/product_category/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type ProductCategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewProductCategoryRepository(db *database.DatabaseClient) ProductCategoryRepositoryImpl {
	return ProductCategoryRepositoryImpl{DB: db.DB}
}

var (
	COUNT_CATEGORY  = sq.Select("COUNT(*)").From("kategoris")
	SELECT_CATEGORY = sq.Select("id", "nama_kategori", "pajak", "status", "deleted_at", "created_at", "updated_at").From("kategoris")
	INSERT_CATEGORY = sq.Insert("kategoris").Columns("nama_kategori", "pajak", "status", "created_at", "updated_at")
	UPDATE_CATEGORY = sq.Update("kategoris")
)

func (cr ProductCategoryRepositoryImpl) Count(ctx context.Context, keyword string) (uint64, error) {
	stmt, params, err := COUNT_CATEGORY.Where(sq.And{sq.Eq{"deleted_at": nil}, sq.Like{"nama_kategori": fmt.Sprintf("%%%s%%", keyword)}}).ToSql()
	if err != nil {
		log.Printf("[Category.Count] error: %v\n", err)
		return 0, err
	}

	var categoryCount uint64
	queryErr := cr.DB.QueryRowContext(ctx, stmt, params...).Scan(&categoryCount)
	if queryErr != nil {
		log.Printf("[Category.Count] error: %v\n", queryErr)
		return 0, err
	}

	return categoryCount, nil
}

func (cr ProductCategoryRepositoryImpl) GetAll(ctx context.Context, keyword string, limit uint64, offset uint64) (entity.ProductCategories, error) {
	stmt, params, err := SELECT_CATEGORY.Where(sq.And{sq.Eq{"deleted_at": nil}, sq.Like{"nama_kategori": fmt.Sprintf("%%%s%%", keyword)}}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[Category.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := cr.DB.QueryContext(ctx, stmt, params...)
	if err != nil {
		log.Printf("[Category.GetAll] error: %v\n", err)
		return nil, err
	}

	categories := entity.ProductCategories{}

	for rows.Next() {
		var category entity.ProductCategory
		err := rows.Scan(&category.Id, &category.Name, &category.Pajak, &category.Status, &category.Deleted, &category.Created, &category.Updated)
		if err != nil {
			log.Printf("[Category.GetAll] error: %v\n", err)
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cr ProductCategoryRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.ProductCategory, error) {
	stmt, params, err := SELECT_CATEGORY.Where(sq.And{sq.Eq{"id": id}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, err)
		return nil, err
	}

	category := &entity.ProductCategory{}

	rows := cr.DB.QueryRowContext(ctx, stmt, params...)
	queryErr := rows.Scan(&category.Id, &category.Name, &category.Pajak, &category.Status, &category.Deleted, &category.Created, &category.Updated)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, queryErr)
		return nil, queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, queryErr)
		return nil, nil
	}

	return category, nil
}

func (cr ProductCategoryRepositoryImpl) Store(ctx context.Context, category entity.ProductCategory) (*entity.ProductCategory, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_CATEGORY.Values(category.Name, category.Pajak, category.Status, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Category.Store] error: %v\n", err)
		return nil, err
	}

	res, err := cr.DB.ExecContext(ctx, stmt, params...)
	if err != nil {
		log.Printf("[Category.Store] error: %v\n", err)
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[Category.Store] error: %v\n", err)
		return nil, err
	}

	return cr.GetById(ctx, lid)
}

func (cr ProductCategoryRepositoryImpl) Update(ctx context.Context, category entity.ProductCategory) (*entity.ProductCategory, error) {
	updateMap := map[string]interface{}{
		"nama_kategori": category.Name,
		"pajak":         category.Pajak,
		"status":        category.Status,
		"updated_at":    time.Now(),
	}

	stmt, params, err := UPDATE_CATEGORY.SetMap(updateMap).Where(sq.Eq{"id": category.Id}).ToSql()
	if err != nil {
		log.Printf("[Category.Update] error: %v\n", err)
		return nil, err
	}

	if _, err := cr.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[Category.Update] error: %v\n", err)
		return nil, err
	}

	return cr.GetById(ctx, category.Id)
}

func (cr ProductCategoryRepositoryImpl) Delete(ctx context.Context, id int64) error {
	currTime := time.Now()

	updateMap := map[string]interface{}{
		"status":     false,
		"updated_at": currTime,
		"deleted_at": currTime,
	}

	stmt, params, err := UPDATE_CATEGORY.SetMap(updateMap).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Printf("[Category.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	if _, err := cr.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[Category.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	return nil
}
