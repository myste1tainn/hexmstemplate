package service

import (
	"github.com/gin-gonic/gin"
	"github.com/myste1tainn/hexmstemplate/internal/core/domain"
	"github.com/myste1tainn/hexmstemplate/internal/core/port"
	"github.com/myste1tainn/msfnd"
)

type createGameService struct {
	gameRepo port.GameProfileRepo
}

func NewCreateGameService(gameRepo port.GameProfileRepo) domain.CreateGameService {
	return &createGameService{
		gameRepo: gameRepo,
	}
}

func (s *createGameService) Execute(req domain.CreateGameRequest, ctx *gin.Context, rct *msfnd.RouteContext) (*domain.CreateGameResponse, error) {
	s.gameRepo.CreateGameProfile(port.GameProfileCreateGameProfileRequest{})
	return nil, nil
}
