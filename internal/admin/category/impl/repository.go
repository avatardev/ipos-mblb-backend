package impl

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/category/entity"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

var (
	COUNT_CATEGORY  = sq.Select("COUNT(*)").From("kategoris")
	SELECT_CATEGORY = sq.Select("id", "nama_kategori", "pajak", "status", "deleted_at", "created_at", "updated_at").From("kategoris")
	INSERT_CATEGORY = sq.Insert("kategoris").Columns("nama_kategori", "pajak", "status", "created_at", "updated_at")
	UPDATE_CATEGORY = sq.Update("kategoris")
)

// func NewCategoryRepository(db *database.DatabaseClient) *CategoryRepositoryImpl {
// 	return &CategoryRepositoryImpl{db: db.DB}
// }

func (cr CategoryRepositoryImpl) Count(ctx context.Context) uint64 {
	stmt, _, err := COUNT_CATEGORY.ToSql()
	if err != nil {
		log.Printf("[Category.Count] error: %v\n", err)
		return 0
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.Count] error: %v\n", err)
		return 0
	}

	var categoryCount uint64
	queryErr := prpd.QueryRowContext(ctx).Scan(&categoryCount)
	if queryErr != nil {
		log.Printf("[Category.Count] error: %v\n", queryErr)
		return 0
	}

	return categoryCount
}

func (cr CategoryRepositoryImpl) GetAll(ctx context.Context) entity.Categories {
	stmt, _, err := SELECT_CATEGORY.Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[Category.GetAll] error: %v\n", err)
		return nil
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.GetAll] error: %v\n", err)
		return nil
	}

	rows, err := prpd.QueryContext(ctx)
	if err != nil {
		log.Printf("[Category.GetAll] error: %v\n", err)
	}

	categories := entity.Categories{}

	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.Id, &category.Name, &category.Pajak, &category.Status, &category.Deleted, &category.Created, &category.Updated)
		if err != nil {
			log.Printf("[Category.GetAll] error: %v\n", err)
			return nil
		}
		categories = append(categories, &category)
	}
	return categories
}

func (cr CategoryRepositoryImpl) GetById(ctx context.Context, id int64) *entity.Category {
	stmt, params, err := SELECT_CATEGORY.Where(sq.And{sq.Eq{"id": id}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, err)
		return nil
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, err)
		return nil
	}

	category := &entity.Category{}

	rows := prpd.QueryRowContext(ctx, params...)
	queryErr := rows.Scan(&category.Id, &category.Name, &category.Pajak, &category.Status, &category.Deleted, &category.Created, &category.Updated)
	if queryErr != nil {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, queryErr)
		return nil
	}

	return category
}

func (cr CategoryRepositoryImpl) Store(ctx context.Context, category entity.Category) *entity.Category {
	currTime := time.Now()
	stmt, params, err := INSERT_CATEGORY.Values(category.Name, category.Pajak, category.Status, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Status, category.Status, err)
		return nil
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Pajak, category.Status, err)
		return nil
	}

	res, err := prpd.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Pajak, category.Status, err)
		return nil
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Pajak, category.Status, err)
		return nil
	}

	return cr.GetById(ctx, lid)
}

func (cr CategoryRepositoryImpl) Update(ctx context.Context, category entity.Category) *entity.Category {
	updateMap := map[string]interface{}{
		"nama_kategori": category.Name,
		"pajak":         category.Pajak,
		"status":        category.Status,
		"updated_at":    time.Now(),
	}

	stmt, params, err := UPDATE_CATEGORY.SetMap(updateMap).Where(sq.Eq{"id": category.Id}).ToSql()
	if err != nil {
		log.Printf("[Category.Update] name: %v, pajak: %v, status: %v\n", category.Name, category.Pajak, category.Status)
		return nil
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.Update] name: %v, pajak: %v, status: %v\n", category.Name, category.Pajak, category.Status)
		return nil
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Category.Update] name: %v, pajak: %v, status: %v\n", category.Name, category.Pajak, category.Status)
		return nil
	}

	return cr.GetById(ctx, category.Id)
}

func (cr CategoryRepositoryImpl) Delete(ctx context.Context, id int64) bool {
	currTime := time.Now()

	updateMap := map[string]interface{}{
		"status":     false,
		"updated_at": currTime,
		"deleted_at": currTime,
	}

	stmt, params, err := UPDATE_CATEGORY.SetMap(updateMap).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Printf("[Category.Delete] id: %v, error: %v\n", id, err)
		return false
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.Delete] id: %v, error: %v\n", id, err)
		return false
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Category.Delete] id: %v, error: %v\n", id, err)
		return false
	}

	return true
}
