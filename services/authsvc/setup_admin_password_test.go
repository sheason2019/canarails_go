package authsvc_test

import (
	"context"
	"testing"

	"canarails.dev/database"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
)

func TestSetupAdminPassword(t *testing.T) {
	password := "test password"
	query.SetDefault(database.GetDb())
	admin := authsvc.SetupAdminPassword(context.Background(), password)
	if !authsvc.ComparePassword(password, admin) {
		t.Errorf("password not equal")
	}
}
