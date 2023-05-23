package usecase

import (
	"gb-ui-core/config"
	authInterface "gb-ui-core/internal/auth/interface"
	"gb-ui-core/internal/auth/model"
	"gb-ui-core/pkg/terrors"
	"gb-ui-core/pkg/thttp/server"
	tlogger "gb-ui-core/pkg/tlogger"
	"gb-ui-core/pkg/tsecure"
	"gb-ui-core/pkg/tutils/ptr"
	"github.com/fiorix/go-redis/redis"
	"github.com/sarulabs/di"
	"time"
)

type AuthUC struct {
	config   *config.Config
	logger   tlogger.ILogger
	authRepo authInterface.RelationalRepository
	fernet   *tsecure.FernetCrypto
	redis    *redis.Client
}

func BuildAuthUsecase(ctn di.Container) (interface{}, error) {
	return &AuthUC{
		config:   ctn.Get("config").(*config.Config),
		logger:   ctn.Get("logger").(tlogger.ILogger),
		authRepo: ctn.Get("authRepo").(authInterface.RelationalRepository),
		fernet:   ctn.Get("fernet").(*tsecure.FernetCrypto),
		redis:    ctn.Get("redis").(*redis.Client),
	}, nil
}

func (as *AuthUC) ValidateService(params *model.AuthHeadersLogic) (bool, error) {
	service, err := as.authRepo.FindServiceByPublicKey(params.PublicKey)
	if err != nil {
		return false, err
	}
	decryptedPrivateKey, err := as.fernet.Decrypt(service.PrivateKey)
	if err != nil {
		return false, terrors.Raise(err, 200003)
	}
	signature := tsecure.CalcSignature(
		decryptedPrivateKey,
		string(params.Body),
		tsecure.SHA512,
	)
	if params.Signature == signature {
		return true, nil
	}

	return false, terrors.Raise(nil, 100003)
}

func (as *AuthUC) GenerateAccessToken(refreshToken string, params *model.CreateAuthTokensLogic) (accessToken string, err error) {
	duration := time.Now().Add(time.Second * time.Duration(as.config.HttpConfig.AccessExpireTime))
	accessToken, err = server.CreateJwtToken(&server.JwtParams{
		Salt:     &as.config.HttpConfig.JWTSalt,
		UserId:   &params.UserId,
		Duration: &duration,
		Type:     ptr.String("access"),
	})

	if err != nil {
		return "", terrors.Raise(err, 100012)
	}
	err = as.redis.Set(accessToken, refreshToken)
	if err != nil {
		return "", terrors.Raise(err, 300002)
	}
	ok, err := as.redis.ExpireAt(accessToken, int(duration.Unix()))
	if !ok || err != nil {
		return "", terrors.Raise(err, 300002)
	}
	return accessToken, nil
}

func (as *AuthUC) GenerateTokensPair(params *model.CreateAuthTokensLogic) (accessToken, refreshToken string, err error) {
	refreshDuration := time.Now().Add(time.Minute * time.Duration(as.config.HttpConfig.RefreshExpireTime))
	refreshToken, err = server.CreateJwtToken(&server.JwtParams{
		Salt:     &as.config.HttpConfig.JWTSalt,
		UserId:   &params.UserId,
		Duration: &refreshDuration,
		Type:     ptr.String("refresh"),
	})
	if err != nil {
		return "", "", terrors.Raise(err, 100012)
	}
	accessToken, err = as.GenerateAccessToken(refreshToken, params)
	if err != nil {
		return "", "", err
	}
	err = as.redis.Set(refreshToken, "")
	if err != nil {
		return "", "", terrors.Raise(err, 300002)
	}
	ok, err := as.redis.ExpireAt(refreshToken, int(refreshDuration.Unix()))
	if !ok || err != nil {
		return "", "", terrors.Raise(err, 300002)
	}
	return accessToken, refreshToken, nil
}

func (as *AuthUC) ValidateUser(userId int64) (ok bool, err error) {
	user, err := as.authRepo.FindUserByIdShort(userId)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, terrors.Raise(err, 100010)
	}

	return true, nil
}
