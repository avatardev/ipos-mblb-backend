package activity_log

import (
	"context"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type LogService interface {
	GetLogs(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*dto.LogResponseJSON, error)
	Store(ctx context.Context, logInfo *dto.LogInfo) error
}

func NewLogService(lr impl.LogRepositoryImpl) LogService {
	return &impl.LogServiceImpl{Lr: lr}
}
