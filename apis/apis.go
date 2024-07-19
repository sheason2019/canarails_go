package apis

import (
	"canarails.dev/apis/authapi"
	"canarails.dev/apis/genapi"
)

type Impl struct {
	authapi.Impl
}

func New() genapi.StrictServerInterface {
	return &Impl{}
}
