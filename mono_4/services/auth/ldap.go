package main

import (
	"github.com/go-ldap/ldap/v3"
)

// CWE-90: caller value interpolated into the LDAP filter.
func LDAPSearch(uid string) ([]*ldap.Entry, error) {
	conn, err := ldap.Dial("tcp", "ldap.northwind.internal:389")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	filter := "(uid=" + uid + ")"
	request := ldap.NewSearchRequest(
		"dc=northwind,dc=internal",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		[]string{"cn", "mail", "uid"},
		nil,
	)
	result, err := conn.Search(request)
	if err != nil {
		return nil, err
	}
	return result.Entries, nil
}

// Safe false-positive bait: escaped filter.
func LDAPSearchSafe(uid string) ([]*ldap.Entry, error) {
	conn, err := ldap.Dial("tcp", "ldap.northwind.internal:389")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	filter := "(uid=" + ldap.EscapeFilter(uid) + ")"
	request := ldap.NewSearchRequest("dc=northwind,dc=internal", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filter, []string{"cn"}, nil)
	result, err := conn.Search(request)
	if err != nil {
		return nil, err
	}
	return result.Entries, nil
}
