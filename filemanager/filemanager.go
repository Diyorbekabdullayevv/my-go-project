package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"example.com/report_finance/userwindow"
)

type FinaceReporter struct {
	Income    float64 `json:"income_data"`
	Outcome   float64 `json:"outcome_data"`
	CreatedAt time.Time `json:"created_at"`
}

func UseFinanceReporter(income, outcome float64) *FinaceReporter {
	financeReporter := FinaceReporter{
		Income:    income,
		Outcome:   outcome,
		CreatedAt: time.Now(),
	}
	return &financeReporter
}

func GetUserIncomeOutcome() *FinaceReporter {
	income, outcome := userwindow.GetUserInput2()

	myFinanceReporter := UseFinanceReporter(income, outcome)
	return myFinanceReporter
}

func WriteToFile() {
	myStruct := GetUserIncomeOutcome()

	var myNewStruct []FinaceReporter

	jsonFile, err := os.ReadFile("finances.json")

	if err == nil && len(jsonFile) > 0 {
		err = json.Unmarshal(jsonFile, &myNewStruct)
		if err != nil {
			fmt.Println("Failed to Unmarshal data!")
		}
	}

	myNewStruct = append(myNewStruct, *myStruct)

	jsonData, err := json.MarshalIndent(myNewStruct, "", " ")

	if err != nil {
		fmt.Println("Failed to Marshal data!")
	}

	err = os.WriteFile("finances.json", jsonData, 0644)

	if err != nil {
		fmt.Println("Failed to create JSON file!")
	}
}

func GetJsonData() []FinaceReporter{
	data, err := os.ReadFile("finances.json")
	if err != nil{
		fmt.Println("Failed to read file!")
		return nil
	}

	var myStruct []FinaceReporter
	err = json.Unmarshal(data, &myStruct)
	if err != nil{
		fmt.Println("Failed to unmarshal file!")
		return nil
	}
	return myStruct
}
