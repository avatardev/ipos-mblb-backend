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

func (b *BuyerServiceImpl) GetBuyer(ctx context.Context, limit uint64, offset uint64) (*dto.BuyersResponse, error) {
	count, err := b.Br.Count(ctx)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrInvalidResources
	}

	buyers, err := b.Br.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(buyers) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewBuyersResponse(buyers, count, limit, offset), nil
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
