package subject

import (
	"strings"
)

type Contact struct {
	Name, Email string
	Age         int
}

func BuildContacts(name, email string, age int) (Contact, struct {
	Name, Email string
	Age         int
}) {
	named := Contact{Name: strings.TrimSpace(name), Email: email, Age: age}
	anonymous := struct {
		Name, Email string
		Age         int
	}{Name: strings.TrimSpace(name), Email: email, Age: age}
	return named, anonymous
}
