package usertokensapi

import (
	"context"
	"math"
	"time"

	"canarails.dev/apis/genapi"
	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"canarails.dev/services/authsvc/tokensvc"
	"github.com/labstack/echo/v4"
)

func (Impl) UserTokensCreate(
	ctx context.Context,
	request genapi.UserTokensCreateRequestObject,
) (genapi.UserTokensCreateResponseObject, error) {
	usr := authsvc.GetCurrentUser(ctx)
	if usr == nil {
		return nil, echo.ErrUnauthorized
	}

	exp := time.UnixMilli(request.Body.ExpiredAt)
	record := models.UserToken{
		Title:       request.Body.Title,
		Description: request.Body.Description,
		LastUsedAt:  time.Now(),
		ExpiredAt:   exp,
		UserID:      usr.ID,
	}
	err := query.UserToken.WithContext(ctx).Create(&record)
	if err != nil {
		return nil, err
	}

	forever := time.UnixMilli(math.MaxInt64)
	claims := tokensvc.
		New(ctx, usr).
		WithExpiredAt(forever).
		WithUserTokenId(record.ID)
	tokenString, err := claims.ToString(ctx)
	if err != nil {
		return nil, err
	}

	return genapi.UserTokensCreate200JSONResponse(genapi.CreateUserTokenRes{
		TokenString: tokenString,
		Id:          int32(record.ID),
		Title:       record.Title,
		Description: record.Description,
		LastUsedAt:  record.LastUsedAt.UnixMilli(),
		ExpiredAt:   record.ExpiredAt.UnixMilli(),
	}), nil
}
