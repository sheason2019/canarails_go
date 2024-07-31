package usertokensapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"github.com/labstack/echo/v4"
)

func (Impl) UserTokensList(
	ctx context.Context,
	request genapi.UserTokensListRequestObject,
) (genapi.UserTokensListResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	list, err := query.UserToken.WithContext(ctx).
		Where(query.UserToken.UserID.Eq(usr.ID)).
		Find()
	if err != nil {
		return nil, err
	}

	userTokens := make([]genapi.UserToken, len(list))
	for i, v := range list {
		userTokens[i] = genapi.UserToken{
			Id:          int32(v.ID),
			Title:       v.Title,
			Description: v.Description,
			ExpiredAt:   v.ExpiredAt.UnixMilli(),
			LastUsedAt:  v.LastUsedAt.UnixMilli(),
		}
	}

	return genapi.UserTokensList200JSONResponse(userTokens), nil
}
