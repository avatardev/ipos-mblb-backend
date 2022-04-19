package privutil

import (
	"context"
	"log"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

const (
	USER_ADMIN   int64 = 1
	USER_SELLER  int64 = 2
	USER_BUYER   int64 = 3
	USER_CHECKER int64 = 4
)

// CheckUserPrivilege is a function that recieve a context injected with necessary user data, and required priivlege level
// and return a boolean value whether current user has sufficient privilege level
func CheckUserPrivilege(ctx context.Context, privLevel ...int64) bool {
	authLevel, ok := ctx.Value(entity.AuthLevelCtxKey("user-auth")).(*dto.AuthUserLevel)
	if !ok {
		log.Printf("[CheckUserPrivilege] invalid auth-level (needed %+v, got None)\n", privLevel)
		return false
	}

	for _, priv := range privLevel {
		if authLevel.Role == priv {
			return true
		}
	}

	log.Printf("[CheckUserPrivilege] invalid auth-level (needed %+v, got %v)\n", privLevel, authLevel.Role)
	return false
}

func CheckSellerID(ctx context.Context, target int64) (same bool) {
	authLevel, ok := ctx.Value(entity.AuthLevelCtxKey("user-auth")).(*dto.AuthUserLevel)
	if !ok || authLevel.SellerID == nil {
		// No need to verify admin's serller id
		if ok && authLevel.Role == USER_ADMIN {
			return true
		}

		log.Printf("[CheckSellerID] invalid sellerID (needed %+v, got None)\n", target)
		return false
	}

	if *authLevel.SellerID != target {
		log.Printf("[CheckUserPrivilege] invalid sellerID (needed %+v, got %v)\n", target, *authLevel.SellerID)
		return false
	}

	return true
}
