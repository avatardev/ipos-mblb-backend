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
	COUNT_TRX    = sq.Select("count(*)", "sum(case when d.pajak_update != 0 then d.pajak_update else d.total_pajak end)").From("order_details d")
)

func (r *DashboardRepositoryImpl) GetInfo(ctx context.Context) (*entity.DashboardInfo, error) {
	info := &entity.DashboardInfo{}

	buyer, err := r.countBuyer(ctx)
	if err != nil {
		return nil, err
	} else {
		info.BuyerCount = buyer
	}

	seller, err := r.countSeller(ctx)
	if err != nil {
		return nil, err
	} else {
		info.SellerCount = seller
	}

	trxCount, taxSum, err := r.countTrx(ctx)
	if err != nil {
		return nil, err
	} else {
		info.TotalTax = taxSum
		info.TrxCount = trxCount
	}

	return info, nil
}

func (r *DashboardRepositoryImpl) countBuyer(ctx context.Context) (int64, error) {
	stmt, params, err := COUNT_BUYER.Where(sq.Eq{"b.status": 1}).ToSql()
	if err != nil {
		log.Printf("[Dashboard.CountBuyer] err: %v\n", err)
		return 0, err
	}

	prpd, err := r.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Dashboard.CountBuyer] err: %v\n", err)
		return 0, err
	}

	var buyers int64
	row := prpd.QueryRowContext(ctx, params...)
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

	prpd, err := r.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Dashboard.CountSeller] err: %v\n", err)
		return 0, err
	}

	var sellers int64
	row := prpd.QueryRowContext(ctx, params...)
	if err := row.Scan(&sellers); err != nil {
		log.Printf("[Dashboard.CountSeller] err: %v\n", err)
		return 0, err
	}

	return sellers, nil
}

func (r *DashboardRepositoryImpl) countTrx(ctx context.Context) (int64, int64, error) {
	stmt, params, err := COUNT_TRX.ToSql()
	if err != nil {
		log.Printf("[Dashboard.CountTrx] err: %v\n", err)
		return 0, 0, err
	}

	prpd, err := r.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Dastboard.CountTrx] err: %v\n", err)
		return 0, 0, err
	}

	var trx, tax int64
	row := prpd.QueryRowContext(ctx, params...)
	if err := row.Scan(&trx, &tax); err != nil {
		log.Printf("[Dashboard.CountTrx] err: %v\n", err)
		return 0, 0, err
	}

	return trx, tax, nil
}
