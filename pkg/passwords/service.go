package passwords

type PasswordStore struct {
	Passwords []Password
}

func NewPasswordStore() *PasswordStore {
	return &PasswordStore{}
}
