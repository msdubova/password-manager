package utils

import (
	"bufio"
	"fmt"
	"os"
	"password-manager/pkg/passwords"
	"strings"
)

const filePath = "./internal/passwordsStorage/passwords.txt"

func StorePasswords(store *passwords.PasswordStore) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, password := range store.Passwords {
		fmt.Fprintf(writer, "%s:%s\n", password.Name, password.Value)
	}
	writer.Flush()
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
		if name != "" {
			break
		} else {
			fmt.Printf("Назва не може бути порожньею")
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
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("🟢  Паролі, що збережені в список: ")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func GetPassword() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("ВВедіть назву пароля, що хочете отримати")
	if !scanner.Scan() {
		fmt.Printf("Помилка введення даних", scanner.Err())
		return
	}

	name := strings.TrimSpace(scanner.Text())

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу: %v\n", err)
		return
	}

	defer file.Close()

	found := false

	scannerFile := bufio.NewScanner(file)
	for scannerFile.Scan() {
		line := strings.TrimSpace(scannerFile.Text())
		parts := strings.Split(line, ":")
		if len(parts) == 2 && parts[0] == name {
			fmt.Printf("Пароль для %s: %s\n", name, parts[1])
			found = true
			break
		}
	}

	if found == false {
		fmt.Printf("Пароль з назвою '%s' не знайдено\n", name)
	}
}
