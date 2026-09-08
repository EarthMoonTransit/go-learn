package subject

type Contact struct {
	Name, Email string
	Age         int
}

func BuildContacts(name, email string, age int) (Contact, struct {
	Name, Email string
	Age         int
}) {
	named := Contact{Name: name, Email: email, Age: age}
	anonymous := struct {
		Name, Email string
		Age         int
	}{Name: name, Email: email, Age: age}
	return named, anonymous
}
