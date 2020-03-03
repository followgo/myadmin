package ldapclient

import (
	"fmt"
	"strings"

	"github.com/go-ldap/ldap"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/util"
	"github.com/followgo/myadmin/util/errors"
)

// search 从服务器搜索用户信息，指定 RDN 搜索过程中还会匹配 RDN 对应的属性
func Search(rdn string) (users []map[string]string, err error) {
	conn, err := dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	filter := Cfg.LDAP.SearchFilter
	if rdn != "" {
		filter = filter[:len(filter)-1] + fmt.Sprintf("(%s=%s)", Cfg.LDAP.RDNAttr, rdn) + ")"
	}
	return search(conn, Cfg.LDAP.SearchBaseDN, filter)
}

func search(conn *ldap.Conn, baseDN, filter string) (users []map[string]string, err error) {
	// First bind with a read only user
	err = conn.Bind(Cfg.LDAP.BindSearcherDN, Cfg.LDAP.BindSearcherDNPassword)
	if err != nil {
		return nil, errors.Wrap(err, "bind searcherDN failed")
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
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 10, false,
		filter, attrs, nil,
	)
	result, err := conn.Search(searchRequest)
	if err != nil {
		return nil, errors.Wrap(err, "search failed")
	}

	// 读取结果，并返回
	if len(result.Entries) == 0 {
		return nil, ErrNoFound
	}
	users = make([]map[string]string, 0, len(result.Entries))
	for _, e := range result.Entries {
		user := make(map[string]string)
		user["dn"] = e.DN
		for _, attr := range attrs {
			user[attr] = strings.Join(e.GetAttributeValues(attr), ",")
		}

		users = append(users, user)
	}

	return users, nil
}
