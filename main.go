package main

import (
	"fmt"

	"example.com/report_finance/reportfinance"
	"example.com/report_finance/userwindow"
)

func main() {
	userwindow.UserWindow()
	reportfinance.IfUserTask()
	fmt.Println("Updated main file")
}
