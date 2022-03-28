package privutil

import (
	"context"
	"log"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

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
