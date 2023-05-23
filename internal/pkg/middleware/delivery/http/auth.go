package http

import (
	"gb-auth-gate/internal/auth/model"
	mdwModel "gb-auth-gate/internal/pkg/middleware/model"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
	"time"
)

// SignatureMiddleware Validates request by HMAC512
func (mdw *MiddlewareManager) SignatureMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		body := c.Body()
		ok, err := mdw.AuthUC.ValidateService(&model.AuthHeadersLogic{
			Signature: headers[mdwModel.SignatureKey],
			PublicKey: headers[mdwModel.PublicKey],
			Body:      body,
		})
		if !ok {
			if err == nil {
				return terrors.Raise(nil, 300001)
			}
			return err
		}
		return c.Next()
	}
}

// JWTMiddleware Validates request by access token
func (mdw *MiddlewareManager) JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		accessToken := c.Cookies(mdwModel.AccessKey)
		if accessToken == "" {
			refreshToken := c.Cookies(mdwModel.RefreshKey)
			if refreshToken == "" {
				return terrors.Raise(nil, 100004)
			}
			claims, err := server.ParseJwtToken(refreshToken, mdw.Config.HttpConfig.JWTSalt)

			accessToken, err = mdw.AuthUC.GenerateAccessToken(refreshToken, &model.CreateAuthTokensLogic{
				UserId: claims["userId"].(int64),
			})
			if err != nil {
				return err
			}

			c.Cookie(&fiber.Cookie{
				Name:     mdwModel.AccessKey,
				Value:    accessToken,
				Domain:   mdw.Config.BaseConfig.Service.URL,
				Expires:  time.Now().Add(time.Duration(mdw.Config.HttpConfig.AccessExpireTime) * time.Minute),
				Secure:   true,
				HTTPOnly: true,
			})
		}
		claims, err := server.ParseJwtToken(accessToken, mdw.Config.HttpConfig.JWTSalt)
		if err != nil {
			return err
		}
		userId := claims["userId"].(int64)
		ok, err := mdw.AuthUC.ValidateUser(userId)
		if err != nil {
			return err
		}
		if !ok {
			return terrors.Raise(nil, 100011)
		}

		return c.Next()
	}
}

// FiberSessionMiddleware Validates request by fiber's session are saved in redis' storage
func (mdw *MiddlewareManager) FiberSessionMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Next()
	}
}
