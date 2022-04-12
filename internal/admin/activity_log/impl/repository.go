package impl

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type LogRepositoryImpl struct {
	DB *sql.DB
}

var (
	SELECT_LOGS = sq.Select("l.id", "l.id_user", "u.username", "l.keterangan", "l.created_at").From("log_aktifitas l").LeftJoin("users u ON l.id_user = u.id")
)

func NewLogRepository(db *database.DatabaseClient) LogRepositoryImpl {
	return LogRepositoryImpl{DB: db.DB}
}

func (l *LogRepositoryImpl) GetAll(ctx context.Context) (entity.Logs, error) {
	stmt, params, err := SELECT_LOGS.OrderBy("l.created_at").ToSql()
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