package appvariantsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"canarails.dev/services/gatewaysvc"
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
	if err != nil {
		return nil, err
	}

	err = gatewaysvc.Reconciliation(ctx)
	if err != nil {
		return nil, err
	}

	return genapi.AppVariantsDelete200JSONResponse(request.Id), nil
}
