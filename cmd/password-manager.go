package main

import "fmt"

func main() {
	fmt.Println("Вітаю у менеджері паролів. Оберіть опцію да введіть її , натисніть enter щоб продоавжити")

	var userChoice int

	fmt.Scan(&userChoice)

	fmt.Printf("User's choice -  %v", userChoice)
}
