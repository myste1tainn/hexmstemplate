package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myste1tainn/hexmstemplate/internal/core/domain"
	"github.com/myste1tainn/msfnd"
)

type CreateGameHandler struct {
	svc domain.CreateGameService
}

func NewCreateGameHandler(service domain.CreateGameService) *CreateGameHandler {
	return &CreateGameHandler{
		svc: service,
	}
}

func (hdl *CreateGameHandler) Handle(ctx *gin.Context, rctx *msfnd.RouteContext) {
	userRefID := rctx.UserRefId
	fmt.Println("userRefID:", userRefID)

	var req = domain.CreateGameRequest{}

	res, err := hdl.svc.Execute(req, ctx, rctx)
	if err != nil {
		//TODO
		ctx.Error(err)
		return
	}

	//call service
	ctx.JSON(http.StatusOK, &res)

}
