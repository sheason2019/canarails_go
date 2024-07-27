package appsapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppsDelete(
	ctx context.Context,
	request genapi.AppsDeleteRequestObject,
) (genapi.AppsDeleteResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	// 级联删除
	err := query.Q.Transaction(func(tx *query.Query) error {
		// app variants
		appVariants, err := tx.AppVariant.WithContext(ctx).
			Where(query.AppVariant.AppID.Eq(uint(request.Id))).
			Find()
		if err != nil {
			return fmt.Errorf("find app variants error: %w", err)
		}
		if len(appVariants) > 0 {
			_, err := tx.AppVariant.WithContext(ctx).Delete(appVariants...)
			if err != nil {
				return fmt.Errorf("delete app variants error: %w", err)
			}
		}

		// app
		_, err = tx.App.WithContext(ctx).
			Where(query.App.ID.Eq(uint(request.Id))).
			Delete()
		return err
	})

	return genapi.AppsDelete200JSONResponse(request.Id), err
}
