package tokensvc

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"canarails.dev/database/models"
	"canarails.dev/query"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserClaims struct {
	jwt.RegisteredClaims

	UserId uint
}

func New(ctx context.Context, usr *models.User) *UserClaims {
	return &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		UserId: usr.ID,
	}
}

func (claim *UserClaims) ToString(ctx context.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret, err := getSecret(ctx)
	if err != nil {
		return "", err
	}

	return token.SignedString(secret)
}

func Parse(ctx context.Context, tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return getSecret(ctx)
		},
	)
	if err != nil {
		return nil, fmt.Errorf("parse with claims error: %w", err)
	}

	if userClaims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return userClaims, nil
	}

	return nil, fmt.Errorf("parse token error by unknown reason")
}

func getSecret(ctx context.Context) ([]byte, error) {
	tx := query.Q.Begin()
	defer tx.Rollback()

	secret, err := tx.PersistData.
		WithContext(ctx).
		Where(tx.PersistData.Key.Eq("JWT_SECRET")).
		First()
	if err == nil {
		return []byte(secret.Value), nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("query jwt secret error: %w", err)
	}

	newSecret := make([]byte, 32)
	rand.Read(newSecret)

	err = tx.PersistData.
		WithContext(ctx).
		Create(&models.PersistData{
			Key:   "JWT_SECRET",
			Value: base64.StdEncoding.EncodeToString(newSecret),
		})
	if err != nil {
		return nil, fmt.Errorf("create persist data JWT_SECRET error: %w", err)
	}

	tx.Commit()
	return newSecret, nil
}
