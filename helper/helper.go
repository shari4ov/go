package helper

func UserInputValid(firstName string,lastName string,remainingTicket uint,userTicket uint) (bool,bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	checkTicket := userTicket <= remainingTicket && userTicket > 0
	return isValidName,checkTicket
}
