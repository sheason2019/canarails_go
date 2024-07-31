package apis

import (
	"canarails.dev/apis/appsapi"
	"canarails.dev/apis/appvariantsapi"
	"canarails.dev/apis/authapi"
	"canarails.dev/apis/genapi"
	"canarails.dev/apis/usertokensapi"
)

type authImpl = authapi.Impl
type appsImpl = appsapi.Impl
type appVariantsImpl = appvariantsapi.Impl
type userTokensImpl = usertokensapi.Impl

type Impl struct {
	authImpl
	appsImpl
	appVariantsImpl
	userTokensImpl
}

func New() genapi.StrictServerInterface {
	return &Impl{}
}
