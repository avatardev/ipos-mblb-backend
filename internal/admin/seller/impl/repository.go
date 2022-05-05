package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	mdEntity "github.com/avatardev/ipos-mblb-backend/internal/admin/product/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
)

type SellerRepositoryImpl struct {
	DB *sql.DB
}

func NewSellerRepository(db *database.DatabaseClient) SellerRepositoryImpl {
	return SellerRepositoryImpl{DB: db.DB}
}

var (
	COUNT_SELLER       = sq.Select(("COUNT(*)")).From("sellers s")
	SELECT_SELLER      = sq.Select("s.id", "s.perusahaan", "s.telp", "s.alamat", "s.kecamatan", "s.email", "s.name_pic", "s.hp_pic", "s.npwp", "s.ktp", "s.no_iup", "s.masa_berlaku", "s.keterangan", "s.status").From("sellers s")
	SELECT_SELLER_NAME = sq.Select("s.id", "s.perusahaan").From("sellers s")
	INSERT_SELLER      = sq.Insert("sellers").Columns("perusahaan", "telp", "alamat", "kecamatan", "email", "name_pic", "hp_pic", "npwp", "ktp", "no_iup", "masa_berlaku", "keterangan", "status", "created_at", "updated_at")
	UPDATE_SELLER      = sq.Update("sellers")
	DELETE_USER        = sq.Delete("users")
)

var (
	INSERT_MERCHANT_ITEM = sq.Insert("produk_sellers").Columns("id_produk", "harga", "status", "id_seller", "updated_at", "created_at")
	SELECT_MASTER_DATA   = sq.Select("p.id", "p.harga_std_m3").From("produks p")
)

func (sr SellerRepositoryImpl) Count(ctx context.Context, keyword string) (uint64, error) {
	stmt, params, err := COUNT_SELLER.Where(sq.And{sq.Like{"perusahaan": fmt.Sprintf("%%%s%%", keyword)}, sq.Eq{"s.deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Seller.Count] error: %v\n", err)
		return 0, err
	}

	prpd, err := sr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Seller.Count] error: %v\n", err)
		return 0, err
	}

	var sellerCount uint64
	queryErr := prpd.QueryRowContext(ctx, params...).Scan(&sellerCount)
	if err != nil {
		log.Printf("[Seller.Count] error: %v\n", queryErr)
		return 0, queryErr
	}

	return sellerCount, nil
}

func (sr SellerRepositoryImpl) GetSelerName(ctx context.Context) (entity.CompanySellers, error) {
	stmt, params, err := SELECT_SELLER_NAME.Where(sq.Eq{"s.deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[Seller.GetSelerName] error: %v\n", err)
		return nil, err
	}

	prpd, err := sr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Seller.GetSelerName] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Seller.GetSelerName] error: %v\n", err)
		return nil, err
	}

	cs := entity.CompanySellers{}

	for rows.Next() {
		seller := &entity.CompanySeller{}
		if err := rows.Scan(&seller.Id, &seller.Company); err != nil {
			log.Printf("[Seller.GetSelerName] error: %v\n", err)
			return nil, err
		}
		cs = append(cs, seller)
	}

	return cs, nil
}

func (sr SellerRepositoryImpl) GetAll(ctx context.Context, keyword string, limit uint64, offset uint64) (entity.Sellers, error) {
	stmt, params, err := SELECT_SELLER.Where(sq.And{sq.Like{"perusahaan": fmt.Sprintf("%%%s%%", keyword)}, sq.Eq{"s.deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Seller.GetAll] error: %v\n", err)
		return nil, err
	}

	prpd, err := sr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Seller.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Seller.GetAll] error: %v\n", err)
		return nil, err
	}

	sellers := entity.Sellers{}

	for rows.Next() {
		var seller entity.Seller
		err := rows.Scan(&seller.Id, &seller.Company, &seller.Phone, &seller.Address, &seller.District, &seller.Email, &seller.PICName, &seller.PICPhone, &seller.NPWP, &seller.KTP, &seller.NoIUP, &seller.ValidPeriod, &seller.Description, &seller.Status)
		if err != nil {
			log.Printf("[Seller.GetAll] error: %v\n", err)
			return nil, err
		}
		sellers = append(sellers, &seller)
	}

	return sellers, nil
}

func (sr SellerRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Seller, error) {
	stmt, params, err := SELECT_SELLER.Where(sq.And{sq.Eq{"id": id}, sq.Eq{"s.deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Seller.GetById] id: %v, error: %v\n", id, err)
		return nil, err
	}

	prpd, err := sr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Seller.GetById] id: %v, error: %v\n", id, err)
		return nil, err
	}

	seller := &entity.Seller{}

	rows := prpd.QueryRowContext(ctx, params...)
	queryErr := rows.Scan(&seller.Id, &seller.Company, &seller.Phone, &seller.Address, &seller.District, &seller.Email, &seller.PICName, &seller.PICPhone, &seller.NPWP, &seller.KTP, &seller.NoIUP, &seller.ValidPeriod, &seller.Description, &seller.Status)
	if queryErr != nil && &queryErr != &sql.ErrNoRows {
		log.Printf("[Seller.GetById] id: %v, error: %v\n", id, queryErr)
		return nil, queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Seller.GetById] id: %v, error: %v\n", id, queryErr)
		return nil, nil
	}

	return seller, nil
}

func (sr SellerRepositoryImpl) Store(ctx context.Context, seller entity.Seller) (*entity.Seller, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_SELLER.Values(seller.Company, seller.Phone, seller.Address, seller.District, seller.Email, seller.PICName, seller.PICPhone, seller.NPWP, seller.KTP, seller.NoIUP, seller.ValidPeriod, seller.Description, seller.Status, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Seller.Store] error: %v\n", err)
		return nil, err
	}

	prpd, err := sr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Seller.Store] error: %v\n", err)
		return nil, err
	}

	res, err := prpd.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("[Seller.Store] error: %v\n", err)
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[Seller.Store] error: %v\n", err)
		return nil, err
	}

	return sr.GetById(ctx, lid)
}

func (sr SellerRepositoryImpl) Update(ctx context.Context, seller entity.Seller) (*entity.Seller, error) {
	updateMap := map[string]interface{}{
		"perusahaan": seller.Company,
		"telp":       seller.Phone,
		"alamat":     seller.Address,
		"kecamatan":  seller.District,
		"email":      seller.Email,
		"name_pic":   seller.PICName,
		"hp_pic":     seller.PICPhone,
		"npwp":       seller.NPWP,
		"ktp":        seller.KTP,
		"no_iup":     seller.NoIUP,
		"keterangan": seller.Description,
		"status":     seller.Status,
		"updated_at": time.Now(),
	}

	stmt, params, err := UPDATE_SELLER.SetMap(updateMap).Where(sq.Eq{"id": seller.Id}).ToSql()
	if err != nil {
		log.Printf("[Seller.Update] id: %v, error: %v\n", seller.Id, err)
		return nil, err
	}

	prpd, err := sr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Seller.Update] id: %v, error: %v\n", seller.Id, err)
		return nil, err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Seller.Update] id: %v, error: %v\n", seller.Id, err)
		return nil, err
	}

	return sr.GetById(ctx, seller.Id)
}

func (sr SellerRepositoryImpl) Delete(ctx context.Context, id int64) error {
	currTime := time.Now()
	updateMap := map[string]interface{}{
		"status":     false,
		"updated_at": currTime,
		"deleted_at": currTime,
	}

	stmt, params, err := UPDATE_SELLER.SetMap(updateMap).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Printf("[Seller.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	prpd, err := sr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Seller.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Seller.Delete] id: %v, error: %v\n", id, err)
		return err
	}

	return nil
}

func (sr SellerRepositoryImpl) GetMasterData(ctx context.Context) (products mdEntity.Products, err error) {
	stmt, args, err := SELECT_MASTER_DATA.Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[GetMasterData] err: %v\n", err)
		return
	}

	rows, err := sr.DB.QueryContext(ctx, stmt, args...)
	if err != nil {
		log.Printf("[GetMasterData] err: %v\n", err)
		return
	}

	return mapMPToEntity(rows)
}

func (sr SellerRepositoryImpl) DeleteUser(ctx context.Context, sellerID int64) error {
	stmt, params, err := DELETE_USER.Where(sq.And{sq.Eq{"id_seller": sellerID}, sq.Eq{"id_role": privutil.USER_SELLER}}).ToSql()
	if err != nil {
		log.Printf("[Seller.DeleteUser] id: %v, err: %v\n", sellerID, err)
		return err
	}

	prpd, err := sr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Seller.DeleteUser] id: %v, err: %v\n", sellerID, err)
		return err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Seller.DeleteUser] id: %v, err: %v\n", sellerID, err)
		return err
	}

	return nil
}

func (sr SellerRepositoryImpl) StoreInitialMerchantItem(ctx context.Context, seller entity.Seller, product *mdEntity.Product) (err error) {
	currTime := time.Now()
	stmt, args, err := INSERT_MERCHANT_ITEM.Values(product.Id, product.Price, 0, seller.Id, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[StoreInitialMerchantItem] err: %v\n", err)
		return
	}

	_, err = sr.DB.ExecContext(ctx, stmt, args...)
	if err != nil {
		log.Printf("[StoreInitialMerchantItem] err: %v\n", err)
		return
	}

	return
}

func mapMPToEntity(rows *sql.Rows) (p mdEntity.Products, err error) {
	p = mdEntity.Products{}

	for rows.Next() {
		temp := &mdEntity.Product{}
		if err = rows.Scan(&temp.Id, &temp.Price); err != nil {
			p = nil
			break
		}

		p = append(p, temp)
	}

	return
}
