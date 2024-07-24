package authapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc/tokensvc"
)

func (Impl) AuthAuth(
	ctx context.Context,
	request genapi.AuthAuthRequestObject,
) (genapi.AuthAuthResponseObject, error) {
	tokenString := request.Params.Authorization
	claim, err := tokensvc.Parse(ctx, tokenString)
	if err != nil {
		return nil, fmt.Errorf("parse token string error: %w", err)
	}

	usr, err := query.User.WithContext(ctx).
		Where(query.User.ID.Eq(claim.UserId)).
		First()
	if err != nil {
		return nil, fmt.Errorf("query user by id error: %w", err)
	}

	return genapi.AuthAuth200JSONResponse(genapi.AuthRes{
		Id:       int32(usr.ID),
		Username: usr.Username,
	}), nil
}
