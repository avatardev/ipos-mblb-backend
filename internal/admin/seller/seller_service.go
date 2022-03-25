package seller

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type SellerService interface {
	GetSeller(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.SellersResponse, error)
}

func NewSellerService(Sr impl.SellerRepositoryImpl) SellerService {
	return &impl.SellerServiceImpl{Sr: Sr}
}
