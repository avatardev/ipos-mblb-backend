package impl

import (
	"bytes"
	"context"
	"encoding/csv"
	"log"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type OrderServiceImpl struct {
	Or OrderRepositoryImpl
}

func (o *OrderServiceImpl) GenerateDetailTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*bytes.Buffer, error) {
	trxData, err := o.Or.GetAll(ctx, dateStart, dateEnd, 0, 0)
	if err != nil {
		return nil, err
	}

	if len(trxData) == 0 {
		return nil, errors.ErrInvalidResources
	}

	for _, trx := range trxData {
		productId := trx.ProductId
		if trx.ProductIdUpdate != 0 {
			productId = trx.ProductIdUpdate
		}

		pName, err := o.Or.GetProductName(ctx, productId)
		if err != nil || pName == "" {
			return nil, err
		}

		trx.Product = pName
	}

	res := dto.NewTrxDetails(trxData)

	csvData := new(bytes.Buffer)
	w := csv.NewWriter(csvData)

	w.WriteAll(res.ToCSV(dateStart, dateEnd))
	if err := w.Error(); err != nil {
		log.Printf("[GenerateDetailTrx] error: %v", err)
		return nil, err
	}

	return csvData, nil
}

func (o *OrderServiceImpl) DetailTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time, offset uint64, limit uint64) (*dto.TrxDetailsJSON, error) {
	count, err := o.Or.Count(ctx, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	trxData, err := o.Or.GetAll(ctx, dateStart, dateEnd, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(trxData) == 0 {
		return nil, errors.ErrInvalidResources
	}

	for _, trx := range trxData {
		productId := trx.ProductId
		if trx.ProductIdUpdate != 0 {
			productId = trx.ProductIdUpdate
		}

		pName, err := o.Or.GetProductName(ctx, productId)
		if err != nil || pName == "" {
			return nil, err
		}

		trx.Product = pName
	}

	return dto.NewTrxDetailsJSON(trxData, limit, offset, count), nil
}

func (o *OrderServiceImpl) GenerateBriefTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*bytes.Buffer, error) {
	trxData, err := o.Or.GetAll(ctx, dateStart, dateEnd, 0, 0)
	if err != nil {
		return nil, err
	}

	if len(trxData) == 0 {
		return nil, errors.ErrInvalidResources
	}

	for _, trx := range trxData {
		productId := trx.ProductId
		if trx.ProductIdUpdate != 0 {
			productId = trx.ProductIdUpdate
		}

		pName, err := o.Or.GetProductName(ctx, productId)
		if err != nil || pName == "" {
			return nil, err
		}

		trx.Product = pName
	}

	res := dto.NewTrxBriefs(trxData)

	csvData := new(bytes.Buffer)
	w := csv.NewWriter(csvData)

	w.WriteAll(res.ToCSV(dateStart, dateEnd))
	if err := w.Error(); err != nil {
		log.Printf("[GenerateBriefTrx] error: %v", err)
		return nil, err
	}

	return csvData, nil
}

func (o *OrderServiceImpl) BriefTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time, offset uint64, limit uint64) (*dto.TrxBriefsJSON, error) {
	count, err := o.Or.Count(ctx, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	trxData, err := o.Or.GetAll(ctx, dateStart, dateEnd, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(trxData) == 0 {
		return nil, errors.ErrInvalidResources
	}

	for _, trx := range trxData {
		productId := trx.ProductId
		if trx.ProductIdUpdate != 0 {
			productId = trx.ProductIdUpdate
		}

		pName, err := o.Or.GetProductName(ctx, productId)
		if err != nil || pName == "" {
			return nil, err
		}

		trx.Product = pName
	}

	return dto.NewTrxBriefsJSON(trxData, limit, offset, count), nil
}

func (o *OrderServiceImpl) InsertNote(ctx context.Context, orderId int64, note string) (*dto.TrxDetail, error) {
	exists, err := o.Or.GetById(ctx, orderId)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	data, err := o.Or.InsertNote(ctx, orderId, note)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewTrxDetail(data), nil
}
