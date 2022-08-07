package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"portfolio-api/container"
	"portfolio-api/middlewares/keys"
	"portfolio-api/pkg/utils"
)

const (
	authorizationHeader = "Authorization"
	localeHeader        = "Locale"
)

func MiddleWare() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			token := request.Header.Get(authorizationHeader)
			locale := request.Header.Get(localeHeader)

			ctx := context.WithValue(request.Context(), keys.LocaleCtxKey, locale)

			if &token == nil {
				return
			} else {
				if token != "" {
					userId, err := utils.ParseToken(token)
					res, err := container.UserRepository.GetOne(ctx, userId)
					if err != nil {
						fmt.Errorf(err.Error())
					}

					ctx = context.WithValue(ctx, keys.UserCtxKey, res)
				}
			}

			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		})
	}
}
