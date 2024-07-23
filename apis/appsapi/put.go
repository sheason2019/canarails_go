package appsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppsPut(
	ctx context.Context,
	request genapi.AppsPutRequestObject,
) (genapi.AppsPutResponseObject, error) {
	usr, err := authsvc.GetUserByToken(ctx, request.Params.Authorization)
	if err != nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	record, err := query.App.WithContext(ctx).
		Where(query.App.ID.Eq(uint(request.Body.Id))).
		First()
	if err != nil {
		return nil, err
	}

	record.Title = request.Body.Title
	record.Description = request.Body.Description
	record.Hostnames = request.Body.Hostnames

	err = query.App.WithContext(ctx).Save(record)
	return genapi.AppsPut200JSONResponse(request.Body.Id), err
}
