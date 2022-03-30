package buyer

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	pkgDto "github.com/avatardev/ipos-mblb-backend/pkg/dto"
)

type BuyerService interface {
	// Services Controller goes here
	GetBuyer(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.BuyersResponse, error)
	GetBuyerById(ctx context.Context, plate string) (*dto.BuyerResponse, error)
	GetBuyerName(ctx context.Context) (*dto.BuyersCompanyResponse, error)
	StoreBuyer(ctx context.Context, req *dto.BuyerRequest) (*dto.BuyerResponse, error)
	UpdateBuyer(ctx context.Context, plate string, req *dto.BuyerRequest) (*dto.BuyerResponse, error)
	DeleteBuyer(ctx context.Context, plate string) error
	Ping(ctx context.Context) pkgDto.PingResponse
	PingError(ctx context.Context) error
}

func NewBuyerService(br impl.BuyerRepositoryImpl) BuyerService {
	return &impl.BuyerServiceImpl{Br: br}
}
