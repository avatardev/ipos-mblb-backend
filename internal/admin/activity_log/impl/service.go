package impl

import (
	"context"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type LogServiceImpl struct {
	Lr LogRepositoryImpl
}

func (l *LogServiceImpl) GetLogs(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*dto.LogResponseJSON, error) {
	logs, err := l.Lr.GetAll(ctx, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewLogResponseJSON(logs), nil
}

func (l *LogServiceImpl) Store(ctx context.Context, logInfo *dto.LogInfo) error {
	return l.Lr.Store(ctx, logInfo.UserId, logInfo.Message)
}
