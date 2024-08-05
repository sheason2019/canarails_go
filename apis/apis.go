package apis

import (
	"canarails.dev/apis/appsapi"
	"canarails.dev/apis/appvariantsapi"
	"canarails.dev/apis/authapi"
	"canarails.dev/apis/genapi"
	"canarails.dev/apis/usertokensapi"
	"canarails.dev/apis/versionsapi"
)

type authImpl = authapi.Impl
type appsImpl = appsapi.Impl
type appVariantsImpl = appvariantsapi.Impl
type userTokensImpl = usertokensapi.Impl
type versionsImpl = versionsapi.Impl

type Impl struct {
	authImpl
	appsImpl
	appVariantsImpl
	userTokensImpl
	versionsImpl
}

func New() genapi.StrictServerInterface {
	return &Impl{}
}
