package router

import (
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer"
	buyerPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/impl"
	bCategory "github.com/avatardev/ipos-mblb-backend/internal/admin/buyer_category"
	bCategoryPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/buyer_category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/product"
	productPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"
	pCategory "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category"
	pCategoryPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *database.DatabaseClient) {
	buyerRepository := buyerPkg.NewBuyerRepository(db)
	buyerService := buyer.NewBuyerService(buyerRepository)
	buyerHandler := buyer.NewBuyerHandler(buyerService)

	r.HandleFunc(AdminPing, buyerHandler.PingBuyer()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminPingError, buyerHandler.PingError()).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc(AdminBuyer, buyerHandler.GetBuyer()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminBuyer, buyerHandler.StoreBuyer()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(AdminBuyerId, buyerHandler.UpdateBuyer()).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc(AdminBuyerId, buyerHandler.DeleteBuyer()).Methods(http.MethodDelete, http.MethodOptions)

	buyerCategoryRepository := bCategoryPkg.NewBuyerCategoryRepository(db)
	buyerCategoryService := bCategory.NewBuyerCategoryService(buyerCategoryRepository)
	buyerCategoryHandler := bCategory.NewBuyerCategoryHandler(buyerCategoryService)

	r.HandleFunc(AdminBuyerCategory, buyerCategoryHandler.GetCategory()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminBuyerCategoryId, buyerCategoryHandler.GetCategoryById()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminBuyerCategory, buyerCategoryHandler.StoreCategory()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(AdminBuyerCategoryId, buyerCategoryHandler.UpdateCategory()).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc(AdminBuyerCategoryId, buyerCategoryHandler.DeleteCategory()).Methods(http.MethodDelete, http.MethodOptions)

	productRepository := productPkg.NewProductRepository(db)
	productService := product.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)

	r.HandleFunc(AdminProduct, productHandler.GetProduct()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminProduct, productHandler.StoreProduct()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.GetProductById()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.UpdateProduct()).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.DeleteProduct()).Methods(http.MethodDelete, http.MethodOptions)

	productCategoryRepository := pCategoryPkg.NewProductCategoryRepository(db)
	productCategoryService := pCategory.NewProductCategoryService(productCategoryRepository)
	productCategoryHandler := pCategory.NewProductCategoryHandler(productCategoryService)

	r.HandleFunc(AdminProductCategory, productCategoryHandler.GetCategory()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminProductCategory, productCategoryHandler.StoreCategory()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(AdminProductCategoryId, productCategoryHandler.GetCategoryById()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminProductCategoryId, productCategoryHandler.UpdateCategory()).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc(AdminProductCategoryId, productCategoryHandler.DeleteCategory()).Methods(http.MethodDelete, http.MethodOptions)
}
