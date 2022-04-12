package activity_log

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type LogService interface {
	GetLogs(ctx context.Context) (*dto.LogResponseJSON, error)
	Store(ctx context.Context, logInfo *dto.LogInfo) error
}

func NewLogService(lr impl.LogRepositoryImpl) LogService {
	return &impl.LogServiceImpl{Lr: lr}
}
