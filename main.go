package main

import (
	"go/helper"
	"fmt" 
	"strings"
)
var remainingTicket uint = 10
var users[]string
func main() {
	greetUsers()
	for remainingTicket>0{	
		firstName,lastName,userTicket := getUserInput()
		isValidName,checkTicket := helper.UserInputValid(firstName,lastName,remainingTicket,userTicket)
		if isValidName && checkTicket{
			remainingTicket=remainingTicket-uint(userTicket)
			users=append(users,firstName + " " +lastName)
			firstNames := printFirstNames()
			fmt.Printf("Tickets are boguth by %v",firstNames)
			fmt.Printf("Remaining tickets: %v\n",remainingTicket)
		} else {
			fmt.Printf("We only have %v tickets. ",remainingTicket)
			continue
		}
	
	}
}
func greetUsers(){
	fmt.Printf("We have %v tickets for you, enter how much you want \n" ,remainingTicket)
}
func printFirstNames() []string{
	firstNames := []string{}
	for _,ticket := range users{
		var names = strings.Fields(ticket)
		firstNames=append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput()(string,string,uint){
	var firstName string
	var lastName string
	var userTicket uint
	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter how much ticket you want: ")
	fmt.Scan(&userTicket)
	return firstName,lastName,userTicket
}