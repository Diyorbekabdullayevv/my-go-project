package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
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
	var income float64
	fmt.Print("Kirim: ")
	fmt.Scan(&income)

	var outcome float64
	fmt.Print("Chiqim: ")
	fmt.Scan(&outcome)

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

func OutputJsonData(){
	data, err := os.ReadFile("finances.json")
	if err != nil{
		fmt.Println("Failed to read file!")
		return
	}

	var myStruct []FinaceReporter
	err = json.Unmarshal(data, &myStruct)
	if err != nil{
		fmt.Println("Failed to unmarshal file!")
		return
	}

	for i, entry := range myStruct {
		fmt.Printf("Entry %d:\n", i+1)
		fmt.Printf("  Income: %.2f\n", entry.Income)
		fmt.Printf("  Outcome: %.2f\n", entry.Outcome)
		fmt.Printf("  Date: %s\n", entry.CreatedAt.Format("02/01/2006"))
	}
}
