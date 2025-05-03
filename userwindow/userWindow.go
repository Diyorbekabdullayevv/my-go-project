package userwindow

import "fmt"

func UserWindow() {
	fmt.Println("")
	fmt.Println("-----KIRIM-CHIQIM DASTURI-----")
	fmt.Println("")
	fmt.Println("1. Kirim-chiqimlarni ko`rish")
	fmt.Println("2. Kirim-chiqim hisoboti yozish")
	fmt.Println("")
}

func UserWindow2(path1, path2 string) {
	fmt.Println(path1)
	fmt.Println(path2)
	fmt.Println("")
}

func GetUserInput(path1, path2 string, userTask int) int {
	fmt.Println(path1)
	fmt.Print(path2)
	fmt.Scan(&userTask)
	fmt.Println("")
	fmt.Println("Assalomu alaykum")
	return userTask
}
