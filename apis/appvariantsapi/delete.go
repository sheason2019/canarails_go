package appvariantsapi

import (
	"context"
	"fmt"

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

	err := query.Q.Transaction(func(tx *query.Query) error {
		_, err := query.AppDeploy.WithContext(ctx).
			Where(query.AppDeploy.AppVariantID.Eq(uint(request.Id))).
			Delete()
		if err != nil {
			return fmt.Errorf("delete app deploy error: %w", err)
		}

		_, err = query.AppVariant.WithContext(ctx).
			Where(query.AppVariant.ID.Eq(uint(request.Id))).
			Delete()
		return err
	})

	return genapi.AppVariantsDelete200JSONResponse(request.Id), err
}
