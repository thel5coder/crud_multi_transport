package middleware

import (
	jwt2 "crud_multi_transport/helpers/jwt"
	"crud_multi_transport/helpers/messages"
	"crud_multi_transport/http_server/handlers"
	"crud_multi_transport/usecase"
	"crud_multi_transport/usecase/viewmodel"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"strings"
	"time"
)

type JwtVerify struct {
	*usecase.UcContract
}



func (jwtVerify JwtVerify) JWTWithConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		claims := &jwt2.CustomClaims{}
		apiHandler := handlers.Handler{UseCaseContract: jwtVerify.UcContract}

		tokenAuthHeader := ctx.Request().Header.Get("Authorization")
		if !strings.Contains(tokenAuthHeader, "Bearer") {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.AuthHeaderNotPresent))
		}

		tokenAuth := strings.Replace(tokenAuthHeader, "Bearer ", "", -1)
		_, err = jwt.ParseWithClaims(tokenAuth, claims, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			secret := jwtVerify.JwtConfig.SigningKey
			return secret, nil
		})
		if err != nil {
			return apiHandler.SendResponseUnauthorized(ctx, err)
		}

		if claims.ExpiresAt < time.Now().Unix() {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.ExpiredToken))
		}

		IDDecrypted, err := jwtVerify.Jwe.Rollback(claims.Id)
		if err != nil {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.FailedLoadPayload))
		}
		if IDDecrypted == "" {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.FailedLoadPayload))
		}
		claims.Id = fmt.Sprintf("%v",IDDecrypted)

		emailDecrypted, err := jwtVerify.Jwe.Rollback(claims.Email)
		if err != nil {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.FailedLoadPayload))
		}
		if emailDecrypted == "" {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.FailedLoadPayload))
		}
		claims.Email = fmt.Sprintf("%v",emailDecrypted)

		sessionData := viewmodel.UserSessionVm{}
		jwtVerify.UcContract.GetFromRedis("session-"+claims.Id, &sessionData)
		if sessionData.Session != claims.Session {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.InvalidSession))
		}
		ctx.Set("user",claims)

		return next(ctx)
	}
}

