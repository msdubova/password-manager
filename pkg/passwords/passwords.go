package passwords

import "fmt"

type Password struct {
	Name  string
	Value string
}

type PasswordService interface {
	ShowPasswords() ([]Password, error)
	SavePassword(name, value string) error
	GetPassword(name string) (Password, error)
}

type PasswordStore struct {
	Passwords []Password
}

func NewPasswordStore() *PasswordStore {
	return &PasswordStore{}
}

func (ps *PasswordStore) ShowPasswords() ([]Password, error) {
	fmt.Println(ps.Passwords)
	return ps.Passwords, nil
}

func (ps *PasswordStore) SavePassword(name, value string) error {
	newPassword := Password{Name: name, Value: value}

	ps.Passwords = append(ps.Passwords, newPassword)
	fmt.Printf("Додали новий пароль %v з назвою %v", value, name)
	return nil
}

func (ps *PasswordStore) GetPassword(name string) (Password, error) {
	return Password{}, fmt.Errorf("Пароль з такою назвою не знайдено")
}
