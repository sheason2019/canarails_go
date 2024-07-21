package authsvc

import (
	"bytes"

	"canarails.dev/database/models"
)

func ComparePassword(
	password string,
	user *models.User,
) bool {
	currentHash := createPasswordHash(password, user.PasswordSalt)
	return bytes.Equal(currentHash, user.PasswordHash)
}
