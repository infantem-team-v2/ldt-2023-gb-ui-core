package authInterface

import "gb-ui-core/internal/auth/model"

type UseCase interface {
	ValidateService(params *model.AuthHeadersLogic) (bool, error)
	GenerateAccessToken(refreshToken string, params *model.CreateAuthTokensLogic) (accessToken string, err error)
	GenerateTokensPair(params *model.CreateAuthTokensLogic) (accessToken, refreshToken string, err error)
	ValidateUser(userId int64) (ok bool, err error)
}
