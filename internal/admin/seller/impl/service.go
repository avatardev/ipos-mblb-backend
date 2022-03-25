package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type SellerServiceImpl struct {
	Sr SellerRepositoryImpl
}

func (s SellerServiceImpl) GetSeller(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.SellersResponse, error) {
	sellerCount, err := s.Sr.Count(ctx, keyword)
	if err != nil {
		return nil, err
	}

	if sellerCount == 0 {
		return nil, errors.ErrInvalidResources
	}

	sellers, err := s.Sr.GetAll(ctx, keyword, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(sellers) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewSellersResponse(sellers, limit, offset, sellerCount), nil
}

func (s SellerServiceImpl) GetSellerById(ctx context.Context, id int64) (*dto.SellerResponse, error) {
	seller, err := s.Sr.GetById(ctx, id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if seller == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewSellerResponse(seller), nil
}

func (s SellerServiceImpl) StoreSeller(ctx context.Context, req *dto.SellerRequest) (*dto.SellerResponse, error) {
	seller, err := req.ToEntity()
	if err != nil {
		return nil, err
	}

	data, err := s.Sr.Store(ctx, *seller)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewSellerResponse(data), nil
}
