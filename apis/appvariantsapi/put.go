package appvariantsapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"canarails.dev/services/gatewaysvc"
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

	err := query.Q.Transaction(func(tx *query.Query) error {
		record, err := tx.AppVariant.WithContext(ctx).
			Where(tx.AppVariant.ID.Eq(uint(request.Id))).
			First()
		if err != nil {
			return fmt.Errorf("find app variant by id error: %w", err)
		}

		record.Title = body.Title
		record.Description = body.Description
		record.ExposePort = uint(body.ExposePort)
		record.ImageName = body.ImageName
		record.Replicas = uint(body.Replicas)
		matches := make([]models.AppVariantMatch, len(body.Matches))
		for i, v := range body.Matches {
			matches[i] = models.AppVariantMatch{
				Header: v.Header,
				Value:  v.Value,
			}
		}
		record.Matches = matches

		err = tx.AppVariant.WithContext(ctx).Save(record)
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

	return genapi.AppVariantsPut200JSONResponse(request.Id), nil
}
