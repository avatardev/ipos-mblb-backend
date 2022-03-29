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
