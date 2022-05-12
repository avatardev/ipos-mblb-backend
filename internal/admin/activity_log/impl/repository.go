package impl

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type LogRepositoryImpl struct {
	DB *sql.DB
}

var (
	SELECT_LOGS = sq.Select("l.id", "l.id_user", "u.username", "l.keterangan", "l.created_at").From("log_aktifitas l").LeftJoin("users u ON l.id_user = u.id")
	STORE_LOG   = sq.Insert("log_aktifitas").Columns("id_user", "keterangan", "created_at", "updated_at")
)

func NewLogRepository(db *database.DatabaseClient) LogRepositoryImpl {
	return LogRepositoryImpl{DB: db.DB}
}

func (l *LogRepositoryImpl) GetAll(ctx context.Context, dateStart time.Time, dateEnd time.Time) (entity.Logs, error) {
	stmt, params, err := SELECT_LOGS.Where(sq.And{sq.GtOrEq{"l.created_at": dateStart}, sq.LtOrEq{"l.created_at": dateEnd}}).OrderBy("l.created_at desc").ToSql()
	if err != nil {
		log.Printf("[Logs.GetAll] error: %v", err)
		return nil, err
	}

	prpd, err := l.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Logs.GetAll] error: %v", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Logs.GetAll] error: %v", err)
		return nil, err
	}

	lg, err := entity.NewLogs(rows)
	if err != nil {
		log.Printf("[Logs.GetAll] error: %v", err)
		return nil, err
	}

	return lg, nil
}

func (l *LogRepositoryImpl) Store(ctx context.Context, userId int64, msg string) error {
	currTime := time.Now()
	stmt, params, err := STORE_LOG.Values(userId, msg, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Logs.Store] error: %v\n", err)
		return err
	}

	prpd, err := l.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Logs.Store] error: %v\n", err)
		return err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Logs.Store] error: %v\n", err)
		return err
	}

	return nil
}
