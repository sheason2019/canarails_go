package authapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
	"canarails.dev/services/authsvc/loginsvc"
	"canarails.dev/services/authsvc/tokensvc"
)

func (Impl) AuthLogin(
	ctx context.Context,
	request genapi.AuthLoginRequestObject,
) (genapi.AuthLoginResponseObject, error) {
	usr, err := query.User.
		WithContext(ctx).
		Where(query.User.Username.Eq(request.Body.Username)).
		First()
	if err != nil {
		return nil, fmt.Errorf("find user by username error: %w", err)
	}

	if !loginsvc.ComparePassword(request.Body.Password, usr) {
		return nil, fmt.Errorf("login error: invalid password")
	}

	token := tokensvc.New(ctx, usr)

	tokenString, err := token.ToString(ctx)
	if err != nil {
		return nil, fmt.Errorf("create token string error: %w", err)
	}

	return genapi.AuthLogin200JSONResponse(genapi.LoginRes{
		Token: tokenString,
	}), nil
}
