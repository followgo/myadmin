package ldapauth

import (
	"crypto/tls"
	"strings"

	"github.com/go-ldap/ldap"

	"github.com/followgo/myadmin/util"
	"github.com/followgo/myadmin/util/errors"
)

// ErrNoFound 没有找到任何用户
var ErrNoFound = errors.New("no users found")

// LDAPSearch 从服务器搜索用户信息
func LDAPSearch(server, bindDN, bindDNPwd, baseDN, filter, rdn string, attrs []string, useTLS bool) (users []map[string]string, err error) {
	var l *ldap.Conn

	// 连接服务器
	if useTLS {
		l, err = ldap.DialTLS("tcp", server, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return nil, errors.Wrap(err, "dial with tls failed")
		}
		defer l.Close()
	} else {
		l, err = ldap.Dial("tcp", server)
		if err != nil {
			return nil, errors.Wrap(err, "dial without tls failed")
		}
		defer l.Close()

		// Reconnect with TLS
		err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return nil, errors.Wrap(err, "reconnect with TLS")
		}
	}

	// First bind with a read only user
	err = l.Bind(bindDN, bindDNPwd)
	if err != nil {
		return nil, errors.Wrap(err, "bind user failed")
	}

	// Search for the given infos of users
	if !util.HasStringSlice(rdn, attrs, false) {
		attrs = append(attrs, rdn)
	}
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, attrs, nil,
	)

	result, err := l.Search(searchRequest)
	if err != nil {
		return nil, errors.Wrap(err, "search failed")
	}
	if len(result.Entries) == 0 {
		return nil, ErrNoFound
	}

	// 读取结果
	users = make([]map[string]string, 0, len(result.Entries))
	for _, e := range result.Entries {
		user := make(map[string]string)
		user["DN"] = e.DN
		for _, attr := range attrs {
			user[attr] = strings.Join(e.GetAttributeValues(attr), ",")
		}
		users = append(users, user)
	}

	return users, nil
}
