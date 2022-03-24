package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	pkgDto "github.com/avatardev/ipos-mblb-backend/pkg/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
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

func (b *BuyerServiceImpl) StoreBuyer(ctx context.Context, req *dto.BuyerRequest) (*dto.BuyerResponse, error) {
	buyer := req.ToEntity()

	data, err := b.Br.Store(ctx, buyer)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewBuyerResponse(*data), nil
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
