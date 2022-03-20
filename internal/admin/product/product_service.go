package product

import "github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"

type ProductService interface {
}

func NewProductService(Pr *impl.ProductRepositoryImpl) ProductService {
	return &impl.ProductServiceImpl{Pr: Pr}
}
