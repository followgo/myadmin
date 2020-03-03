package ldapclient

import (
	"crypto/tls"
	"time"

	"github.com/go-ldap/ldap"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/util/errors"
)

// dial 连接 LDAP 服务器
func dial() (conn *ldap.Conn, err error) {
	// LDAPs
	if Cfg.LDAP.UseTLS {
		conn, err = ldap.DialTLS("tcp", Cfg.LDAP.ServerAddr, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return nil, errors.Wrap(err, "dial with tls failed")
		}

		conn.SetTimeout(10 * time.Second)
		return conn, nil
	}

	// LDAP
	conn, err = ldap.Dial("tcp", Cfg.LDAP.ServerAddr)
	if err != nil {
		return nil, errors.Wrap(err, "dial without tls failed")
	}

	if Cfg.LDAP.StartTLS {
		// Reconnect with TLS
		err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return nil, errors.Wrap(err, "reconnect with TLS")
		}
	}

	conn.SetTimeout(10 * time.Second)
	return conn, nil
}
