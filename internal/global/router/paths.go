package router

const (
	AdminPing              = "/admin/v1/ping"
	AdminPingError         = "/admin/v1/ping/error"
	AdminProduct           = "/admin/v1/products"
	AdminProductId         = "/admin/v1/products/{productId}"
	AdminProductCategory   = "/admin/v1/products/categories"
	AdminProductCategoryId = "/admin/v1/products/categories/{categoryId}"
	AdminBuyer             = "/admin/v1/buyers"
	AdminBuyerId           = "/admin/v1/buyers/{buyerId}"
	AdminBuyerCategory     = "/admin/v1/buyers/categories"
	AdminBuyerCategoryId   = "/admin/v1/buyers/categories/{categoryId}"
	AdminSeller            = "/admin/v1/sellers"
	AdminSellerId          = "/admin/v1/sellers/{sellerId}"
	AdminUserAdmin         = "/admin/v1/users/admin"
	AdminUserAdminId       = "/admin/v1/users/admin/{userId}"
	AdminUserChecker       = "/admin/v1/users/checker"
	AdminUserCheckerId     = "/admin/v1/users/checker/{userId}"
)
