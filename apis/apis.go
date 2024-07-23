package apis

import (
	"canarails.dev/apis/appsapi"
	"canarails.dev/apis/authapi"
	"canarails.dev/apis/genapi"
)

type authImpl = authapi.Impl
type appsImpl = appsapi.Impl

type Impl struct {
	authImpl
	appsImpl
}

func New() genapi.StrictServerInterface {
	return &Impl{}
}
