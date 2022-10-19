package httpserv

import (
	"github.com/gin-gonic/gin"
	"github.com/myste1tainn/hexmstemplate/internal/adaptor/handler"
	"github.com/myste1tainn/hexmstemplate/internal/adaptor/repo"
	"github.com/myste1tainn/hexmstemplate/internal/core/service"
	"github.com/myste1tainn/msfnd"
	"github.com/myste1tainn/msnet"
)

func bindCreateGameRoute(a *gin.Engine) {
	gameRepo := repo.NewGameProfileRepo(msnet.NewClient(msnet.ClientConfig{}), repo.NewGameProfileRepoConfig())
	svc := service.NewCreateGameService(gameRepo)
	hdl := handler.NewCreateGameHandler(svc)

	a.GET(
		"/v1/game",
		func(ctx *gin.Context) {
			hdl.Handle(ctx, &msfnd.RouteContext{})
		},
	)
}

func bindHealthRoute(a *gin.Engine) {
	a.GET(
		"/health",
		func(ctx *gin.Context) {
			ctx.JSON(200, map[string]string{"status": "healthy"})
		},
	)
}
