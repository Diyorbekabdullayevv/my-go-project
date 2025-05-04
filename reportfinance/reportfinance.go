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
		userwindow.UserWindow2("1. Hozirgacha bo`lgan umumiy kirim-chiqimlarni ko`rish", "2. Muayyan sanadagi kirim-chiqimlarni ko`rish")

		var userTask int
		userCurrentTask := userwindow.GetUserInput("Qaysi amalni bajarishni xohlaysiz?", "Raqamni kiriting: ", userTask)

		if userCurrentTask == 1 {

			fmt.Println("Hozircha tamirlash ishlari ketmoqda!")
		} else {
			var dateTime string
			var formattedDate string

			fmt.Print("Sanani kiriting(kk/oo/yyyy): ")
			fmt.Scan(&dateTime)

			myStruct := filemanager.GetJsonData()

			for _, entry := range myStruct {
				formattedDate = entry.CreatedAt.Format("02/01/2006")
			}

			if dateTime == formattedDate {
				for i, entry := range myStruct {
					if i == 2{
					fmt.Printf("  Kirim: %.2f\n", entry.Income)
					fmt.Printf("  Chiqim: %.2f\n", entry.Outcome)
					fmt.Printf("  Date: %s\n", entry.CreatedAt.Format("02/01/2006"))
				}}
			}
			// for i, entry := range myStruct {
			// 	fmt.Printf("Entry %d:\n", i+1)
			// 	fmt.Printf("  Income: %.2f\n", entry.Income)
			// 	fmt.Printf("  Outcome: %.2f\n", entry.Outcome)
			// 	fmt.Printf("  Date: %s\n", entry.CreatedAt.Format("02/01/2006"))
			// }
		}
	} else if userCurrentTask == 2 {
		filemanager.WriteToFile()
	}
}
