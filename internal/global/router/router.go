package router

import (
	"net/http"

	authSys "github.com/avatardev/ipos-mblb-backend/internal/admin/auth"
	authSysPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/auth/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth/middleware"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer"
	buyerPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/impl"
	bCategory "github.com/avatardev/ipos-mblb-backend/internal/admin/buyer_category"
	bCategoryPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/buyer_category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/merchant"
	merchantPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/merchant/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/product"
	productPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"
	pCategory "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category"
	pCategoryPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller"
	sellerPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/seller/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/user"
	userPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/user/impl"

	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *database.DatabaseClient) {
	authRepository := authSysPkg.NewAuthRepository(db)
	authService := authSys.NewAuthService(authRepository)
	authHandler := authSys.NewAuthHandler(authService)

	authRouter := r.NewRoute().Subrouter()
	protectedRouter := r.NewRoute().Subrouter()
	protectedRouter.Use(middleware.AuthMiddleware(authService))

	authRouter.HandleFunc(AdminAuth, authHandler.Login()).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc(AdminAuthRefresh, authHandler.RefreshToken()).Methods(http.MethodPost, http.MethodOptions)

	buyerCategoryRepository := bCategoryPkg.NewBuyerCategoryRepository(db)
	buyerCategoryService := bCategory.NewBuyerCategoryService(buyerCategoryRepository)
	buyerCategoryHandler := bCategory.NewBuyerCategoryHandler(buyerCategoryService)

	protectedRouter.HandleFunc(AdminBuyerCategory, buyerCategoryHandler.GetCategory()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBuyerCategoryId, buyerCategoryHandler.GetCategoryById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBuyerCategory, buyerCategoryHandler.StoreCategory()).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBuyerCategoryId, buyerCategoryHandler.UpdateCategory()).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBuyerCategoryId, buyerCategoryHandler.DeleteCategory()).Methods(http.MethodDelete, http.MethodOptions)

	buyerRepository := buyerPkg.NewBuyerRepository(db)
	buyerService := buyer.NewBuyerService(buyerRepository)
	buyerHandler := buyer.NewBuyerHandler(buyerService)

	protectedRouter.HandleFunc(AdminPing, buyerHandler.PingBuyer()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminPingError, buyerHandler.PingError()).Methods(http.MethodGet, http.MethodOptions)

	protectedRouter.HandleFunc(AdminBuyer, buyerHandler.GetBuyer()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBuyer, buyerHandler.StoreBuyer()).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBuyerId, buyerHandler.GetBuyerById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBuyerId, buyerHandler.UpdateBuyer()).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBuyerId, buyerHandler.DeleteBuyer()).Methods(http.MethodDelete, http.MethodOptions)

	merchanRepository := merchantPkg.NewMerchantRepository(db)
	merchantService := merchant.NewMerchantService(merchanRepository)
	merchantHandler := merchant.NewMerchantHandler(merchantService)

	protectedRouter.HandleFunc(AdminMerchantItem, merchantHandler.GetMerchant()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminMerchantItemId, merchantHandler.GetMerchantById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminMerchantItemId, merchantHandler.UpdateMerchant()).Methods(http.MethodPut, http.MethodOptions)

	sellerRepository := sellerPkg.NewSellerRepository(db)
	sellerService := seller.NewSellerService(sellerRepository)
	sellerHandler := seller.NewSellerHandler(sellerService)

	protectedRouter.HandleFunc(AdminSeller, sellerHandler.GetSeller()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminSellerId, sellerHandler.GetSellerById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminSeller, sellerHandler.StoreSeller()).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminSellerId, sellerHandler.UpdateSeller()).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminSellerId, sellerHandler.DeleteSeller()).Methods(http.MethodDelete, http.MethodOptions)

	productCategoryRepository := pCategoryPkg.NewProductCategoryRepository(db)
	productCategoryService := pCategory.NewProductCategoryService(productCategoryRepository)
	productCategoryHandler := pCategory.NewProductCategoryHandler(productCategoryService)

	protectedRouter.HandleFunc(AdminProductCategory, productCategoryHandler.GetCategory()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminProductCategory, productCategoryHandler.StoreCategory()).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminProductCategoryId, productCategoryHandler.GetCategoryById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminProductCategoryId, productCategoryHandler.UpdateCategory()).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminProductCategoryId, productCategoryHandler.DeleteCategory()).Methods(http.MethodDelete, http.MethodOptions)

	productRepository := productPkg.NewProductRepository(db)
	productService := product.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)

	protectedRouter.HandleFunc(AdminProduct, productHandler.GetProduct()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminProduct, productHandler.StoreProduct()).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminProductId, productHandler.GetProductById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminProductId, productHandler.UpdateProduct()).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminProductId, productHandler.DeleteProduct()).Methods(http.MethodDelete, http.MethodOptions)

	userRepository := userPkg.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	protectedRouter.HandleFunc(AdminUserAdmin, userHandler.GetUserAdmin()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserAdmin, userHandler.StoreUserAdmin()).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserAdminId, userHandler.GetUserAdminById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserAdminId, userHandler.UpdateUserAdmin()).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserAdminId, userHandler.DeleteUserAdmin()).Methods(http.MethodDelete, http.MethodOptions)

	protectedRouter.HandleFunc(AdminUserChecker, userHandler.GetUserChecker()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserChecker, userHandler.StoreUserChecker()).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserCheckerId, userHandler.GetUserCheckerById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserCheckerId, userHandler.UpdateUserChecker()).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserCheckerId, userHandler.DeleteUserChecker()).Methods(http.MethodDelete, http.MethodOptions)
}
