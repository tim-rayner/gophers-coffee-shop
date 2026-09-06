package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var order Order
	serving := true

	CafeMenu.PrintMenu()
	fmt.Println("Welcome to Gophers Coffee Shop!")

	if scanner.Err() != nil {
		return
	}

	for serving {
		var currentItem = OrderItem{}
		var usersChoice string

		for currentItem.drink.name == "" {
			fmt.Println("Lets start with your drink, what can I get you?")
			if scanner.Scan() {
				usersChoice = strings.TrimSpace(scanner.Text())
			}

			if ok := ValdidateItem(&currentItem, usersChoice, "drink"); !ok {
				fmt.Printf("We don't have %v, sorry\ntry again\n", usersChoice)
			}
		}

		for currentItem.syrup.name == "" {
			fmt.Println("And which syrup would you like?")
			if scanner.Scan() {
				usersChoice = strings.TrimSpace(scanner.Text())
			}

			if ok := ValdidateItem(&currentItem, usersChoice, "syrup"); !ok {
				fmt.Println("We don't have that syrup")
			}
		}

		for currentItem.size.name == "" {
			fmt.Println("And what size is that in?")
			if scanner.Scan() {
				usersChoice = strings.TrimSpace(scanner.Text())
			}

			if ok := ValdidateItem(&currentItem, usersChoice, "size"); !ok {
				fmt.Println("We don't sell drinks in that size")
			}
		}

		for currentItem.temperature.name == "" {
			fmt.Println("And is that Hot, Warm, or Over Ice?")
			if scanner.Scan() {
				usersChoice = strings.TrimSpace(scanner.Text())
			}

			if ok := ValdidateItem(&currentItem, usersChoice, "temperature"); !ok {
				fmt.Println("We don't do that temperature, sorry")
			}
		}

		order.Items = append(order.Items, currentItem)
		fmt.Printf("\n%v coming right up!\n", currentItem)

		fmt.Println("Anything else?")
		if scanner.Scan() {
			usersChoice = strings.TrimSpace(scanner.Text())
		}

		if usersChoice == "no" {
			serving = false
			break
		}
	}

	fmt.Printf("Your order comes to £%.2f\n", CalculateTotal(order))
	order.PrintReceipt()

	fmt.Printf("your order will be with you in about %d seconds, speedy service I know", rand.IntN(4) + 2)

}
