package dashboard

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/dashboard/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type DashboardService interface {
	GetStatistics(ctx context.Context) (*dto.DashboardInfoResponse, error)
}

func NewDashboardService(Dr impl.DashboardRepositoryImpl) DashboardService {
	return &impl.DashboardServiceImpl{Dr: Dr}
}
