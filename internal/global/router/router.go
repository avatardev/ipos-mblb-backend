package router

import (
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer"
	buyerPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/product"
	productPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"
	category "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category"
	productCategoryPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category/impl"
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

	productCategoryRepository := productCategoryPkg.NewProductCategoryRepository(db)
	productCategoryService := category.NewProductCategoryService(productCategoryRepository)
	productCategoryHandler := category.NewProductCategoryHandler(productCategoryService)

	r.HandleFunc(AdminCategory, productCategoryHandler.GetCategory()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminCategory, productCategoryHandler.StoreCategory()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(AdminCategoryId, productCategoryHandler.GetCategoryById()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminCategoryId, productCategoryHandler.UpdateCategory()).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc(AdminCategoryId, productCategoryHandler.DeleteCategory()).Methods(http.MethodDelete, http.MethodOptions)

	productRepository := productPkg.NewProductRepository(db)
	productService := product.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)

	r.HandleFunc(AdminProduct, productHandler.GetProduct()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminProduct, productHandler.StoreProduct()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.GetProductById()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.UpdateProduct()).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.DeleteProduct()).Methods(http.MethodDelete, http.MethodOptions)
}
