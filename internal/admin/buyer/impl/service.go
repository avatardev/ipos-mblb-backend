package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	pkgDto "github.com/avatardev/ipos-mblb-backend/pkg/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/logutil"
)

type BuyerServiceImpl struct {
	Br BuyerRepositoryImpl
}

func (b *BuyerServiceImpl) GetBuyer(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.BuyersResponse, error) {
	count, err := b.Br.Count(ctx, keyword)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrInvalidResources
	}

	buyers, err := b.Br.GetAll(ctx, keyword, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(buyers) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewBuyersResponse(buyers, count, limit, offset), nil
}

func (b *BuyerServiceImpl) GetBuyerName(ctx context.Context) (*dto.BuyersCompanyResponse, error) {
	buyers, err := b.Br.GetBuyerName(ctx)
	if err != nil {
		return nil, err
	}

	if len(buyers) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewBuyersCompanyResponse(buyers), nil
}

func (b *BuyerServiceImpl) GetBuyerById(ctx context.Context, plate string) (*dto.BuyerResponse, error) {
	buyer, err := b.Br.GetById(ctx, plate)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if buyer == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewBuyerResponse(*buyer), nil
}

func (b *BuyerServiceImpl) StoreBuyer(ctx context.Context, req *dto.BuyerRequest) (*dto.BuyerResponse, error) {
	buyer := req.ToEntity()

	exists, err := b.Br.GetById(ctx, buyer.VehiclePlate)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if exists != nil {
		return nil, errors.ErrUserExisted
	}

	data, err := b.Br.Store(ctx, buyer)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("added new buyer with vehicle plate %s", req.VehiclePlate))
	return dto.NewBuyerResponse(*data), nil
}

func (b *BuyerServiceImpl) UpdateBuyer(ctx context.Context, plate string, req *dto.BuyerRequest) (*dto.BuyerResponse, error) {
	buyer := req.ToEntity()

	exists, err := b.Br.GetById(ctx, plate)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	data, err := b.Br.Update(ctx, plate, buyer)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("changed buyer data with vehicle plate %s", plate))
	return dto.NewBuyerResponse(*data), nil
}

func (b *BuyerServiceImpl) DeleteBuyer(ctx context.Context, plate string) error {
	exists, err := b.Br.GetById(ctx, plate)
	if err != nil {
		return errors.ErrUnknown
	}

	if exists == nil {
		return errors.ErrNotFound
	}

	if err := b.Br.Delete(ctx, plate); err != nil {
		return errors.ErrUnknown
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("deleted buyer data with vehicle plate %s", plate))
	if err := b.Br.DeleteUser(ctx, plate); err != nil {
		return err
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("deleted user(s) with vehicle plate %s", plate))
	return nil
}

func (b *BuyerServiceImpl) Ping(ctx context.Context) pkgDto.PingResponse {
	return pkgDto.PingResponse{
		Message:         "pong",
		ServerTimestamp: fmt.Sprint(time.Now().UnixMilli()),
	}
}

func (b *BuyerServiceImpl) PingError(ctx context.Context) error {
	return errors.ErrBuyerPing
}
