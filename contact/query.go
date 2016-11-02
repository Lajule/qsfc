package contact

import (
	"github.com/nimajalali/go-force/force"
	"github.com/nimajalali/go-force/sobjects"
	"log"
)

type ContactQueryResponse struct {
	sobjects.BaseQuery
	Records []*Contact `force:"records"`
}

func Contacts(forceApi *force.ForceApi) *ContactQueryResponse {
	contacts := &ContactQueryResponse{}
	query := force.BuildQuery("Id, Name, Email, MobilePhone, Title, Department", "Contact", nil)
	err := forceApi.Query(query, contacts)
	if err != nil {
		log.Fatal(err)
	}
	return contacts
}

func NextContacts(forceApi *force.ForceApi, contacts *ContactQueryResponse) *ContactQueryResponse {
	err := forceApi.QueryNext(contacts.NextRecordsUri, contacts)
	if err != nil {
		log.Fatal(err)
	}
	return contacts
}
