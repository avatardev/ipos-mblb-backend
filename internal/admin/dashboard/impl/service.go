package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
)

type DashboardServiceImpl struct {
	Dr DashboardRepositoryImpl
}

func (d *DashboardServiceImpl) GetStatistics(ctx context.Context) (*dto.DashboardInfoResponse, error) {
	var sellerID int64 
	
	meta := privutil.GetAuthMetadata(ctx)
	if meta.SellerID != nil {
		sellerID = *meta.SellerID
	}

	res, err := d.Dr.GetInfo(ctx, sellerID)
	if err != nil {
		return nil, err
	}

	return dto.NewDashboardInfoResponse(res), nil
}
