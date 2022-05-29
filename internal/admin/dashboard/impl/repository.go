package impl

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/dashboard/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type DashboardRepositoryImpl struct {
	DB *sql.DB
}

func NewDashboardRepository(db *database.DatabaseClient) DashboardRepositoryImpl {
	return DashboardRepositoryImpl{DB: db.DB}
}

var (
	COUNT_BUYER  = sq.Select("count(*)").From("buyer_truks b")
	COUNT_SELLER = sq.Select("count(*)").From("sellers s")
	COUNT_TRX    = sq.Select("count(*)", "sum(case when d.pajak_update != 0 then d.pajak_update else d.total_pajak end)").
			From("orders o").LeftJoin("order_details d ON o.id = d.id_order")
)

func (r *DashboardRepositoryImpl) GetInfo(ctx context.Context, sellerID int64) (info *entity.DashboardInfo, err error) {
	info = &entity.DashboardInfo{}

	buyer, err := r.countBuyer(ctx)
	if err != nil {
		return nil, err
	} else {
		info.BuyerCount = buyer
	}

	if sellerID != 0 {
		info.SellerCount = 1
	} else {
		seller, err := r.countSeller(ctx)
		if err != nil {
			return nil, err
		} else {
			info.SellerCount = seller
		}
	}

	info.TrxCount, info.TotalTax, err = r.countTrx(ctx, sellerID)
	if err != nil {
		info = nil
	}

	return
}

func (r *DashboardRepositoryImpl) countBuyer(ctx context.Context) (int64, error) {
	stmt, params, err := COUNT_BUYER.Where(sq.And{sq.Eq{"b.status": 1}, sq.Eq{"b.deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Dashboard.CountBuyer] err: %v\n", err)
		return 0, err
	}

	var buyers int64
	row := r.DB.QueryRowContext(ctx, stmt, params...)
	if err := row.Scan(&buyers); err != nil {
		log.Printf("[Dashboard.CountBuyer] err: %v\n", err)
		return 0, err
	}

	return buyers, nil
}

func (r *DashboardRepositoryImpl) countSeller(ctx context.Context) (int64, error) {
	stmt, params, err := COUNT_SELLER.Where(sq.Eq{"s.deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[Dashboard.CountSeller] err: %v\n", err)
		return 0, err
	}

	var sellers int64
	row := r.DB.QueryRowContext(ctx, stmt, params...)
	if err := row.Scan(&sellers); err != nil {
		log.Printf("[Dashboard.CountSeller] err: %v\n", err)
		return 0, err
	}

	return sellers, nil
}

func (r *DashboardRepositoryImpl) countTrx(ctx context.Context, sellerID int64) (int64, int64, error) {
	baseQuery := COUNT_TRX

	if sellerID != 0 {
		baseQuery = baseQuery.Where(sq.And{sq.Eq{"o.id_seller": sellerID}, sq.Eq{"o.deleted_at": nil}})
	} else {
		baseQuery = baseQuery.Where(sq.Eq{"o.deleted_at": nil})
	}

	stmt, params, err := baseQuery.ToSql()
	if err != nil {
		log.Printf("[Dashboard.CountTrx] err: %v\n", err)
		return 0, 0, err
	}

	var trx, tax *int64
	row := r.DB.QueryRowContext(ctx, stmt, params...)
	if err := row.Scan(&trx, &tax); err != nil {
		log.Printf("[Dashboard.CountTrx] err: %v\n", err)
		return 0, 0, err
	}

	if trx == nil || tax == nil {
		return 0, 0, nil
	}

	return *trx, *tax, nil
}
