package sqlstore

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	databaseURL := "host=localhost dbname=restapi_test sslmode=disable"
	/*	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable"
	}*/

	os.Exit(m.Run())
}
