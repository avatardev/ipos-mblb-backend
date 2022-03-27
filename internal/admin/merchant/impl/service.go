package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type MerchantServiceImpl struct {
	Mr MerchantRepositoryImpl
}

func (m *MerchantServiceImpl) GetMerchant(ctx context.Context, sellerId int64, keyword string, limit uint64, offset uint64) (*dto.MerchantsResponse, error) {
	count, err := m.Mr.Count(ctx, sellerId, keyword)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if count == 0 {
		return nil, errors.ErrInvalidResources
	}

	items, err := m.Mr.GetAll(ctx, sellerId, keyword, limit, offset)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if len(items) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewMerchantsResponse(items, limit, offset, count), nil
}

func (m *MerchantServiceImpl) GetMerchantById(ctx context.Context, sellerId int64, itemId int64) (*dto.MerchantResponse, error) {
	item, err := m.Mr.GetById(ctx, sellerId, itemId)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewMerchantResponse(item), nil
}

func (m *MerchantServiceImpl) UpdateMerchant(ctx context.Context, sellerId int64, itemid int64, req *dto.MerchantRequest) (*dto.MerchantResponse, error) {
	item := req.ToEntity()
	item.Id = itemid

	exist, err := m.Mr.GetById(ctx, sellerId, item.Id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if exist == nil {
		return nil, errors.ErrNotFound
	}

	data, err := m.Mr.Update(ctx, sellerId, item)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewMerchantResponse(data), nil
}
