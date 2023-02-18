package main

import (
	"fmt"

	"github.com/homayoonalimohammadi/splitshare/events"
	"github.com/homayoonalimohammadi/splitshare/member"
)

func main() {

	// Load from file or Input event manually
	var err error
	event := &events.Event{}
	for {
		fmt.Println("Choose an option:\n(1) Load from disk\n" +
			"(2) Input events manually")
		var opt string
		fmt.Scanln(&opt)
		if opt == "1" {
			var fileName string
			fmt.Println("File name (without extension): ")
			fmt.Scanln(&fileName)

			event, err = event.LoadFromDisk(fileName)
			if err != nil {
				fmt.Println("failed to load from events from disk:", err)
			}
			fmt.Println("Successfully loaded events from disk.")
			break
		} else if opt == "2" {
			event, err = enterEvent()
			if err != nil {
				fmt.Println("failed to enter events manually:", err)
			}
			break
		} else {
			fmt.Println("Enter a number from the available options.")
		}
	}

	event.Describe()
}

func enterEvent() (*events.Event, error) {

	// Enter members names
	event := events.New()
	for {
		fmt.Println("Choose an option:\n(1) Enter Member name\n(2) Proceed to assigning shares")
		var opt string
		fmt.Scanln(&opt)
		if opt == "1" {
			fmt.Println("Member Name: ")
			var name string
			fmt.Scanln(&name)
			member := member.New(name)
			event.Members = append(event.Members, member)
		} else if opt == "2" {
			break
		} else {
			fmt.Println("Enter a number from the available options.")
		}
	}

	// Enter members spendings
	for {
		fmt.Println("Choose an option:\n(1) Enter Member and the amount they've spent\n" +
			"(2) Proceed to Saving step")
		var opt string
		fmt.Scanln(&opt)
		if opt == "1" {
			var (
				spender string
				amount  float64
			)
			fmt.Println("Member Name: ")
			fmt.Scanln(&spender)
			if !event.HasMember(spender) {
				fmt.Println("Invalid name. Member is not registered in event.")
			} else {
				fmt.Println("Amount: ")
				fmt.Scanln(&amount)
				event.MemberSpent(spender, amount)
			}
		} else if opt == "2" {
			break
		} else {
			fmt.Println("Enter a number from the available options.")
		}
	}

	// Save to disk
	for {
		fmt.Println("Choose an option:\n(1) Save to disk\n" +
			"(2) Describe split results")
		var opt string
		fmt.Scanln(&opt)
		if opt == "1" {
			var fileName string
			fmt.Println("File name (without extension): ")
			fmt.Scanln(&fileName)

			err := event.SaveToDisk(fileName)
			if err != nil {
				return nil, err
			}
		} else if opt == "2" {
			break
		} else {
			fmt.Println("Enter a number from the available options.")
		}
	}

	return event, nil
}
