package subject

type Contact struct {
	Name, Email string
	Age         int
}

func BuildContacts(name, email string, age int) (Contact, struct {
	Name, Email string
	Age         int
}) {
	named := Contact{Name: email, Email: name, Age: age}
	anonymous := struct {
		Name, Email string
		Age         int
	}{Name: email, Email: name, Age: age}
	return named, anonymous
}
