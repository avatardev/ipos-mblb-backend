package impl

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type OrderServiceImpl struct {
	Or OrderRepositoryImpl
}

func (o *OrderServiceImpl) GenerateDetailTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*bytes.Buffer, error) {
	trxData, err := o.Or.GetAll(ctx, dateStart, dateEnd)
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

func (o *OrderServiceImpl) DetailTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*dto.TrxDetailsJSON, error) {
	trxData, err := o.Or.GetAll(ctx, dateStart, dateEnd)
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

	return dto.NewTrxDetailsJSON(trxData), nil
}

func (o *OrderServiceImpl) GenerateBriefTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*bytes.Buffer, error) {
	trxData, err := o.Or.GetAll(ctx, dateStart, dateEnd)
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

func (o *OrderServiceImpl) BriefTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*dto.TrxBriefsJSON, error) {
	trxData, err := o.Or.GetAll(ctx, dateStart, dateEnd)
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

	return dto.NewTrxBriefsJSON(trxData), nil
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

func (o *OrderServiceImpl) DailyTrx(ctx context.Context, sellerId int64) (*dto.TrxDailiesJSON, error) {
	d := time.Now()

	data, err := o.Or.GetAllDaily(ctx, sellerId, FirstDayOfMonth(d), LastDayOfMonth(d))
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.ErrNotFound
	}

	return dto.NewTrxDailiesJSON(data), nil
}

func (o *OrderServiceImpl) GenerateDailyTrx(ctx context.Context, sellerId int64) (*bytes.Buffer, error) {
	d := time.Now()
	data, err := o.Or.GetAllDaily(ctx, sellerId, FirstDayOfMonth(d), LastDayOfMonth(d))
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.ErrNotFound
	}

	res := dto.NewTrxDailies(data)
	csvData := new(bytes.Buffer)
	w := csv.NewWriter(csvData)

	company, npwp, err := o.Or.GetCompanyName(ctx, sellerId)
	if err != nil {
		return nil, err
	}

	w.WriteAll(res.ToCSV(*company, *npwp, time.Now().Month().String(), time.Now().Year()))
	if err := w.Error(); err != nil {
		log.Printf("[GenerateDailyTrx] error: %v", err)
		return nil, err
	}

	return csvData, nil
}

func (o *OrderServiceImpl) MonitorTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*dto.TrxMonitorJSON, error) {
	trxData, err := o.Or.GetAllMonitored(ctx, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	if len(trxData) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewTrxMonitorJSON(trxData), nil
}

func (o *OrderServiceImpl) GenerateMonitorTrx(ctx context.Context, dateStart time.Time, dateEnd time.Time) (*bytes.Buffer, error) {
	trxData, err := o.Or.GetAllMonitored(ctx, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	if len(trxData) == 0 {
		return nil, errors.ErrInvalidResources
	}

	res := dto.NewTrxMonitors(trxData)
	csvData := new(bytes.Buffer)
	w := csv.NewWriter(csvData)

	w.WriteAll(res.ToCSV(dateStart, dateEnd))
	if err := w.Error(); err != nil {
		log.Printf("[GenerateMonitorTrx] error: %v", err)
		return nil, err
	}

	return csvData, nil
}

func FirstDayOfMonth(date time.Time) time.Time {
	y, m, d := date.AddDate(0, 0, -date.Day()+1).Date()
	t, err := time.Parse("2-1-2006", fmt.Sprintf("%v-%v-%v", d, int(m), y))
	log.Println(err)
	return t
}

func LastDayOfMonth(date time.Time) time.Time {
	y, m, d := date.AddDate(0, 1, -date.Day()).Date()
	t, _ := time.Parse("2-1-2006", fmt.Sprintf("%v-%v-%v", d, int(m), y))
	return t
}
