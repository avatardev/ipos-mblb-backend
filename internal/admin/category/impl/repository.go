package impl

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/category/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type CategoryRepositoryImpl struct {
	db *sql.DB
}

var (
	COUNT_CATEGORY  = sq.Select("COUNT(*)").From("kategoris")
	SELECT_CATEGORY = sq.Select("id", "nama_kategori", "pajak", "status", "deleted_at", "created_at", "updated_at").From("kategoris")
	INSERT_CATEGORY = sq.Insert("kategoris").Columns("nama_kategori", "pajak", "status")
)

func NewCategoryRepository(db *database.DatabaseClient) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db: db.DB}
}

func (cr CategoryRepositoryImpl) Count(ctx context.Context) uint64 {
	stmt, _, err := COUNT_CATEGORY.ToSql()
	if err != nil {
		log.Printf("[Category.Count] error: %v\n", err)
		return 0
	}

	prpd, err := cr.db.PrepareContext(ctx, stmt)
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

	prpd, err := cr.db.PrepareContext(ctx, stmt)
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
	stmt, params, err := SELECT_CATEGORY.Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Printf("[Category.GetById] id: %v, error: %v\n", id, err)
		return nil
	}

	prpd, err := cr.db.PrepareContext(ctx, stmt)
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
	stmt, params, err := INSERT_CATEGORY.Values(category.Name, category.Pajak, category.Status).ToSql()
	if err != nil {
		log.Printf("[Category.Store] name: %v, pajak: %v, status: %v, error: %v\n", category.Name, category.Status, category.Status, err)
		return nil
	}

	prpd, err := cr.db.PrepareContext(ctx, stmt)
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
