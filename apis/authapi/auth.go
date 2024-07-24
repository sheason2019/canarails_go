package authapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AuthAuth(
	ctx context.Context,
	request genapi.AuthAuthRequestObject,
) (genapi.AuthAuthResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	return genapi.AuthAuth200JSONResponse(genapi.AuthRes{
		Id:       int32(usr.ID),
		Username: usr.Username,
	}), nil
}
