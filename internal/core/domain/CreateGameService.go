package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/myste1tainn/msfnd"
)

type CreateGameService interface {
	Execute(request CreateGameRequest, ctx *gin.Context, rctx *msfnd.RouteContext) (*CreateGameResponse, error)
}

type CreateGameRequest struct {
	Txt  string `json:"txt"`
	Type string `json:"type" `
}

type CreateGameResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
