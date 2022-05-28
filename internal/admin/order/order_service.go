package order

import (
	"bytes"
	"context"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/order/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type OrderService interface {
	GenerateDetailTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time, companyName string) (*bytes.Buffer, error)
	DetailTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time, companyName string) (*dto.TrxDetailsJSON, error)
	GenerateBriefTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time, companyName string) (*bytes.Buffer, error)
	BriefTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time, companyName string) (*dto.TrxBriefsJSON, error)
	DailyTrx(ctx context.Context, sellerId int64) (*dto.TrxDailiesJSON, error)
	GenerateDailyTrx(ctx context.Context, sellerId int64) (*bytes.Buffer, error)
	MonitorTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time, companyName string) (*dto.TrxMonitorJSON, error)
	GenerateMonitorTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time, companyName string) (*bytes.Buffer, error)
	InsertNote(ctx context.Context, orderId int64, note string) (*dto.TrxDetail, error)
}

func NewOrderService(Or impl.OrderRepositoryImpl) OrderService {
	return &impl.OrderServiceImpl{Or: Or}
}
