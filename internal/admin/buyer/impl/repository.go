package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type BuyerRepositoryImpl struct {
	DB *sql.DB
}

var (
	COUNT_BUYER  = sq.Select("COUNT(*)").From("buyer_truks b").LeftJoin("kategori_kendaraans k ON b.kategori = k.id")
	SELECT_BUYER = sq.Select("b.plat_truk", "k.nama_kategori", "b.perusahaan", "b.telp", "b.alamat", "b.email", "b.name_pic", "b.hp_pic", "b.keterangan", "b.status").
			From("buyer_truks b").LeftJoin("kategori_kendaraans k ON b.kategori = k.id")
	INSERT_BUYER = sq.Insert("buyer_truks").Columns("plat_truk", "kategori", "perusahaan", "telp", "alamat", "email", "name_pic", "hp_pic", "keterangan", "status", "created_at", "updated_at")
)

func NewBuyerRepository(db *database.DatabaseClient) BuyerRepositoryImpl {
	return BuyerRepositoryImpl{DB: db.DB}
}

func (b *BuyerRepositoryImpl) Count(ctx context.Context, keyword string) (uint64, error) {
	stmt, params, err := COUNT_BUYER.Where(sq.Like{"b.plat_truk": fmt.Sprintf("%%%s%%", keyword)}).ToSql()
	if err != nil {
		log.Printf("[Buyer.Count] error: %v\n", err)
		return 0, err
	}

	prpd, err := b.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Buyer.Count] error: %v\n", err)
		return 0, err
	}

	var buyerCount uint64
	row := prpd.QueryRowContext(ctx, params...)
	queryErr := row.Scan(&buyerCount)
	if queryErr != nil {
		log.Printf("[Buyer.Count] error: %v\n", err)
		return 0, err
	}

	return buyerCount, nil
}

func (b *BuyerRepositoryImpl) GetAll(ctx context.Context, keyword string, limit uint64, offset uint64) (entity.Buyers, error) {
	stmt, params, err := SELECT_BUYER.Where(sq.Like{"b.plat_truk": fmt.Sprintf("%%%s%%", keyword)}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[Buyer.GetAll] error: %v\n", err)
		return nil, err
	}

	prpd, err := b.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Buyer.GetAll] error: %v\n", err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[Buyer.GetAll] error: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	buyers := entity.Buyers{}

	for rows.Next() {
		temp := &entity.Buyer{}
		err := rows.Scan(&temp.VehiclePlate, &temp.VehicleCategoryName, &temp.Company, &temp.Phone, &temp.Address, &temp.Email, &temp.PICName, &temp.PICPhone, &temp.Description, &temp.Status)
		if err != nil {
			log.Printf("[Buyer.GetAll] error: %v\n", err)
			return nil, err
		}

		buyers = append(buyers, temp)
	}

	return buyers, nil
}

func (b *BuyerRepositoryImpl) GetById(ctx context.Context, plate string) (*entity.Buyer, error) {
	stmt, params, err := SELECT_BUYER.Where(sq.Eq{"b.plat_truk": plate}).ToSql()
	if err != nil {
		log.Printf("[Buyer.GetById] plate: %s, error: %v\n", plate, err)
		return nil, err
	}

	prpd, err := b.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Buyer.GetById] plate: %s, error: %v\n", plate, err)
		return nil, err
	}

	rows := prpd.QueryRowContext(ctx, params...)

	buyer := &entity.Buyer{}
	queryErr := rows.Scan(&buyer.VehiclePlate, &buyer.VehicleCategoryName, &buyer.Company, &buyer.Phone, &buyer.Address, &buyer.Email, &buyer.PICName, &buyer.PICPhone, &buyer.Description, &buyer.Status)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[Buyer.GetById] plate: %v, error: %v\n", plate, err)
		return nil, err
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Buer.GetById] plate: %v, error: %v\n", plate, err)
		return nil, nil
	}

	return buyer, nil
}

func (b *BuyerRepositoryImpl) Store(ctx context.Context, buyer entity.Buyer) (*entity.Buyer, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_BUYER.Values(buyer.VehiclePlate, buyer.VehicleCategoryId, buyer.Company, buyer.Phone, buyer.Address, buyer.Email, buyer.PICName, buyer.PICPhone, buyer.Description, buyer.Status, currTime, currTime).ToSql()
	if err != nil {
		log.Printf("[Buyer.Store] error: %v\n", err)
		return nil, err
	}

	prpd, err := b.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[Buyer.Store] error: %v\n", err)
		return nil, err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[Buyer.Store] error: %v\n", err)
		return nil, err
	}

	return b.GetById(ctx, buyer.VehiclePlate)
}
