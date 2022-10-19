package repo

import (
	"errors"
	"strings"

	"github.com/myste1tainn/hexmstemplate/internal/core/port"
	"github.com/myste1tainn/msnet"
)

type GameProfileRepo struct {
	client msnet.Client
	config *msnet.Config
}

func NewGameProfileRepo(client msnet.Client, config *msnet.Config) port.GameProfileRepo {
	return GameProfileRepo{
		client: client,
		config: config,
	}
}

func (r GameProfileRepo) CreateGameProfile(request port.GameProfileCreateGameProfileRequest) (*port.GameProfileCreateGameProfileResponse, error) {
	//var l = ctx.GetLogger().NewSpanLogger()
	//defer l.Destroy()

	var configKey = strings.ToLower("CreateGameProfile")
	var result port.GameProfileCreateGameProfileResponse
	var error msnet.ErrorResponse
	res, err := r.client.
		//WithLogger(l).
		//RequestWithContext(rctx, r.config, configKey).
		Request(r.config, configKey).
		Call(&result, &error)
	if err != nil {
		return nil, err
	} else if res.IsError() {
		return nil, errors.New("there should be something here")
	} else {
		return &result, nil
	}
}
