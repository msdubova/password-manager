package main

import "fmt"

func main() {
	fmt.Println("Вітаю у менеджері паролів")
	fmt.Println("Якщо хочете вивести назви всіх збережених паролів - натисніть 1")
	fmt.Println("Якщо хочете зберегти новий пароль - натисніть 2")
	fmt.Println("Якщо хочете дістати збережений пароль - натисніть 3")
	fmt.Println("Якщо хочете покинути програму  - натисніть 4")
	var userChoice int

	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		fmt.Printf("Ви обрали -  %v", userChoice)
		ShowPasswords()

	case 2:
		fmt.Printf("Ви обрали -  %v", userChoice)
		SavePassword()
	case 3:
		fmt.Printf("Ви обрали -  %v", userChoice)
		GetPassword
	case 4:
		fmt.Printf("Ви обрали -  %v", userChoice)
		ExitProgram()
	default:
		fmt.Printf("Oберіть один с запропонованих варіанті від 1 до 4")
	}
}
