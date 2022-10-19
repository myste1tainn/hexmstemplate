package repo

import (
	"github.com/myste1tainn/hexmstemplate/internal/core/port"
)

type gameRepo struct {
}

func NewGameRepo() port.GameRepo {
	return &gameRepo{}
}

func (repo *gameRepo) Get(req port.GetGameRequest) (*port.GetGameResponse, error) {
	//TODO implement me
	return &port.GetGameResponse{}, nil

}
