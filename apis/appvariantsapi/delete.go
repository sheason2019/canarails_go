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

	err := query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.AppVariant.WithContext(ctx).
			Where(tx.AppVariant.ID.Eq(uint(request.Id))).
			Delete()
		if err != nil {
			return err
		}

		err = gatewaysvc.Reconciliation(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return genapi.AppVariantsDelete200JSONResponse(request.Id), nil
}
