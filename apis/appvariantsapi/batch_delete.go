package appvariantsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppVariantsBatchDelete(
	ctx context.Context,
	request genapi.AppVariantsBatchDeleteRequestObject,
) (genapi.AppVariantsBatchDeleteResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	info, err := query.AppVariant.WithContext(ctx).
		Where(query.AppVariant.Title.Eq(request.Params.Title)).
		Delete()
	if err != nil {
		return nil, err
	}

	return genapi.AppVariantsBatchDelete200JSONResponse(info.RowsAffected), nil
}
