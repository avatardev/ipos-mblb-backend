package dto

type PingResponse struct {
	Message         string `json:"message"`
	ServerTimestamp string `json:"serverTimestamp"`
}
