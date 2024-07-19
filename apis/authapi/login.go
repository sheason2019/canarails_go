package authapi

import (
	"context"
	"log"

	"canarails.dev/apis/genapi"
)

func (Impl) AuthLogin(
	ctx context.Context,
	request genapi.AuthLoginRequestObject,
) (genapi.AuthLoginResponseObject, error) {
	log.Printf("receive login, username: %s, password: %s\n", request.Body.Username, request.Body.Password)

	return genapi.AuthLogin200JSONResponse(genapi.LoginRes{}), nil
}
