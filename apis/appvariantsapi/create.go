package appvariantsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"canarails.dev/services/gatewaysvc"
	"canarails.dev/services/recordsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppVariantsCreate(
	ctx context.Context,
	request genapi.AppVariantsCreateRequestObject,
) (genapi.AppVariantsCreateResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	appVar := recordsvc.MapAppVar(request.Body)
	err := query.Q.Transaction(func(tx *query.Query) error {
		err := tx.AppVariant.WithContext(ctx).Create(appVar)
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

	return genapi.AppVariantsCreate200JSONResponse(appVar.ID), nil
}
