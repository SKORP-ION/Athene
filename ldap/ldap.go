package ldap

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var RegExp = regexp.MustCompile("[А-ЯЁа-яё]+\\s[А-ЯЁа-яё]+\\s[А-ЯЁа-яё]*")

var (
	DN           []string
	ServiceHost  string
	ldapUsername string
	ldapPassword string
	ldapAddr     string
	conn         *ldap.Conn
	connected    = false
	lastQuery    = time.Now()
)

func init() {
	ldapUsername = os.Getenv("ldapUsername")
	ldapPassword = os.Getenv("ldapPassword")
	ldapAddr = os.Getenv("ldapAddr")
	ServiceHost = os.Getenv("serviceHost")
	DN = strings.Split(os.Getenv("DN"), ";")
	go DisconnectAfterMin()
}

func Connect() {
	var err error
	conn, err = ldap.Dial("tcp", ldapAddr)

	if err != nil {
		log.Fatalln(err)
	}

	err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})

	if err != nil {
		log.Fatalln(err)
	}

	err = conn.Bind(ldapUsername, ldapPassword)

	if err != nil {
		log.Fatalln(err)
	}
	connected = true
}

func DisconnectAfterMin() {
	for {
		if time.Now().Sub(lastQuery).Minutes() >= 1 && connected {
			CloseConnection()
			connected = false
		}
		time.Sleep(time.Second * 30)
	}
}

func CloseConnection() {
	conn.Close()
}

func IsConnected() bool {
	return connected
}

func Search(DNs []string, searchWord string) (*ldap.SearchResult, error) {
	lastQuery = time.Now()
	result := &ldap.SearchResult{}
	for _, dn := range DNs {
		searchRequest := ldap.NewSearchRequest(dn, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0,
			false, fmt.Sprintf("(&(objectClass=user)(objectCategory=person)(CN=%s*))", searchWord),
			[]string{"CN", "Manager", "title", "EmployeeNumber", "extensionAttribute2"},
			nil)
		localResult, err := conn.Search(searchRequest)

		if err != nil {
			return nil, err
		}
		result.Entries = append(result.Entries, localResult.Entries...)
	}

	return result, nil
}
