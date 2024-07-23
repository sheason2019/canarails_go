package appsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppsCreate(
	ctx context.Context,
	request genapi.AppsCreateRequestObject,
) (genapi.AppsCreateResponseObject, error) {
	usr, err := authsvc.GetUserByToken(ctx, request.Params.Authorization)
	if err != nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	app := &models.App{
		Title:       request.Body.Title,
		Description: request.Body.Description,
		Hostnames:   request.Body.Hostnames,
	}
	err = query.App.WithContext(ctx).Create(app)
	if err != nil {
		return nil, err
	}

	return genapi.AppsCreate200JSONResponse(app.ID), nil
}
