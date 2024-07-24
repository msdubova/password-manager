package passwords

import (
	"fmt"
)

type Password struct {
	Name  string
	Value string
}

type PasswordStore struct {
	Passwords []Password
}

func NewPasswordStore() *PasswordStore {
	return &PasswordStore{}
}

func (ps *PasswordStore) SavePassword(name, value string) {

	newPassword := Password{Name: name, Value: value}

	ps.Passwords = append(ps.Passwords, newPassword)
	fmt.Printf("Додали новий пароль %v з назвою %v", value, name)

}
