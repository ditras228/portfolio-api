package keys

import (
	"context"
	"portfolio-api/graph/model"
)

var (
	UserCtxKey   = &contextKey{"user"}
	LocaleCtxKey = &contextKey{"locale"}
)

type contextKey struct {
	name string
}

func ForContext(ctx context.Context) model.User {
	raw := ctx.Value(UserCtxKey).(model.User)
	return raw
}
func LocaleForContext(ctx context.Context) int {
	raw := ctx.Value(LocaleCtxKey).(string)
	return GetLocaleId(raw)
}

func GetLocaleId(locale string) int {
	switch locale {
	case "en-US":
		return 1
	case "ru":
		return 2
	}
	return 1
}
