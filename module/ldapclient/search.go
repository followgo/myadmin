package ldapclient

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-ldap/ldap"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/util"
	"github.com/followgo/myadmin/util/errors"
)

// Search 从服务器搜索用户信息
// 不指定 RDN 则返回所有成员
func Search(rdn string) (users []map[string]string, err error) {
	var l *ldap.Conn
	filter := Cfg.LDAP.SearchFilter
	if rdn != "" {
		filter = filter[:len(filter)-1] + fmt.Sprintf("(%s=%s)", Cfg.LDAP.RDNAttr, rdn) + ")"
	}

	// 连接服务器
	if Cfg.LDAP.UseTLS {
		l, err = ldap.DialTLS("tcp", Cfg.LDAP.ServerAddr, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return nil, errors.Wrap(err, "dial with tls failed")
		}
		defer l.Close()
	} else {
		l, err = ldap.Dial("tcp", Cfg.LDAP.ServerAddr)
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
	err = l.Bind(Cfg.LDAP.BindSearcherDN, Cfg.LDAP.BindSearcherDNPassword)
	if err != nil {
		return nil, errors.Wrap(err, "bind user failed")
	}

	// Search for the given infos of users
	attrs := []string{
		Cfg.LDAP.UserAttributes.CommonName,
		Cfg.LDAP.UserAttributes.Surname,
		Cfg.LDAP.UserAttributes.Description,
		Cfg.LDAP.UserAttributes.Telephone,
	}
	if !util.HasStringSlice(Cfg.LDAP.RDNAttr, attrs, false) {
		attrs = append(attrs, Cfg.LDAP.RDNAttr)
	}
	searchRequest := ldap.NewSearchRequest(
		Cfg.LDAP.SearchBaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, attrs, nil,
	)
	result, err := l.Search(searchRequest)
	if err != nil {
		return nil, errors.Wrap(err, "search failed")
	}

	// 读取结果
	if len(result.Entries) == 0 {
		return nil, ErrNoFound
	}
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
