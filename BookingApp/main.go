package main

import "fmt"

func main() {
	confName := "Go Conference"
	const confTickets = 50
	var remainingTickets = confTickets

	fmt.Printf("Welcome to %v booking application!\n", confName)
	fmt.Println("We have a total of ", confTickets, "tackets and ", remainingTickets, "are available")
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	// ask for user name
	fmt.Scan(&userName)
	fmt.Scan(&userTickets)

	fmt.Print(userName)
	fmt.Print(&userName)

	fmt.Printf("User %v, tickets %v", userName, userTickets)
}
