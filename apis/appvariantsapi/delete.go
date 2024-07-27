package appvariantsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppVariantsDelete(
	ctx context.Context,
	request genapi.AppVariantsDeleteRequestObject,
) (genapi.AppVariantsDeleteResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	_, err := query.AppVariant.WithContext(ctx).
		Where(query.AppVariant.ID.Eq(uint(request.Id))).
		Delete()

	return genapi.AppVariantsDelete200JSONResponse(request.Id), err
}
