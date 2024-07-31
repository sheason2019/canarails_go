package appvariantsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppVariantsBatchPut(
	ctx context.Context,
	request genapi.AppVariantsBatchPutRequestObject,
) (genapi.AppVariantsBatchPutResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	matches := make([]models.AppVariantMatch, len(request.Body.Matches))
	for i, v := range request.Body.Matches {
		matches[i] = models.AppVariantMatch{
			Header: v.Header,
			Value:  v.Value,
		}
	}

	record := models.AppVariant{
		Title:       request.Params.Title,
		Description: request.Body.Description,
		Matches:     matches,
		ImageName:   request.Body.ImageName,
		ExposePort:  uint(request.Body.ExposePort),
		Replicas:    uint(request.Body.Replicas),
	}

	info, err := query.AppVariant.WithContext(ctx).
		Where(query.AppVariant.Title.Eq(request.Params.Title)).
		Updates(record)
	if err != nil {
		return nil, err
	}

	return genapi.AppVariantsBatchPut200JSONResponse(info.RowsAffected), nil
}
