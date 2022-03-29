package impl

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/order/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type OrderRepositoryImpl struct {
	DB *sql.DB
}

var (
	SELECT_ORDER_DETAIL = sq.Select("o.order_date", "s.perusahaan", "o.plat_truk", "p.pembayaran", "d.id_ps", "d.id_ps_update", "d.qty", "d.qty_update", "s.status", "d.disc", "d.total_pajak",
		"d.pajak_update", "d.total_tagihan", "d.created_at", "d.updated_at").From("orders o").
		LeftJoin("sellers s ON s.id = o.id_seller").LeftJoin("method_payments p ON p.id = o.id_payment").
		LeftJoin("order_details d ON d.id_order = o.id")

	SELECT_PRODUCT_NAME = sq.Select("p.nama_produk").From("produk_sellers ps").LeftJoin("produks p ON p.id = ps.id_produk")
)

func NewOrderRepository(db *database.DatabaseClient) OrderRepositoryImpl {
	return OrderRepositoryImpl{DB: db.DB}
}

func (o *OrderRepositoryImpl) GetAll(ctx context.Context, start time.Time, end time.Time) (entity.TrxDetails, error) {
	stmt, params, err := SELECT_ORDER_DETAIL.Where(sq.And{sq.GtOrEq{"o.order_date": start}, sq.LtOrEq{"o.order_date": end}}).ToSql()
	if err != nil {
		log.Printf("[Trx.GetAll] error: %v\n", err)
		return nil, err
	}

	prpd, err := o.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Trx.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Trx.GetAll] error: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	data := entity.TrxDetails{}

	for rows.Next() {
		temp := &entity.TrxDetail{}
		err := temp.FromSql(rows)
		if err != nil {
			log.Printf("[Trx.GetAll] error: %v\n", err)
			return nil, err
		}

		data = append(data, temp)
	}

	return data, nil
}

func (o *OrderRepositoryImpl) GetProductName(ctx context.Context, ps int64) (string, error) {
	stmt, params, err := SELECT_PRODUCT_NAME.Where(sq.And{sq.Eq{"ps.id": ps}, sq.Eq{"ps.deleted_at": nil}, sq.Eq{"p.deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Trx.GetProductName] ps: %v, error: %v\n", ps, err)
		return "", err
	}

	prpd, err := o.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Trx.GetProductName] ps: %v, error: %v\n", ps, err)
		return "", err
	}

	row := prpd.QueryRowContext(ctx, params...)

	var pName string

	queryErr := row.Scan(&pName)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[Trx.GetProductName] ps: %v, error: %v\n", ps, queryErr)
		return "", err
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Trx.GetProductName] ps: %v, error: %v\n", ps, queryErr)
		return "", nil
	}

	return pName, nil
}