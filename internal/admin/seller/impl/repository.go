package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type SellerRepositoryImpl struct {
	DB *sql.DB
}

func NewSellerRepository(db *database.DatabaseClient) SellerRepositoryImpl {
	return SellerRepositoryImpl{DB: db.DB}
}

var (
	COUNT_SELLER  = sq.Select(("COUNT(*)")).From("sellers s")
	SELECT_SELLER = sq.Select("s.id", "s.perusahaan", "s.telp", "s.alamat", "s.kecamatan", "s.email", "s.name_pic", "s.hp_pic", "s.npwp", "s.ktp", "s.no_iup", "s.masa_berlaku", "s.keterangan", "s.status").From("sellers s")
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
