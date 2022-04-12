package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log/entity"
)

type LogInfo struct {
	UserId  int64  `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func (l *LogInfo) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(l)
}

func (l *LogInfo) ToEntity() entity.LogInfo {
	return entity.LogInfo{
		UserId:  l.UserId,
		Message: l.Message,
	}
}
