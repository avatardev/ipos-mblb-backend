package router

import (
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/category"
	categoryPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/product"
	productPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *database.DatabaseClient) {
	buyerService := buyer.NewBuyerService()
	buyerHandler := buyer.NewBuyerHandler(buyerService)

	r.HandleFunc(AdminPing, buyerHandler.PingBuyer()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminPingError, buyerHandler.PingError()).Methods(http.MethodGet, http.MethodOptions)

	categoryRepostory := categoryPkg.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepostory)
	categoryHandler := category.NewCategoryHandler(categoryService)

	r.HandleFunc(AdminCategory, categoryHandler.GetCategory()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminCategory, categoryHandler.StoreCategory()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(AdminCategoryId, categoryHandler.GetCategoryById()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminCategoryId, categoryHandler.UpdateCategory()).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc(AdminCategoryId, categoryHandler.DeleteCategory()).Methods(http.MethodDelete, http.MethodOptions)

	productRepository := productPkg.NewProductRepository(db)
	productService := product.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)

	r.HandleFunc(AdminProduct, productHandler.GetProduct()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminProduct, productHandler.StoreProduct()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.GetProductById()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.UpdateProduct()).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc(AdminProductId, productHandler.DeleteProduct()).Methods(http.MethodDelete, http.MethodOptions)
}
