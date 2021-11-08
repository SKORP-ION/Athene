package service

import (
	"Athena/ldap"
	"Athena/ldap/api"
	. "Athena/log"
	"context"
)

type GrpcService struct {
	api.UnimplementedLdapServer
}

func (GrpcService) SearchEmployee(c context.Context, Name *api.SearchRequest) (*api.SearchResponse, error) {
	if !ldap.IsConnected() {
		ldap.Connect()
	}

	word := Name.Name

	result, err := ldap.Search(ldap.DN, word)

	if err != nil {
		Error.Printf("Failed to search %s\n", word)
		return nil, err
	}

	SearchResponse := &api.SearchResponse{}

	for _, entry := range result.Entries {

		cn := entry.GetAttributeValue("cn")
		title := entry.GetAttributeValue("title")
		department := entry.GetAttributeValue("extensionAttribute2")
		employeeNumber := entry.GetAttributeValue("employeeNumber")

		manager := entry.GetAttributeValue("manager")

		managerList := ldap.RegExp.FindStringSubmatch(manager)
		managerParse := ""

		if len(managerList) >= 1 {
			managerParse = managerList[0]
		} else {
			managerParse = manager
		}

		employee := &api.LdapEmployee{
			EmployeeNumber: employeeNumber,
			Name:           cn,
			Manager:        managerParse,
			Department:     department,
			JobTitle:       title,
		}
		SearchResponse.Staff = append(SearchResponse.Staff, employee)
	}
	Info.Printf("Send result for search %s\n", word)
	return SearchResponse, nil
}
