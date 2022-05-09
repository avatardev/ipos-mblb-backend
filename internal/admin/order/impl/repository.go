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
	COUNT_ORDER_DETAIL  = sq.Select("COUNT(*)").From("orders o")
	SELECT_ORDER_DETAIL = sq.Select("o.id", "o.order_date", "s.perusahaan", "o.plat_truk", "p.pembayaran", "d.id_ps", "d.id_ps_update", "d.qty", "d.qty_update", "s.status", "d.disc", "d.total_pajak",
		"d.pajak_update", "d.total_tagihan", "d.catatan_transaksi", "d.created_at", "d.updated_at").From("orders o").
		LeftJoin("sellers s ON s.id = o.id_seller").LeftJoin("method_payments p ON p.id = o.id_payment").
		LeftJoin("order_details d ON d.id_order = o.id")
	SELECT_ORDER_MONITORING = sq.Select("o.id", "o.order_date", "s.perusahaan", "o.plat_truk", "p_a.nama_produk", "p_b.nama_produk", "d.qty", "p_a.harga_std_m3", "ps_a.harga", "d.qty_update", "d.total_pajak",
		"d.pajak_update", "d.created_at", "d.updated_at").From("orders o").
		LeftJoin("sellers s ON s.id = o.id_seller").LeftJoin("order_details d ON d.id_order = o.id").
		LeftJoin("produk_sellers ps_a ON ps_a.id = d.id_ps").LeftJoin("produks p_a ON p_a.id = ps_a.id_produk").
		LeftJoin("produk_sellers ps_b ON ps_b.id = d.id_ps_update").LeftJoin("produks p_b ON p_b.id = ps_b.id_produk")
	SELECT_SELLER_DATA  = sq.Select("s.perusahaan", "s.npwp").From("sellers s")
	SELECT_PRODUCT_NAME = sq.Select(`(case when p.deleted_at IS NOT NULL then concat(p.nama_produk, " (deleted)") else p.nama_produk end)`).
				From("produk_sellers ps").LeftJoin("produks p ON p.id = ps.id_produk")
	// SELECT_DAILY_TRX = sq.Select("date(o.order_date) order_date", "(case when d.qty_update != 0 then d.qty_update else d.qty end) volume", "count(*)").
	// 			From("orders o").LeftJoin("order_details d ON d.id_order=o.id").GroupBy("volume", "order_date")
	SELECT_DAILY_TRX = sq.Select("date(o.order_date) order_date", "d.qty volume", "count(*)").
				From("orders o").LeftJoin("order_details d ON d.id_order=o.id").GroupBy("volume", "order_date")
	INSERT_NOTE = sq.Update("order_details")
)

func NewOrderRepository(db *database.DatabaseClient) OrderRepositoryImpl {
	return OrderRepositoryImpl{DB: db.DB}
}

func (o *OrderRepositoryImpl) Count(ctx context.Context, start time.Time, end time.Time) (uint64, error) {
	stmt, params, err := COUNT_ORDER_DETAIL.Where(sq.And{sq.GtOrEq{"o.order_date": start}, sq.LtOrEq{"o.order_date": end}}).OrderBy("o.order_date").ToSql()
	if err != nil {
		log.Printf("[Trx.Count] error: %v\n", err)
		return 0, err
	}

	prpd, err := o.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Trx.Count] error: %v\n", err)
		return 0, err
	}

	var orderCount uint64
	row := prpd.QueryRowContext(ctx, params...)
	queryErr := row.Scan(&orderCount)
	if queryErr != nil {
		log.Printf("[Trx.Count] error: %v\n", err)
		return 0, err
	}

	return orderCount, nil
}

func (o *OrderRepositoryImpl) GetById(ctx context.Context, orderId int64) (*entity.TrxDetail, error) {
	stmt, params, err := SELECT_ORDER_DETAIL.Where(sq.Eq{"o.id": orderId}).ToSql()
	if err != nil {
		log.Printf("[Trx.GetById] error: %v\n", err)
		return nil, err
	}

	prpd, err := o.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Trx.GetById] error: %v\n", err)
		return nil, err
	}

	rows := prpd.QueryRowContext(ctx, params...)

	item := &entity.TrxDetail{}
	queryErr := item.FromSingleSql(rows)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[Trx.GetById] error: %v\n", queryErr)
		return nil, queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Trx.GetById] error: %v\n", queryErr)
		return nil, nil
	}

	return item, nil
}

func (o *OrderRepositoryImpl) GetAll(ctx context.Context, start time.Time, end time.Time, id int64) (entity.TrxDetails, error) {
	query := SELECT_ORDER_DETAIL
	if id != 0 {
		query = query.Where(sq.And{sq.GtOrEq{"o.order_date": start}, sq.LtOrEq{"o.order_date": end}, sq.Eq{"o.id_seller": id}})
	} else {
		query = query.Where(sq.And{sq.GtOrEq{"o.order_date": start}, sq.LtOrEq{"o.order_date": end}})
	}

	stmt, params, err := query.OrderBy("o.order_date").ToSql()
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

func (o *OrderRepositoryImpl) GetAllDaily(ctx context.Context, sellerId int64, start time.Time, end time.Time) (entity.TrxDailies, error) {
	stmt, params, err := SELECT_DAILY_TRX.Where(sq.And{sq.GtOrEq{"o.order_date": start}, sq.LtOrEq{"o.order_date": end}, sq.Eq{"o.id_seller": sellerId}}).OrderBy("o.order_date").ToSql()
	if err != nil {
		log.Printf("[Trx.GetAllDaily] error: %v\n", err)
		return nil, err
	}

	prpd, err := o.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Trx.GetAllDaily] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Trx.GetAllDaily] error: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	data := entity.TrxDailies{}

	for rows.Next() {
		temp := &entity.TrxDaily{}
		err := temp.FromSql(rows)
		if err != nil {
			log.Printf("[Trx.GetAllDaily] error: %v\n", err)
			return nil, err
		}
		data = append(data, temp)
	}

	return data, nil
}

func (o *OrderRepositoryImpl) GetAllMonitored(ctx context.Context, start time.Time, end time.Time, id int64) (entity.TrxMonitors, error) {
	query := SELECT_ORDER_MONITORING
	if id != 0 {
		query = query.Where(sq.And{sq.GtOrEq{"o.order_date": start}, sq.LtOrEq{"o.order_date": end}, sq.Eq{"o.id_seller": id}})
	} else {
		query = query.Where(sq.And{sq.GtOrEq{"o.order_date": start}, sq.LtOrEq{"o.order_date": end}})
	}

	stmt, params, err := query.OrderBy("o.order_date").ToSql()
	if err != nil {
		log.Printf("[Trx.GetAllMonitor] error: %v\n", err)
		return nil, err
	}

	prpd, err := o.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Trx.GetAllMonitor] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Trx.GetAllMonitor] error: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	data := entity.TrxMonitors{}

	for rows.Next() {
		temp := &entity.TrxMonitor{}
		err := temp.FromOrderSql(rows)
		if err != nil {
			log.Printf("[Trx.GetAllMonitor] error: %v\n", err)
			return nil, err
		}
		data = append(data, temp)
	}

	return data, nil
}

func (o *OrderRepositoryImpl) GetProductName(ctx context.Context, ps int64) (string, error) {
	stmt, params, err := SELECT_PRODUCT_NAME.Where(sq.And{sq.Eq{"ps.id": ps}, sq.Eq{"ps.deleted_at": nil}}).ToSql()
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
		return "", queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Trx.GetProductName] ps: %v, error: %v\n", ps, queryErr)
		return "", nil
	}

	return pName, nil
}

func (o *OrderRepositoryImpl) GetCompanyName(ctx context.Context, s int64) (*string, *int64, error) {
	stmt, params, err := SELECT_SELLER_DATA.Where(sq.And{sq.Eq{"s.id": s}, sq.Eq{"s.deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Trx.GetCompanyName] s: %v, error: %v\n", s, err)
		return nil, nil, err
	}

	prpd, err := o.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Trx.GetCompanyName] s: %v, error: %v\n", s, err)
		return nil, nil, err
	}

	row := prpd.QueryRowContext(ctx, params...)

	var company string
	var npwp int64

	queryErr := row.Scan(&company, &npwp)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[Trx.GetCompanyName] s: %v, error: %v\n", s, queryErr)
		return nil, nil, queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Trx.GetCompanyName] s: %v, error: %v\n", s, queryErr)
		return nil, nil, queryErr
	}

	return &company, &npwp, nil
}

func (o *OrderRepositoryImpl) InsertNote(ctx context.Context, orderId int64, note string) (*entity.TrxDetail, error) {
	updateMap := map[string]interface{}{
		"catatan_transaksi": note,
		"updated_at":        time.Now(),
	}

	stmt, params, err := INSERT_NOTE.SetMap(updateMap).Where(sq.Eq{"order_details.id_order": orderId}).ToSql()
	if err != nil {
		log.Printf("[Trx.InsertNote] error: %v\n", err)
		return nil, err
	}

	prpd, err := o.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Trx.InsertNote] error: %v\n", err)
		return nil, err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Trx.InsertNote] error: %v\n", err)
		return nil, err
	}

	return o.GetById(ctx, orderId)
}
