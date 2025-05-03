package reportfinance

import (
	"fmt"

	"example.com/report_finance/filemanager"
	"example.com/report_finance/userwindow"
)

func IfUserTask() {
	var userTask int
	userCurrentTask := userwindow.GetUserInput("Qaysi amalni bajarishni xohlaysiz?", "Raqamni kiriting: ", userTask)

	if userCurrentTask == 1 {
		userwindow.UserWindow2("1. Hozirgacha bo`lgan umumiy kirim-chiqimlarni ko`rish", "2. Har kungi kirim-chiqimlarni ko`rish")

		var userTask int
		userCurrentTask := userwindow.GetUserInput("Qaysi amalni bajarishni xohlaysiz?", "Raqamni kiriting: ", userTask)

		if userCurrentTask == 1 {
			filemanager.OutputJsonData()
		} else {
			fmt.Println("Hozircha tamirlash ishlari ketmoqda!")
		}
	} else if userCurrentTask == 2 {
		filemanager.WriteToFile()
	}
}
