package main

import (
	"fmt"
	"time"
)

func main() {
	// Get current date
	now := time.Now()

	// Get the first day of the current month
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	// Determine the weekday of the first day of the month (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	startingDay := firstOfMonth.Weekday()

	// Print the month and year header
	fmt.Printf("\t[%s %d]\n", now.Month().String(), now.Year())

	// Print the days of the week header
	fmt.Println("Su Mo Tu We Th Fr Sa")

	// Determine number of days in the current month
	daysInMonth := daysIn(now.Month(), now.Year())

	// Print leading spaces before the first day of the month
	for i := 0; i < int(startingDay); i++ {
		fmt.Printf("   ")
	}

	// Print each day of the month
	for day := 1; day <= daysInMonth; day++ {
		if (day != now.Day()) {
			fmt.Printf("%2d ", day)
		} else {
			fmt.Printf("%2d<", day) // CURRENT DAY
		}
		if (int(startingDay)+day)%7 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
}

// Function to determine number of days in a given month and year
func daysIn(month time.Month, year int) int {
	switch month {
	case time.February:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return 29 // Leap year
		}
		return 28
	case time.April, time.June, time.September, time.November:
		return 30
	default:
		return 31
	}
}

