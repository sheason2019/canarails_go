package appsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) AppsDelete(
	ctx context.Context,
	request genapi.AppsDeleteRequestObject,
) (genapi.AppsDeleteResponseObject, error) {
	usr, err := authsvc.GetUserByToken(ctx, request.Params.Authorization)
	if err != nil {
		return nil, echo.ErrUnauthorized
	}

	if usr.Username != "admin" {
		return nil, echo.ErrForbidden
	}

	// 级联删除
	info, err := query.AppDeploy.WithContext(ctx).
		LeftJoin(query.AppVariant, query.AppVariant.ID.EqCol(query.AppDeploy.AppVariantID)).
		LeftJoin(query.App, query.App.ID.EqCol(query.AppVariant.AppID)).
		Where(query.App.ID.Eq(uint(request.Body.Id))).
		Delete()

	return genapi.AppsDelete200JSONResponse(info.RowsAffected), err
}
