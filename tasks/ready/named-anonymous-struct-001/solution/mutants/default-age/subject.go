package subject

type Contact struct {
	Name, Email string
	Age         int
}

func BuildContacts(name, email string, age int) (Contact, struct {
	Name, Email string
	Age         int
}) {
	if age == 0 {
		age = 18
	}
	named := Contact{Name: name, Email: email, Age: age}
	anonymous := struct {
		Name, Email string
		Age         int
	}{Name: name, Email: email, Age: age}
	return named, anonymous
}
