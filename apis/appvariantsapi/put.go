package appvariantsapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppVariantsPut(
	ctx context.Context,
	request genapi.AppVariantsPutRequestObject,
) (genapi.AppVariantsPutResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	body := request.Body

	record, err := query.AppVariant.WithContext(ctx).
		Where(query.AppVariant.ID.Eq(uint(body.Id))).
		First()
	if err != nil {
		return nil, fmt.Errorf("find app variant by id error: %w", err)
	}

	record.Title = body.Title
	record.Description = body.Description
	record.ExposePort = int(body.ExposePort)
	matches := make([]models.AppVariantMatch, len(record.Matches))
	for i, v := range body.Matches {
		matches[i] = models.AppVariantMatch{
			Header: v.Header,
			Value:  v.Value,
		}
	}

	err = query.AppVariant.WithContext(ctx).Save(record)

	return genapi.AppVariantsPut200JSONResponse(body.Id), err
}
