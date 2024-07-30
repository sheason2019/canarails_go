package appvariantsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"canarails.dev/services/gatewaysvc"
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

	body := request.Body

	appVar := &models.AppVariant{
		Title:       body.Title,
		Description: body.Description,
		ExposePort:  uint(body.ExposePort),
		ImageName:   body.ImageName,
		Replicas:    uint(body.Replicas),
		AppID:       uint(body.AppId),
	}

	err := query.Q.Transaction(func(tx *query.Query) error {
		err := query.AppVariant.WithContext(ctx).Create(appVar)
		if err != nil {
			return err
		}

		err = gatewaysvc.Reconciliation(ctx)
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
