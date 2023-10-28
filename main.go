package main // This line organises the application into a package

import "fmt" // All functions in Go come from a package, in this case the Print function comes from the fmt package (fmt is a core Go formatting package)

func main() { // This function tells the compiler where to start running the application's code. There can only be one of these 'main' functions.
	var conferenceName = "Go Conference" // This is an example of a variable
	var remainingTickets = 50
	const conferenceTickets = 50 // This is an example of a constant

	fmt.Printf("Welcome to %v booking application\n", conferenceName) // You must add the package name before a funtion to tell the app where that function comes from.
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v has booked %v tickets\n", userName, userTickets)

}
