package impl

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/product/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/internal/global/config"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/logutil"
)

type ProductServiceImpl struct {
	Pr ProductRepositoryImpl
}

func (p *ProductServiceImpl) GetProduct(ctx context.Context, query string, limit uint64, offset uint64) (*dto.ProductsResponse, error) {
	productCount, err := p.Pr.Count(ctx, query)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if productCount == 0 {
		return nil, errors.ErrInvalidResources
	}

	products, err := p.Pr.GetAll(ctx, query, limit, offset)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if len(products) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewProductsResponse(products, limit, offset, productCount), nil
}

func (p *ProductServiceImpl) GetProductById(ctx context.Context, id int64) (*dto.ProductResponse, error) {
	product, err := p.Pr.GetById(ctx, id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if product == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewProductResponse(*product), nil
}

func (p *ProductServiceImpl) StoreProduct(ctx context.Context, req *dto.ProductRequest) (res *dto.ProductResponse, err error) {
	product := req.ToEntity()

	data, err := p.Pr.Store(ctx, product)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	conf := config.GetConfig()
	fName := fmt.Sprintf("img-product-%d.png", data.Id)

	img, err := os.ReadFile("./static/default_product.png")
	if err != nil {
		log.Printf("[StoreProduct] failed to open default pics, err => %+v\n", err)
		return
	}

	data, err = p.Pr.UpdateImage(ctx, data.Id, fmt.Sprintf("%s/private/nota/%s", conf.BaseURL, fName))
	if err != nil {
		log.Printf("[StoreProduct] failed to insert new img, err => %+v\n", err)
		return
	}

	err = os.WriteFile(filepath.Join(conf.LocalRepo, fName), img, 0666)
	if err != nil {
		log.Printf("[StoreProduct] failed to write data to local storage, err => %+v\n", err)
		return
	}

	sellers, err := p.Pr.FindActiveSeller(ctx)
	if err != nil {
		res = nil
		return
	}

	p.insertNewMerchantItem(ctx, sellers, *data)

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("added new product %s", req.Name))
	return dto.NewProductResponse(*data), nil
}

func (p *ProductServiceImpl) UpdateProduct(ctx context.Context, id int64, req *dto.ProductRequest) (*dto.ProductResponse, error) {
	product := req.ToEntity()
	product.Id = id

	exists, err := p.Pr.GetById(ctx, id)
	if err != nil {
		return nil, errors.ErrUnknown
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	data, err := p.Pr.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("changed product data %s", req.Name))
	return dto.NewProductResponse(*data), nil
}

func (p *ProductServiceImpl) EditProductImage(ctx context.Context, id int64, img *bytes.Buffer, fName string) (res *dto.ProductResponse, err error) {
	exists, err := p.Pr.GetById(ctx, id)
	if err != nil {
		err = errors.ErrUnknown
		return
	} else if exists == nil {
		err = errors.ErrInvalidResources
		return
	}

	conf := config.GetConfig()
	fName = fmt.Sprintf("img-product-%d%s", id, filepath.Ext(fName))

	data, err := p.Pr.UpdateImage(ctx, id, fmt.Sprintf("%s/private/nota/%s", conf.BaseURL, fName))
	if err != nil {
		log.Printf("[EditProductImage] failed to insert new img, err => %+v\n", err)
		return
	}

	err = os.WriteFile(filepath.Join(conf.LocalRepo, fName), img.Bytes(), 0666)
	if err != nil {
		log.Printf("[EditProductImage] failed to write data to local storage, err => %+v\n", err)
		return
	}

	return dto.NewProductResponse(*data), nil
}

func (p *ProductServiceImpl) DeleteProduct(ctx context.Context, id int64) error {
	exists, err := p.Pr.GetById(ctx, id)
	if err != nil {
		return errors.ErrUnknown
	}

	if exists == nil {
		return errors.ErrNotFound
	}

	if err := p.Pr.Delete(ctx, id); err != nil {
		return errors.ErrUnknown
	}

	logutil.GenerateActivityLog(ctx, fmt.Sprintf("deleted product data %s", exists.Name))
	return nil
}

func (p *ProductServiceImpl) insertNewMerchantItem(ctx context.Context, sellers []int64, product entity.Product) (err error) {
	for _, seller := range sellers {
		err = p.Pr.StoreNewMerchantItem(ctx, seller, product)
		if err != nil {
			err = errors.ErrUnknown
			return
		}
	}

	return
}
