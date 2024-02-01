package dto

import (
	"configservice/models"
)

type PingResponse struct {
	Message string `json:"message"`
}

type GetPingResponse struct {
	Pings []models.PingPong `json:"pings"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type PingRequest struct {
	Text string `json:"text"`
}
