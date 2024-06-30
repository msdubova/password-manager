package main

import (
	"bufio"
	"fmt"
	"os"
	"password-manager/internal/utils"
	"password-manager/pkg/passwords"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	store := passwords.NewPasswordStore()

	fmt.Println("🟢  Вітаю у менеджері паролів. Оберіть опцію 1, 2, 3, 4. \n❗️  Користувач може обрати лише цифру")
	fmt.Println("1️⃣  Якщо хочете вивести назви всіх збережених паролів - натисніть 1")
	fmt.Println("2️⃣  Якщо хочете зберегти новий пароль - натисніть 2")
	fmt.Println("3️⃣  Якщо хочете дістати збережений пароль - натисніть 3")
	fmt.Println("4️⃣  Якщо хочете покинути програму  - натисніть 4")

	var userChoice int
	for scanner.Scan() {
		input := scanner.Text()
		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("ПОмилка при читанні даних користувача", err)
			continue
		}

		userChoice = choice
		break
	}

	switch userChoice {
	case 1:
		utils.ShowPasswords()
	case 2:

		utils.SavePassword(store)
	case 3:

		utils.GetPassword()

	default:
		fmt.Printf("Oберіть один с запропонованих варіанті від 1 до 4")
	}
}
