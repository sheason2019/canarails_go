package authsvc

import (
	"context"
	"fmt"

	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc/tokensvc"
)

func GetUserByToken(
	ctx context.Context,
	tokenString string,
) (*models.User, error) {
	claim, err := tokensvc.Parse(ctx, tokenString)
	if err != nil {
		return nil, fmt.Errorf("parse token string error: %w", err)
	}

	return query.User.WithContext(ctx).
		Where(query.User.ID.Eq(claim.UserId)).
		First()
}
