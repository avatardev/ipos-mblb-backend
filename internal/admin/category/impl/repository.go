package impl

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/category/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *database.DatabaseClient) CategoryRepositoryImpl {
	return CategoryRepositoryImpl{DB: db.DB}
}

var (
	COUNT_CATEGORY  = sq.Select("COUNT(*)").From("kategoris")
	SELECT_CATEGORY = sq.Select("id", "nama_kategori", "pajak", "status", "deleted_at", "created_at", "updated_at").From("kategoris")
	INSERT_CATEGORY = sq.Insert("kategoris").Columns("nama_kategori", "pajak", "status", "created_at", "updated_at")
	UPDATE_CATEGORY = sq.Update("kategoris")
)

func (cr CategoryRepositoryImpl) Count(ctx context.Context) (uint64, error) {
	stmt, _, err := COUNT_CATEGORY.ToSql()
	if err != nil {
		log.Printf("[Category.Count] error: %v\n", err)
		return 0, err
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.Count] error: %v\n", err)
		return 0, err
	}

	var categoryCount uint64
	queryErr := prpd.QueryRowContext(ctx).Scan(&categoryCount)
	if queryErr != nil {
		log.Printf("[Category.Count] error: %v\n", queryErr)
		return 0, err
	}

	return categoryCount, nil
}

func (cr CategoryRepositoryImpl) GetAll(ctx context.Context) (entity.Categories, error) {
	stmt, _, err := SELECT_CATEGORY.Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[Category.GetAll] error: %v\n", err)
		return nil, err
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx)
	if err != nil {
		log.Printf("[Category.GetAll] error: %v\n", err)
		return nil, err
	}

	categories := entity.Categories{}

	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.Id, &category.Name, &category.Pajak, &category.Status, &category.Deleted, &category.Created, &category.Updated)
		if err != nil {
			log.Printf("[Category.GetAll] error: %v\n", err)
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cr CategoryRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Category, error) {
	stmt, params, err := SELECT_CATEGORY.Where(sq.And{sq.Eq{"id": id}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, err)
		return nil, err
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, err)
		return nil, err
	}

	category := &entity.Category{}

	rows := prpd.QueryRowContext(ctx, params...)
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

func (cr CategoryRepositoryImpl) Store(ctx context.Context, category entity.Category) (*entity.Category, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_CATEGORY.Values(category.Name, category.Pajak, category.Status, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Status, category.Status, err)
		return nil, err
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Pajak, category.Status, err)
		return nil, err
	}

	res, err := prpd.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Pajak, category.Status, err)
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Pajak, category.Status, err)
		return nil, err
	}

	return cr.GetById(ctx, lid)
}

func (cr CategoryRepositoryImpl) Update(ctx context.Context, category entity.Category) (*entity.Category, error) {
	updateMap := map[string]interface{}{
		"nama_kategori": category.Name,
		"pajak":         category.Pajak,
		"status":        category.Status,
		"updated_at":    time.Now(),
	}

	stmt, params, err := UPDATE_CATEGORY.SetMap(updateMap).Where(sq.Eq{"id": category.Id}).ToSql()
	if err != nil {
		log.Printf("[Category.Update] name: %v, pajak: %v, status: %v\n", category.Name, category.Pajak, category.Status)
		return nil, err
	}

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.Update] name: %v, pajak: %v, status: %v\n", category.Name, category.Pajak, category.Status)
		return nil, err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Category.Update] name: %v, pajak: %v, status: %v\n", category.Name, category.Pajak, category.Status)
		return nil, err
	}

	return cr.GetById(ctx, category.Id)
}

func (cr CategoryRepositoryImpl) Delete(ctx context.Context, id int64) error {
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

	prpd, err := cr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Category.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Category.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	return nil
}
