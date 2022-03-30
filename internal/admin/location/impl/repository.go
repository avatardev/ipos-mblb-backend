package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/location/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type LocationRepositoryImpl struct {
	DB *sql.DB
}

func NewLocationRepostiory(db *database.DatabaseClient) LocationRepositoryImpl {
	return LocationRepositoryImpl{DB: db.DB}
}

var (
	COUNT_LOCATION  = sq.Select("COUNT(*)").From("lokasis l")
	SELECT_LOCATION = sq.Select("l.id", "l.nama_lokasi").From("lokasis l")
	INSERT_LOCATION = sq.Insert("lokasis").Columns("nama_lokasi", "created_at", "updated_at")
	UPDATE_LOCATION = sq.Update("lokasis")
	DELETE_LOCATION = sq.Delete("lokasis")
)

func (lr *LocationRepositoryImpl) Count(ctx context.Context, keyword string) (uint64, error) {
	stmt, params, err := COUNT_LOCATION.Where(sq.Like{"l.nama_lokasi": fmt.Sprintf("%%%s%%", keyword)}).ToSql()
	if err != nil {
		log.Printf("[Location.Count] error: %v\n", err)
		return 0, err
	}

	prpd, err := lr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Location.Count] error: %v\n", err)
		return 0, err
	}

	var locationCount uint64
	row := prpd.QueryRowContext(ctx, params...)
	queryErr := row.Scan(&locationCount)
	if queryErr != nil {
		log.Printf("[Location.Count] error:%v\n", err)
		return 0, err
	}

	return locationCount, nil
}

func (lr *LocationRepositoryImpl) GetAll(ctx context.Context, keyword string, limit uint64, offset uint64) (entity.Locations, error) {
	stmt, params, err := SELECT_LOCATION.Where(sq.Like{"l.nama_lokasi": fmt.Sprintf("%%%s%%", keyword)}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[Location.GetAll] error: %v\n", err)
		return nil, err
	}

	prpd, err := lr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Location.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Location.GetAll] error: %v\n", err)
		return nil, err
	}

	locs, err := entity.NewLocations(rows)
	if err != nil {
		log.Printf("[Location.GetAll] error: %v", err)
		return nil, err
	}

	return locs, nil
}

func (lr *LocationRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Location, error) {
	stmt, params, err := SELECT_LOCATION.Where(sq.Eq{"l.id": id}).ToSql()
	if err != nil {
		log.Printf("[Location.GetById] error: %v\n", err)
		return nil, err
	}

	prpd, err := lr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Location.GetById] error: %v\n", err)
		return nil, err
	}

	loc := &entity.Location{}
	row := prpd.QueryRowContext(ctx, params...)
	if err := loc.FromSql(row); err != nil && err != sql.ErrNoRows {
		log.Printf("[Location.GetById] error:%v\n", err)
		return nil, err
	} else if err == sql.ErrNoRows {
		log.Printf("[Location.GetById] error:%v\n", err)
		return nil, nil
	}

	return loc, nil
}

func (lr *LocationRepositoryImpl) Store(ctx context.Context, location entity.Location) (*entity.Location, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_LOCATION.Values(location.Name, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Location.Store] error: %v\n", err)
		return nil, err
	}

	prpd, err := lr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Location.Store] error: %v\n", err)
		return nil, err
	}

	res, err := prpd.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("[Location.Store] error: %v\n", err)
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[Location.Store] error: %v\n", err)
		return nil, err
	}

	return lr.GetById(ctx, lid)
}

func (lr *LocationRepositoryImpl) Update(ctx context.Context, location entity.Location) (*entity.Location, error) {
	updateMap := map[string]interface{}{
		"nama_lokasi": location.Name,
		"updated_at":  time.Now(),
	}
	stmt, params, err := UPDATE_LOCATION.SetMap(updateMap).Where(sq.Eq{"id": location.Id}).ToSql()
	if err != nil {
		log.Printf("[Location.Update] err: %v", err)
		return nil, err
	}

	prpd, err := lr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Location.Update] err: %v", err)
		return nil, err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Location.Update] err: %v", err)
		return nil, err
	}

	return lr.GetById(ctx, location.Id)
}

func (lr *LocationRepositoryImpl) Delete(ctx context.Context, id int64) error {
	stmt, params, err := DELETE_LOCATION.Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Printf("[Location.Delete] err: %v\n", err)
		return err
	}

	prpd, err := lr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Location.Delete] err: %v\n", err)
		return err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Location.Delete] err: %v\n", err)
		return err
	}

	return nil
}
