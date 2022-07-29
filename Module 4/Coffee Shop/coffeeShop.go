/*
Jordan Satterfield
6/30/22
Write a program that will calculate the cost of a custom cup of coffee at a gourmet coffee shop, based on the size of the cup, the type of coffee selected, and flavors that can be added to the coffee. It should complete the following steps:

    Ask the user what size cup they want, choosing between small, medium, and large.
    Ask the user what kind of coffee they want, choosing between brewed, espresso, and cold brew.
    Ask the user what flavoring they want, if any. Choices include hazelnut, vanilla, and caramel.
    Calculate the price of the cup using the following values:
        Size:
            small: $2
            medium: $3
            large: $4
        Type:
            brewed: no additional cost
            espresso: 50 cents
            cold brew: $1
        Flavoring:
            None: no additional cost
            All other options: 50 cents
    Display a statement that summarizes what the user ordered.
    Display the total cost of the cup of coffee as well as the cost with a 15% tip, in phrases that explain the values to the user. Round the cost with tip to two decimal places.
*/

package main

import (
	"fmt"
	"strings"
)

func main(){
	var size, kind, flavoring string
	var cost float32 = 0

	// Size
	var validInput bool = false
	fmt.Print("What size of coffee (small, medium, or large): ")
	for validInput == false {
		fmt.Scan(&size)
		if  strings.ToLower(size) == "small" || strings.ToLower(size) == "medium" || strings.ToLower(size) == "large"{
			validInput = true
			break
		}
		fmt.Print("You must enter small, medium, or large: ")
	}

	// kind
	validInput = false
	fmt.Print("What kind of coffee (brewed, espresso, and cold brew): ")
	for validInput == false {
		fmt.Scan(&kind)
		if  strings.ToLower(kind) == "brewed" || strings.ToLower(kind) == "espresso" || strings.ToLower(kind) == "cold brew"{
			validInput = true
			break
		}
		fmt.Print("You must enter brewed, espresso, and cold brew: ")
	}


	// flavoring
	validInput = false
	fmt.Print("What kind of flavoring (hazelnut, vanilla, caramel, or none): ")
	for validInput == false {
		fmt.Scan(&flavoring)
		if  strings.ToLower(flavoring) == "hazelnut" || strings.ToLower(flavoring) == "vanilla" || strings.ToLower(flavoring) == "caramel" || strings.ToLower(flavoring) == "none"{
			validInput = true
			break
		}
		fmt.Print("You must enter hazelnut, vanilla, caramel, or none: ")
	}

	// size calc
	if size == "small"{
		cost += 2
	}else if size == "medium"{
		cost += 3
	}else {
		cost += 4
	}

	// kind calc
	if kind == "espresso"{
		cost+= 0.5
	}else if kind == "cold brew"{
		cost += 1
	}

	// flavoring calc
	if flavoring != "none"{
		cost += 0.5
	}

	fmt.Printf("\nYour order: %s, %s, %s coffee\nYour total is $%.2f\nYour total with 15 percent tip is $%.2f", size, kind, flavoring,cost, cost*1.15)

}