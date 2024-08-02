package appvariantsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"canarails.dev/services/gatewaysvc"
	"canarails.dev/services/recordsvc"
	"github.com/labstack/echo/v4"
	"gorm.io/gen/field"
)

func (Impl) AppVariantsBatchPatch(
	ctx context.Context,
	request genapi.AppVariantsBatchPatchRequestObject,
) (genapi.AppVariantsBatchPatchResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	queryRecord := &models.AppVariant{
		AppID: uint(request.Params.AppId),
	}
	if request.Params.Title != nil {
		queryRecord.Title = *request.Params.Title
	}
	if request.Params.Id != nil {
		queryRecord.ID = uint(*request.Params.Id)
	}

	targetRecord := recordsvc.MapAppVar(request.Body)

	var rowsAffected int64

	err := query.Q.Transaction(func(tx *query.Query) error {
		info, err := tx.AppVariant.WithContext(ctx).
			Where(field.Attrs(queryRecord)).
			Updates(targetRecord)
		if err != nil {
			return err
		}

		err = gatewaysvc.Reconciliation(ctx, tx)
		if err != nil {
			return err
		}

		rowsAffected = info.RowsAffected
		return nil
	})

	if err != nil {
		return nil, err
	}

	return genapi.AppVariantsBatchPatch200JSONResponse(rowsAffected), nil
}
