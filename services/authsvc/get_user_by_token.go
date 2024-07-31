package authsvc

import (
	"context"
	"fmt"
	"time"

	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc/tokensvc"
)

func getUserByToken(
	ctx context.Context,
	tokenString string,
) (*models.User, error) {
	claim, err := tokensvc.Parse(ctx, tokenString)
	if err != nil {
		return nil, fmt.Errorf("parse token string error: %w", err)
	}

	usr, err := query.User.WithContext(ctx).
		Where(query.User.ID.Eq(claim.UserId)).
		First()
	if err != nil {
		return nil, err
	}

	if claim.UserTokenId == 0 {
		return usr, nil
	}

	userToken, err := query.UserToken.WithContext(ctx).
		Where(query.UserToken.ID.Eq(claim.UserTokenId)).
		First()
	if err != nil {
		return nil, err
	}
	if userToken.ExpiredAt.Before(time.Now()) {
		return nil, fmt.Errorf("token already expired")
	}

	return usr, err
}
