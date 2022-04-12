package impl

import (
	"context"
	"log"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type LogServiceImpl struct {
	Lr LogRepositoryImpl
}

func (l *LogServiceImpl) GetLogs(ctx context.Context) (*dto.LogResponseJSON, error) {
	logs, err := l.Lr.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewLogResponseJSON(logs), nil
}

func (l *LogServiceImpl) Store(ctx context.Context, logInfo *dto.LogInfo) error {
	// return l.Lr.Store(ctx, logInfo.UserId, logInfo.Message)
	log.Println(logInfo.UserId, logInfo.Message)
	return nil
}
