package buyer

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/impl"
	"github.com/avatardev/ipos-mblb-backend/pkg/dto"
)

type BuyerService interface {
	Ping(ctx context.Context) dto.PingResponse
	PingError(ctx context.Context) error
}

func NewBuyerService() BuyerService {
	return &impl.BuyerServiceImpl{}
}
