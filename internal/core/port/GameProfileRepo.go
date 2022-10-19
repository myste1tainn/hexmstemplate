package port

type GameProfileRepo interface {
	CreateGameProfile(request GameProfileCreateGameProfileRequest) (*GameProfileCreateGameProfileResponse, error)
}

type GameProfileGetGameProfileRequest struct {
}

type GameProfileGetGameProfileResponse struct {
}

type GameProfileCreateGameProfileRequest struct {
}

type GameProfileCreateGameProfileResponse struct {
}
