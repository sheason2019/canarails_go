package authsvc

import (
	"context"

	"canarails.dev/database/models"
	"github.com/labstack/echo/v4"
)

type key int

const keyUserContext key = iota

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		ctx := c.Request().Context()

		usr, _ := getUserByToken(ctx, tokenString)

		usrCtx := context.WithValue(ctx, keyUserContext, usr)

		c.SetRequest(c.Request().WithContext(usrCtx))

		return next(c)
	}
}

func GetCurrentUser(ctx context.Context) *models.User {
	usr, ok := ctx.Value(keyUserContext).(*models.User)
	if ok {
		return usr
	}

	return nil
}
