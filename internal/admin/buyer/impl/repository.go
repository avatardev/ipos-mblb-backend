package impl

import (
	"context"
	"database/sql"
	"log"

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
)

func (b *BuyerRepositoryImpl) Count(ctx context.Context) (uint64, error) {
	stmt, params, err := COUNT_BUYER.ToSql()
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

func (b *BuyerRepositoryImpl) GetAll(ctx context.Context, limit uint64, offset uint64) (entity.Buyers, error) {
	stmt, params, err := SELECT_BUYER.Limit(limit).Offset(offset).ToSql()
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

func NewBuyerRepository(db *database.DatabaseClient) BuyerRepositoryImpl {
	return BuyerRepositoryImpl{DB: db.DB}
}
