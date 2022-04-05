package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type DashboardServiceImpl struct {
	Dr DashboardRepositoryImpl
}

func (d *DashboardServiceImpl) GetStatistics(ctx context.Context) (*dto.DashboardInfoResponse, error) {
	res, err := d.Dr.GetInfo(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewDashboardInfoResponse(res), nil
}
