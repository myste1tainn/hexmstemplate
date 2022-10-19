package port

type GameRepo interface {
	Get(req GetGameRequest) (*GetGameResponse, error)
}

type GetGameRequest struct {
	Txt string
}

type GetGameResponse struct {
	Code string
}
