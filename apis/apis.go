package apis

import (
	"canarails.dev/apis/appsapi"
	"canarails.dev/apis/appvariantsapi"
	"canarails.dev/apis/authapi"
	"canarails.dev/apis/genapi"
)

type authImpl = authapi.Impl
type appsImpl = appsapi.Impl
type appVariantsImpl = appvariantsapi.Impl

type Impl struct {
	authImpl
	appsImpl
	appVariantsImpl
}

func New() genapi.StrictServerInterface {
	return &Impl{}
}
