package buyer

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	pkgDto "github.com/avatardev/ipos-mblb-backend/pkg/dto"
)

type BuyerService interface {
	// Services Controller goes here
	GetBuyer(ctx context.Context, limit uint64, offset uint64) (*dto.BuyersResponse, error)
	Ping(ctx context.Context) pkgDto.PingResponse
	PingError(ctx context.Context) error
}

func NewBuyerService(br impl.BuyerRepositoryImpl) BuyerService {
	return &impl.BuyerServiceImpl{Br: br}
}
