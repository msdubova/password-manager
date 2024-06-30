package passwords

type Password struct {
	Name  string
	Value string
}

type PasswordService interface {
	ShowPasswords() ([]string, error)
	SavePassword(name, value string) error
	GetPassword(name string) (string, error)
}
