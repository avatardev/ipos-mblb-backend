package impl

import (
	"context"
	"fmt"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/logutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
)

type SellerServiceImpl struct {
	Sr SellerRepositoryImpl
}

func (s SellerServiceImpl) GetSeller(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.SellersResponse, error) {
	if meta := privutil.GetAuthMetadata(ctx); meta != nil {
		if meta.SellerID != nil {
			sellerID := *meta.SellerID
			data, err := s.Sr.GetById(ctx, sellerID)
			if err != nil {
				return nil, err
			}

			item := entity.Sellers{}
			item = append(item, data)

			return dto.NewSellersResponse(item, 1, 0, 1), nil
		}
	}

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

func (s SellerServiceImpl) GetSellerName(ctx context.Context) (*dto.SellersCompanyResponse, error) {
	sellers, err := s.Sr.GetSelerName(ctx)
	if err != nil {
		return nil, err
	}

	if len(sellers) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewSellersCompanyResponse(sellers), nil
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

func (s SellerServiceImpl) StoreSeller(ctx context.Context, req *dto.SellerRequest) (res *dto.SellerResponse, err error) {
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

	if err = s.storeInitialMerchantItem(ctx, *data); err != nil {
		res = nil
		return
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("added new seller %s", seller.Company))
	return dto.NewSellerResponse(data), nil
}

func (s SellerServiceImpl) UpdateSeller(ctx context.Context, id int64, req *dto.SellerRequest) (*dto.SellerResponse, error) {
	seller, err := req.ToEntity()
	if err != nil {
		return nil, err
	}

	seller.Id = id
	exists, err := s.Sr.GetById(ctx, seller.Id)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	data, err := s.Sr.Update(ctx, *seller)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("changed seller data %s", seller.Company))
	return dto.NewSellerResponse(data), nil
}

func (s SellerServiceImpl) DeleteSeller(ctx context.Context, id int64) error {
	exists, err := s.Sr.GetById(ctx, id)
	if err != nil {
		return err
	}

	if exists == nil {
		return errors.ErrNotFound
	}

	if err := s.Sr.Delete(ctx, id); err != nil {
		return err
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("deleted seller data %s", exists.Company))
	return nil
}

func (s SellerServiceImpl) storeInitialMerchantItem(ctx context.Context, seller entity.Seller) (err error) {
	masterProducts, err := s.Sr.GetMasterData(ctx)
	if err != nil {
		return
	}

	for _, product := range masterProducts {
		err = s.Sr.StoreInitialMerchantItem(ctx, seller, product)
		if err != nil {
			err = errors.ErrUnknown
			return
		}
	}

	return
}
