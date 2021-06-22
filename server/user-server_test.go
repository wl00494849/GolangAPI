package server_test

import (
	"golang-api/server"
	"testing"
)

func Test_UsersList(t *testing.T) {
	u := server.UsersList()

	if u != nil {
		t.Logf("Have %d row data", len(u))
	} else {
		t.Error("Noting")
	}
}
