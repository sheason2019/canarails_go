package tokensvc

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.RegisteredClaims

	UserId      uint
	UserTokenId uint
}

func (c *UserClaims) WithExpiredAt(t time.Time) *UserClaims {
	c.ExpiresAt = jwt.NewNumericDate(t)
	return c
}

func (c *UserClaims) WithUserTokenId(id uint) *UserClaims {
	c.UserTokenId = id
	return c
}
