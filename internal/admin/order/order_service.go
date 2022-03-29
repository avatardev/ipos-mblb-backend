package order

import (
	"bytes"
	"context"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/order/impl"
)

type OrderService interface {
	GenerateDetailTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*bytes.Buffer, error)
}

func NewOrderService(Or impl.OrderRepositoryImpl) OrderService {
	return &impl.OrderServiceImpl{Or: Or}
}
