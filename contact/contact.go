package contact

import (
	"fmt"
	"github.com/nimajalali/go-force/sobjects"
)

type Contact struct {
	sobjects.BaseSObject
	Email       string `force:"Email"`
	MobilePhone string `force:"MobilePhone"`
	Title       string `force:"Title"`
	Department  string `force:"Department"`
}

func Header() string {
	return "Id\tName\tEmail\tMobilePhone\tTitle\tDepartment"
}

func (c *Contact) ApiName() string {
	return "Contact"
}

func (c *Contact) CSV() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s", c.Id, c.Name, c.Email, c.MobilePhone, c.Title, c.Department)
}
