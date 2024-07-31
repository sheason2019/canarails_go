package usertokensapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) UserTokensDelete(
	ctx context.Context,
	request genapi.UserTokensDeleteRequestObject,
) (genapi.UserTokensDeleteResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	info, err := query.UserToken.WithContext(ctx).
		Where(query.UserToken.ID.Eq(uint(request.Id))).
		Where(query.UserToken.UserID.Eq(usr.ID)).
		Delete()
	if err != nil {
		return nil, err
	}

	if info.RowsAffected == 0 {
		return nil, fmt.Errorf("no data")
	}

	return genapi.UserTokensDelete200JSONResponse(request.Id), nil
}
