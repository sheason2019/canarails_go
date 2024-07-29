package clientsvc_test

import (
	"testing"

	"canarails.dev/services/gatewaysvc/clientsvc"
)

func TestNew(t *testing.T) {
	clientsvc.New()
}

func TestDynamic(t *testing.T) {
	clientsvc.NewDynamic()
}
