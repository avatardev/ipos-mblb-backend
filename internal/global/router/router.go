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
	"github.com/avatardev/ipos-mblb-backend/internal/admin/dashboard"
	dashboardPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/dashboard/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/location"
	locationPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/location/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/merchant"
	merchantPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/merchant/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/order"
	orderPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/order/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/product"
	productPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product/impl"
	pCategory "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category"
	pCategoryPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller"
	sellerPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/seller/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/user"
	userPkg "github.com/avatardev/ipos-mblb-backend/internal/admin/user/impl"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"

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

	dashboardRepository := dashboardPkg.NewDashboardRepository(db)
	dashboardService := dashboard.NewDashboardService(dashboardRepository)
	dashboardHandler := dashboard.NewDashboardHandler(dashboardService)

	protectedRouter.HandleFunc(AdminDashboardStatistics, dashboardHandler.GetStatistics()).Methods(http.MethodGet, http.MethodOptions)

	orderRepository := orderPkg.NewOrderRepository(db)
	orderService := order.NewOrderService(orderRepository)
	orderHandler := order.NewOrderHandler(orderService)

	protectedRouter.HandleFunc(AdminGenerateDetailTrx, orderHandler.GenerateDetailTrx()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminGenerateBriefTrx, orderHandler.GenerateBriefTrx()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminGenerateMonitor, orderHandler.GenerateMonitorTrx()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminGenerateDaily, orderHandler.GenerateDailyTrx()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminBriefTrx, orderHandler.BriefTrx()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminDetailTrx, orderHandler.DetailTrx()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminMonitor, orderHandler.MonitorTrx()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminDaily, orderHandler.DailyTrx()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminAddNote, orderHandler.InsertNote()).Methods(http.MethodPost, http.MethodOptions)

	locationRepository := locationPkg.NewLocationRepostiory(db)
	locationService := location.NewLocationService(locationRepository)
	locationHandler := location.NewLocationHandler(locationService)

	protectedRouter.HandleFunc(AdminLocation, locationHandler.GetLocation()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminLocation, locationHandler.StoreLocation()).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminLocationId, locationHandler.GetLocationById()).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminLocationId, locationHandler.UpdateLocation()).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminLocationId, locationHandler.DeleteLocation()).Methods(http.MethodDelete, http.MethodOptions)

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

	protectedRouter.HandleFunc(AdminBuyerName, buyerHandler.GetBuyerName()).Methods(http.MethodGet, http.MethodOptions)
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

	protectedRouter.HandleFunc(AdminSellerName, sellerHandler.GetSellerName()).Methods(http.MethodGet, http.MethodOptions)
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

	protectedRouter.HandleFunc(AdminUserAdmin, userHandler.GetUser(privutil.USER_ADMIN)).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserAdmin, userHandler.StoreUser(privutil.USER_ADMIN)).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserAdminId, userHandler.GetUserById(privutil.USER_ADMIN)).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserAdminId, userHandler.UpdateUser(privutil.USER_ADMIN)).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserAdminId, userHandler.DeleteUser(privutil.USER_ADMIN)).Methods(http.MethodDelete, http.MethodOptions)

	protectedRouter.HandleFunc(AdminUserBuyer, userHandler.GetUserBuyer(privutil.USER_BUYER)).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserBuyer, userHandler.StoreUser(privutil.USER_BUYER)).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserBuyerId, userHandler.GetUserById(privutil.USER_BUYER)).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserBuyerId, userHandler.UpdateUser(privutil.USER_BUYER)).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserBuyerId, userHandler.DeleteUser(privutil.USER_BUYER)).Methods(http.MethodDelete, http.MethodOptions)

	protectedRouter.HandleFunc(AdminUserSeller, userHandler.GetUserSeller(privutil.USER_SELLER)).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserSeller, userHandler.StoreUser(privutil.USER_SELLER)).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserSellerId, userHandler.GetUserById(privutil.USER_SELLER)).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserSellerId, userHandler.UpdateUser(privutil.USER_SELLER)).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserSellerId, userHandler.DeleteUser(privutil.USER_SELLER)).Methods(http.MethodDelete, http.MethodOptions)

	protectedRouter.HandleFunc(AdminUserChecker, userHandler.GetUser(privutil.USER_CHECKER)).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserChecker, userHandler.StoreUser(privutil.USER_CHECKER)).Methods(http.MethodPost, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserCheckerId, userHandler.GetUserById(privutil.USER_CHECKER)).Methods(http.MethodGet, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserCheckerId, userHandler.UpdateUser(privutil.USER_CHECKER)).Methods(http.MethodPut, http.MethodOptions)
	protectedRouter.HandleFunc(AdminUserCheckerId, userHandler.DeleteUser(privutil.USER_CHECKER)).Methods(http.MethodDelete, http.MethodOptions)
}
