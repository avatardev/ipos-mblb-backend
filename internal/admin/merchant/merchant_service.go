package merchant

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/merchant/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type MerchantService interface {
	GetMerchant(ctx context.Context, sellerId int64, keyword string, limit uint64, offset uint64) (*dto.MerchantsResponse, error)
	GetMerchantById(ctx context.Context, sellerId int64, itemId int64) (*dto.MerchantResponse, error)
	UpdateMerchant(ctx context.Context, sellerId int64, itemId int64, req *dto.MerchantRequest) (*dto.MerchantResponse, error)
}

func NewMerchantService(Mr impl.MerchantRepositoryImpl) MerchantService {
	return &impl.MerchantServiceImpl{Mr: Mr}
}
