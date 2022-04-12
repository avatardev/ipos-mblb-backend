package impl

import (
	"context"

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