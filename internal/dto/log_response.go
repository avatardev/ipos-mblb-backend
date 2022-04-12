package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log/entity"

type LogResponse struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
	Msg       string `json:"message"`
}

type LogResponseJSON struct {
	Logs []*LogResponse `json:"logs"`
}

func NewLogResponse(log *entity.LogInfo) *LogResponse {
	var username string
	if log.Username == nil {
		username = "deleted_user"
	} else {
		username = *log.Username
	}

	return &LogResponse{
		Timestamp: log.Timestamp.UTC().Format("2006-01-02 15:04"),
		Id:        log.Id,
		UserId:    log.UserId,
		Username:  username,
		Msg:       log.Message,
	}
}

func NewLogResponseJSON(logs entity.Logs) *LogResponseJSON {
	res := &LogResponseJSON{}

	for _, log := range logs {
		l := NewLogResponse(log)
		res.Logs = append(res.Logs, l)
	}

	return res
}
