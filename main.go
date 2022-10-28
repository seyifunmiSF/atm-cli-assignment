package main

import (
	"fmt"
	"os"
)

var ATM = []string{"Default ATM"} 
var defaultPin = "1234"
var accountBalance = 1000000

func main(){
  welcome()
  for {
  fmt.Println("Enter your 4-digit pin")
  var inputPin string 
  _, err := fmt.Scan(&inputPin)
  if err != nil {
   fmt.Println("Error: please enter correct pin")
  }

 pin := confirmPin(inputPin) 
  if pin == false{
    continue 
  } 


  displayMenu()
  var menuNumber int
  _, error := fmt.Scan(&menuNumber)
  if error != nil {
   fmt.Println("Error: please input correct pin ")
  }


  switch menuNumber {
	case 1:
		changePin()
	case 2:
		checkAccountBalance()
	case 3:
		withdrawFunds()
	case 4:
		depositFunds()
	case 5:
		exit()
	default:
		fmt.Println("Error: Invalid input")
  }
  value := anotherOperation() 
  if value == true {
    continue 
  } else {
    exit()
  }
}
}

func welcome() {
	//pin  ="1234"
	fmt.Println("******{Welcome, user}******")
    newLine(1)
}

func confirmPin(pin string) bool{
	if pin != defaultPin{
		fmt.Println("Invalid Pin")
    return false
	} else {
    return true
  }

}
func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Println("\n")
		i++
	}
}

func printAtm() {
	fmt.Println("List of items in your ATM are:")
	newLine(2)
 for item, index := range ATM {
	 fmt.Printf("ATM #%v : %v", index, item)
 }
 newLine(2)
}

func displayMenu() {
	fmt.Println("Select Operation:")
	fmt.Println("1. Change Pin \t\t\t 2. Account Balance")
	fmt.Println("3. Withdraw funds \t\t 4. Deposit Funds")
	fmt.Println("5. Exit \t")
}

func anotherOperation() bool{
 fmt.Println("Would you like to perform another operation? Yes or No")
 var input string
 _, err := fmt.Scan(&input)
 if err != nil {
  fmt.Println("Error: Invalid input!")
 }

 if input == "Yes" || input == "YES" || input == "yes" {
   return true
 } else {
     return false
   }
 
  }

func changePin() {
	fmt.Println("Enter your old pin")
  var oldPin string 
  _, err := fmt.Scan(&oldPin)
  if err != nil {
   fmt.Println("Error: please enter correct pin ")
  }
  
  fmt.Println("Enter your new 4-digit pin")
  var newPin string 
  _, error := fmt.Scan(&newPin)
  if error != nil {
   fmt.Println("Error: please enter correct pin")
  }
 if oldPin != defaultPin {
	 fmt.Println("Invalid pin, try again")
	 exit()
 }
 defaultPin = newPin
 fmt.Println("Pin changed successfully") 
}

func checkAccountBalance() {
	fmt.Println("Your Account Balance is:", accountBalance)
}

func withdrawFunds() {
	fmt.Println("Enter amount")
  var amount int
  _, error := fmt.Scan(&amount)
  if error != nil {
   fmt.Println("Error: please input correct amount")
  }

  if amount > accountBalance {
	  fmt.Println("Insufficient Balance")
	  exit()
  }

  accountBalance = accountBalance - amount 
  fmt.Println("withdrawal successful")
  fmt.Println("Your balance is:", accountBalance)
}

func depositFunds() {
	fmt.Println("Enter amount")
  var amount int
  _, error := fmt.Scan(&amount)
  if error != nil {
   fmt.Println("Error: please insert correct amount ")
  }
  
  accountBalance = accountBalance + amount 
  fmt.Println("deposit successful")
  fmt.Println("Your balance is:", accountBalance)
}

func exit() {
	fmt.Println("Goodbye, user 👋🏽")
	os.Exit(5)
}