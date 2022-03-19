package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/avatardev/ipos-mblb-backend/pkg/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type BuyerServiceImpl struct {
	// Repository GOes here
}

func (b *BuyerServiceImpl) Ping(ctx context.Context) dto.PingResponse {
	return dto.PingResponse{
		Message:         "pong",
		ServerTimestamp: fmt.Sprint(time.Now().UnixMilli()),
	}
}

func (b *BuyerServiceImpl) PingError(ctx context.Context) error {
	return errors.ErrBuyerPing
}
