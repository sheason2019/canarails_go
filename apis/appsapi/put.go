package appsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"canarails.dev/services/gatewaysvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppsPut(
	ctx context.Context,
	request genapi.AppsPutRequestObject,
) (genapi.AppsPutResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	record, err := query.App.WithContext(ctx).
		Where(query.App.ID.Eq(uint(request.Id))).
		First()
	if err != nil {
		return nil, err
	}

	record.Title = request.Body.Title
	record.Description = request.Body.Description
	record.Hostnames = request.Body.Hostnames
	record.DefaultVariantID = uint(request.Body.DefaultVariantId)

	err = query.App.WithContext(ctx).Save(record)
	if err != nil {
		return nil, err
	}

	err = gatewaysvc.Reconciliation(ctx)
	if err != nil {
		return nil, err
	}

	return genapi.AppsPut200JSONResponse(request.Id), nil
}
