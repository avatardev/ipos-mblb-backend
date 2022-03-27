package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/merchant/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type MerchantRepositoryImpl struct {
	DB *sql.DB
}

var (
	COUNT_ITEM  = sq.Select("COUNT(*)").From("produk_sellers").From("produk_sellers m").LeftJoin("produks p ON m.id_produk = p.id")
	SELECT_ITEM = sq.Select("m.id", "m.id_produk", "p.nama_produk", "m.harga", "m.keterangan", "m.status").From("produk_sellers m").LeftJoin("produks p ON m.id_produk = p.id")
	UPDATE_ITEM = sq.Update("produk_sellers")
)

func NewMerchantRepository(db *database.DatabaseClient) MerchantRepositoryImpl {
	return MerchantRepositoryImpl{DB: db.DB}
}

func (m *MerchantRepositoryImpl) Count(ctx context.Context, sellerId int64, keyword string) (uint64, error) {
	stmt, params, err := COUNT_ITEM.Where(sq.And{sq.Eq{"m.deleted_at": nil}, sq.Eq{"m.id_seller": sellerId}, sq.Like{"p.nama_produk": fmt.Sprintf("%%%s%%", keyword)}}).ToSql()
	if err != nil {
		log.Printf("[MerchantItem.Count] error: %v\n", err)
		return 0, err
	}

	prpd, err := m.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[MerchantItem.Count] error: %v\n", err)
		return 0, err
	}

	var itemCount uint64
	row := prpd.QueryRowContext(ctx, params...)
	queryErr := row.Scan(&itemCount)
	if queryErr != nil {
		log.Printf("[MerchantItem.Count] error: %v\n", err)
		return 0, queryErr
	}

	return itemCount, nil
}

func (m *MerchantRepositoryImpl) GetAll(ctx context.Context, sellerId int64, keyword string, limit uint64, offset uint64) (entity.MerchantItems, error) {
	stmt, params, err := SELECT_ITEM.Where(sq.And{sq.Eq{"m.deleted_at": nil}, sq.Eq{"m.id_seller": sellerId}, sq.Like{"p.nama_produk": fmt.Sprintf("%%%s%%", keyword)}}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[MerchantItem.GetAll] error: %v\n", err)
		return nil, err
	}

	prpd, err := m.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[MerchantItem.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[MerchantItem.GetAll] error: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	items := entity.MerchantItems{}

	for rows.Next() {
		temp := &entity.MerchantItem{}
		err := rows.Scan(&temp.Id, &temp.ProductId, &temp.Name, &temp.Price, &temp.Description, &temp.Status)
		if err != nil {
			log.Printf("[MerchantItem.GetAll] error: %v\n", err)
			return nil, err
		}

		items = append(items, temp)
	}

	return items, nil
}

func (m *MerchantRepositoryImpl) GetById(ctx context.Context, sellerId int64, itemId int64) (*entity.MerchantItem, error) {
	stmt, params, err := SELECT_ITEM.Where(sq.And{sq.Eq{"m.deleted_at": nil}, sq.Eq{"m.id_seller": sellerId}, sq.Eq{"m.id": itemId}}).ToSql()
	if err != nil {
		log.Printf("[MerchantItem.GetById] item: %v, error: %v\n", itemId, err)
		return nil, err
	}

	prpd, err := m.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[MerchantItem.GetById] item: %v, error: %v\n", itemId, err)
		return nil, err
	}

	rows := prpd.QueryRowContext(ctx, params...)

	item := &entity.MerchantItem{}
	queryErr := rows.Scan(&item.Id, &item.ProductId, &item.Name, &item.Price, &item.Description, &item.Status)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[MerchantItem.GetById] item: %v, error: %v\n", itemId, queryErr)
		return nil, queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[MerchantItem.GetById] item: %v, error: %v\n", itemId, queryErr)
		return nil, nil
	}

	return item, nil
}

func (m *MerchantRepositoryImpl) Update(ctx context.Context, sellerId int64, item entity.MerchantItem) (*entity.MerchantItem, error) {
	updateMap := map[string]interface{}{
		"id_produk":  item.ProductId,
		"harga":      item.Price,
		"keterangan": item.Description,
		"status":     item.Status,
		"updated_at": time.Now(),
	}

	log.Println(item.Price)

	stmt, params, err := UPDATE_ITEM.SetMap(updateMap).Where(sq.And{sq.Eq{"produk_sellers.id": item.Id}, sq.Eq{"produk_sellers.id_seller": sellerId}}).ToSql()
	if err != nil {
		log.Printf("[MerchantItem.Update] error: %v\n", err)
		return nil, err
	}

	prpd, err := m.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[MerchantItem.Update] error: %v\n", err)
		return nil, err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Buyer.Update] errror: %v\n", err)
		return nil, err
	}

	return m.GetById(ctx, sellerId, item.Id)
}
