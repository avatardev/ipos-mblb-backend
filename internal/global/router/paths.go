package router

const (
	AdminPing      = "/admin/v1/ping"
	AdminPingError = "/admin/v1/ping/error"

	AdminDashboardStatistics = "/admin/v1/dashboard/statistics"

	AdminGenerateDetailTrx = "/admin/v1/orders/report/generateDetail"
	AdminGenerateBriefTrx  = "/admin/v1/orders/report/generateBrief"
	AdminGenerateMonitor   = "/admin/v1/orders/report/generateComparison"
	AdminGenerateDaily     = "/admin/v1/orders/report/generateDaily/{sellerId}"
	AdminBriefTrx          = "/admin/v1/orders/report/brief"
	AdminDetailTrx         = "/admin/v1/orders/report/detail"
	AdminAddNote           = "/admin/v1/orders/report/note/{orderId}"
	AdminMonitor           = "/admin/v1/orders/report/comparison"
	AdminDaily             = "/admin/v1/orders/report/daily/{sellerId}"

	AdminProduct           = "/admin/v1/products"
	AdminProductId         = "/admin/v1/products/{productId}"
	AdminProductIdImg      = "/admin/v1/products/{productId}/img"
	AdminProductCategory   = "/admin/v1/products/categories"
	AdminProductCategoryId = "/admin/v1/products/categories/{categoryId}"

	AdminBuyer           = "/admin/v1/buyers"
	AdminBuyerId         = "/admin/v1/buyers/{buyerId}"
	AdminBuyerCategory   = "/admin/v1/buyers/categories"
	AdminBuyerCategoryId = "/admin/v1/buyers/categories/{categoryId}"
	AdminBuyerName       = "/admin/v1/buyers/companies"

	AdminSeller         = "/admin/v1/sellers"
	AdminSellerId       = "/admin/v1/sellers/{sellerId}"
	AdminMerchantItem   = "/admin/v1/sellers/{sellerId}/items"
	AdminMerchantItemId = "/admin/v1/sellers/{sellerId}/items/{itemId}"
	AdminSellerName     = "/admin/v1/sellers/companies"

	AdminUserAdmin     = "/admin/v1/user/admins"
	AdminUserAdminId   = "/admin/v1/user/admins/{userId}"
	AdminUserSeller    = "/admin/v1/user/sellers"
	AdminUserSellerId  = "/admin/v1/user/sellers/{userId}"
	AdminUserBuyer     = "/admin/v1/user/buyers"
	AdminUserBuyerId   = "/admin/v1/user/buyers/{userId}"
	AdminUserChecker   = "/admin/v1/user/checkers"
	AdminUserCheckerId = "/admin/v1/user/checkers/{userId}"

	AdminAuth        = "/admin/v1/auth"
	AdminAuthRefresh = "/admin/v1/auth/refreshToken"

	AdminLocation   = "/admin/v1/locations"
	AdminLocationId = "/admin/v1/locations/{locationId}"

	AdminLogInfo = "/admin/v1/logs"
)
