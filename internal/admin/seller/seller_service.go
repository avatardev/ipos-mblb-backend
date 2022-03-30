package seller

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type SellerService interface {
	GetSeller(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.SellersResponse, error)
	GetSellerName(ctx context.Context) (*dto.SellersCompanyResponse, error)
	GetSellerById(ctx context.Context, id int64) (*dto.SellerResponse, error)
	StoreSeller(ctx context.Context, req *dto.SellerRequest) (*dto.SellerResponse, error)
	UpdateSeller(ctx context.Context, id int64, req *dto.SellerRequest) (*dto.SellerResponse, error)
	DeleteSeller(ctx context.Context, id int64) error
}

func NewSellerService(Sr impl.SellerRepositoryImpl) SellerService {
	return &impl.SellerServiceImpl{Sr: Sr}
}
