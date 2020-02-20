package ldapauth

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-ldap/ldap"

	"github.com/followgo/myadmin/util/errors"
)

// Login 登陆
func Login(addr, username, password string, useTLS bool) (err error) {
	var l *ldap.Conn

	// 连接服务器
	if useTLS {
		l, err = ldap.DialTLS("tcp", addr, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return errors.Wrap(err, "dial with tls failed")
		}
		defer l.Close()
	} else {
		l, err = ldap.Dial("tcp", addr)
		if err != nil {
			return errors.Wrap(err, "dial without tls failed")
		}
		defer l.Close()

		// Reconnect with TLS
		err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return err
		}
	}

	// First bind with a read only user
	err = l.Bind(username, password)
	if err != nil {
		return errors.Wrap(err, "bind user failed")
	}

	// Search for the given username
	var (
		baseDn = username[strings.Index(username, "dc"):]
		leaf   = strings.TrimSpace(strings.Split(username, ",")[0])
	)
	searchRequest := ldap.NewSearchRequest(
		baseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=person)(%s))", leaf),
		nil, nil,
	)

	result, err := l.Search(searchRequest)
	if err != nil {
		return errors.Wrap(err, "search failed")
	}

	if len(result.Entries) == 0 {
		return errors.New("user does not exist or too many entries returned")
	}

	result.PrettyPrint(2)

	return nil
}
