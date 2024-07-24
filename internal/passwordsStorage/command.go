package command

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	passwords "password-manager/pkg/passwords"

	"strings"
)

const FilePath = "./internal/passwordsStorage/passwords.json"

func StorePasswords(store *passwords.PasswordStore) {
	file, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу: %v", err)
		return
	}

	defer file.Close()

	var passwordsInStore []passwords.Password

	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("помилка читання інфо : %v /n", err)
		return
	}

	if stat.Size() != 0 {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&passwordsInStore)
		if err != nil {
			fmt.Printf("Помилка декодування json: %v\n", err)
			return
		}
	}

	for _, password := range store.Passwords {
		passwordsInStore = append(passwordsInStore, passwords.Password{
			Name:  password.Name,
			Value: password.Value,
		})
	}
	file.Truncate(0)
	file.Seek(0, 0)

	data, err := json.MarshalIndent(passwordsInStore, "", " ")
	if err != nil {
		fmt.Printf("Помилка парсінгу джейсон %v", err)
		return
	}

	_, err = file.Write(data)

	if err != nil {
		fmt.Printf("помилка запису в файл %v ", err)
		return

	}

}

func SavePassword(store *passwords.PasswordStore) {
	scanner := bufio.NewScanner(os.Stdin)
	var name, password string
	for {
		fmt.Println("Введіть назву для пароля ====>")

		if !scanner.Scan() {
			fmt.Printf("Помилка введення: %v", scanner.Err())
			return
		}

		name = strings.TrimSpace(scanner.Text())
		if PasswordExists(name) {
			fmt.Printf("Пароль с назвою вже існує - %v\n", name)
			return
		}

		if name != "" {
			break
		} else {
			fmt.Printf("Назва не може бути порожньою \n")
		}

	}

	for {
		fmt.Println("Введіть пароль ====>")

		if !scanner.Scan() {
			fmt.Printf("Помилка введення: %v", scanner.Err())
			return
		}

		password = strings.TrimSpace(scanner.Text())

		if password != "" {
			break
		} else {
			fmt.Printf("пароль не може бути порожнім")
		}

	}

	store.SavePassword(name, password)
	StorePasswords(store)
	fmt.Println("\n✅   Пароль успіщно збережнео")
}

func ShowPasswords() {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу: %v\n", err)
		return
	}

	defer file.Close()
	var passwords []passwords.Password
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&passwords)
	if err != nil {
		fmt.Printf("Помилка декодування JSON: %v\n", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("🟢  Порожній список ")
	} else {
		fmt.Println("🟢  Паролі, що збережені в список: ")
		for _, password := range passwords {
			fmt.Printf("Name: %s, Password: %s\n", password.Name, password.Value)
		}
	}

}

func GetPassword() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("ВВедіть назву пароля, що хочете отримати")
	if !scanner.Scan() {
		fmt.Printf("Помилка введення даних: %v", scanner.Err())
		return
	}

	name := strings.TrimSpace(scanner.Text())

	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу: %v\n", err)
		return
	}

	defer file.Close()
	var passwords []passwords.Password
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&passwords)
	if err != nil {
		fmt.Printf("Помилка декодування JSON: %v\n", err)
		return
	}

	found := false
	for _, password := range passwords {
		if password.Name == name {
			fmt.Printf("Пароль для %s: %s\n", name, password.Value)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Пароль з назвою '%s' не знайдено\n", name)
	}
}

func PasswordExists(name string) bool {

	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу: %v\n", err)
		return false
	}

	defer file.Close()

	var passwordsInStore []passwords.Password
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&passwordsInStore)
	if err != nil {
		fmt.Printf("Помилка декодування JSON: %v\n", err)
		return false
	}

	for _, p := range passwordsInStore {
		if p.Name == name {
			return true
		}
	}
	return false
}
