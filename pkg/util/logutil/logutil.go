package logutil

import (
	"context"
	"log"

	logEntity "github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log/entity"
	authEntity "github.com/avatardev/ipos-mblb-backend/internal/admin/auth/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

func GenerateActivityLog(ctx context.Context, msg string) {
	authinfo, ok := ctx.Value(authEntity.AuthLevelCtxKey("user-auth")).(*dto.AuthUserLevel)
	if !ok {
		log.Printf("[GenerateActivityLog] no user-auth found!")
		return
	}

	info, ok := ctx.Value(logEntity.LogCtxKey("log-info")).(*dto.LogInfo)
	if !ok {
		log.Printf("[GenereateActivityLog] no log-info found!")
		return
	}

	info.Message = msg
	if info.UserId == 0 {
		info.UserId = authinfo.Id
	}
}

func GenerateActivityLogNoAuth(ctx context.Context, userId int64, msg string) {
	info, ok := ctx.Value(logEntity.LogCtxKey("log-info")).(*dto.LogInfo)
	if !ok {
		log.Printf("[GenereateAuthActivityLog] no log-info found!")
		return
	}

	info.Message = msg
	info.UserId = userId
}
