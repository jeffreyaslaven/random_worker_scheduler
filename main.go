package main

import (
	"fmt"
	"math/rand"
)

type worker struct {
	firstName string
	lastName  string
	position  string
}

type schedule struct {
	worker       worker
	scheduledDay string
}

func contains(sliceToCheck []string, value string) bool {
	for _, sliceValue := range sliceToCheck {
		if sliceValue == value {
			return true
		}
	}
	return false
}

func main() {
	var howManyWorkers int
	var firstName string
	var lastName string
	var position string
	var workers []worker
	var scheduledWorkers []schedule
	var daysCovered []string
	daysOfWeek := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	fmt.Println("How many workers do you have on staff: ")
	fmt.Scanf("%v", &howManyWorkers)

	//Gathering worker information
	for i := 0; i < howManyWorkers; i++ {
		fmt.Println("Please enter the first name of the worker: ")
		fmt.Scanf("%v", &firstName)
		fmt.Println("Please enter the last name of the worker: ")
		fmt.Scanf("%v", &lastName)
		fmt.Println("Please enter the position of the worker: ")
		fmt.Scanf("%v", &position)
		fmt.Println("_________________________________________________")

		workers = append(workers, worker{firstName: firstName, lastName: lastName, position: position})
	}

	// Assigning a day of the week to each worker
	for i := 0; i < howManyWorkers; i++ {
		// Clearing days covered
		if len(daysCovered) == 5 {
			daysCovered = nil
		}
		for {
			dayValue := daysOfWeek[rand.Intn(5)]
			if !contains(daysCovered, dayValue) {
				scheduledWorkers = append(scheduledWorkers, schedule{worker: workers[i], scheduledDay: dayValue})
				daysCovered = append(daysCovered, dayValue)
				break
			}
		}
	}

	fmt.Println("The following schedule was generated: ")
	fmt.Println("_________________________________________________")
	for i := 0; i < howManyWorkers; i++ {
		fmt.Print("First Name: ")
		fmt.Println(scheduledWorkers[i].worker.firstName)
		fmt.Print("Last Name: ")
		fmt.Println(scheduledWorkers[i].worker.lastName)
		fmt.Print("Position: ")
		fmt.Println(scheduledWorkers[i].worker.position)
		fmt.Print("Day Scheduled To Work: ")
		fmt.Println(scheduledWorkers[i].scheduledDay)
		fmt.Println("_________________________________________________")
	}
}
