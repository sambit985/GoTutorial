package main

import (
	"fmt"
	"time"
)

func main() {
	liveDateTime := time.Now()
	fmt.Println("Current time is:--", liveDateTime)
	fmt.Printf("live time type is:- %T\n", liveDateTime)

	//format date
	var formatted = liveDateTime.Format("02-01-2006")
	fmt.Println("Formatted Time is:", formatted)

	//get current day
	formattedDay := liveDateTime.Format("02-01-2006, Monday")
	fmt.Println("The formatted day is:--", formattedDay)

	//get live time
	formattedTime := liveDateTime.Format("02-01-2006, 15:04:05")
	fmt.Println("Live time is:-", formattedTime)

	// 1 day add
	new_day := liveDateTime.Add(24 * time.Hour)
	fmt.Println("After one day:--", new_day)
	formatted_newDay := new_day.Format("2006/01/02")
	fmt.Println("After One day:--", formatted_newDay)
}
