package ldapauth

import (
	"testing"
)

func TestLogin(t *testing.T) {
	var (
		server   = "192.168.50.131:636"
		username = "cn=Join,ou=People,dc=firstmile,dc=cn"
		password = "123456"
		useTLS   = true
	)
	if err := Login(server, username, password, useTLS); err != nil {
		t.Error(err)
	}

}
